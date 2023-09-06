package voucherorder

import (
	"fmt"
	"strconv"
)

type SaveOrderPrice struct {
	PayMoneyDomestic             int     `json:"payMoneyDomestic"`
	OrderPayMoneyDomestic        int     `json:"orderPayMoneyDomestic"`
	OrderPayMoneyOrigTaxFree     float64 `json:"orderPayMoneyOrigTaxfree"`
	PayMoneyOrigTaxFree          float64 `json:"payMoneyOrigTaxfree"` // 合计无税金额
	TotalMoneyOrigTaxFree        float64 `json:"totalMoneyOrigTaxfree"`
	Currency                     string  `json:"currency"`
	AgentTaxItem                 string  `json:"agentTaxItem"`
	ExchRate                     float64 `json:"exchRate"`
	ExchangeRateType             string  `json:"exchangeRateType"`
	NatCurrency                  string  `json:"natCurrency"`
	TaxInclusive                 string  `json:"taxInclusive"`
	TotalOriTax                  float64 `json:"totalOriTax"` // 税额
	TotalNatTax                  float64 `json:"totalNatTax"` // 本币总税额
	PayMoneyDomesticTaxfree      float64 `json:"payMoneyDomesticTaxfree"`
	OrderPayMoneyDomesticTaxfree float64 `json:"orderPayMoneyDomesticTaxfree"`
}

// 平铺的字段，保存的时候需要这种结构（singleSave）
type FlatSaveOrderPrice struct {
	PayMoneyDomestic             int     `json:"orderPrices!payMoneyDomestic"`
	OrderPayMoneyDomestic        int     `json:"orderPrices!orderPayMoneyDomestic"`
	OrderPayMoneyOrigTaxFree     float64 `json:"orderPrices!orderPayMoneyOrigTaxfree"`
	PayMoneyOrigTaxFree          float64 `json:"orderPrices!payMoneyOrigTaxfree"` // 合计无税金额
	TotalMoneyOrigTaxFree        float64 `json:"orderPrices!totalMoneyOrigTaxfree"`
	Currency                     string  `json:"orderPrices!currency"`
	AgentTaxItem                 string  `json:"orderPrices!agentTaxItem"`
	ExchRate                     float64 `json:"orderPrices!exchRate"`
	ExchangeRateType             string  `json:"orderPrices!exchangeRateType"`
	NatCurrency                  string  `json:"orderPrices!natCurrency"`
	TaxInclusive                 string  `json:"orderPrices!taxInclusive"`
	TotalOriTax                  float64 `json:"orderPrices!totalOriTax"` // 税额
	TotalNatTax                  float64 `json:"orderPrices!totalNatTax"` // 本币总税额
	PayMoneyDomesticTaxfree      float64 `json:"orderPrices!payMoneyDomesticTaxfree"`
	OrderPayMoneyDomesticTaxfree float64 `json:"orderPrices!orderPayMoneyDomesticTaxfree"`
}

