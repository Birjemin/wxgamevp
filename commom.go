package wxgamevp

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	wechatDomain = "https://api.weixin.qq.com"

	getBalanceURI        = "/cgi-bin/midas/getbalance"
	getSandboxBalanceURI = "/cgi-bin/midas/sandbox/getbalance"

	getCancelPayURI        = "/cgi-bin/midas/cancelpay"
	getSandboxCancelPayURI = "/cgi-bin/midas/sandbox/cancelpay"

	getPayURI        = "/cgi-bin/midas/pay"
	getSandboxPayURI = "/cgi-bin/midas/sandbox/pay"

	getPresentURI        = "/cgi-bin/midas/present"
	getSandboxPresentURI = "/cgi-bin/midas/sandbox/present"
)

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

// CommonError model
type CommonError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
