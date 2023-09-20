package voucherorder

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type SaveOrderDetailPrice struct {
	OriTaxUnitPrice float64 `json:"oriTaxUnitPrice"` // 原币含税单价
	OriSum          float64 `json:"oriSum"`          // 原币含税金额
	OriTax          float64 `json:"oriTax"`          // 原币税额
	OriUnitPrice    float64 `json:"oriUnitPrice"`    // 原币无税单价
	OriMoney        float64 `json:"oriMoney"`        // 原币无税金额

	NatTaxUnitPrice float64 `json:"natTaxUnitPrice"` // 本币含税单价
	NatSum          float64 `json:"natSum"`          // 本币含税金额
	NatTax          float64 `json:"natTax"`          // 本币税额
	NatUnitPrice    float64 `json:"natUnitPrice"`    // 本币无税单价
	NatMoney        float64 `json:"natMoney"`        //本币无税金额
}

// 平铺的字段，保存的时候需要这种结构（singleSave）
type FlatSaveOrderDetailPrice struct {
	OriTaxUnitPrice float64 `json:"orderDetailPrices!oriTaxUnitPrice"` // 原币含税单价
	OriSum          float64 `json:"orderDetailPrices!oriSum"`          // 原币含税金额
	OriTax          float64 `json:"orderDetailPrices!oriTax"`          // 原币税额
	OriUnitPrice    float64 `json:"orderDetailPrices!oriUnitPrice"`    // 原币无税单价
	OriMoney        float64 `json:"orderDetailPrices!oriMoney"`        // 原币无税金额

	NatTaxUnitPrice float64 `json:"orderDetailPrices!natTaxUnitPrice"` // 本币含税单价
	NatSum          float64 `json:"orderDetailPrices!natSum"`          // 本币含税金额
	NatTax          float64 `json:"orderDetailPrices!natTax"`          // 本币税额
	NatUnitPrice    float64 `json:"orderDetailPrices!natUnitPrice"`    // 本币无税单价
	NatMoney        float64 `json:"orderDetailPrices!natMoney"`        //本币无税金额
}

type SaveOrderDetail struct {
	ProductUnitName    string `json:"productUnitName"`
	IsExpiryDateManage bool   `json:"isExpiryDateManage"`
	TaxItems           string `json:"taxItems"`
	// OrdRealMoney            float64              `json:"ordRealMoney"`
	IsBatchManage         bool    `json:"isBatchManage"`
	ProductId             string  `json:"productId"`
	SKUId                 string  `json:"skuId"`
	OrderProductType      string  `json:"orderProductType"`
	UnitExchangeType      int     `json:"unitExchangeType"`
	UnitExchangeTypePrice float64 `json:"unitExchangeTypePrice"` // 浮动（销售）
	ProductAuxUnitName    string  `json:"productAuxUnitName"`
	InvExchRate           float64 `json:"invExchRate"`      // 销售换算率
	InvPriceExchRate      float64 `json:"invPriceExchRate"` // 计价换算率
	PriceQty              float64 `json:"priceQty"`
	SubQty                float64 `json:"subQty"`
	Qty                   float64 `json:"qty"`
	QtyName               string  `json:"qtyName"`
	StockId               string  `json:"stockId"`
	ConsignTime           string  `json:"consignTime"`
	StockOrgId            string  `json:"stockOrgId"`
	SettlementOrgId       string  `json:"settlementOrgId"`
	OriTaxUnitPrice       float64 `json:"oriTaxUnitPrice"` // 原币含税成交价
	OriSum                float64 `json:"oriSum"`          // 原币含税金额
	TaxRate               float64 `json:"taxRate"`         // 税率
	OriTax                float64 `json:"oriTax"`          // 税额
	OriUnitPrice          float64 `json:"oriUnitPrice"`    // 原币无税单价
	OriMoney              float64 `json:"oriMoney"`        // 原币无税金额

	TaxId string `json:"taxId"`
	// TaxCode                 string               `json:"taxCode"`                 // 税目税率编码
	TransactionTypeId string `json:"transactionTypeId"`
	SalesOrgId        string `json:"salesOrgId"`
	// OrderDetailPrice        float64              `json:"orderDetailPrice"` // 订单金额
	IProductAuxUnitId string               `json:"iProductAuxUnitId"`
	IProductUnitId    string               `json:"iProductUnitId"`
	MasterUnitId      string               `json:"masterUnitId"`
	BodyItem          struct{}             `json:"bodyItem"`
	Status            string               `json:"_status"`
	NatTaxUnitPrice   float64              `json:"natTaxUnitPrice"` // 本币含税单价
	NatSum            float64              `json:"natSum"`          // 本币含税金额
	NatTax            float64              `json:"natTax"`          // 本币税额
	NatUnitPrice      float64              `json:"natUnitPrice"`    // 本币无税单价
	NatMoney          float64              `json:"natMoney"`        // 本币无税金额
	OrderDetailPrices SaveOrderDetailPrice `json:"orderDetailPrices"`
	// SalePrice               float64              `json:"salePrice"`
	// SaleCost                float64              `json:"saleCost"`
	// PayMoneyDomesticTaxfree float64              `json:"payMoneyDomesticTaxfree"` // 本币无税金额
	FlatSaveOrderDetailPrice
}

