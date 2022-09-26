package returnorder

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type ModifyRequest struct {
	AppKey       string
	AppSecret    string
	Ids          string
	PartParam    request.Values
	ExternalData struct {
		Head     []request.Values
		Added    []request.Values
		Modified []request.Values
	}
}

func (req ModifyRequest) ToValues() request.Values {
	values := request.Values{
		"ids":       req.Ids,
		"partParam": req.PartParam,
		"externalData": request.Values{
			"Head":     req.ExternalData.Head,
			"Added":    req.ExternalData.Added,
			"Modified": req.ExternalData.Modified,
		},
	}

	return values
}

/*
{
    "code": "200",
    "message": "操作成功",
    "data": [
        {
            "code": "1",
            "isShowMsg": true,
            "externalMap": {},
            "failCount": "0",
            "sucIdAndPubts": {
                "1510014852506583050": "2022-09-26 11:44:15"
            },
            "successCount": "1",
            "isExcuteAction": true,
            "actionName": "退换货修改"
        },
        {
            "code": "1",
            "isShowMsg": false,
            "externalMap": {},
            "failCount": "0",
            "sucIdAndPubts": {
                "1510014852506583050": "2022-09-26 11:44:15"
            },
            "successCount": "0",
            "isExcuteAction": true,
            "actionName": "提交存量"
        }
    ]
}
*/

type ModifyResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    []request.Values `json:"data"`
}

func Modify(req ModifyRequest) (ModifyResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/returnorder/modify", req.ToValues())
	if err != nil {
		return ModifyResponse{}, err
	}

	res, err := vals.GetResult(ModifyResponse{})
	if err != nil {
		return ModifyResponse{}, err
	}

	resp, ok := res.(*ModifyResponse)
	if !ok {
		return ModifyResponse{}, fmt.Errorf("%w error: response is not type of ModifyResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
