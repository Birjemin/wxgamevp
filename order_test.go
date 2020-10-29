package wxgamevp

import (
	"errors"
	"github.com/birjemin/wxgamevp/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getOrderURI {
			t.Fatalf("path is invalid: %s, %s'", getOrderURI, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("param %v can not be empty", v)
			}
		}

		body, _ := ioutil.ReadAll(r.Body)
		if string(body) == "" {
			t.Fatal("body is empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{"errcode":0,"errmsg":"","payfororder":{"out_trade_no":"1000","order_no":"1001","openid":"1002","create_time":10,"amount":10,"status":10,"zone_id":"10","env":1,"pay_time":10}}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	defer ts.Close()

	httpClient := &utils.HTTPClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	order := Order{
		AppID:       "wx1234567",
		OrderNo:     "1001",
		OutTradeNo:  "1000",
		AccessToken: "ACCESSTOKEN",
		Secret:      "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
		HTTPRequest: httpClient,
	}

	if ret, err := order.doGetOrder(ts.URL); err != nil {
		t.Error(err)
	} else {
		if ret.ErrCode != 0 {
			t.Error(errors.New("msg: " + ret.ErrMsg))
		}
	}

}
