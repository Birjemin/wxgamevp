package wxgamevp

import (
	"fmt"
	"github.com/birjemin/wxgamevp/utils"
)

// GenerateSign generate sign
func GenerateSign(uri, method, secretName, secret string, params map[string]string) string {
	uriStr, paramsStr := generateURIString(uri, method, secretName, secret), generateParamsString(params)
	return utils.GenerateSha256(secret, fmt.Sprintf("%s%s", paramsStr, uriStr))
}

// generateURIString
func generateURIString(uri, method, secretName, secret string) string {
	return fmt.Sprintf("&org_loc=%s&method=%s&%s=%s", uri, method, secretName, secret)
}

// generateParamsString
func generateParamsString(params map[string]string) string {
	return utils.QuerySortByKeyStr(params)
}
