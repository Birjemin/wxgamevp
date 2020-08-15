package wxgamevp

import (
	"fmt"
	"github.com/birjemin/wxgamevp/utils"
	"github.com/spf13/cast"
	"log"
)

// Present model
type Present struct {
	OpenID        string
	AppID         string
	OfferID       string
	Ts            int
	ZoneID        string
	Pf            string
	UserIP        string
	BillNo        string
	PresentCounts int
	AccessToken   string
	Secret        string
	SessionToken  string
	HTTPRequest   *utils.HTTPClient
	Debug         bool
}

// RespPresent response
type RespPresent struct {
	CommonError
	Balance string `json:"balance"`
	BillNo  string `json:"bill_no"`
}

// Present pay
func (p *Present) Present() (*RespPay, error) {
	return p.doPresent(wechatDomain)
}

// doPresent
func (p *Present) doPresent(domain string) (*RespPay, error) {
	params := p.getQueryParams()
	jsonStr, err := jsonIter.Marshal(params)
	if err != nil {
		log.Println("[pay]Pay, json marshal failed", err, string(jsonStr))
		return nil, err
	}
	url := fmt.Sprintf("%s%s?access_token=%s", domain, p.getPresentURI(), p.AccessToken)

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
func (p *Present) getQueryParams() map[string]string {
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
	params["present_counts"] = cast.ToString(p.PresentCounts)
	params["sig"] = GenerateSign(p.getPresentURI(), "POST", "secret", p.Secret, params)
	params["access_token"] = p.AccessToken
	params["mp_sig"] = GenerateSign(p.getPresentURI(), "POST", "session_key", p.SessionToken, params)
	delete(params, "access_token")
	return params
}

func (p *Present) getPresentURI() string {
	if p.Debug {
		return getSandboxPresentURI
	}
	return getPresentURI
}
