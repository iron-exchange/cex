package task

import (
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HandleOwnCoinStart 更改发币状态 (每天00:01执行一次)
func (s *sTask) HandleOwnCoinStart(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: 平台币发行开始状态流转")

	// 1. 查询表 T_OWN_COIN 中 begin_time <= 当前时间且 status == 1 的平台币发行项目
	var coins []*entity.OwnCoin
	now := gtime.Now()
	err := dao.OwnCoin.Ctx(ctx).Where("status", 1).Where("begin_time <=", now).Scan(&coins)
	if err != nil {
		g.Log().Error(ctx, "查询待发行平台币失败:", err)
		return err
	}

	for _, coin := range coins {
		// 2. 将这些项目的 status 更新为 2 (正在发行中)
		_, err = dao.OwnCoin.Ctx(ctx).Where("id", coin.Id).Update(g.Map{"status": 2, "update_time": now})
		if err != nil {
			g.Log().Errorf(ctx, "更新平台币 %s 发行状态失败: %v", coin.Coin, err)
		} else {
			g.Log().Infof(ctx, "平台币 %s (ID:%d) 已进入发行认购期", coin.Coin, coin.Id)
		}
	}

	return nil
}

// HandleOwnCoinEnd 发币结束 申购资产发送 (每天00:01执行一次)
func (s *sTask) HandleOwnCoinEnd(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: 平台币发币结束与资产下发")

	// 1. 查询表 T_OWN_COIN 中 end_time <= 当前时间且 status == 2 的项目
	var endedCoins []*entity.OwnCoin
	now := gtime.Now()
	err := dao.OwnCoin.Ctx(ctx).Where("status", 2).Where("end_time <=", now).Scan(&endedCoins)
	if err != nil {
		g.Log().Error(ctx, "查询结束发行的平台币失败:", err)
		return err
	}

	for _, coin := range endedCoins {
		func(c *entity.OwnCoin) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "处理平台币 %s 结束结算异常: %v", c.Coin, r)
				}
			}()

			// 2. 将这些项目 status 更新为 3 (已结束)
			_, err = dao.OwnCoin.Ctx(ctx).Where("id", c.Id).Update(g.Map{"status": 3, "update_time": now})
			if err != nil {
				g.Log().Error(ctx, "更新平台币失败:", err)
				return
			}
			g.Log().Infof(ctx, "平台币 %s 众筹期满结束，开启空投结算...", c.Coin)

			// 3. 遍历查询对应的申购订单 T_OWN_COIN_ORDER (status == "1") 注意这里是字符串
			var orders []*entity.OwnCoinOrder
			_ = dao.OwnCoinOrder.Ctx(ctx).Where("status", "1").Where("own_id", c.Id).Scan(&orders)

			for _, order := range orders {
				// 4. 将订单状态更新为 2 (已结算)
				_, _ = dao.OwnCoinOrder.Ctx(ctx).Where("id", order.Id).Update(g.Map{"status": "2", "update_time": now})

				// 5. 检查并为用户创建该申购币种的 T_APP_ASSET 钱包记录
				var asset entity.AppAsset
				// Type=1 现货户
				_ = dao.AppAsset.Ctx(ctx).Where("user_id", order.UserId).Where("type", 1).Where("symbol", c.Coin).Scan(&asset)

				if asset.Id == 0 {
					asset = entity.AppAsset{
						UserId:          order.UserId,
						Symbol:          c.Coin,
						Type:            "1",
						Amout:           0,
						AvailableAmount: 0,
						CreateTime:      now,
						UpdateTime:      now,
					}
					insAs, _ := dao.AppAsset.Ctx(ctx).Insert(asset)
					assetId, _ := insAs.LastInsertId()
					asset.Id = int(assetId)
				}

				// 6. 将用户买到的数量加入 T_APP_ASSET
				beforeAmout := asset.AvailableAmount
				addNum := float64(order.Number)
				asset.AvailableAmount += addNum
				asset.Amout += addNum
				asset.UpdateTime = now
				_, _ = dao.AppAsset.Ctx(ctx).Where("id", asset.Id).Update(asset)

				// 写入流水 OWN_COIN_BUY -> 假设对应枚举编号是 40
				record := entity.AppWalletRecord{
					UserId:       order.UserId,
					Amount:       addNum,
					Type:         40, // OWN_COIN_BUY
					SerialId:     order.OrderId,
					Symbol:       c.Coin,
					Remark:       "IEO平台币认购发放到账",
					BeforeAmount: beforeAmout,
					AfterAmount:  asset.AvailableAmount,
					CreateTime:   now,
				}
				_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(record)
			}

			// 7. 发币后自动上币: 将该币种写入 KLINE_SYMBOL 列表
			cnt, _ := dao.KlineSymbol.Ctx(ctx).Where("symbol", c.Coin).Count()
			if cnt == 0 {
				kline := entity.KlineSymbol{
					Market:      c.ReferMarket,
					Symbol:      c.Coin,
					Slug:        c.ShowSymbol,
					Status:      1,
					Logo:        c.Logo,
					ReferMarket: c.ReferMarket,
					ReferCoin:   c.ReferCoin,
					Proportion:  c.Proportion,
					CreateTime:  now,
				}
				_, _ = dao.KlineSymbol.Ctx(ctx).Insert(kline)
			}

			// 8. 自动上秒合约 (SecondCoinConfig，对应原系统的 T_SECOND_COIN_CONFIG)
			secCnt, _ := dao.SecondCoinConfig.Ctx(ctx).Where("symbol", c.Coin).Count()
			if secCnt == 0 {
				secConf := entity.SecondCoinConfig{
					Symbol:     c.Coin,
					Status:     1,
					Market:     c.ReferMarket,
					ShowSymbol: c.ShowSymbol,
					Coin:       c.Coin,
					Type:       2, // 2 虚拟币
					CreateTime: now,
					Logo:       c.Logo,
				}
				_, _ = dao.SecondCoinConfig.Ctx(ctx).Insert(secConf)
			}

		}(coin)
	}

	return nil
}
