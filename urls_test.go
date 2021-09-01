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
	}

	for _, test := range tests {
		assert.Equal(t, test.want, test.given)
	}
}
