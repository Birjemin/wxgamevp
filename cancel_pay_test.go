package wxgamevp

import (
	"errors"
	"github.com/birjemin/wxgamevp/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestCancelPayQueryParams test query string
func TestCancelPayQueryParams(t *testing.T) {
	ast := assert.New(t)
	pay := CancelPay{
		OpenID:       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		AppID:        "wx1234567",
		OfferID:      "12345678",
		Ts:           1507530737,
		ZoneID:       "1",
		Pf:           "android",
		BillNo:       "BillNo_123",
		AccessToken:  "ACCESSTOKEN",
		Secret:       "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
		SessionToken: "V7Q38/i2KXaqrQyl2Yx9Hg==",
	}

	query := pay.getQueryParams()
	ast.Equal("cff702559f26433de1df7e20921d5798bf4dc1c7636472a0bec82369a8bb6ba8", query["sig"])
	ast.Equal("35b37a9192e7bbb595627b371ac8dc6ccf879f3d8b9486927708a1de13232219", query["mp_sig"])
	ast.Equal(9, len(query))
}

// TestCancelPay test cancel_pay
func TestCancelPay(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getCancelPayURI {
			t.Fatalf("path is invalid: %s, %s'", getCancelPayURI, path)
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

		raw := `{"errcode":0,"errmsg":"","bill_no":"1"}`
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

	pay := CancelPay{
		OpenID:       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		AppID:        "wx1234567",
		OfferID:      "12345678",
		Ts:           1507530737,
		ZoneID:       "1",
		BillNo:       "BillNo_123",
		Pf:           "android",
		AccessToken:  "ACCESSTOKEN",
		Secret:       "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
		SessionToken: "V7Q38/i2KXaqrQyl2Yx9Hg==",
		HTTPRequest:  httpClient,
	}

	if ret, err := pay.doCancelPay(ts.URL); err != nil {
		t.Error(err)
	} else {
		if ret.ErrCode != 0 {
			t.Error(errors.New("msg: " + ret.ErrMsg))
		}
	}

}
