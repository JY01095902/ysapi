package tradeorder

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

/*
{
	"ids": "1489313582600421394",
	"partParam": {
		"1489313582600421394": "2022-10-25 13:48:02"
	}
}
*/

type AuditRequest struct {
	AppKey    string
	AppSecret string
	Ids       string
	PartParam request.Values
}

func (req AuditRequest) ToValues() request.Values {
	values := request.Values{
		"ids":       req.Ids,
		"partParam": req.PartParam,
	}

	return values
}

/*
ERROR

	{
	    "code": "200",
	    "message": "操作成功",
	    "data": [
	        {
	            "exceptionMsg": "TEST20221025001企业发货且非虚拟商品的订单快递公司不能为空,请自动匹配!",
	            "code": "1",
	            "isShowMsg": true,
	            "failCount": "1",
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "订单客审"
	        }
	    ]
	}

SUCCESS

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
	                "1575944756563279907": 1666677696000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "订单客审"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1575944756563279907": 1666677696000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "获取菜鸟序列号"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1575944756563279907": 1666677696000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "提交存量"
	        },
	        {
	            "exceptionMsg": "该店铺未设置免审策略！",
	            "code": "1",
	            "isShowMsg": false,
	            "failCount": "1",
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "财审免审检查"
	        }
	    ]
	}
*/
type AuditResponse struct {
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

func (resp AuditResponse) IsSuccessed(id string) (bool, string) {
	for _, action := range resp.Data {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "订单客审" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有订单客审的事件，把data的内容都返回
	b, _ := json.Marshal(resp.Data)
	return false, string(b)
}

func (resp AuditResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	for _, action := range resp.Data {
		if action.ActionName == "订单客审" {
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

func Audit(req AuditRequest) (AuditResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/tradeorder/audit", req.ToValues())
	if err != nil {
		return AuditResponse{}, err
	}

	res, err := vals.GetResult(AuditResponse{})
	if err != nil {
		return AuditResponse{}, err
	}

	resp, ok := res.(*AuditResponse)
	if !ok {
		return AuditResponse{}, fmt.Errorf("%w error: response is not type of AuditResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
