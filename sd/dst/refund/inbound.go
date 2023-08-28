package refund

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jy01095902/ysapi/request"
)

type InboundRequest struct {
	AppKey     string
	AppSecret  string
	Id         string
	Ts         string
	Stocks     []request.Values
	StocksInfo []request.Values
}

func (req InboundRequest) ToValues() request.Values {
	values := request.Values{
		"id":         req.Id,
		"ts":         req.Ts,
		"stocks":     req.Stocks,
		"stocksInfo": req.StocksInfo,
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
	                "1564232234717675521": 1665300958000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "确认入库"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1564232234717675521": 1665300958000
	            },
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "确认入库通知AG退款"
	        },
	        {
	            "code": "1",
	            "isShowMsg": false,
	            "externalMap": {},
	            "failCount": "0",
	            "sucIdAndPubts": {
	                "1564232234717675521": 1665300958000
	            },
	            "successCount": "1",
	            "isExcuteAction": true,
	            "actionName": "提交存量"
	        },
	        {
	            "exceptionMsg": "退货换单没有换货行，无需换货执行！",
	            "code": "1",
	            "isShowMsg": true,
	            "failCount": "1",
	            "successCount": "0",
	            "isExcuteAction": true,
	            "actionName": "换货执行"
	        }
	    ]
	}
*/
type InboundResponse struct {
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

func (resp InboundResponse) IsSuccessed(id string) (bool, string) {
	for _, action := range resp.Data {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "确认入库" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有确认入库的事件，把data的内容都返回
	return false, resp.Message
}

func (resp InboundResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	for _, action := range resp.Data {
		if action.ActionName == "确认入库" {
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

func Inbound(req InboundRequest) (InboundResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/refund/inbound", req.ToValues())
	if err != nil {
		return InboundResponse{}, err
	}

	res, err := vals.GetResult(InboundResponse{})
	if err != nil {
		return InboundResponse{}, err
	}

	resp, ok := res.(*InboundResponse)
	if !ok {
		return InboundResponse{}, fmt.Errorf("%w error: response is not type of InboundResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
