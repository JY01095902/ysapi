package tradeorder

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/jy01095902/ysapi/request"
	"github.com/stretchr/testify/assert"
)

func TestShipRequestToValues(t *testing.T) {
	orderId := "1481740807928021008"
	tid := "TEST0220620001"
	shipCode := "TESTFH20220620001"
	expressCode := "SF20220607001"
	req := ShipRequest{
		AppKey:    "1",
		AppSecret: "2",
		PartParam: map[string]interface{}{
			orderId: orderId,
		},
		Path: 2,
		Ids:  orderId,
	}

	req.ExternalData.ShipDetails = []request.Values{
		{
			"warehouseCode": "XNCX-BJ007",
			"iquantity":     1,
			"cshipcode":     shipCode,
			"skuCode":       "LG0020000CN3004",
			"parentid":      orderId,
		},
	}
	req.ExternalData.Option = request.Values{"pushFlag": true}
	req.ExternalData.Modify = []request.Values{
		{
			"cExpressCode": expressCode,
			"id":           orderId,
		},
	}
	req.ExternalData.ExpressLists = []request.Values{
		{
			"cshipcode":    shipCode,
			"tid":          tid,
			"parentid":     orderId,
			"cexpresscode": expressCode,
		},
	}

	params := req.ToValues()
	assert.Equal(t, orderId, params["ids"])
	assert.Equal(t, orderId, params["partParam"].(request.Values)[orderId])
	shipDetails := params["externalData"].(request.Values)["shipdetails"].([]request.Values)
	assert.Equal(t, "XNCX-BJ007", shipDetails[0]["warehouseCode"])
	option := params["externalData"].(request.Values)["option"].(request.Values)
	assert.Equal(t, true, option["pushFlag"])
	modify := params["externalData"].(request.Values)["Modify"].([]request.Values)
	assert.Equal(t, orderId, modify[0]["id"])
	expressLists := params["externalData"].(request.Values)["expresslists"].([]request.Values)
	assert.Equal(t, tid, expressLists[0]["tid"])
}

func TestShip(t *testing.T) {
	orderId := "1481740807928021008"
	tid := "TEST0220620001"
	shipCode := "TESTFH20220620001"
	expressCode := "SF20220607001"
	req := ShipRequest{
		AppKey:    "4097e5845ca34c138f70bc7d2132825a",
		AppSecret: "b0126be656234f8db2c3de6f9906304f",
		PartParam: map[string]interface{}{
			orderId: orderId,
		},
		Path: 2,
		Ids:  orderId,
	}

	req.ExternalData.ShipDetails = []request.Values{
		{
			"warehouseCode": "XNCX-BJ007",
			"iquantity":     1,
			"cshipcode":     shipCode,
			"skuCode":       "LG0020000CN3004",
			"parentid":      orderId,
		},
	}
	req.ExternalData.Option = request.Values{"pushFlag": true}
	req.ExternalData.Modify = []request.Values{
		{
			"cExpressCode": expressCode,
			"id":           orderId,
		},
	}
	req.ExternalData.ExpressLists = []request.Values{
		{
			"cshipcode":    shipCode,
			"tid":          tid,
			"parentid":     orderId,
			"cexpresscode": expressCode,
		},
	}

	resp, err := Ship(req)
	assert.Nil(t, err)
	assert.Equal(t, "200", resp.Code)
	b, _ := json.Marshal(resp)
	log.Printf("resp: %s", string(b))
}
