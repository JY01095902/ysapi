package salesout

type DetailDto struct {
	ReceiverDetail                   string                 `json:"cReceiveAddress"`
	NativeCurrencyCode               string                 `json:"natCurrencyCode"`
	CurrencyPriceDigit               int                    `json:"currency_priceDigit"`
	InvoiceOrgId                     string                 `json:"invoiceOrg"`
	MasterOrgKeyField                string                 `json:"masterOrgKeyField"`
	BusinessTypeName                 string                 `json:"bustype_name"`
	SourceOrderId                    string                 `json:"srcBill"`
	NativeCurrencyId                 string                 `json:"natCurrency"`
	SourceSystem                     string                 `json:"sourcesys"`
	ReceiveContactPhoneNumber        string                 `json:"receiveContacterPhone"`
	Id                               int64                  `json:"id"`
	BusinessTypeExtendAttributesJson string                 `json:"bustype_extend_attrs_json"`
	NativeCurrencyName               string                 `json:"natCurrencyName"`
	LogisticName                     string                 `json:"iLogisticId_name"`
	SalesOrgName                     string                 `json:"salesOrg_name"`
	ContactName                      string                 `json:"contactName"`
	SourceOrderNumber                string                 `json:"srcBillNO"`
	WarehouseId                      int64                  `json:"warehouse"`
	SourceOrderType                  string                 `json:"srcBillType"`
	NativeCurrencyMoneyDigit         int                    `json:"natCurrency_moneyDigit"`
	NativeCurrencyPriceDigit         int                    `json:"natCurrency_priceDigit"`
	AccountOrgName                   string                 `json:"accountOrg_name"`
	ExchangeRateType                 string                 `json:"exchRateType"`
	InvoiceCustomerId                int64                  `json:"invoiceCust"`
	Status                           int                    `json:"status"`
	IsUpdateCost                     bool                   `json:"isUpdateCost"`
	CurrencyMoneyDigit               int                    `json:"currency_moneyDigit"`
	OrderNumber                      string                 `json:"code"`
	LogisticsOrderNumber             string                 `json:"cLogisticsBillNo"`
	DefWMSLogisticsOrderNumber       string                 `json:"headItem!define1"` // 旧的自定义项
	ReceiveAccountingBasis           string                 `json:"receiveAccountingBasis"`
	InvoiceCustomerName              string                 `json:"invoiceCust_name"`
	ExchangeRate                     float64                `json:"exchRate"`
	SalesOrgId                       string                 `json:"salesOrg"`
	InvoiceOrgName                   string                 `json:"invoiceOrg_name"`
	StockAccount                     string                 `json:"stockAccount"`
	Date                             string                 `json:"vouchdate"`
	CurrencyName                     string                 `json:"currencyName"`
	CustomerName                     string                 `json:"cust_name"`
	CurrencyId                       string                 `json:"currency"`
	OrgName                          string                 `json:"org_name"`
	DepartmentId                     string                 `json:"department"`
	PublishedTime                    string                 `json:"pubts"`
	ReceiverMobileNumber             string                 `json:"cReceiveMobile"`
	Creator                          string                 `json:"creator"`
	ReceiveZipCode                   string                 `json:"cReceiveZipCode"`
	OrgId                            string                 `json:"org"`
	DepartmentName                   string                 `json:"department_name"`
	ExchangeRateName                 string                 `json:"exchRateType_name"`
	AccountOrgId                     string                 `json:"accountOrg"`
	TransactionTypeKeyField          string                 `json:"transTypeKeyField"`
	BusinessTypeId                   string                 `json:"bustype"`
	IsRetailInvestors                bool                   `json:"retailInvestors"`
	ReceiverName                     string                 `json:"cReceiver"`
	CreatedTime                      string                 `json:"createTime"`
	LogisticId                       int64                  `json:"iLogisticId"`
	CustomerId                       int64                  `json:"cust"`
	CurrencyCode                     string                 `json:"currencyCode"`
	Memo                             string                 `json:"memo"`
	SalesmanId                       string                 `json:"operator"`
	SalesmanName                     string                 `json:"operator_name"`
	Modifier                         string                 `json:"modifier"`
	ModifiedTime                     string                 `json:"modifyTime"`
	TotalQty                         float64                `json:"totalQuantity"`
	InvoiceTitleType                 string                 `json:"invoiceTitleType"`
	WMSCancelStatus                  int                    `json:"wmsCancelStatus"`
	DeliverStatus                    string                 `json:"diliverStatus"`
	TotalPieces                      float64                `json:"totalPieces"`
	WMSInStatus                      int                    `json:"wmsInStatus"`
	ReceiveId                        int64                  `json:"receiveId"`
	InvoiceType                      string                 `json:"invoiceUpcType"`
	SalesOutDefineCharacter          map[string]interface{} `json:"salesOutDefineCharacter"` // 特征
	Items                            []DetailItemDto        `json:"details"`
	WarehouseName                    string                 `json:"-"`
	// 用友说不用传这些字段 2024-10-12 11:28:38
	// IsCountCostWarehouse             bool                   `json:"warehouse_countCost"`
	// IsSerialNumberManageWarehouse    bool                   `json:"warehouse_iSerialManage"`
	// IsGoodsPositionWarehouse         bool                   `json:"warehouse_isGoodsPosition"`
	// IsGoodsPositionStockWarehouse    bool                   `json:"warehouse_isGoodsPositionStock"`
	// WarehouseName                    string                 `json:"warehouse_name"`
}

