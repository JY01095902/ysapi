package ysapi

var tokenURL = "https://open.yonyoucloud.com/open-auth/selfAppAuth/getAccessToken"

var urlRoot = "https://api.diwork.com/yonbip"

// digitalModel
var digitalModelRoot = urlRoot + "/digitalModel"

var currencyRoot = digitalModelRoot + "/currency"
var CurrencyListURL = currencyRoot + "/currency/list"
var CurrencyDetailURL = currencyRoot + "/currency/detail"

var merchantRoot = digitalModelRoot + "/merchant"
var MerchantListURL = merchantRoot + "/list"
var MerchantDetailURL = merchantRoot + "/detail"
var MerchantProductPriceURL = sdRoot + "/goods/price/customerproduct/list"

var AdminDeptRoot = digitalModelRoot + "/admindept"
var AdminDeptTreeURL = AdminDeptRoot + "/tree"

var exchangeRateTypeRoot = digitalModelRoot + "/exchangeratetype"
var ExchangeRateTypeListURL = exchangeRateTypeRoot + "/list"

var orgUnitRoot = digitalModelRoot + "/orgunit"
var OrgUnitTreeURL = orgUnitRoot + "/querytree"

var productRoot = digitalModelRoot + "/product"
var ProductListURL = productRoot + "/list"
var ProductDetailURL = productRoot + "/detail"

var staffRoot = digitalModelRoot + "/staff"
var StaffListURL = staffRoot + "/list"
var StaffDetailURL = staffRoot + "/detail"

var transactionRoot = digitalModelRoot + "/transtype"
var TransactionListURL = transactionRoot + "/list"
var TransactionTreeURL = transactionRoot + "/tree"

var warehouseRoot = digitalModelRoot + "/warehouse"
var WarehouseListURL = warehouseRoot + "/list"
var WarehouseDetailURL = warehouseRoot + "/detail"

// scm
var scmRoot = urlRoot + "/scm"

var othInRecordRoot = scmRoot + "/othinrecord"
var OthInRecordListURL = othInRecordRoot + "/list"
var OthInRecordDetailURL = othInRecordRoot + "/detail"

var othOutRecordRoot = scmRoot + "/othoutrecord"
var OthOutRecordListURL = othOutRecordRoot + "/list"
var OthOutRecordDetailURL = othOutRecordRoot + "/detail"

var purInRecord = scmRoot + "/purinrecord"
var PurInRecordListURL = purInRecord + "/list"
var PurInRecordDetailURL = purInRecord + "/detail"

var salesOutRoot = scmRoot + "/salesout"
var SalesOutListURL = salesOutRoot + "/list"
var SalesOutDetailURL = salesOutRoot + "/detail"

var snFlowDirectionRoot = scmRoot + "/snflowdirection"
var SNFlowDirectionListURL = snFlowDirectionRoot + "/list"

var snQuerySNStateRoot = scmRoot + "/snQuerysnstate"
var SNQuerySNStateListURL = snQuerySNStateRoot + "/list"

var stockRoot = scmRoot + "/stock"
var QueryCurrentStocksByConditionURL = stockRoot + "/QueryCurrentStocksByCondition"

var transferApplyRoot = scmRoot + "/transferapply"
var TransferApplySaveURL = transferApplyRoot + "/save"
var TransferApplyBatchAuditURL = transferApplyRoot + "/batchaudit"

var storeOutRoot = scmRoot + "/storeout"
var StoreOutListURL = storeOutRoot + "/list"
var StoreOutDetailURL = storeOutRoot + "/detail"

var storeInRoot = scmRoot + "/storein"
var StoreInListURL = storeInRoot + "/list"
var StoreInDetailURL = storeInRoot + "/detail"

var storeTransferRoot = scmRoot + "/storetransfer"
var StoreTransferListURL = storeTransferRoot + "/list"

// sd
var sdRoot = urlRoot + "/sd"

var voucherOrderRoot = sdRoot + "/voucherorder"
var VoucherOrderListURL = voucherOrderRoot + "/list"
var VoucherOrderDetailURL = voucherOrderRoot + "/detail"
var VoucherOrderSaveURL = voucherOrderRoot + "/save"
var VoucherOrderBatchAuditURL = voucherOrderRoot + "/batchaudit"

var voucherSaleReturnRoot = sdRoot + "/vouchersalereturn"
var VoucherSaleReturnListURL = voucherSaleReturnRoot + "/list"
var VoucherSaleReturnDetailURL = voucherSaleReturnRoot + "/detail"

// dst
var dstRoot = sdRoot + "/dst"

var tradeVouchRoot = dstRoot + "/tradevouch"
var TradeVouchImportURL = tradeVouchRoot + "/import"
var TradeVouchQueryURL = tradeVouchRoot + "/query"