type SaveOrder struct {
	Code                 string                 `json:"code"`
	SalesOrgId           string                 `json:"salesOrgId"`
	TransactionTypeId    string                 `json:"transactionTypeId"`
	Vouchdate            string                 `json:"vouchdate"`
	SettlementOrgId      string                 `json:"settlementOrgId"`
	SaleDepartmentId     string                 `json:"saleDepartmentId"`
	CorpContactId        string                 `json:"corpContact"`         // 销售业务员Id
	CorpContactName      string                 `json:"corpContactUserName"` // 销售业务员名
	AgentId              string                 `json:"agentId"`
	AgentRelationId      string                 `json:"agentRelationId"`
	InvoiceMoney         float64                `json:"invoiceMoney"`
	OrderPayType         string                 `json:"orderPayType"`
	Settlement           string                 `json:"settlement"`
	ShippingChoiceId     string                 `json:"shippingChoiceId"`
	InvoiceAgentId       string                 `json:"invoiceAgentId"`
	ModifyInvoiceType    bool                   `json:"modifyInvoiceType"`
	InvoiceUpcType       string                 `json:"invoiceUpcType"`
	InvoiceTitleType     int                    `json:"invoiceTitleType"`
	TotalMoney           float64                `json:"totalMoney"`     // 总金额
	PromotionMoney       float64                `json:"promotionMoney"` // 总优惠金额
	RebateMoney          float64                `json:"rebateMoney"`
	RebateCashMoney      float64                `json:"rebateCashMoney"`
	PointsMoney          float64                `json:"pointsMoney"`
	Reight               float64                `json:"reight"`
	PayMoney             float64                `json:"payMoney"` // 合计含税金额
	OrderPayMoney        float64                `json:"orderPayMoney"`
	RealMoney            float64                `json:"realMoney"`
	OrderRealMoney       float64                `json:"orderRealMoney"`
	ParticularlyMoney    float64                `json:"particularlyMoney"`
	Status               string                 `json:"_status"`
	BizFlow              string                 `json:"bizFlow"`
	Prices               SaveOrderPrice         `json:"orderPrices"`
	Details              []SaveOrderDetail      `json:"orderDetails"`
	ReceiverName         string                 `json:"receiver"`
	ReceiverMobileNo     string                 `json:"receiveMobile"`
	ReceiverAddress      string                 `json:"receiveAddress"`
	Creator              string                 `json:"creator"`
	Memo                 string                 `json:"memo"`
	RetailAgentName      string                 `json:"retailAgentName"`
	HeadFreeItem         map[string]string      `json:"headFreeItem"`
	OrderDefineCharacter map[string]interface{} `json:"orderDefineCharacter"` // 特征
	// Version           string            `json:"version"`
	// PubuTs            string            `json:"pubuts"`
	FlatSaveOrderPrice
}

func (order *SaveOrder) AddDetails(details ...SaveOrderDetail) {
	order.Details = append(order.Details, details...)
}

func (order *SaveOrder) CalculateAmount() {
	fomatPrice := func(value float64) float64 {
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
		return value
	}

	totalAmount := 0
	totalFreeTaxAmount := 0
	toalTax := 0
	for _, item := range order.Details {
		totalAmount += int(item.OriSum * 100)
		totalFreeTaxAmount += int(item.OrderDetailPrices.OriMoney * 100)
		toalTax += int(item.OrderDetailPrices.OriTax * 100)
	}

	order.TotalMoney = fomatPrice(float64(totalAmount) / 100)
	order.PayMoney = fomatPrice(float64(totalAmount) / 100)
	order.Prices.PayMoneyOrigTaxFree = fomatPrice(float64(totalFreeTaxAmount) / 100)
	order.Prices.TotalOriTax = fomatPrice(float64(toalTax) / 100)

	order.FlatSaveOrderPrice.PayMoneyDomestic = order.Prices.PayMoneyDomestic
	order.FlatSaveOrderPrice.OrderPayMoneyDomestic = order.Prices.OrderPayMoneyDomestic
	order.FlatSaveOrderPrice.OrderPayMoneyOrigTaxFree = order.Prices.OrderPayMoneyOrigTaxFree
	order.FlatSaveOrderPrice.PayMoneyOrigTaxFree = order.Prices.PayMoneyOrigTaxFree
	order.FlatSaveOrderPrice.TotalMoneyOrigTaxFree = order.Prices.TotalMoneyOrigTaxFree
	order.FlatSaveOrderPrice.Currency = order.Prices.Currency
	order.FlatSaveOrderPrice.AgentTaxItem = order.Prices.AgentTaxItem
	order.FlatSaveOrderPrice.ExchRate = order.Prices.ExchRate
	order.FlatSaveOrderPrice.ExchangeRateType = order.Prices.ExchangeRateType
	order.FlatSaveOrderPrice.NatCurrency = order.Prices.NatCurrency
	order.FlatSaveOrderPrice.TaxInclusive = order.Prices.TaxInclusive
	order.FlatSaveOrderPrice.TotalOriTax = order.Prices.TotalOriTax
	order.FlatSaveOrderPrice.TotalNatTax = order.Prices.TotalNatTax
	order.FlatSaveOrderPrice.PayMoneyDomesticTaxfree = order.Prices.PayMoneyDomesticTaxfree
	order.FlatSaveOrderPrice.OrderPayMoneyDomesticTaxfree = order.Prices.OrderPayMoneyDomesticTaxfree
}
