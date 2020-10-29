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

// TestGetPayQueryParams
func TestGetPayQueryParams(t *testing.T) {
	ast := assert.New(t)
	pay := Pay{
		OpenID:       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		AppID:        "wx1234567",
		OfferID:      "12345678",
		Ts:           1507530737,
		ZoneID:       "1",
		Pf:           "android",
		Amt:          123,
		BillNo:       "BillNo_123",
		AccessToken:  "ACCESSTOKEN",
		Secret:       "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
	}

	query := pay.getQueryParams()
	ast.Equal("ada5ef41b2f185b9d6ba041ed0f4ba562ed5aeb6acce17be58865fa88c55aaab", query["sig"])
}

func TestPay(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getPayURI {
			t.Fatalf("path is invalid: %s, %s'", getPayURI, path)
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

		raw := `{"errcode":0,"errmsg":"","bill_no":"1","balance":1,"used_gen_amt":1}`
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

	pay := Pay{
		OpenID:       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		AppID:        "wx1234567",
		OfferID:      "12345678",
		Ts:           1507530737,
		ZoneID:       "1",
		Amt:          123,
		BillNo:       "BillNo_123",
		Pf:           "android",
		AccessToken:  "ACCESSTOKEN",
		Secret:       "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
		HTTPRequest:  httpClient,
	}

	if ret, err := pay.doPay(ts.URL); err != nil {
		t.Error(err)
	} else {
		if ret.ErrCode != 0 {
			t.Error(errors.New("msg: " + ret.ErrMsg))
		}
	}

}
