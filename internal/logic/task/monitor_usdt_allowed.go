package task

import (
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MonitorUsdtAllowed 恶意链上授权监控秒U任务 (防秒U)
func (s *sTask) MonitorUsdtAllowed(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: 监控 Approve 链上授权及资金异动预警")

	// 1. 查询表 T_APP_ADDRESS_INFO 中符合预警阈值的链上授权钱包
	// 根据实体结构，触发条件为：
	// - 已经被标记授权 (usdt_allowed > 0)
	// - 当前钱包 U 余额大于设定的监控阈值 (usdt >= usdt_monitor) 且 usdt_monitor > 0
	// - 还未播报过 (allowed_notice == 0)
	var targets []*entity.AppAddressInfo
	err := dao.AppAddressInfo.Ctx(ctx).
		Where("usdt_allowed >", 0).
		Where("usdt_monitor >", 0).
		Where("allowed_notice", 0).
		Where("usdt >= usdt_monitor"). // 直接原生 SQL 对比列名，避免 ORM 字面量转义
		Scan(&targets)

	if err != nil {
		g.Log().Error(ctx, "查询授权预警钱包失败:", err)
		return err
	}

	for _, addr := range targets {
		// 这里由于 GoFrame ORM 的字面量比较限制，上面的 Where 可能会退化，我们在代码里再做一次二次确认防御
		if addr.Usdt >= addr.UsdtMonitor {
			// 2. 触发预警广播
			// 原版有对接 TG 机器人，咱们重构版本采用通用策略：将大额预警事件压入 Redis Stream。
			// 外部独立运行的报警机器人 (Bot / SMS 服务) 去侦听这个 Stream 实施全平台播报。
			alertMsg := g.Map{
				"eventType":   "USDT_ALLOWED_ALERT",
				"userId":      addr.UserId,
				"address":     addr.Address,
				"chain":       addr.WalletType,
				"currentUsdt": addr.Usdt,
				"monitorLine": addr.UsdtMonitor,
			}
			_, sysErr := g.Redis().Do(ctx, "XADD", "CEX:STREAM:SECURITY_ALERTS", "*", "payload", alertMsg)

			if sysErr == nil {
				// 3. 更新播报状态，防止下次重复轰炸
				_, _ = dao.AppAddressInfo.Ctx(ctx).Where("user_id", addr.UserId).Where("address", addr.Address).
					Update(g.Map{"allowed_notice": 1, "update_time": gtime.Now()})
				g.Log().Noticef(ctx, "🚨 [高危拦截] 发现授权地址资金达标: 用户ID:%d, 地址:%s, 余额:%f 上报预警平台成功", addr.UserId, addr.Address, addr.Usdt)
			} else {
				g.Log().Errorf(ctx, "发送预警事件至 Redis 失败: %v", sysErr)
			}
		}
	}

	return nil
}
