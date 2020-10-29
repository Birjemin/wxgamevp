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

// TestGetQueryParams
func TestGetQueryParams(t *testing.T) {
	ast := assert.New(t)
	balance := Balance{
		OpenID:       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		AppID:        "wx1234567",
		OfferID:      "12345678",
		Ts:           1507530737,
		ZoneID:       "1",
		Pf:           "android",
		AccessToken:  "ACCESSTOKEN",
		Secret:       "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
	}

	query := balance.getQueryParams()
	ast.Equal("1ad64e8dcb2ec1dc486b7fdf01f4a15159fc623dc3422470e51cf6870734726b", query["sig"])
}

func TestGetBalance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getBalanceURI {
			t.Fatalf("path is invalid: %s, %s'", getBalanceURI, path)
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

		raw := `{"errcode":0,"errmsg":"","balance":1,"gen_balance":1,"first_save":1,"save_amt":1,"save_sum":1,"cost_sum":1,"present_sum":1}`
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

	balance := Balance{
		OpenID:       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		AppID:        "wx1234567",
		OfferID:      "12345678",
		Ts:           1507530737,
		ZoneID:       "1",
		Pf:           "android",
		AccessToken:  "ACCESSTOKEN",
		Secret:       "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
		HTTPRequest:  httpClient,
	}

	if ret, err := balance.doGetBalance(ts.URL); err != nil {
		t.Error(err)
	} else {
		if ret.ErrCode != 0 {
			t.Error(errors.New("msg: " + ret.ErrMsg))
		}
	}

}
