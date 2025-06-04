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

type CloseRequest struct {
	AppKey    string
	AppSecret string
	Ids       string
	PartParam request.Values
}

func (req CloseRequest) ToValues() request.Values {
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
	            "exceptionMsg": "没有需要操作的单据",
	            "code": "1",
	            "isShowMsg": true,
	            "failCount": "1",
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "订单关闭"
	        },
	        {
	            "exceptionMsg": "没有需要操作的单据",
	            "code": "1",
	            "isShowMsg": true,
	            "failCount": "1",
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "订单关闭"
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
	                "1648795880186183721": 1675321224000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "订单关闭"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1648795880186183721": 1675321224000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "取消顺丰WMS出库单"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1648795880186183721": 1675321224000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "提交存量"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1648795880186183721": 1675321224000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "取消奇门单据"
	        },
	        {
	            "code": "1",
	            "isShowMsg": true,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1648795880186183727": 1675321225000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "订单关闭"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1648795880186183727": 1675321225000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "取消顺丰WMS出库单"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1648795880186183727": 1675321225000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "提交存量"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1648795880186183727": 1675321225000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "取消奇门单据"
	        }
	    ]
	}
*/
type CloseResponse struct {
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

func (resp CloseResponse) IsSuccessed(id string) (bool, string) {
	for _, action := range resp.Data {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "订单关闭" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有订单关闭的事件，把data的内容都返回
	b, _ := json.Marshal(resp.Data)
	return false, string(b)
}

func (resp CloseResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	for _, action := range resp.Data {
		if action.ActionName == "订单关闭" {
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

func Close(req CloseRequest) (CloseResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/tradeorder/orderclose", req.ToValues())
	if err != nil {
		return CloseResponse{}, err
	}

	res, err := vals.GetResult(CloseResponse{})
	if err != nil {
		return CloseResponse{}, err
	}

	resp, ok := res.(*CloseResponse)
	if !ok {
		return CloseResponse{}, fmt.Errorf("%w error: response is not type of CloseResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
