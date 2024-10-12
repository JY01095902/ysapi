package single

type OtherInOrder struct {
	TotalPieces                      float64                `json:"totalPieces"`
	OrderNumber                      string                 `json:"code"`
	MasterOrgKeyField                string                 `json:"masterOrgKeyField"`
	BusinessTypeId                   string                 `json:"bustype"`
	BusinessTypeName                 string                 `json:"bustype_name"`
	Date                             string                 `json:"vouchdate"`
	NativeCurrencyId                 string                 `json:"natCurrency"`
	TotalQty                         float64                `json:"totalQuantity"`
	Id                               int64                  `json:"id"`
	OrgId                            string                 `json:"org"`
	OrgName                          string                 `json:"org_name"`
	BusinessTypeExtendAttributesJson string                 `json:"bustype_extend_attrs_json,omitempty"`
	PublishedTime                    string                 `json:"pubts"`
	AccountOrgId                     string                 `json:"accountOrg"`
	WarehouseId                      int64                  `json:"warehouse"`
	TransactionTypeKeyField          string                 `json:"transTypeKeyField"`
	NativeCurrencyMoneyDigit         int                    `json:"natCurrency_moneyDigit"`
	NativeCurrencyPriceDigit         int                    `json:"natCurrency_priceDigit"`
	AccountOrgName                   string                 `json:"accountOrg_name"`
	Status                           int                    `json:"status"`
	IsUpdateCost                     bool                   `json:"isUpdateCost"`
	IsCostStart                      bool                   `json:"costStart"`
	Memo                             string                 `json:"memo"`
	SalesmanId                       string                 `json:"operator"`
	SalesmanName                     string                 `json:"operator_name"`
	StockManagerId                   string                 `json:"stockMgr"`
	StockManagerName                 string                 `json:"stockMgr_name"`
	DepartmentId                     string                 `json:"department"`
	DepartmentName                   string                 `json:"department_name"`
	CreatorId                        int64                  `json:"creatorId"`
	CreatorName                      string                 `json:"creator"`
	CreatedDate                      string                 `json:"createDate"`
	CreatedTime                      string                 `json:"createTime"`
	ModifierId                       int64                  `json:"modifierId"`
	ModifierName                     string                 `json:"modifier"`
	ModifiedDate                     string                 `json:"modifyDate"`
	ModifiedTime                     string                 `json:"modifyTime"`
	PaymentNumber                    string                 `json:"defines!define1,omitempty"`  // 读取，但是测试和运营环境可能不同，这个字段会覆盖defines里的值
	Defines                          map[string]interface{} `json:"defines"`                    // 写入
	OthInRecordDefineCharacter       map[string]interface{} `json:"othInRecordDefineCharacter"` // 新版 特征组
	Details                          []OtherInOrderDetail   `json:"othInRecords,omitempty"`
	WarehouseName                    string                 `json:"-"`
	// 用友说不用传这些字段 2024-10-12 11:28:38
	// IsCountCostWarehouse             bool                   `json:"warehouse_countCost"`
	// IsSerialNumberManageWarehouse    bool                   `json:"warehouse_iSerialManage"`
	// IsGoodsPositionWarehouse         bool                   `json:"warehouse_isGoodsPosition"`
	// IsGoodsPositionStockWarehouse    bool                   `json:"warehouse_isGoodsPositionStock"`
	// WarehouseName                    string                 `json:"warehouse_name"`
}

type OtherInOrderDetail struct {
	RowNumber            int64          `json:"rowno"`
	StockUnitId          int64          `json:"stockUnitId"`
	StockUnitName        string         `json:"stockUnit_name"`
	IsAutoCalculateCost  bool           `json:"autoCalcCost"`
	ReserveId            int64          `json:"reserveid"`
	ProductClassId       int64          `json:"productClass"`
	ProductClassName     string         `json:"productClassName"`
	ProductId            int64          `json:"product"`
	ProductCode          string         `json:"product_cCode"`
	ProductName          string         `json:"product_cName"`
	ProductUnitName      string         `json:"product_unitName"`
	ProductSKUId         int64          `json:"productsku"`
	ProductSKUName       string         `json:"productsku_cName"`
	StockStatusDocId     int64          `json:"stockStatusDoc"`
	ProductSKUCode       string         `json:"productsku_cCode"`
	StockStatusDocName   string         `json:"stockStatusDoc_name"`
	IsBatchManage        bool           `json:"isBatchManage"`
	IsExpiryDateManage   bool           `json:"isExpiryDateManage"`
	StockUnitPrecision   int            `json:"stockUnitId_Precision"`
	Id                   int64          `json:"id"`
	MainId               int64          `json:"mainid"`
	PublishedTime        string         `json:"pubts"`
	IsSerialNumberManage bool           `json:"isSerialNoManage"`
	UnitExchangeType     int            `json:"unitExchangeType"`
	RecordDate           string         `json:"recorddate"`
	IsScrap              bool           `json:"isScrap"`
	TaxRate              float64        `json:"taxRate"`
	UnitId               int64          `json:"unit"`
	Qty                  float64        `json:"qty"`
	SubQty               float64        `json:"subQty"`
	ContactsQty          float64        `json:"contactsQuantity"`
	ContactsPieces       float64        `json:"contactsPieces"`
	UnitPrecision        int            `json:"unit_Precision"`
	ExchangeRate         float64        `json:"invExchRate"`
	SerialNumbers        []SerialNumber `json:"othInRecordsSNs"`
	PositionId           interface{}    `json:"goodsposition"` // 保存时是string，查询时是int64
	UpCode               string         `json:"upcode"`
}

type SerialNumber struct {
	OrderDetailId int64  `json:"detailid"`
	Id            int64  `json:"id"`
	SerialNumber  string `json:"sn"`
	PublishedTime string `json:"pubts"`
	OrderId       int64  `json:"grandpaid"`
}
