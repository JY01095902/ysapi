package refund

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type InboundRequest struct {
	AppKey     string
	AppSecret  string
	Id         string
	Ts         string
	StocksInfo []request.Values
}

func (req InboundRequest) ToValues() request.Values {
	values := request.Values{
		"id":         req.Id,
		"ts":         req.Ts,
		"stocks":     []request.Values{},
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

func Inbound(req InboundRequest) (InboundResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/refund/inbound", req.ToValues())
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
