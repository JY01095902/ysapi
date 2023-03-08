package morphologyconversion

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type DetailRequest struct {
	AppKey    string
	AppSecret string
	Id        string
	Code      string
}

type DetailResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func Detail(req DetailRequest) (DetailResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	param := map[string]string{
		"code": req.Code,
	}
	if req.Id != "" {
		param = map[string]string{
			"id": req.Id,
		}
	}
	vals, err := apiReq.Get(request.NewURLRoot+"/scm/morphologyconversion/detail", param)
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
