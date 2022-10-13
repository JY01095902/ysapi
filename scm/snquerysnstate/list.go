package snquerysnstate

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type SimpleVO struct {
	Field  string
	Op     string
	Value1 string
	Value2 string
}

type ListRequest struct {
	AppKey    string
	AppSecret string
	PageIndex int
	PageSize  int
	Params    request.Values
	SimpleVOs []SimpleVO
}

func (req ListRequest) ToValues() request.Values {
	values := request.Values{
		"pageIndex": req.PageIndex,
		"pageSize":  req.PageSize,
	}

	for k, v := range req.Params {
		values.Set(k, v)
	}

	if len(req.SimpleVOs) > 0 {
		ovs := []request.Values{}
		for _, ov := range req.SimpleVOs {
			ovs = append(ovs, request.Values{
				"field":  ov.Field,
				"op":     ov.Op,
				"value1": ov.Value1,
				"value2": ov.Value2,
			})
		}

		values.Set("simpleVOs", ovs)
	}

	return values
}

type ListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PageIndex   int `json:"pageIndex"`
		PageSize    int `json:"pageSize"`
		RecordCount int `json:"recordCount"`
		RecordList  []struct {
			Product                int64  `json:"product"`
			BatchNo                string `json:"batchno"`
			Org                    string `json:"org"`
			ProductCode            string `json:"product_cCode"`
			StockStatusDoc         int64  `json:"stockStatusDoc"`
			StockStatusDocName     string `json:"stockStatusDoc_name"`
			Warehouse              int64  `json:"warehouse"`
			WarehouseName          string `json:"warehouse_name"`
			LocationCode           string `json:"location_code"`
			LocationName           string `json:"location_name"`
			ProductManageClassCode string `json:"product_ManageClass_code"`
			ProductSKUName         string `json:"productsku_cName"`
			ProductManageClassName string `json:"product_ManageClass_name"`
			UpdateCount            int    `json:"updatecount"`
			ProductManageClass     int64  `json:"product_ManageClass"`
			ProductSKU             int64  `json:"productsku"`
			ProductSKUCode         string `json:"productsku_cCode"`
			ProductName            string `json:"product_cName"`
			SNState                string `json:"snstate"`
			Id                     int64  `json:"id"`
			SN                     string `json:"sn"`
			PubTs                  string `json:"pubts"`
			OrgName                string `json:"org_name"`
		} `json:"recordList"`
		PageCount      int    `json:"pageCount"`
		BeginPageIndex int    `json:"beginPageIndex"`
		EndPageIndex   int    `json:"endPageIndex"`
		PubTs          string `json:"pubts"`
	} `json:"data"`
}

func (resp ListResponse) Total() int {
	return resp.Data.RecordCount
}

func (resp ListResponse) PageCount() int {
	return resp.Data.PageCount
}

func List(req ListRequest) (ListResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.SNQuerySNStateListURL, req.ToValues())
	if err != nil {
		return ListResponse{}, err
	}

	res, err := vals.GetResult(ListResponse{})
	if err != nil {
		return ListResponse{}, err
	}

	resp, ok := res.(*ListResponse)
	if !ok {
		return ListResponse{}, fmt.Errorf("%w error: response is not type of ListResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	if len(resp.Data.RecordList) == 0 {
		return *resp, fmt.Errorf("%w error: not found sn state", request.ErrYonSuiteAPIBizError)
	}

	return *resp, nil
}
