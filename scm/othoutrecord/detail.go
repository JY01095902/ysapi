package othoutrecord

import (
	"fmt"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type DetailRequest struct {
	AppKey    string
	AppSecret string
	Id        string
}

type DetailResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func Detail(req DetailRequest) (DetailResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Get(request.URLRoot+"/yonbip/scm/othoutrecord/detail", map[string]string{
		"id": req.Id,
	})
	if err != nil {
		return DetailResponse{}, err
	}

	resp, err := vals.GetResult(DetailResponse{})
	if err != nil {
		return DetailResponse{}, err
	}

	res, ok := resp.(*DetailResponse)
	if !ok {
		return DetailResponse{}, fmt.Errorf("%w error: response is not type of DetailResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if res.Code != "200" {
		return *res, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, res.Message)
	}

	return *res, nil
}