/*
本币含税金额=（原币）含税金额*汇率； orderDetailPrices!natSum = oriSum * orderPrices!exchRate；
（原币）含税金额=无税金额+税额； oriSum = orderDetailPrices!oriMoney + orderDetailPrices!oriTax；
本币含税金额=本币无税金额+税额；orderDetailPrices!natSum = orderDetailPrices!natMoney + orderDetailPrices!natTax
*/
func (detail SaveOrderDetail) Check(exchRate float64) error {
	// 因为精度问题 金额都要保留2位小数后再比较验证
	fixed2 := func(value float64) float64 {
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
		return value
	}

	if fixed2(detail.FlatSaveOrderDetailPrice.NatSum) != fixed2(detail.OriSum*exchRate) {
		return fmt.Errorf("本币含税金额[%v] 应该等于（原币）含税金额[%v] * 汇率[%v]", detail.FlatSaveOrderDetailPrice.NatSum, detail.OriSum, exchRate)
	}

	if fixed2(detail.OriSum) != fixed2(detail.FlatSaveOrderDetailPrice.OriMoney+detail.FlatSaveOrderDetailPrice.OriTax) {
		return fmt.Errorf("（原币）含税金额[%v] 应该等于 无税金额[%v] + 税额[%v]", detail.OriSum, detail.FlatSaveOrderDetailPrice.OriMoney, detail.FlatSaveOrderDetailPrice.OriTax)

	}

	if fixed2(detail.FlatSaveOrderDetailPrice.NatSum) != fixed2(detail.FlatSaveOrderDetailPrice.NatMoney+detail.FlatSaveOrderDetailPrice.NatTax) {
		return fmt.Errorf("本币含税金额[%v] 应该等于 本币无税金额[%v] + 税额[%v]", detail.FlatSaveOrderDetailPrice.NatSum, detail.FlatSaveOrderDetailPrice.NatMoney, detail.FlatSaveOrderDetailPrice.NatTax)
	}

	return nil
}

// exchRate 汇率
func (detail *SaveOrderDetail) SetTaxUnitPrice(price, exchRate float64) {
	fomatPrice := func(value float64) float64 {
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
		return value
	}

	rate := detail.TaxRate / 100

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

	detail.OrderDetailPrices = SaveOrderDetailPrice{
		OriTaxUnitPrice: detail.OriTaxUnitPrice,
		OriSum:          detail.OriSum,
		OriTax:          detail.OriTax,
		OriUnitPrice:    detail.OriUnitPrice,
		OriMoney:        detail.OriMoney,
		NatTaxUnitPrice: detail.NatTaxUnitPrice,
		NatSum:          detail.NatSum,
		NatTax:          detail.NatTax,
		NatUnitPrice:    detail.NatUnitPrice,
		NatMoney:        detail.NatMoney,
	}

	detail.FlatSaveOrderDetailPrice.OriTaxUnitPrice = detail.OriTaxUnitPrice
	detail.FlatSaveOrderDetailPrice.OriSum = detail.OriSum
	detail.FlatSaveOrderDetailPrice.OriTax = detail.OriTax
	detail.FlatSaveOrderDetailPrice.OriUnitPrice = detail.OriUnitPrice
	detail.FlatSaveOrderDetailPrice.OriMoney = detail.OriMoney
	detail.FlatSaveOrderDetailPrice.NatTaxUnitPrice = detail.NatTaxUnitPrice
	detail.FlatSaveOrderDetailPrice.NatSum = detail.NatSum
	detail.FlatSaveOrderDetailPrice.NatTax = detail.NatTax
	detail.FlatSaveOrderDetailPrice.NatUnitPrice = detail.NatUnitPrice
	detail.FlatSaveOrderDetailPrice.NatMoney = detail.NatMoney

	data, _ := json.Marshal(detail)
	log.Printf("detail: %s", string(data))
}
