package wxgamevp

import (
	"fmt"
	"github.com/birjemin/wxgamevp/utils"
	"github.com/spf13/cast"
	"log"
)

// Balance model
type Balance struct {
	OpenID       string
	AppID        string
	OfferID      string
	Ts           int
	ZoneID       string
	Pf           string
	UserIP       string
	AccessToken  string
	Secret       string
	HTTPRequest  *utils.HTTPClient
	Debug        bool
}

// RespBalance response
type RespBalance struct {
	CommonError
	Balance    int `json:"balance"`
	GenBalance int `json:"gen_balance"`
	FirstSave  int `json:"first_save"`
	SaveAmt    int `json:"save_amt"`
	SaveSum    int `json:"save_sum"`
	CostSum    int `json:"cost_sum"`
	PresentSum int `json:"present_sum"`
}

// GetBalance get balance
func (b *Balance) GetBalance() (*RespBalance, error) {
	return b.doGetBalance(wechatDomain)
}

// doGetBalance
func (b *Balance) doGetBalance(domain string) (*RespBalance, error) {
	params := b.getQueryParams()
	jsonStr, err := jsonIter.Marshal(params)
	if err != nil {
		log.Println("[balance]doGetBalance, json marshal failed", err, string(jsonStr))
		return nil, err
	}

	url := fmt.Sprintf("%s%s?access_token=%s", domain, b.getBalanceURI(), b.AccessToken)

	// log.Println("post url: ", url)
	// log.Println("post str: ", string(jsonStr))

	if err := b.HTTPRequest.HTTPPostJSON(url, string(jsonStr)); err != nil {
		log.Println("[balance]doGetBalance, post failed", err)
		return nil, err
	}

	var respBalance = new(RespBalance)
	if err = b.HTTPRequest.GetResponseJSON(respBalance); err != nil {
		log.Println("[balance]doGetBalance, response json failed", err)
		return nil, err
	}
	return respBalance, nil
}

// getQueryParams
func (b *Balance) getQueryParams() map[string]string {
	params := make(map[string]string, 8)
	params["openid"] = b.OpenID
	params["appid"] = b.AppID
	params["offer_id"] = b.OfferID
	params["ts"] = cast.ToString(b.Ts)
	params["zone_id"] = b.ZoneID
	params["pf"] = b.Pf
	if b.UserIP != "" {
		params["user_ip"] = b.UserIP
	}
	params["sig"] = GenerateSign(b.getBalanceURI(), "POST", "secret", b.Secret, params)
	return params
}

// getBalanceURI
func (b *Balance) getBalanceURI() string {
	if b.Debug {
		return getSandboxBalanceURI
	}
	return getBalanceURI
}
