package returnorder

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.libratone.com/internet/ysapi.git/request"
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
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		ExceptionMsg   string         `json:"exceptionMsg"`
		Code           string         `json:"code"`
		IsShowMsg      bool           `json:"isShowMsg"`
		ExternalMap    request.Values `json:"externalMap"`
		FailCount      string         `json:"failCount"`
		SucIdAndPubts  request.Values `json:"sucIdAndPubts"`
		SuccessCount   string         `json:"successCount"`
		IsExcuteAction bool           `json:"isExcuteAction"`
		ActionName     string         `json:"actionName"`
	} `json:"data"`
}

func (resp ModifyResponse) IsSuccessed(id string) (bool, string) {
	for _, action := range resp.Data {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "退换货修改" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有退换货修改的事件，把data的内容都返回
	return false, resp.Message
}

func (resp ModifyResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	for _, action := range resp.Data {
		if action.ActionName == "退换货修改" {
			switch val := action.SucIdAndPubts[id].(type) {
			case int:
				return strconv.Itoa(val)
			case int64:
				return strconv.FormatInt(val, 10)
			case float64:
				return strconv.FormatInt(int64(val), 10)
			case string:
				return val
			case json.Number:
				return val.String()
			default:
				return ""
			}
		}
	}

	return ""
}

func Modify(req ModifyRequest) (ModifyResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/returnorder/modify", req.ToValues())
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
