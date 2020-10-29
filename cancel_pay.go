package wxgamevp

import (
	"fmt"
	"github.com/birjemin/wxgamevp/utils"
	"github.com/spf13/cast"
	"log"
)

// CancelPay model
type CancelPay struct {
	OpenID       string
	AppID        string
	OfferID      string
	Ts           int
	ZoneID       string
	Pf           string
	UserIP       string
	BillNo       string
	PayItem      string
	AccessToken  string
	Secret       string
	HTTPRequest  *utils.HTTPClient
	Debug        bool
}

// RespCancelPay response
type RespCancelPay struct {
	CommonError
	BillNo string `json:"bill_no"`
}

// CancelPay pay
func (p *CancelPay) CancelPay() (*RespPay, error) {
	return p.doCancelPay(wechatDomain)
}

// doCancelPay
func (p *CancelPay) doCancelPay(domain string) (*RespPay, error) {
	params := p.getQueryParams()
	jsonStr, err := jsonIter.Marshal(params)
	if err != nil {
		log.Println("[pay]Pay, json marshal failed", err, string(jsonStr))
		return nil, err
	}
	url := fmt.Sprintf("%s%s?access_token=%s", domain, p.getCancelPayURI(), p.AccessToken)

	// log.Println("post url: ", url)
	// log.Println("post str: ", string(jsonStr))

	if err := p.HTTPRequest.HTTPPostJSON(url, string(jsonStr)); err != nil {
		log.Println("[pay]Pay, post failed", err)
		return nil, err
	}

	var respPay = new(RespPay)
	if err = p.HTTPRequest.GetResponseJSON(respPay); err != nil {
		log.Println("[pay]Pay, response json failed", err)
		return respPay, err
	}
	return respPay, nil
}

// getQueryParams
func (p *CancelPay) getQueryParams() map[string]string {
	params := make(map[string]string, 9)
	params["openid"] = p.OpenID
	params["appid"] = p.AppID
	params["offer_id"] = p.OfferID
	params["ts"] = cast.ToString(p.Ts)
	params["zone_id"] = p.ZoneID
	params["pf"] = p.Pf
	params["bill_no"] = p.BillNo
	if p.UserIP != "" {
		params["user_ip"] = p.UserIP
	}
	if p.PayItem != "" {
		params["pay_item"] = p.PayItem
	}
	params["sig"] = GenerateSign(p.getCancelPayURI(), "POST", "secret", p.Secret, params)
	return params
}

func (p *CancelPay) getCancelPayURI() string {
	if p.Debug {
		return getSandboxCancelPayURI
	}
	return getCancelPayURI
}
