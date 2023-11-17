package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLs(t *testing.T) {
	tests := []struct {
		given string
		want  string
	}{
		{given: WarehouseListURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/digitalModel/warehouse/list"},
		{given: QueryCurrentStocksByConditionURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/stock/QueryCurrentStocksByCondition"},
		{given: VoucherOrderBatchAuditURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/voucherorder/batchaudit"},
		{given: TradeVouchQueryURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/dst/tradevouch/query"},
		{given: TransferApplySaveURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/transferapply/save"},
		{given: TransferApplyListURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/transferapply/list"},
		{given: StoreOutDetailURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/storeout/detail"},
		{given: StoreInListURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/storein/list"},
		{given: StoreTransferListURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/storetransfer/list"},
		{given: CurrencyListURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/digitalModel/currency/list"},
		{given: ExchangeRateListURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/digitalModel/exchangerate/list"},
		{given: GoodsPositionTree, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/digitalModel/goodsposition/tree"},
		{given: LocationStockAnalysisList, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/locationstockanalysis/list"},
		{given: StockAnalysisList, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/stockanalysis/list"},
		{given: GoodsProductSKUProList, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/digitalModel/goodsproductskupro/list"},
		{given: TransferReqSaveURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/transferreq/save"},
		{given: TransferReqBatchAuditURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/transferreq/batchaudit"},
		{given: MorphologyConversionSaveURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/morphologyconversion/save"},
		{given: MorphologyConversionBatchAuditURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/scm/morphologyconversion/batchaudit"},
		{given: VoucherDeliveryListURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/voucherdelivery/list"},
		{given: VoucherDeliveryUnauditURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/voucherdelivery/unaudit"},
		{given: VoucherDeliveryBatchDeleteURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/voucherdelivery/batchdelete"},
		{given: SuitGoodsQueryURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/dst/suitgoods/query"},
		{given: TradeOrderQueryURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/dst/tradeorder/query"},
		{given: TradeOrderHoldURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/dst/tradeorder/tradehold"},
		{given: VoucherOrderCloseURL, want: "https://c2.yonyoucloud.com/iuap-api-gateway/yonbip/sd/voucherorder/close"},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, test.given)
	}
}
