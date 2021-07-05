package ysapi

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

func convertTimestampToTime(timestamp int64) time.Time {
	return time.Unix(0, timestamp*1e6)
}

func convertTimeToTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func saveToken(token string, expireIn int) {
	os.Setenv("YS_TOKEN", token)
	validityDuration := time.Duration(float32(expireIn)*0.8) * time.Second // 应该是 7200 秒，这里缩短过期时间(*0.8)，以免在临界点有问题
	expTime := time.Now().Add(validityDuration)
	expTs := convertTimeToTimestamp(expTime)
	os.Setenv("YS_TOKEN_EXPIRED_TIMESTAMP", strconv.FormatInt(expTs, 10))
}

func fetchToken(appKey, appSecret, tokenURL string) (string, int, error) {
	type Response struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Token    string `json:"access_token"`
			ExpireIn int    `json:"expire"`
		} `json:"data"`
	}

	params := url.Values{}
	params.Add("appKey", appKey)
	params.Add("timestamp", strconv.FormatInt(getTimestamp(), 10))

	cli := resty.New()
	resp, err := cli.R().
		SetQueryString(getQueryString(params, appSecret)).
		SetResult(Response{}).
		Get(tokenURL)

	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode() != 200 {
		return "", 0, fmt.Errorf("%w error: %s", ErrCallYonSuiteAPIFailed, resp.String())
	}

	result, ok := resp.Result().(*Response)
	if !ok {
		return "", 0, fmt.Errorf("%w: result is not token response", ErrYonSuiteAPIBizError)
	}

	if result.Code != "00000" {
		return "", 0, fmt.Errorf("%w: %s", ErrYonSuiteAPIBizError, result.Message)
	}

	return result.Data.Token, result.Data.ExpireIn, nil
}

func isTokenExpired() bool {
	expTs, err := strconv.ParseInt(os.Getenv("YS_TOKEN_EXPIRED_TIMESTAMP"), 10, 64)
	if err != nil {
		return true
	}

	expTime := convertTimestampToTime(expTs)

	return time.Now().After(expTime)
}

func getToken(appKey, appSecret, tokenURL string) (string, error) {
	if isTokenExpired() {
		token, expireIn, err := fetchToken(appKey, appSecret, tokenURL)
		if err != nil {
			return "", err
		}

		saveToken(token, expireIn)

		return token, nil
	}

	return os.Getenv("YS_TOKEN"), nil
}
