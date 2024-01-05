package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	neturl "net/url"

	"github.com/go-resty/resty/v2"
)

type Request struct {
	appKey    string
	appSecret string
}

func New(appKey, appSecret string) Request {
	req := Request{
		appKey:    appKey,
		appSecret: appSecret,
	}

	return req
}

func (req Request) execute(r *resty.Request, method, url string) (Values, error) {
	token, err := req.getToken()
	if err != nil {
		return nil, err
	}

	resp, err := r.
		SetQueryParam("access_token", neturl.QueryEscape(token)).
		Execute(method, url)

	if err != nil {
		return nil, fmt.Errorf("%w error: %s", ErrCallYonSuiteAPIFailed, err.Error())
	}

	err = checkResponse(resp.StatusCode(), resp.Body())
	if err != nil {
		return nil, err
	}

	values := Values{}

	// body是后端的http返回结果
	d := json.NewDecoder(bytes.NewReader(resp.Body()))
	d.UseNumber()
	if err := d.Decode(&values); err != nil {
		return Values{}, fmt.Errorf("%w: decode response body failed, error: %v", ErrYonSuiteAPIBizError, err)
	}

	return values, nil
}

func (req Request) Post(url string, body interface{}) (Values, error) {
	r := resty.New().R().
		EnableTrace().
		SetBody(body)

	return req.execute(r, resty.MethodPost, url)
}

func (req Request) Get(url string, params map[string]string) (Values, error) {
	r := resty.New().R().
		EnableTrace().
		SetQueryParams(params)

	return req.execute(r, resty.MethodGet, url)
}

func (req Request) getToken() (string, error) {
	return getToken(req.appKey, req.appSecret, tokenURL)
}
