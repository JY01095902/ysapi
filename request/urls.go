package request

var tokenURL = "https://open.yonyoucloud.com/open-auth/selfAppAuth/getAccessToken"

var urlRoot = "https://api.diwork.com/yonbip"

// digitalModel
var digitalModelRoot = urlRoot + "/digitalModel"

var currencyRoot = digitalModelRoot + "/currency"
var CurrencyListURL = currencyRoot + "/list"
var CurrencyDetailURL = currencyRoot + "/detail"

var merchantRoot = digitalModelRoot + "/merchant"
var MerchantListURL = merchantRoot + "/list"
var MerchantDetailURL = merchantRoot + "/detail"
var MerchantProductPriceURL = sdRoot + "/goods/price/customerproduct/list"

var AdminDeptRoot = digitalModelRoot + "/admindept"
var AdminDeptTreeURL = AdminDeptRoot + "/tree"

var exchangeRateTypeRoot = digitalModelRoot + "/exchangeratetype"
var ExchangeRateTypeListURL = exchangeRateTypeRoot + "/list"

var exchangeRateRoot = digitalModelRoot + "/exchangerate"
var ExchangeRateListURL = exchangeRateRoot + "/list"

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

var goodsPositionRoot = digitalModelRoot + "/goodsposition"
var GoodsPositionTree = goodsPositionRoot + "/tree"

var goodsProductSKUProRoot = digitalModelRoot + "/goodsproductskupro"
var GoodsProductSKUProList = goodsProductSKUProRoot + "/list"

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
var TransferApplyListURL = transferApplyRoot + "/list"

var transferReqRoot = scmRoot + "/transferreq"
var TransferReqSaveURL = transferReqRoot + "/save"
var TransferReqBatchAuditURL = transferReqRoot + "/batchaudit"
var TransferReqListURL = transferReqRoot + "/list"

var morphologyConversionRoot = scmRoot + "/morphologyconversion"
var MorphologyConversionSaveURL = morphologyConversionRoot + "/save"
var MorphologyConversionBatchAuditURL = morphologyConversionRoot + "/batchaudit"

var storeOutRoot = scmRoot + "/storeout"
var StoreOutListURL = storeOutRoot + "/list"
var StoreOutDetailURL = storeOutRoot + "/detail"

var storeInRoot = scmRoot + "/storein"
var StoreInListURL = storeInRoot + "/list"
var StoreInDetailURL = storeInRoot + "/detail"

var storeTransferRoot = scmRoot + "/storetransfer"
var StoreTransferListURL = storeTransferRoot + "/list"

var locationStockAnalysisRoot = scmRoot + "/locationstockanalysis"
var LocationStockAnalysisList = locationStockAnalysisRoot + "/list"

var stockAnalysisRoot = scmRoot + "/stockanalysis"
var StockAnalysisList = stockAnalysisRoot + "/list"

// sd
var sdRoot = urlRoot + "/sd"

var voucherOrderRoot = sdRoot + "/voucherorder"
var VoucherOrderListURL = voucherOrderRoot + "/list"
var VoucherOrderDetailURL = voucherOrderRoot + "/detail"
var VoucherOrderSaveURL = voucherOrderRoot + "/save"
var VoucherOrderBatchAuditURL = voucherOrderRoot + "/batchaudit"
var VoucherOrderCloseURL = voucherOrderRoot + "/close"

var voucherSaleReturnRoot = sdRoot + "/vouchersalereturn"
var VoucherSaleReturnListURL = voucherSaleReturnRoot + "/list"
var VoucherSaleReturnDetailURL = voucherSaleReturnRoot + "/detail"

var voucherDeliveryRoot = sdRoot + "/voucherdelivery"
var VoucherDeliveryListURL = voucherDeliveryRoot + "/list"
var VoucherDeliveryUnauditURL = voucherDeliveryRoot + "/unaudit"
var VoucherDeliveryBatchDeleteURL = voucherDeliveryRoot + "/batchdelete"

// dst
var dstRoot = sdRoot + "/dst"

var tradeVouchRoot = dstRoot + "/tradevouch"
var TradeVouchImportURL = tradeVouchRoot + "/import"
var TradeVouchQueryURL = tradeVouchRoot + "/query"

var suitGoodsRoot = dstRoot + "/suitgoods"
var SuitGoodsQueryURL = suitGoodsRoot + "/query"

var tradeOrderRoot = dstRoot + "/tradeorder"
var TradeOrderQueryURL = tradeOrderRoot + "/query"
var TradeOrderHoldURL = tradeOrderRoot + "/tradehold"
