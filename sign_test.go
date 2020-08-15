package wxgamevp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGenerateParamsString
func TestGenerateParamsString(t *testing.T) {
	ast := assert.New(t)
	ast.Equal("appid=wx1234567&offer_id=12345678&openid=odkx20ENSNa2w5y3g_qOkOvBNM1g&pf=android&ts=1507530737&zone_id=1",
		generateParamsString(map[string]string{
			"openid":   "odkx20ENSNa2w5y3g_qOkOvBNM1g",
			"appid":    "wx1234567",
			"offer_id": "12345678",
			"ts":       "1507530737",
			"zone_id":  "1",
			"pf":       "android",
		}))

	ast.Equal("access_token=ACCESSTOKEN&appid=wx1234567&offer_id=12345678&openid=odkx20ENSNa2w5y3g_qOkOvBNM1g&pf=android&sig=1ad64e8dcb2ec1dc486b7fdf01f4a15159fc623dc3422470e51cf6870734726b&ts=1507530737&zone_id=1",
		generateParamsString(map[string]string{
			"access_token": "ACCESSTOKEN",
			"openid":       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
			"appid":        "wx1234567",
			"offer_id":     "12345678",
			"ts":           "1507530737",
			"zone_id":      "1",
			"pf":           "android",
			"sig":          "1ad64e8dcb2ec1dc486b7fdf01f4a15159fc623dc3422470e51cf6870734726b",
		}))
}

// TestGenerateUriString
func TestGenerateUriString(t *testing.T) {
	ast := assert.New(t)
	ast.Equal("&org_loc=/cgi-bin/midas/getbalance&method=POST&secret=zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
		generateURIString("/cgi-bin/midas/getbalance", "POST", "secret", "zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u"),
	)

	ast.Equal("&org_loc=/cgi-bin/midas/getbalance&method=POST&session_key=V7Q38/i2KXaqrQyl2Yx9Hg==",
		generateURIString("/cgi-bin/midas/getbalance", "POST", "session_key", "V7Q38/i2KXaqrQyl2Yx9Hg=="),
	)
}

// TestGenerateSign
func TestGenerateSign(t *testing.T) {
	ast := assert.New(t)
	ast.Equal("1ad64e8dcb2ec1dc486b7fdf01f4a15159fc623dc3422470e51cf6870734726b",
		GenerateSign(
			"/cgi-bin/midas/getbalance",
			"POST",
			"secret",
			"zNLgAGgqsEWJOg1nFVaO5r7fAlIQxr1u",
			map[string]string{
				"openid":   "odkx20ENSNa2w5y3g_qOkOvBNM1g",
				"appid":    "wx1234567",
				"offer_id": "12345678",
				"ts":       "1507530737",
				"zone_id":  "1",
				"pf":       "android",
			},
		),
	)

	ast.Equal("ff4c5bb39dea1002a8f03be0438724e1a8bcea5ebce8f221f9b9fea3bcf3bf76",
		GenerateSign(
			"/cgi-bin/midas/getbalance",
			"POST",
			"session_key",
			"V7Q38/i2KXaqrQyl2Yx9Hg==",
			map[string]string{
				"access_token": "ACCESSTOKEN",
				"openid":       "odkx20ENSNa2w5y3g_qOkOvBNM1g",
				"appid":        "wx1234567",
				"offer_id":     "12345678",
				"ts":           "1507530737",
				"zone_id":      "1",
				"pf":           "android",
				"sig":          "1ad64e8dcb2ec1dc486b7fdf01f4a15159fc623dc3422470e51cf6870734726b",
			},
		),
	)
}
