package request

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (r response) String() string {
	return "[" + r.Code + "]" + r.Message
}

func checkResponseBody(body []byte) error {
	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("%w: type of body is not response", ErrYonSuiteAPIBizError)
	}

	if resp.Code == "200" {
		return nil
	}

	// API 限流
	if resp.Code == "310046" {
		return fmt.Errorf("%w: %s", ErrAPILimit, resp.String())
	}

	return fmt.Errorf("%w: %s", ErrYonSuiteAPIBizError, resp.String())
}

func checkResponse(status int, body []byte) error {
	// API返回200 - 299 的状态码后才会进一步验证body
	if status >= 200 && status <= 299 {
		return checkResponseBody(body)
	}
	// 看body中是否是response类型，是的话就返回message，不是的话返回整个body
	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("%w: %s", ErrCallYonSuiteAPIFailed, string(body))
	}

	if status == 429 {
		return fmt.Errorf("%w: %s", ErrAPILimit, resp.String())
	}

	return fmt.Errorf("%w: %s", ErrCallYonSuiteAPIFailed, resp.String())
}
