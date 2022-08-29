package report

import (
	"log"
	"strings"
	"testing"

	"github.com/jy01095902/ysapi/request"
	"github.com/stretchr/testify/assert"
)

func TestListRequestToValues(t *testing.T) {
	req := BalanceRequest{
		AppKey:     "",
		AppSecret:  "",
		Groups:     []string{"accentity", "customer"},
		Conditions: []Condition{},
	}
	req.Params.Bz = request.Values{
		"periodStart": "2022-01",
		"querytype":   "1",
		"periodEnd":   "2022-03",
	}

	params := req.ToValues()
	assert.Equal(t, "accentity", params["groups"].([]request.Values)[0]["name"])
	assert.Equal(t, "customer", params["groups"].([]request.Values)[1]["name"])
	assert.Equal(t, "2022-01", params["params"].(request.Values)["bz"].(request.Values)["periodStart"])
	assert.Equal(t, "1", params["params"].(request.Values)["bz"].(request.Values)["querytype"])
	assert.Equal(t, "2022-03", params["params"].(request.Values)["bz"].(request.Values)["periodEnd"])
	assert.Nil(t, params["conditions"])

	req.Conditions = []Condition{
		{
			Op: "and",
			Items: []ConditionItem{
				{
					Op:   "in",
					Name: "accentity.code",
					V1:   []string{"001"},
				},
				{
					Op:   "eq",
					Name: "accentity",
					V1:   []string{"1"},
				},
			},
		},
		{
			Op: "or",
			Items: []ConditionItem{
				{
					Op:   "in",
					Name: "customer.code",
					V1:   []string{"C001"},
				},
				{
					Op:   "eq",
					Name: "customer",
					V1:   []string{"C1"},
				},
			},
		},
	}
	params = req.ToValues()
	conds := params["conditions"].([]request.Values)
	assert.Equal(t, "and", conds[0]["op"])
	assert.Equal(t, "in", conds[0]["items"].([]request.Values)[0]["op"])
	assert.Equal(t, "accentity.code", conds[0]["items"].([]request.Values)[0]["name"])
	assert.Contains(t, conds[0]["items"].([]request.Values)[0]["v1"], "001")
	assert.Equal(t, "eq", conds[0]["items"].([]request.Values)[1]["op"])
	assert.Equal(t, "accentity", conds[0]["items"].([]request.Values)[1]["name"])
	assert.Contains(t, conds[0]["items"].([]request.Values)[1]["v1"], "1")
	assert.Equal(t, "or", conds[1]["op"])
	assert.Equal(t, "in", conds[1]["items"].([]request.Values)[0]["op"])
	assert.Equal(t, "customer.code", conds[1]["items"].([]request.Values)[0]["name"])
	assert.Contains(t, conds[1]["items"].([]request.Values)[0]["v1"], "C001")
	assert.Equal(t, "eq", conds[1]["items"].([]request.Values)[1]["op"])
	assert.Equal(t, "customer", conds[1]["items"].([]request.Values)[1]["name"])
	assert.Contains(t, conds[1]["items"].([]request.Values)[1]["v1"], "C1")
}

func TestBalance(t *testing.T) {
	req := BalanceRequest{
		AppKey:     "",
		AppSecret:  "",
		Groups:     []string{"accentity", "customer"},
		Conditions: []Condition{},
	}
	req.Params.Bz = request.Values{
		"periodStart": "2021-01",
		"querytype":   "1",
		"periodEnd":   "2022-03",
	}
	req.Conditions = []Condition{
		{
			Op: "and",
			Items: []ConditionItem{
				{
					Op:   "in",
					Name: "accentity.code",
					V1:   []string{"001"},
				},
			},
		},
	}
	for i := 0; i < 20; i++ {
		resp, err := Balance(req)
		log.Printf("err: %v", err)

		if err != nil {
			if resp.Code == "200" && len(resp.Data.Result) == 0 {
				assert.ErrorIs(t, err, request.ErrYonSuiteAPIBizError)
			} else {
				if strings.Contains(err.Error(), "310046") {
					assert.ErrorIs(t, err, request.ErrAPILimit)
				}

				assert.ErrorIs(t, err, request.ErrCallYonSuiteAPIFailed)
			}
		} else {
			val := resp.Data.Result[0].Get("end_local_amount_total")
			assert.GreaterOrEqual(t, val.(float64), float64(0))
		}
	}
}
