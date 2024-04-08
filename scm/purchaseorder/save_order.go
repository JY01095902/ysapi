package purchaseorder

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type SaveOrderDetail struct {
	BusType                       string                 `json:"bustype"`
	DemandOrgCode                 string                 `json:"demandOrg_code"`
	InInvoiceOrgCode              string                 `json:"inInvoiceOrg_code"`
	Id                            int64                  `json:"id,omitempty"`
	InOrgCode                     string                 `json:"inOrg_code"`
	InvExchRate                   float64                `json:"invExchRate"`
	InvoiceVendor                 string                 `json:"invoiceVendor"`
	Source                        string                 `json:"source"`
	MakeRuleCode                  string                 `json:"makeRuleCode"`
	Upcode                        string                 `json:"upcode"`
	SourceAutoId                  int64                  `json:"sourceautoid,omitempty"`
	SourceId                      int64                  `json:"sourceid,omitempty"`
	FirstSourceId                 int64                  `json:"firstsourceid,omitempty"`
	FirstSourceAutoId             int64                  `json:"firstsourceautoid,omitempty"`
	FirstSource                   string                 `json:"firstsource"`
	FirstUpCode                   string                 `json:"firstupcode"`
	BatchNumber                   int64                  `json:"batchno"`
	Define1                       string                 `json:"define1"`
	Define10                      string                 `json:"define10"`
	InvalidDate                   string                 `json:"invaliddate"`
	ProduceDate                   string                 `json:"producedate"`
	MainId                        int64                  `json:"mainid,omitempty"`
	MoneySum                      float64                `json:"moneysum"`
	NatMoney                      float64                `json:"natMoney"`
	NatSum                        float64                `json:"natSum"`
	NatTax                        float64                `json:"natTax"`
	NatTaxUnitPrice               float64                `json:"natTaxUnitPrice"`
	NatUnitPrice                  float64                `json:"natUnitPrice"`
	OriMoney                      float64                `json:"oriMoney"`
	OriSum                        float64                `json:"oriSum"`
	OriTax                        float64                `json:"oriTax"`
	DiscountTaxType               string                 `json:"discountTaxType"`
	OriTaxUnitPrice               float64                `json:"oriTaxUnitPrice"`
	OriUnitPrice                  float64                `json:"oriUnitPrice"`
	TaxitemsCode                  string                 `json:"taxitems_code"`
	ExpenseItemCode               string                 `json:"expenseItemId_code"`
	PriceQty                      float64                `json:"priceQty"`
	ProductSKU                    string                 `json:"productsku"`
	ProductCode                   string                 `json:"product_cCode"`
	PriceUOMCode                  string                 `json:"priceUOM_Code"`
	PurUOMCode                    string                 `json:"purUOM_Code"`
	ProjectCode                   string                 `json:"project_code"`
	Qty                           float64                `json:"qty"`
	RowNumber                     string                 `json:"rowno"`
	SubQty                        float64                `json:"subQty"`
	UnitExchangeTypePrice         float64                `json:"unitExchangeTypePrice"`
	UnitExchangeType              float64                `json:"unitExchangeType"`
	InvPriceExchRate              float64                `json:"invPriceExchRate"`
	WarehouseCode                 string                 `json:"warehouse_code"`
	UnitCode                      string                 `json:"unit_code"`
	IsGiftProduct                 bool                   `json:"isGiftProduct"`
	PurchaseOrdersCharacteristics map[string]interface{} `json:"purchaseOrdersCharacteristics"`
	PurchaseOrdersDefineCharacter map[string]interface{} `json:"purchaseOrdersDefineCharacter"`
	Status                        string                 `json:"_status"`
	Signatory                     string                 `json:"signatory"`
	OutForeignBusinessmenCode     string                 `json:"outForeignBusinessmen_code"`
	HasForeignInvestors           bool                   `json:"hasForeignInvestors"`
	BodyParallel                  map[string]interface{} `json:"bodyParallel"`
}

