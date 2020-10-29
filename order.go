package wxgamevp

import (
	"fmt"
	"github.com/birjemin/wxgamevp/utils"
	"log"
)

// Order model
type Order struct {
	AppID       string
	OrderNo     string
	OutTradeNo  string
	AccessToken string
	HTTPRequest *utils.HTTPClient
	Debug       bool
}

// RespOrder response
type RespOrder struct {
	CommonError
	PayForOrder OrderDetail `json:"payfororder"`
}

// OrderDetail struct
type OrderDetail struct {
	OutTradeNo string `json:"out_trade_no"`
	OrderNo    string `json:"order_no"`
	OpenID     string `json:"openid"`
	CreateTime int    `json:"create_time"`
	Amount     int    `json:"amount"`
	Status     int    `json:"status"`
	ZoneID     string `json:"zone_id"`
	Env        int    `json:"env"`
	PayTime    int    `json:"pay_time"`
}

// GetOrder get order
func (b *Order) GetOrder() (*RespOrder, error) {
	return b.doGetOrder(wechatDomain)
}

// doGetOrder
func (b *Order) doGetOrder(domain string) (*RespOrder, error) {
	params := b.getQueryParams()
	jsonStr, err := jsonIter.Marshal(params)
	if err != nil {
		log.Println("[order]doGetOrder, json marshal failed", err, string(jsonStr))
		return nil, err
	}

	url := fmt.Sprintf("%s%s?access_token=%s", domain, b.getOrderURI(), b.AccessToken)

	// log.Println("post url: ", url)
	// log.Println("post str: ", string(jsonStr))

	if err := b.HTTPRequest.HTTPPostJSON(url, string(jsonStr)); err != nil {
		log.Println("[order]doGetOrder, post failed", err)
		return nil, err
	}

	var respOrder = new(RespOrder)
	if err = b.HTTPRequest.GetResponseJSON(respOrder); err != nil {
		log.Println("[order]doGetOrder, response json failed", err)
		return nil, err
	}
	return respOrder, nil
}

// getQueryParams
func (b *Order) getQueryParams() map[string]string {
	params := make(map[string]string, 3)
	params["appid"] = b.AppID
	params["order_no"] = b.OrderNo
	params["out_trade_no"] = b.OutTradeNo
	return params
}

// getOrderURI
func (b *Order) getOrderURI() string {
	if b.Debug {
		return getSandboxOrderURI
	}
	return getOrderURI
}
