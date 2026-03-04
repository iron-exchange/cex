package task

import (
	"GoCEX/internal/dao"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpdateCodingDaily 打码量每日归集计算任务
func (s *sTask) UpdateCodingDaily(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: 打码量每日盘点")

	// 1. 获取基于 PG 内存的分布式锁，防止多台机器跑重 (标识 1001)
	val, err := g.DB().GetValue(ctx, "SELECT pg_try_advisory_lock(1001)")
	lockOk := val.Bool()
	if err != nil || !lockOk {
		g.Log().Warning(ctx, "未获取到定时任务锁，跳过本次 UpdateCodingDaily 执行")
		return nil
	}
	// 执行完释放锁
	defer g.DB().Exec(ctx, "SELECT pg_advisory_unlock(1001)")

	// 2. 批量将全表 AppAsset 的 "今日打码量" 清零，并将当前的 "可用余额" 快照拷贝到 "今日可用余额基数"
	// 对应 SQL 原理: UPDATE t_app_asset SET coding_volume_daily = 0, available_amount_daily = available_amount
	result, err := dao.AppAsset.Ctx(ctx).
		Data(g.Map{
			dao.AppAsset.Columns().CodingVolumeDaily:    0,
			dao.AppAsset.Columns().AvailableAmountDaily: gdb.Raw(dao.AppAsset.Columns().AvailableAmount),
			dao.AppAsset.Columns().UpdateTime:           gtime.Now(),
		}).
		Where("1=1").
		Update()

	if err != nil {
		g.Log().Error(ctx, "重置全站用户打码量失败:", err)
		return err
	}

	affected, _ := result.RowsAffected()
	g.Log().Infof(ctx, "✅ 成功重置且盘点打码量, 影响资产数: %d 条", affected)

	return nil
}