/*
	计算下列字段的值
	NatMoney:              0,
	NatSum:                0,
	NatTax:                0,
	NatTaxUnitPrice:       0,
	NatUnitPrice:          0,
	OriMoney:              0,
	OriSum:                0,
	OriTax:                0,
	OriTaxUnitPrice:       0,
	OriUnitPrice:          0,
*/
// taxRate 税率，exchRate 汇率
func (detail *SaveOrderDetail) SetTaxUnitPrice(price, taxRate, exchRate float64) {
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

	data, _ := json.Marshal(detail)
	log.Printf("detail: %s", string(data))
}

// 因为精度问题 金额都要保留2位小数后再比较验证
func fixed2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

/*
本币含税金额=（原币）含税金额*汇率； natSum = oriSum * exchRate；
（原币）含税金额=无税金额+税额； oriSum = oriMoney + oriTax；
本币含税金额=本币无税金额+税额； natSum = natMoney + natTax
*/
func (detail SaveOrderDetail) Check(exchRate float64) error {
	if fixed2(detail.NatSum) != fixed2(detail.OriSum*exchRate) {
		return fmt.Errorf("本币含税金额[%v] 应该等于（原币）含税金额[%v] * 汇率[%v]", detail.NatSum, detail.OriSum, exchRate)
	}

	if fixed2(detail.OriSum) != fixed2(detail.OriMoney+detail.OriTax) {
		return fmt.Errorf("（原币）含税金额[%v] 应该等于 无税金额[%v] + 税额[%v]", detail.OriSum, detail.OriMoney, detail.OriTax)
	}

	if fixed2(detail.NatSum) != fixed2(detail.NatMoney+detail.NatTax) {
		return fmt.Errorf("本币含税金额[%v] 应该等于 本币无税金额[%v] + 税额[%v]", detail.NatSum, detail.NatMoney, detail.NatTax)
	}

	return nil
}

type SaveOrder struct {
	ResubmitCheckKey             string                 `json:"resubmitCheckKey"`
	BustypeCode                  string                 `json:"bustype_code"`
	Code                         string                 `json:"code"`
	ExchRate                     float64                `json:"exchRate"`
	ExchRateType                 string                 `json:"exchRateType"`
	InvoiceVendorCode            string                 `json:"invoiceVendor_code"`
	NatCurrencyCode              string                 `json:"natCurrency_code"` // 文档里是这个字段，但实际生效的是currency_code
	CurrencyCode                 string                 `json:"currency_code"`
	NatMoney                     float64                `json:"natMoney"`
	NatSum                       float64                `json:"natSum"`
	OrgCode                      string                 `json:"org_code"`
	OriMoney                     float64                `json:"oriMoney"`
	OriSum                       float64                `json:"oriSum"`
	SrcBillId                    int64                  `json:"srcBill"`
	SrcBillNumber                string                 `json:"srcBillNO"`
	Source                       string                 `json:"source"`
	Operator                     string                 `json:"operator"`
	Department                   string                 `json:"department"`
	VendorContact                string                 `json:"vendorcontact"`
	Contact                      string                 `json:"contact"`
	PurchaseOrderDefineCharacter map[string]interface{} `json:"purchaseOrderDefineCharacter"`
	Details                      []SaveOrderDetail      `json:"purchaseOrders"`
	Status                       string                 `json:"_status"`
	HeadParallel                 map[string]interface{} `json:"headParallel"`
	VendorCode                   string                 `json:"vendor_code"`
	VouchDate                    string                 `json:"vouchdate"`
	Creator                      string                 `json:"creator"`

	NatTax float64 `json:"natTax"`
}

func (order SaveOrder) Check() error {
	for _, dtl := range order.Details {
		if err := dtl.Check(order.ExchRate); err != nil {
			return err
		}
	}

	return nil
}

func (order *SaveOrder) CalculateAmount() {
	var totalNatMoney,
		totalNatSum,
		totalOriMoney,
		totalOriSum float64

	for _, item := range order.Details {
		totalNatMoney += item.NatMoney
		totalNatSum += item.NatSum
		totalOriMoney += item.OriMoney
		totalOriSum += item.OriSum
	}

	order.NatMoney = fixed2(totalNatMoney)
	order.NatSum = fixed2(totalNatSum)
	order.NatTax = fixed2(totalNatSum - totalNatMoney)
	order.OriMoney = fixed2(totalOriMoney)
	order.OriSum = fixed2(totalOriSum)
}
