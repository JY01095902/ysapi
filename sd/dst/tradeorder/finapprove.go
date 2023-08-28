package tradeorder

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jy01095902/ysapi/request"
)

type FinApproveRequest struct {
	AppKey    string
	AppSecret string
	Ids       string
	PartParam request.Values
}

func (req FinApproveRequest) ToValues() request.Values {
	values := request.Values{
		"ids":       req.Ids,
		"partParam": req.PartParam,
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
	                "1575944756563279907": 1666678600000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "订单财审"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1575944756563279907": 1666678600000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "发送Wms发货单"
	        }
	    ]
	}
*/
type FinApproveResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		ExceptionMsg   string         `json:"exceptionMsg"`
		Code           string         `json:"code"`
		IsShowMsg      bool           `json:"isShowMsg"`
		FailCount      string         `json:"failCount"`
		SucIdAndPubts  request.Values `json:"sucIdAndPubts"`
		SuccessCount   string         `json:"successCount"`
		IsExcuteAction bool           `json:"isExcuteAction"`
		ActionName     string         `json:"actionName"`
	} `json:"data"`
}

func (resp FinApproveResponse) IsSuccessed(id string) (bool, string) {
	for _, action := range resp.Data {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "订单财审" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有订单财审的事件，把data的内容都返回
	b, _ := json.Marshal(resp.Data)
	return false, string(b)
}

func (resp FinApproveResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	for _, action := range resp.Data {
		if action.ActionName == "订单财审" {
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

func FinApprove(req FinApproveRequest) (FinApproveResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/tradeorder/finapprove", req.ToValues())
	if err != nil {
		return FinApproveResponse{}, err
	}

	res, err := vals.GetResult(FinApproveResponse{})
	if err != nil {
		return FinApproveResponse{}, err
	}

	resp, ok := res.(*FinApproveResponse)
	if !ok {
		return FinApproveResponse{}, fmt.Errorf("%w error: response is not type of FinApproveResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
