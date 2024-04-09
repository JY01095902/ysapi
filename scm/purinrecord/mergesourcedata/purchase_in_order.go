package mergesourcedata

type SerialNumber struct {
	SerialNumber string `json:"sn"`
	Status       string `json:"_status"`
}

type PurchaseInOrderDetail struct {
	MakeRuleCode                string                 `json:"makeRuleCode"`
	SourceId                    string                 `json:"sourceid"`
	SourceAutoId                string                 `json:"sourceautoid"`
	BatchNo                     string                 `json:"batchno"`
	ProduceDate                 string                 `json:"producedate"`
	GoodsPosition               string                 `json:"goodsposition"`
	Qty                         float64                `json:"qty"`
	TaxItems                    string                 `json:"taxitems"`
	OriUnitPrice                float64                `json:"oriUnitPrice"`
	OriTaxUnitPrice             float64                `json:"oriTaxUnitPrice"`
	OriMoney                    float64                `json:"oriMoney"`
	OriSum                      float64                `json:"oriSum"`
	NatUnitPrice                float64                `json:"natUnitPrice"`
	NatTaxUnitPrice             float64                `json:"natTaxUnitPrice"`
	NatMoney                    float64                `json:"natMoney"`
	NatSum                      float64                `json:"natSum"`
	Memo                        string                 `json:"memo"`
	Status                      string                 `json:"_status"`
	PurInRecordsSNs             []SerialNumber         `json:"purInRecordsSNs"`
	PurInRecordsDefineCharacter map[string]interface{} `json:"purInRecordsDefineCharacter"`
	PurInRecordsCharacteristics map[string]interface{} `json:"purInRecordsCharacteristics"`
}

type PurchaseInOrder struct {
	MergeSourceData            string                  `json:"mergeSourceData"`
	NeedCalcLines              bool                    `json:"needCalcLines"`
	CalcLinesKey               string                  `json:"calcLinesKey"`
	Code                       string                  `json:"code"`
	VouchDate                  string                  `json:"vouchdate"`
	BusinessType               string                  `json:"bustype"`
	Warehouse                  string                  `json:"warehouse"`
	Memo                       string                  `json:"memo"`
	Status                     string                  `json:"_status"`
	PurInRecordDefineCharacter map[string]interface{}  `json:"purInRecordDefineCharacter"`
	Details                    []PurchaseInOrderDetail `json:"purInRecords"`
}
