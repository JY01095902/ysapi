package snquerysnstate

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/jy01095902/ysapi/request"
	"github.com/stretchr/testify/assert"
)

func TestListRequestToValues(t *testing.T) {
	req := ListRequest{
		PageIndex: 1,
		PageSize:  10,
		Params:    request.Values{"sn": "12345"},
	}

	params := req.ToValues()
	assert.Equal(t, 1, params["pageIndex"])
	assert.Equal(t, 10, params["pageSize"])
	assert.Equal(t, "12345", params["sn"])
	assert.Nil(t, params["simpleOVs"])

	req.SimpleVOs = []SimpleVO{
		{
			Field:  "id",
			Op:     "eq",
			Value1: "1",
			Value2: "",
		},
		{
			Field:  "name",
			Op:     "between",
			Value1: "a",
			Value2: "z",
		},
	}
	params = req.ToValues()
	ovs := params["simpleVOs"].([]request.Values)
	assert.Equal(t, "id", ovs[0]["field"])
	assert.Equal(t, "eq", ovs[0]["op"])
	assert.Equal(t, "1", ovs[0]["value1"])
	assert.Equal(t, "", ovs[0]["value2"])
	assert.Equal(t, "name", ovs[1]["field"])
	assert.Equal(t, "between", ovs[1]["op"])
	assert.Equal(t, "a", ovs[1]["value1"])
	assert.Equal(t, "z", ovs[1]["value2"])
}

func TestList(t *testing.T) {
	sn := "000112-1"
	req := ListRequest{
		AppKey:    "",
		AppSecret: "",
		PageIndex: 1,
		PageSize:  10,
		Params:    request.Values{"sn": sn},
	}
	for i := 0; i < 20; i++ {
		resp, err := List(req)
		log.Printf("err: %v", err)
		b, _ := json.Marshal(resp)
		log.Printf("resp: %v", string(b))

		if err != nil {
			if resp.Code == "200" && len(resp.Data.RecordList) == 0 {
				assert.ErrorIs(t, err, request.ErrYonSuiteAPIBizError)
			} else {
				if strings.Contains(err.Error(), "限流") {
					assert.ErrorIs(t, err, request.ErrAPILimit)
				} else {
					assert.ErrorIs(t, err, request.ErrCallYonSuiteAPIFailed)
				}
			}
		} else {
			assert.Equal(t, sn, resp.Data.RecordList[0].SN)
		}
	}
}
