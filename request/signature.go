package request

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

func sign(content, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(content))
	if err != nil {
		return ""
	}

	return url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}

func getQueryString(params url.Values, appSecret string) string {
	rawParamSlice := []string{}
	for k := range params {
		val := params.Get(k)
		rawParamSlice = append(rawParamSlice, k+val)
	}
	sort.Strings(rawParamSlice)
	rawString := strings.Join(rawParamSlice, "")
	signature := sign(rawString, appSecret)

	paramSlice := []string{}
	for k := range params {
		val := params.Get(k)
		paramSlice = append(paramSlice, k+"="+url.QueryEscape(val))
	}
	sort.Strings(paramSlice)
	paramSlice = append(paramSlice, "signature="+signature)

	return strings.Join(paramSlice, "&")
}
