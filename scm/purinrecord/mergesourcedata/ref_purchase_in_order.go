package mergesourcedata

type RefPurchaseInOrderDetail struct {
	MakeRuleCode                string                 `json:"makeRuleCode"`
	SourceId                    string                 `json:"sourceid"`
	SourceAutoId                string                 `json:"sourceautoid"`
	GoodsPosition               string                 `json:"goodsposition"`
	Status                      string                 `json:"_status"`
	PurInRecordsSNs             []SerialNumber         `json:"purInRecordsSNs"`
	PurInRecordsDefineCharacter map[string]interface{} `json:"purInRecordsDefineCharacter"`
	PurInRecordsCharacteristics map[string]interface{} `json:"purInRecordsCharacteristics"`
}

// 参照采购订单生成采购入库时用的，省略了不必要的字段（都从采购订单获取）
type RefPurchaseInOrder struct {
	MergeSourceData            bool                       `json:"mergeSourceData"`
	Code                       string                     `json:"code"`
	VouchDate                  string                     `json:"vouchdate"`
	BusinessType               string                     `json:"bustype"`
	Warehouse                  string                     `json:"warehouse"`
	Memo                       string                     `json:"memo"`
	Status                     string                     `json:"_status"`
	PurInRecordDefineCharacter map[string]interface{}     `json:"purInRecordDefineCharacter"`
	Details                    []RefPurchaseInOrderDetail `json:"purInRecords"`
}
