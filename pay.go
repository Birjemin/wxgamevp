package wxgamevp

import (
	"fmt"
	"github.com/birjemin/wxgamevp/utils"
	"github.com/spf13/cast"
	"log"
)

// Pay model
type Pay struct {
	OpenID       string
	AppID        string
	OfferID      string
	Ts           int
	ZoneID       string
	Pf           string
	UserIP       string
	Amt          int
	BillNo       string
	PayItem      string
	AppRemark    string
	AccessToken  string
	Secret       string
	SessionToken string
	HTTPRequest  *utils.HTTPClient
	Debug        bool
}

// RespPay response
type RespPay struct {
	CommonError
	BillNo     string `json:"bill_no"`
	Balance    int    `json:"balance"`
	UsedGenAmt int    `json:"used_gen_amt"`
}

// Pay pay
func (p *Pay) Pay() (*RespPay, error) {
	return p.doPay(wechatDomain)
}

// doGetBalance
func (p *Pay) doPay(domain string) (*RespPay, error) {
	params := p.getQueryParams()
	jsonStr, err := jsonIter.Marshal(params)
	if err != nil {
		log.Println("[pay]Pay, json marshal failed", err, string(jsonStr))
		return nil, err
	}
	url := fmt.Sprintf("%s%s?access_token=%s", domain, p.getPayURI(), p.AccessToken)

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
func (p *Pay) getQueryParams() map[string]string {
	params := make(map[string]string, 9)
	params["openid"] = p.OpenID
	params["appid"] = p.AppID
	params["offer_id"] = p.OfferID
	params["ts"] = cast.ToString(p.Ts)
	params["zone_id"] = p.ZoneID
	params["pf"] = p.Pf
	params["amt"] = cast.ToString(p.Amt)
	params["bill_no"] = p.BillNo
	if p.UserIP != "" {
		params["user_ip"] = p.UserIP
	}
	if p.PayItem != "" {
		params["pay_item"] = p.PayItem
	}
	if p.AppRemark != "" {
		params["app_remark"] = p.AppRemark
	}
	params["sig"] = GenerateSign(p.getPayURI(), "POST", "secret", p.Secret, params)
	params["access_token"] = p.AccessToken
	params["mp_sig"] = GenerateSign(p.getPayURI(), "POST", "session_key", p.SessionToken, params)
	delete(params, "access_token")
	return params
}

func (p *Pay) getPayURI() string {
	if p.Debug {
		return getSandboxPayURI
	}
	return getPayURI
}
