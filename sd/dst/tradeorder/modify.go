package tradeorder

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	}

	extdata := request.Values{
		"Head": req.ExternalData.Head,
	}

	if req.ExternalData.Added != nil {
		extdata["Added"] = req.ExternalData.Added
	}

	if req.ExternalData.Modified != nil {
		extdata["Modified"] = req.ExternalData.Modified
	}

	values["externalData"] = extdata

	return values
}

/*
错误的
{
    "code": "200",
    "message": "[{\"actionName\":\"订单修改\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1632373780662714390\":\"2023-01-09 11:04:34\"},\"successCount\":\"1\"},{\"actionName\":\"自动匹配\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"1\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{},\"successCount\":\"0\"}]",
    "data": null
}

格式化后
[
    {
        "actionName": "订单修改",
        "code": "1",
        "externalMap": {},
        "failCount": "0",
        "isExcuteAction": true,
        "isShowMsg": true,
        "sucIdAndPubts": {
            "1632373780662714390": "2023-01-09 11:04:34"
        },
        "successCount": "1"
    },
    {
        "actionName": "自动匹配",
        "code": "1",
        "externalMap": {},
        "failCount": "1",
        "isExcuteAction": true,
        "isShowMsg": false,
        "sucIdAndPubts": {},
        "successCount": "0"
    }
]

*/

/*
正确的
{
    "code": "200",
    "message": "[{\"actionName\":\"订单修改\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1632373780662714390\":\"2023-01-09 11:09:53\"},\"successCount\":\"1\"},{\"actionName\":\"自动匹配\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1632373780662714390\":\"2023-01-09 11:09:54\"},\"successCount\":\"1\"},{\"actionName\":\"重量重算\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1632373780662714390\":\"2023-01-09 11:09:54\"},\"successCount\":\"1\"},{\"actionName\":\"表头汇总表体仓库\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1632373780662714390\":\"2023-01-09 11:09:54\"},\"successCount\":\"1\"},{\"actionName\":\"提交存量\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1632373780662714390\":\"2023-01-09 11:09:54\"},\"successCount\":\"1\"}]",
    "data": null
}

格式化后
[
    {
        "actionName": "订单修改",
        "code": "1",
        "externalMap": {},
        "failCount": "0",
        "isExcuteAction": true,
        "isShowMsg": true,
        "sucIdAndPubts": {
            "1632366221520273433": "2023-01-09 10:56:58"
        },
        "successCount": "1"
    },
    {
        "actionName": "自动匹配",
        "code": "1",
        "externalMap": {},
        "failCount": "0",
        "isExcuteAction": true,
        "isShowMsg": false,
        "sucIdAndPubts": {
            "1632366221520273433": "2023-01-09 10:56:59"
        },
        "successCount": "1"
    },
    {
        "actionName": "重量重算",
        "code": "1",
        "externalMap": {},
        "failCount": "0",
        "isExcuteAction": true,
        "isShowMsg": false,
        "sucIdAndPubts": {
            "1632366221520273433": "2023-01-09 10:56:59"
        },
        "successCount": "1"
    },
    {
        "actionName": "表头汇总表体仓库",
        "code": "1",
        "externalMap": {},
        "failCount": "0",
        "isExcuteAction": true,
        "isShowMsg": false,
        "sucIdAndPubts": {
            "1632366221520273433": "2023-01-09 10:56:59"
        },
        "successCount": "1"
    },
    {
        "actionName": "提交存量",
        "code": "1",
        "externalMap": {},
        "failCount": "0",
        "isExcuteAction": true,
        "isShowMsg": false,
        "sucIdAndPubts": {
            "1632366221520273433": "2023-01-09 10:56:59"
        },
        "successCount": "1"
    }
]
*/
type ModifyResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func (resp ModifyResponse) IsSuccessed() (bool, string) {
	type action struct {
		ExceptionMsg   string         `json:"exceptionMsg"`
		Code           string         `json:"code"`
		IsShowMsg      bool           `json:"isShowMsg"`
		FailCount      string         `json:"failCount"`
		SucIdAndPubts  request.Values `json:"sucIdAndPubts"`
		SuccessCount   string         `json:"successCount"`
		IsExcuteAction bool           `json:"isExcuteAction"`
		ActionName     string         `json:"actionName"`
	}

	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return false, err.Error()
	}

	actionSuccess := map[string]bool{}
	for _, action := range actions {
		isSuccess := false
		if action.SuccessCount == "1" {
			isSuccess = true
		}

		if action.FailCount == "1" {
			isSuccess = false
		}
		actionSuccess[action.ActionName] = isSuccess
	}

	isModifyOK := false
	if isSuccess, exist := actionSuccess["订单修改"]; exist {
		isModifyOK = isSuccess
	}

	isAutoMatchOK := false
	if isSuccess, exist := actionSuccess["自动匹配"]; exist {
		isAutoMatchOK = isSuccess
	}

	if isModifyOK && isAutoMatchOK {
		return true, ""
	}

	return false, resp.Message
}

func (resp ModifyResponse) Timestamp() string {
	if ok, _ := resp.IsSuccessed(); !ok {
		return ""
	}

	type action struct {
		ExceptionMsg   string         `json:"exceptionMsg"`
		Code           string         `json:"code"`
		IsShowMsg      bool           `json:"isShowMsg"`
		FailCount      string         `json:"failCount"`
		SucIdAndPubts  request.Values `json:"sucIdAndPubts"`
		SuccessCount   string         `json:"successCount"`
		IsExcuteAction bool           `json:"isExcuteAction"`
		ActionName     string         `json:"actionName"`
	}

	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return ""
	}
	for _, action := range actions {
		if action.ActionName == "自动匹配" {
			for _, v := range action.SucIdAndPubts {
				switch val := v.(type) {
				case int:
					return strconv.Itoa(val)
				case int64:
					return strconv.FormatInt(val, 10)
				case float64:
					return strconv.FormatInt(int64(val), 10)
				case string:
					return val
				default:
					return ""
				}
			}
		}
	}

	return ""
}

func Modify(req ModifyRequest) (ModifyResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/tradeorder/modify", req.ToValues())
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
