package ysapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLs(t *testing.T) {
	tests := []struct {
		given string
		want  string
	}{
		{given: WarehouseListURL, want: "https://api.diwork.com/yonbip/digitalModel/warehouse/list"},
		{given: QueryCurrentStocksByConditionURL, want: "https://api.diwork.com/yonbip/scm/stock/QueryCurrentStocksByCondition"},
		{given: VoucherOrderBatchAuditURL, want: "https://api.diwork.com/yonbip/sd/voucherorder/batchaudit"},
		{given: TradeVouchQueryURL, want: "https://api.diwork.com/yonbip/sd/dst/tradevouch/query"},
		{given: TransferApplySaveURL, want: "https://api.diwork.com/yonbip/scm/transferapply/save"},
		{given: TransferApplyListURL, want: "https://api.diwork.com/yonbip/scm/transferapply/list"},
		{given: StoreOutDetailURL, want: "https://api.diwork.com/yonbip/scm/storeout/detail"},
		{given: StoreInListURL, want: "https://api.diwork.com/yonbip/scm/storein/list"},
		{given: StoreTransferListURL, want: "https://api.diwork.com/yonbip/scm/storetransfer/list"},
		{given: CurrencyListURL, want: "https://api.diwork.com/yonbip/digitalModel/currency/list"},
		{given: ExchangeRateListURL, want: "https://api.diwork.com/yonbip/digitalModel/exchangerate/list"},
		{given: GoodsPositionTree, want: "https://api.diwork.com/yonbip/digitalModel/goodsposition/tree"},
		{given: LocationStockAnalysisList, want: "https://api.diwork.com/yonbip/scm/locationstockanalysis/list"},
		{given: StockAnalysisList, want: "https://api.diwork.com/yonbip/scm/stockanalysis/list"},
		{given: GoodsProductSKUProList, want: "https://api.diwork.com/yonbip/digitalModel/goodsproductskupro/list"},
		{given: TransferReqSaveURL, want: "https://api.diwork.com/yonbip/scm/transferreq/save"},
		{given: TransferReqBatchAuditURL, want: "https://api.diwork.com/yonbip/scm/transferreq/batchaudit"},
		{given: MorphologyConversionSaveURL, want: "https://api.diwork.com/yonbip/scm/morphologyconversion/save"},
		{given: MorphologyConversionBatchAuditURL, want: "https://api.diwork.com/yonbip/scm/morphologyconversion/batchaudit"},
		{given: VoucherDeliveryListURL, want: "https://api.diwork.com/yonbip/sd/voucherdelivery/list"},
		{given: VoucherDeliveryUnauditURL, want: "https://api.diwork.com/yonbip/sd/voucherdelivery/unaudit"},
		{given: VoucherDeliveryBatchDeleteURL, want: "https://api.diwork.com/yonbip/sd/voucherdelivery/batchdelete"},
		{given: SuitGoodsQueryURL, want: "https://api.diwork.com/yonbip/sd/dst/suitgoods/query"},
		{given: TradeOrderQueryURL, want: "https://api.diwork.com/yonbip/sd/dst/tradeorder/query"},
		{given: TradeOrderHoldURL, want: "https://api.diwork.com/yonbip/sd/dst/tradeorder/tradehold"},
		{given: VoucherOrderCloseURL, want: "https://api.diwork.com/yonbip/sd/voucherorder/close"},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, test.given)
	}
}
