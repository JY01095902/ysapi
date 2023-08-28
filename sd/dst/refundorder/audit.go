package refundorder

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jy01095902/ysapi/request"
)

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
	{
	    "code": "200",
	    "message": "[{\"code\":\"1\",\"isShowMsg\":true,\"externalMap\":{},\"failCount\":\"0\",\"sucIdAndPubts\":{\"1587276323400712196\":1668047093000},\"successCount\":\"1\",\"isExcuteAction\":true,\"actionName\":\"退换货单审核\"},{\"code\":\"1\",\"isShowMsg\":false,\"externalMap\":{},\"failCount\":\"0\",\"sucIdAndPubts\":{\"1587276323400712196\":1668047093000},\"successCount\":\"0\",\"isExcuteAction\":true,\"actionName\":\"创建顺丰WMS入库单\"},{\"code\":\"1\",\"isShowMsg\":false,\"externalMap\":{},\"failCount\":\"0\",\"sucIdAndPubts\":{\"1587276323400712196\":1668047093000},\"successCount\":\"0\",\"isExcuteAction\":true,\"actionName\":\"提交存量\"}]",
	    "data": null
	}
*/
type AuditResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    []request.Values `json:"data"`
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

func (resp AuditResponse) IsSuccessed(id string) (bool, string) {
	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return false, err.Error()
	}
	for _, action := range actions {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "退换货单审核" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有订单审核的事件，把data的内容都返回
	return false, resp.Message
}

func (resp AuditResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return ""
	}
	for _, action := range actions {
		if action.ActionName == "退换货单审核" {
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

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/refundorder/audit", req.ToValues())
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
