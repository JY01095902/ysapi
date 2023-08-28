package tradevouch

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type ImportRequest struct {
	AppKey       string
	AppSecret    string
	ExternalData []request.Values
	PartParam    request.Values
}

func (req ImportRequest) ToValues() request.Values {
	values := request.Values{
		"partParam":    req.PartParam,
		"externalData": req.ExternalData,
	}

	return values
}

/*
	{
	    "code": "200",
	    "message": "导入原单成功 导入失败0单 ",
	    "data": {
	        "1575944747991171084": "TEST20221025001"
	    }
	}
*/
type ImportResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func (resp ImportResponse) IsSuccessed(tid string) bool {
	if len(resp.Data) == 0 {
		return false
	}

	for _, v := range resp.Data {
		if v == tid {
			return true
		}
	}

	return false
}

func Import(req ImportRequest) (ImportResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/tradevouch/import", req.ToValues())
	if err != nil {
		return ImportResponse{}, err
	}

	res, err := vals.GetResult(ImportResponse{})
	if err != nil {
		return ImportResponse{}, err
	}

	resp, ok := res.(*ImportResponse)
	if !ok {
		return ImportResponse{}, fmt.Errorf("%w error: response is not type of ImportResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
