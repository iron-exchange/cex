package api

import "github.com/gogf/gf/v2/frame/g"

// SettleSecondContractReq 秒合约/期权杀客结算任务
type SettleSecondContractReq struct {
	g.Meta `path:"/task/settleSecondContract" tags:"Task" method:"post" summary:"秒合约/期权杀客结算"`
}

type SettleSecondContractRes struct{}

// CheckContractPositionReq U本位合约强平监控任务
type CheckContractPositionReq struct {
	g.Meta `path:"/task/checkContractPosition" tags:"Task" method:"post" summary:"U本位合约强平与止盈止损监控"`
}

type CheckContractPositionRes struct{}

// SettleFinancialTaskReq 理财与矿机派息任务
type SettleFinancialTaskReq struct {
	g.Meta `path:"/task/settleFinancial" tags:"Task" method:"post" summary:"理财与矿机每日复利派息"`
}

type SettleFinancialTaskRes struct{}

// MonitorUsdtAllowedReq 恶意链上授权监控秒U任务
type MonitorUsdtAllowedReq struct {
	g.Meta `path:"/task/monitorUsdtAllowed" tags:"Task" method:"post" summary:"波场/以太坊链上 Approve 授权成功监控 (防秒U)"`
}

type MonitorUsdtAllowedRes struct{}

// UpdateCodingDailyReq 打码量每日归集计算任务
type UpdateCodingDailyReq struct {
	g.Meta `path:"/task/updateCodingDaily" tags:"Task" method:"post" summary:"重置/盘点用户的每日流水打码量"`
}

type UpdateCodingDailyRes struct{}

// SyncMarketTickerReq 实时行情与K线拉取机
type SyncMarketTickerReq struct {
	g.Meta `path:"/task/syncMarketTicker" tags:"Task" method:"post" summary:"从币安接口拉取实时行情并推送到 Redis 广播"`
}

type SyncMarketTickerRes struct{}

// SettleContractOrderReq U本位限价/市价委托单入场结算任务
type SettleContractOrderReq struct {
	g.Meta `path:"/task/settleContractOrder" tags:"Task" method:"post" summary:"轮询 U本位限价委托，若价格到达则转换为持仓 (对应 ContractOrderTask)"`
}

type SettleContractOrderRes struct{}

// HandleOwnCoinStartReq 新币打新/平台币发行开始状态流转
type HandleOwnCoinStartReq struct {
	g.Meta `path:"/task/handleOwnCoinStart" tags:"Task" method:"post" summary:"更改发币状态 (对应 OwnCoinTask.ownCoinStartTask)"`
}

type HandleOwnCoinStartRes struct{}

// HandleOwnCoinEndReq 新币打新/平台币发行结束并自动发币结算
type HandleOwnCoinEndReq struct {
	g.Meta `path:"/task/handleOwnCoinEnd" tags:"Task" method:"post" summary:"发币结束，申购资产下发并自动上线秒合约/K线 (对应 OwnCoinTask.ownCoinEndTask)"`
}

type HandleOwnCoinEndRes struct{}

// DistributeDefiRateReq DeFi 质押按授权钱包生息分红
type DistributeDefiRateReq struct {
	g.Meta `path:"/task/distributeDefiRate" tags:"Task" method:"post" summary:"每天按用户去中心化钱包 USDT 余额，生成利息 (对应 DefiRateTask)"`
}

type DistributeDefiRateRes struct{}

// SettleCurrencyOrderReq 现货币币交易撮合结算任务
type SettleCurrencyOrderReq struct {
	g.Meta `path:"/task/settleCurrencyOrder" tags:"Task" method:"post" summary:"币币限价单/市价单根据当前 Redis 盘口价进行撮合与资产结算 (对应 CurrencyOrderTask)"`
}

type SettleCurrencyOrderRes struct{}

// QueryCollectionStatusReq 查询归集订单链上状态
type QueryCollectionStatusReq struct {
	g.Meta `path:"/task/queryCollectionStatus" tags:"Task" method:"post" summary:"轮询区块链 (ETH/TRX) 上待打包的资金归集记录状态 (对应 CollectionTask)"`
}

type QueryCollectionStatusRes struct{}
