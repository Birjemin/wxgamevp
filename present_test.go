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

// TestPresentQueryParams test query string
func TestPresentQueryParams(t *testing.T) {
	ast := assert.New(t)
	pay := Present{
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
	ast.Equal("1d3804993d06ae06560eb70ff9b102d7d5a932d99f9d4ef0d0fd07d9a10040d8", query["sig"])
	ast.Equal("3f415a9b2c5e7befebb0db80924c64fb4bdf3456321acae7a8acd6fb85027f43", query["mp_sig"])
	ast.Equal(10, len(query))
}

// TestPresent test cancel_pay
func TestPresent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != getPresentURI {
			t.Fatalf("path is invalid: %s, %s'", getPresentURI, path)
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

		raw := `{"errcode":0,"errmsg":"","bill_no":"1","balance":0}`
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

	pay := Present{
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

	if ret, err := pay.doPresent(ts.URL); err != nil {
		t.Error(err)
	} else {
		if ret.ErrCode != 0 {
			t.Error(errors.New("msg: " + ret.ErrMsg))
		}
	}

}
