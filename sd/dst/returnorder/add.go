package returnorder

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type AddRequest struct {
	AppKey       string
	AppSecret    string
	PartParam    request.Values
	ExternalData struct {
		Head  []request.Values
		Added []request.Values
	}
}

func (req AddRequest) ToValues() request.Values {
	values := request.Values{
		"partParam": req.PartParam,
		"externalData": request.Values{
			"Head":  req.ExternalData.Head,
			"Added": req.ExternalData.Added,
		},
	}

	return values
}

/*
{
    "code": "200",
    "message": "[{\"actionName\":\"订单挂起\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1509293890716303371\":1658905111000},\"successCount\":\"1\"},{\"actionName\":\"取消顺丰WMS出库单\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1509293890716303371\":1658905111000},\"successCount\":\"0\"},{\"actionName\":\"取消奇门单据\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1509293890716303371\":1658905111000},\"successCount\":\"0\"},{\"actionName\":\"退换货单新增\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1509295411134726149\":1658905111000},\"successCount\":\"1\"},{\"actionName\":\"提交存量\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1509295411134726149\":1658905111000},\"successCount\":\"1\"}]",
    "data": null
}
*/
type AddResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func Add(req AddRequest) (AddResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/returnorder/add", req.ToValues())
	if err != nil {
		return AddResponse{}, err
	}

	res, err := vals.GetResult(AddResponse{})
	if err != nil {
		return AddResponse{}, err
	}

	resp, ok := res.(*AddResponse)
	if !ok {
		return AddResponse{}, fmt.Errorf("%w error: response is not type of AddResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
