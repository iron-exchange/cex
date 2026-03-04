package task

import (
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
	"context"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// QueryCollectionStatus 查询区块链上的资金归集进度 (CollectionTask)
func (s *sTask) QueryCollectionStatus(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: (并发优化版) 轮询区块链 ETH/TRX 资金归集状态")

	// 1. 从归集订单表 TCollectionOrder 中查出 `status=1` (进行中) 的未完成链上交易。
	var orders []*entity.CollectionOrder
	err := dao.CollectionOrder.Ctx(ctx).Where("status", "1").Scan(&orders)
	if err != nil {
		g.Log().Error(ctx, "查询进行中的归集订单失败:", err)
		return err
	}
	if len(orders) == 0 {
		return nil
	}

	// [防御强化]: 设定容忍极限，比如 3 天 (72小时)
	thresholdDate := gtime.Now().Add(-72 * time.Hour)

	// 并发控制: 假设最多同时存在 20 个请求
	poolSize := 20
	sem := make(chan struct{}, poolSize)
	var wg sync.WaitGroup

	for _, order := range orders {
		// 如果挂起超过三天，直接报警并置为失败防止死循环堆积
		if order.CreateTime != nil && order.CreateTime.Before(thresholdDate) {
			_, _ = dao.CollectionOrder.Ctx(ctx).Where("id", order.Id).Update(g.Map{"status": "3", "update_time": gtime.Now()})
			g.Log().Warningf(ctx, "🚨 [节点卡死预警] 归集交易 %s 挂起超过3天，已被强制截停", order.Hash)
			// TODO: 可以推送到刚才的 CEX:STREAM:SECURITY_ALERTS 流
			continue
		}

		wg.Add(1)
		sem <- struct{}{} // 占用一个并发槽

		go func(o *entity.CollectionOrder) {
			defer wg.Done()
			defer func() { <-sem }() // 释放槽位
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "轮询 Hash %s 状态崩溃: %v", o.Hash, r)
				}
			}()

			// 创建带超时控制的请求上下文，防止单一恶劣节点拖死整个 Go 程
			reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			var isSuccess bool
			var isFailed bool

			// 3. 按公链 (`Chain`) 调用对应的节点 SDK
			// 目前此微服务尚未引入底层 EthUtils / TronUtils 的 Go 重写版
			// 此处结构为预留对接层，实际业务应调用封装的 RPC 函数
			if strings.EqualFold(o.Chain, "eth") || strings.EqualFold(o.Chain, "erc20") {
				// isSuccess, isFailed = ethUtils.CheckTransaction(reqCtx, o.Hash)
				// 模拟
				g.Log().Debugf(reqCtx, "查询 ETH Hash: %s", o.Hash)
			} else if strings.EqualFold(o.Chain, "trx") || strings.EqualFold(o.Chain, "trc20") {
				// isSuccess, isFailed = tronUtils.CheckTransaction(reqCtx, o.Hash)
				g.Log().Debugf(reqCtx, "查询 TRX Hash: %s", o.Hash)
			} else {
				g.Log().Errorf(reqCtx, "未知公链类型: %s, 无法查询 Hash: %s", o.Chain, o.Hash)
				return
			}

			// 4. 更新订单状态 (模拟逻辑)
			if isSuccess {
				_, _ = dao.CollectionOrder.Ctx(ctx).Where("id", o.Id).Update(g.Map{"status": "2", "update_time": gtime.Now()})
				g.Log().Infof(ctx, "✅ 资金归集成功: %s", o.Hash)
			} else if isFailed {
				_, _ = dao.CollectionOrder.Ctx(ctx).Where("id", o.Id).Update(g.Map{"status": "3", "update_time": gtime.Now()})
				g.Log().Warningf(ctx, "❌ 资金归集在链上遭遇 Revert 或 OutOfGas: %s", o.Hash)

				// [联动修复]: 归集失败，应当通过 Redis 通知热钱包管理员充当 Gas / 解冻
				alertMsg := g.Map{
					"eventType": "COLLECTION_FAILED",
					"orderId":   o.OrderId,
					"hash":      o.Hash,
					"chain":     o.Chain,
					"amount":    o.Amount,
				}
				_, _ = g.Redis().Do(ctx, "XADD", "CEX:STREAM:SECURITY_ALERTS", "*", "payload", alertMsg)
			}
			// 若既没成功也没失败 (依然 Pending)，则不修改状态，等下次定时器扫描

		}(order)
	}

	wg.Wait() // 等待本次并发池内所有 Hash 轮询完毕
	return nil
}
