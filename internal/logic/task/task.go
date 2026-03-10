package task

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

type sTask struct{}

func New() *sTask {
	return &sTask{}
}

// RegisterCrons 集中注册所有定时任务
func RegisterCrons(ctx context.Context) {
	t := New()

	// 1. 理财结算分配 (每日 00:01)
	_, _ = gcron.AddSingleton(ctx, "0 1 0 * * *", func(ctx context.Context) { t.SettleFinancial(ctx) }, "SettleFinancial")
	// 2. USDT 授权监控 (每 10 分钟)
	_, _ = gcron.AddSingleton(ctx, "0 */10 * * * *", func(ctx context.Context) { t.MonitorUsdtAllowed(ctx) }, "MonitorUsdtAllowed")
	// 3. 归集状态查询 (每小时)
	_, _ = gcron.AddSingleton(ctx, "0 0 * * * *", func(ctx context.Context) { t.QueryCollectionStatus(ctx) }, "QueryCollectionStatus")

	// 4. 秒合约结算 (每秒跑一次)
	_, _ = gcron.AddSingleton(ctx, "* * * * * *", func(ctx context.Context) { t.SettleSecondContract(ctx) }, "SettleSecondContract")
	// 5. 永续合约订单与爆仓检查 (每 5 秒跑一次)
	_, _ = gcron.AddSingleton(ctx, "*/5 * * * * *", func(ctx context.Context) { t.SettleContractOrder(ctx) }, "SettleContractOrder")
	_, _ = gcron.AddSingleton(ctx, "*/5 * * * * *", func(ctx context.Context) { t.CheckContractPosition(ctx) }, "CheckContractPosition")
	// 6. 现货订单撮合 (每 30 秒)
	_, _ = gcron.AddSingleton(ctx, "*/30 * * * * *", func(ctx context.Context) { t.SettleCurrencyOrder(ctx) }, "SettleCurrencyOrder")

	// 7. Defi 收益派发 (每日 00:05)
	_, _ = gcron.AddSingleton(ctx, "0 5 0 * * *", func(ctx context.Context) { t.DistributeDefiRate(ctx) }, "DistributeDefiRate")

	// 8. 每日打卡状态更新 (每日 00:00)
	_, _ = gcron.AddSingleton(ctx, "0 0 0 * * *", func(ctx context.Context) { t.UpdateCodingDaily(ctx) }, "UpdateCodingDaily")

	// 9. 自发币大厅状态驱动 (每分钟)
	_, _ = gcron.AddSingleton(ctx, "0 * * * * *", func(ctx context.Context) { t.HandleOwnCoinStart(ctx) }, "HandleOwnCoinStart")
	_, _ = gcron.AddSingleton(ctx, "0 * * * * *", func(ctx context.Context) { t.HandleOwnCoinEnd(ctx) }, "HandleOwnCoinEnd")

	g.Log().Info(ctx, "Successfully registered all background cron tasks.")
}
