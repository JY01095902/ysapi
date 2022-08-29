package report

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type Condition struct {
	Op    string
	Items []ConditionItem
}

type ConditionItem struct {
	Op   string
	Name string
	V1   []string
}

type BalanceRequest struct {
	AppKey    string
	AppSecret string
	Groups    []string
	Params    struct {
		Bz request.Values
	}
	Conditions []Condition
}

func (req BalanceRequest) ToValues() request.Values {
	groups := []request.Values{}
	for _, g := range req.Groups {
		groups = append(groups, request.Values{"name": g})
	}

	fields := []request.Values{}
	for _, f := range []string{"accentity", "customer", "currency", "begin_amount_total", "begin_local_amount_total", "debit_amount_total", "debit_local_amount_total", "credit_amount_total", "credit_local_amount_total", "end_amount_total", "end_local_amount_total"} {
		fields = append(fields, request.Values{"name": f})
	}

	values := request.Values{
		"groups": groups,
		"params": request.Values{
			"bz": req.Params.Bz,
		},
		"fields": fields,
	}

	if len(req.Conditions) > 0 {
		conditions := []request.Values{}
		for _, cond := range req.Conditions {
			items := []request.Values{}
			for _, item := range cond.Items {
				items = append(items, request.Values{
					"op":   item.Op,
					"name": item.Name,
					"v1":   item.V1,
				})
			}
			conditions = append(conditions, request.Values{
				"op":    cond.Op,
				"items": items,
			})
		}

		values.Set("conditions", conditions)
	}

	return values
}

type BalanceResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Result           []request.Values `json:"result"`
		PageIndex        int              `json:"pageIndex"`
		PageSize         int              `json:"pageSize"`
		RecordCount      int              `json:"recordCount"`
		PageCount        int              `json:"pageCount"`
		NeedConvert      bool             `json:"needconvert"`
		DynamicCondition struct {
			Op    string `json:"op"`
			Items []struct {
				Name string   `json:"name"`
				Op   string   `json:"op"`
				V1   []string `json:"v1"`
			} `json:"items"`
		}
	} `json:"data"`
}

func Balance(req BalanceRequest) (BalanceResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/fi/ar/report/balance", req.ToValues())
	if err != nil {
		return BalanceResponse{}, err
	}

	res, err := vals.GetResult(BalanceResponse{})
	if err != nil {
		return BalanceResponse{}, err
	}

	resp, ok := res.(*BalanceResponse)
	if !ok {
		return BalanceResponse{}, fmt.Errorf("%w error: response is not type of BalanceResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	if len(resp.Data.Result) == 0 {
		return *resp, fmt.Errorf("%w error: not found balance", request.ErrYonSuiteAPIBizError)
	}

	return *resp, nil
}
