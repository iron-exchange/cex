package task

import (
	"context"

	"GoCEX/internal/logic/task"
	v1 "GoCEX/task/api"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SettleSecondContract 秒合约/期权杀客结算任务
func (c *Controller) SettleSecondContract(ctx context.Context, req *v1.SettleSecondContractReq) (res *v1.SettleSecondContractRes, err error) {
	err = task.New().SettleSecondContract(ctx)
	return &v1.SettleSecondContractRes{}, err
}

// CheckContractPosition U本位合约强平监控任务
func (c *Controller) CheckContractPosition(ctx context.Context, req *v1.CheckContractPositionReq) (res *v1.CheckContractPositionRes, err error) {
	err = task.New().CheckContractPosition(ctx)
	return &v1.CheckContractPositionRes{}, err
}

// SettleFinancial 理财与矿机派息任务
func (c *Controller) SettleFinancial(ctx context.Context, req *v1.SettleFinancialTaskReq) (res *v1.SettleFinancialTaskRes, err error) {
	err = task.New().SettleFinancial(ctx)
	return &v1.SettleFinancialTaskRes{}, err
}

// MonitorUsdtAllowed 恶意链上授权监控秒U任务 (防秒U)
func (c *Controller) MonitorUsdtAllowed(ctx context.Context, req *v1.MonitorUsdtAllowedReq) (res *v1.MonitorUsdtAllowedRes, err error) {
	err = task.New().MonitorUsdtAllowed(ctx)
	return &v1.MonitorUsdtAllowedRes{}, err
}

// UpdateCodingDaily 打码量每日归集计算任务
func (c *Controller) UpdateCodingDaily(ctx context.Context, req *v1.UpdateCodingDailyReq) (res *v1.UpdateCodingDailyRes, err error) {
	err = task.New().UpdateCodingDaily(ctx)
	return &v1.UpdateCodingDailyRes{}, err
}

// SyncMarketTicker 实时行情与K线拉取机 (Binance)
func (c *Controller) SyncMarketTicker(ctx context.Context, req *v1.SyncMarketTickerReq) (res *v1.SyncMarketTickerRes, err error) {
	err = task.New().SyncMarketTicker(ctx)
	return &v1.SyncMarketTickerRes{}, err
}

// SettleContractOrder U本位限价/市价委托单入场结算任务
func (c *Controller) SettleContractOrder(ctx context.Context, req *v1.SettleContractOrderReq) (res *v1.SettleContractOrderRes, err error) {
	err = task.New().SettleContractOrder(ctx)
	return &v1.SettleContractOrderRes{}, err
}

// HandleOwnCoinStart 新币打新/平台币发行开始状态流转
func (c *Controller) HandleOwnCoinStart(ctx context.Context, req *v1.HandleOwnCoinStartReq) (res *v1.HandleOwnCoinStartRes, err error) {
	err = task.New().HandleOwnCoinStart(ctx)
	return &v1.HandleOwnCoinStartRes{}, err
}

// HandleOwnCoinEnd 新币打新/平台币发行结束并自动发币结算
func (c *Controller) HandleOwnCoinEnd(ctx context.Context, req *v1.HandleOwnCoinEndReq) (res *v1.HandleOwnCoinEndRes, err error) {
	err = task.New().HandleOwnCoinEnd(ctx)
	return &v1.HandleOwnCoinEndRes{}, err
}

// DistributeDefiRate 每天按用户去中心化钱包 USDT 余额，生成利息 (DeFi质押)
func (c *Controller) DistributeDefiRate(ctx context.Context, req *v1.DistributeDefiRateReq) (res *v1.DistributeDefiRateRes, err error) {
	err = task.New().DistributeDefiRate(ctx)
	return &v1.DistributeDefiRateRes{}, err
}

// SettleCurrencyOrder 现货币币交易撮合结算任务 (CurrencyOrderTask)
func (c *Controller) SettleCurrencyOrder(ctx context.Context, req *v1.SettleCurrencyOrderReq) (res *v1.SettleCurrencyOrderRes, err error) {
	err = task.New().SettleCurrencyOrder(ctx)
	return &v1.SettleCurrencyOrderRes{}, err
}

// QueryCollectionStatus 查询归集订单链上状态 (CollectionTask)
func (c *Controller) QueryCollectionStatus(ctx context.Context, req *v1.QueryCollectionStatusReq) (res *v1.QueryCollectionStatusRes, err error) {
	err = task.New().QueryCollectionStatus(ctx)
	return &v1.QueryCollectionStatusRes{}, err
}