type DetailItemDto struct {
	StockUnitName          string         `json:"stockUnit_name"`
	OriginalTax            float64        `json:"oriTax"`
	ProductCode            string         `json:"product_cCode"`
	OrderId                int64          `json:"orderId"`
	PriceUOMPrecision      int            `json:"priceUOM_Precision"`
	NativeTax              float64        `json:"natTax"`
	SourceOrderType        string         `json:"source"`
	StockStatusDocName     string         `json:"stockStatusDoc_name"`
	SubQty                 float64        `json:"subQty"`
	ProductName            string         `json:"product_cName"`
	IsExpiryDateManage     bool           `json:"isExpiryDateManage"`
	IsReservation          bool           `json:"reservation"`
	StockUnitPrecision     int            `json:"stockUnitId_Precision"`
	Id                     int64          `json:"id"`
	MainId                 int64          `json:"mainid"`
	IsSerialNumberManage   bool           `json:"isSerialNoManage"`
	UnitName               string         `json:"unitName"`
	OriginalUnitPrice      float64        `json:"oriUnitPrice"`
	NativeSummary          float64        `json:"natSum"`
	IsScrap                bool           `json:"isScrap"`
	TaxRate                float64        `json:"taxRate"`
	UnitId                 int64          `json:"unit"`
	ProductSKUId           int64          `json:"productsku"`
	ProductSKUCode         string         `json:"productsku_cCode"`
	UnitPrecision          int            `json:"unit_Precision"`
	Qty                    float64        `json:"qty"`
	OriginalTaxUnitPrice   float64        `json:"oriTaxUnitPrice"`
	OriginalMoney          float64        `json:"oriMoney"`
	ExchangeRate           float64        `json:"invExchRate"`
	StockUnitId            int64          `json:"stockUnitId"`
	NativeUnitPrice        float64        `json:"natUnitPrice"`
	IsAutoCalculateCost    bool           `json:"autoCalcCost"`
	ReserveId              int64          `json:"reserveid"`
	MakeRuleCode           string         `json:"makeRuleCode"`
	StockStatusDocId       int64          `json:"stockStatusDoc"`
	ProductSKUName         string         `json:"productsku_cName"`
	PriceUOMId             int64          `json:"priceUOM"`
	PriceExchangeRate      float64        `json:"invPriceExchRate"`
	IsBatchManage          bool           `json:"isBatchManage"`
	PublishedTime          string         `json:"pubts"`
	SourceOrderMainId      int64          `json:"sourceid"`
	ProductId              int64          `json:"product"`
	OriginalSummary        float64        `json:"oriSum"`
	PriceUOMName           string         `json:"priceUOM_name"`
	DetailId               int64          `json:"detailid"`
	IsRebate               bool           `json:"rebateFlag"`
	SourceOrderSubId       int64          `json:"sourceautoid"`
	PriceQty               float64        `json:"priceQty"`
	IsTaxUnitPrice         bool           `json:"taxUnitPriceTag"`
	SourceOrderNumber      string         `json:"upcode"`
	SaleStyle              string         `json:"saleStyle"`
	NotInvoicedQty         float64        `json:"unInvoiceQty"`
	NativeTaxUnitPrice     float64        `json:"natTaxUnitPrice"`
	NativeMoney            float64        `json:"natMoney"`
	ContractQty            float64        `json:"contactsQuantity"`
	FirstSourceOrderType   string         `json:"firstsource"`
	FirstSourceOrderNumber string         `json:"firstupcode"`
	TaxItemName            string         `json:"taxItems"`
	TaxCode                string         `json:"taxCode"`
	RebateMoney            float64        `json:"rebateMoney"`
	TaxId                  string         `json:"taxId"`
	ContractPieces         float64        `json:"contactsPieces"`
	FirstSourceOrderSubId  int64          `json:"firstsourceautoid"`
	FirstSourceOrderMainId int64          `json:"firstsourceid"`
	RowNumber              int            `json:"rowno"`
	UnitExchangeType       int            `json:"unitExchangeType"`
	CashRebateMoney        float64        `json:"cashRebateMoney"`
	OrderDetailId          int64          `json:"orderDetailId"`
	OrderRebateMoney       float64        `json:"orderRebateMoney"`
	SourceOrderRowId       int64          `json:"srcBillRow"`
	OrderNumber            string         `json:"orderCode"`
	SerialNumbers          []SerialNumber `json:"salesOutsSNs"`
}

type SerialNumber struct {
	OrderDetailId int64  `json:"detailid"`
	Id            int64  `json:"id"`
	SerialNumber  string `json:"sn"`
	PublishedTime string `json:"pubts"`
	OrderId       int64  `json:"grandpaid"`
}
