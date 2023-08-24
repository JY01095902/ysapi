package vouchersalereturn

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jy01095902/ysapi/request"
)

type SingleSaveRequest struct {
	AppKey    string
	AppSecret string
	Data      SaleReturnOrder
}

type SaleReturnOrder struct {
	ResubmitCheckKey     string  `json:"resubmitCheckKey"`
	Code                 string  `json:"code"`
	SalesOrgId           string  `json:"salesOrgId"`
	TransactionTypeId    string  `json:"transactionTypeId"`
	AgentId              string  `json:"agentId"`
	Vouchdate            string  `json:"vouchdate"`
	SettlementOrgId      string  `json:"settlementOrgId"`
	Currency             string  `json:"currency"`
	ExchangeRateType     string  `json:"exchangeRateType"`
	NatCurrency          string  `json:"natCurrency"`
	ExchRate             float64 `json:"exchRate"`
	TaxInclusiv          bool    `json:"taxInclusive"`
	SaleReturnSourceType string  `json:"saleReturnSourceType"`
	InvoiceAgentId       string  `json:"invoiceAgentId"`
	TotalMoney           float64 `json:"totalMoney"`
	PayMoney             float64 `json:"payMoney"`
	Status               string  `json:"_status"`
	BizFlow              string  `json:"bizFlow"`
	IsFlowCoreBill       bool    `json:"isFlowCoreBill"`
	SaleReturnMemo       struct {
		Remark string `json:"remark"`
	} `json:"saleReturnMemo"`
	Creator      string             `json:"creator"`
	HeadFreeItem map[string]string  `json:"headFreeItem"`
	Details      []SaleReturnDetail `json:"saleReturnDetails"`
}

func (order SaleReturnOrder) Check() error {
	if order.SaleReturnSourceType == "" {
		return errors.New("[SaleReturnOrder] SaleReturnSourceType is required")
	}

	return nil
}

type SaleReturnDetail struct {
	ProductId             string  `json:"productId"`
	SkuId                 string  `json:"skuId"`
	UnitExchangeType      int     `json:"unitExchangeType"`
	UnitExchangeTypePrice int     `json:"unitExchangeTypePrice"`
	TaxId                 string  `json:"taxId"`
	StockId               string  `json:"stockId"`
	StockOrgId            string  `json:"stockOrgId"`
	ProductAuxUnitId      string  `json:"iProductAuxUnitId"`
	ProductUnitId         string  `json:"iProductUnitId"`
	MasterUnitId          string  `json:"masterUnitId"`
	SalesOrgId            string  `json:"salesOrgId"`
	InvExchRate           float64 `json:"invExchRate"`
	InvPriceExchRate      float64 `json:"invPriceExchRate"`
	PriceQty              float64 `json:"priceQty"`
	SubQty                float64 `json:"subQty"`
	Qty                   float64 `json:"qty"`
	OriTaxUnitPrice       float64 `json:"oriTaxUnitPrice"`
	OriUnitPrice          float64 `json:"oriUnitPrice"`
	OriSum                float64 `json:"oriSum"`
	OriTax                float64 `json:"oriTax"`
	OriMoney              float64 `json:"oriMoney"`
	NatTaxUnitPrice       float64 `json:"natTaxUnitPrice"`
	NatSum                float64 `json:"natSum"`
	NatTax                float64 `json:"natTax"`
	NatUnitPrice          float64 `json:"natUnitPrice"`
	NatMoney              float64 `json:"natMoney"`
	Status                string  `json:"_status"`
}

// exchRate 汇率
func (detail *SaleReturnDetail) SetTaxUnitPrice(price, taxRate, exchRate float64) {
	fomatPrice := func(value float64) float64 {
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
		return value
	}

	rate := taxRate / 100

	detail.OriTaxUnitPrice = fomatPrice(price)
	detail.OriSum = fomatPrice(price * detail.Qty)

	money := detail.OriSum / (1 + rate) // 本币无税金额
	detail.OriMoney = fomatPrice(money)
	detail.OriUnitPrice = fomatPrice(detail.OriMoney / detail.Qty)
	detail.OriTax = fomatPrice(detail.OriSum - detail.OriMoney)

	natSum := detail.OriSum * exchRate
	detail.NatSum = fomatPrice(natSum)
	natMoney := detail.NatSum / (1 + rate)
	detail.NatMoney = fomatPrice(natMoney)
	detail.NatTaxUnitPrice = fomatPrice(detail.NatSum / detail.Qty)
	detail.NatTax = fomatPrice(detail.NatSum - detail.NatMoney)
	detail.NatUnitPrice = fomatPrice(detail.NatMoney / detail.Qty)
}

func (req SingleSaveRequest) ToValues() request.Values {
	values := request.Values{
		"data": req.Data,
	}

	return values
}

type SingleSaveResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func SingleSave(req SingleSaveRequest) (SingleSaveResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/sd/vouchersalereturn/singleSave", req.ToValues())
	if err != nil {
		return SingleSaveResponse{}, err
	}

	res, err := vals.GetResult(SingleSaveResponse{})
	if err != nil {
		return SingleSaveResponse{}, err
	}

	resp, ok := res.(*SingleSaveResponse)
	if !ok {
		return SingleSaveResponse{}, fmt.Errorf("%w error: response is not type of SingleSaveResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
