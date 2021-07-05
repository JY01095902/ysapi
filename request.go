package ysapi

import (
	"fmt"
	neturl "net/url"

	"github.com/go-resty/resty/v2"
)

type request struct {
	appKey    string
	appSecret string
}

func NewRequest(appKey, appSecret string) request {
	req := request{
		appKey:    appKey,
		appSecret: appSecret,
	}

	return req
}

func (req request) execute(r *resty.Request, method, url string) (Values, error) {
	token, err := req.getToken()
	if err != nil {
		return nil, err
	}

	resp, err := r.
		SetQueryParam("access_token", neturl.QueryEscape(token)).
		SetResult(Values{}).
		Execute(method, url)

	if err != nil {
		return nil, fmt.Errorf("%w error: %s", ErrCallYonSuiteAPIFailed, err.Error())
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("%w error: %s", ErrCallYonSuiteAPIFailed, resp.String())
	}

	result, ok := resp.Result().(*Values)
	if !ok {
		return Values{}, fmt.Errorf("%w: type of result is not Values", ErrYonSuiteAPIBizError)
	}

	if err := checkAPIResponse(*result); err != nil {
		return Values{}, err
	}

	return *result, err
}

func (req request) Post(url string, body interface{}) (Values, error) {
	r := resty.New().R().
		EnableTrace().
		SetBody(body)

	return req.execute(r, resty.MethodPost, url)
}

func (req request) Get(url string, params map[string]string) (Values, error) {
	r := resty.New().R().
		EnableTrace().
		SetQueryParams(params)

	return req.execute(r, resty.MethodGet, url)
}

func (req request) getToken() (string, error) {
	return getToken(req.appKey, req.appSecret, tokenURL)
}
