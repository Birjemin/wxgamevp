package utils

import (
	"net/url"
)

// QuerySortByKeyStr query sort by key str2
func QuerySortByKeyStr(params map[string]string) (str string) {
	// or you can create new url.Values struct and encode that like so
	q := url.Values{}
	for key, val := range params {
		q.Add(key, val)
	}
	return q.Encode()
}
