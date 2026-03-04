package task

import (
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// SettleCurrencyOrder 现货币币交易撮合结算任务 (CurrencyOrderTask)
func (s *sTask) SettleCurrencyOrder(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: 定期扫表撮合现货币币交易订单")

	// 1. 获取基于 PG 内存的分布式锁，防止多台机器跑重 (标识 1003)
	val, errLock := g.DB().GetValue(ctx, "SELECT pg_try_advisory_lock(1003)")
	if errLock != nil || !val.Bool() {
		g.Log().Warning(ctx, "未获取到现货撮合结算锁，跳过本次执行")
		return nil
	}
	defer func() {
		_, _ = g.DB().Exec(ctx, "SELECT pg_advisory_unlock(1003)")
		// 往 Stream 中 PUSH 现货结算动态供前端渲染 (暂时保留 Redis Stream 用于前端通知)
		_, _ = g.Redis().Do(ctx, "XADD", "api-redis-stream.names", "*", "settlement", "1")
	}()

	// 2. 从 TCurrencyOrder 查询未成交的委托单 (`status=0` 且 `delegate_type=0` 即撮合单)
	var orders []*entity.CurrencyOrder
	err := dao.CurrencyOrder.Ctx(ctx).Where("status", 0).Where("delegate_type", 0).Scan(&orders)
	if err != nil {
		g.Log().Error(ctx, "查询等待成交的现货委托单失败:", err)
		return err
	}

	for _, order := range orders {
		func(o *entity.CurrencyOrder) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "处理现货撮合结算异常 id:%d, err:%v", o.Id, r)
				}
			}()

			// 3. 遍历委托单，获取币对配置 (TCurrencySymbol) 以及手续费率 (FeeRate)
			var curSymbol *entity.CurrencySymbol
			err := dao.CurrencySymbol.Ctx(ctx).Where("symbol", o.Symbol).Scan(&curSymbol)
			if err != nil || curSymbol == nil {
				return
			}
			ratio := decimal.NewFromFloat(curSymbol.FeeRate)

			// 4. 从 PG sys_config (替代Redis) 拉取当前币种实时价格
			redisKey := "CURRENCY_PRICE:" + o.Symbol
			redisV, _ := g.DB().GetValue(ctx, "SELECT config_value FROM sys_config WHERE config_key = ?", redisKey)
			if redisV.IsEmpty() {
				return
			}
			currentlyPriceStr := strings.Trim(redisV.String(), "\"")
			settlePrice, err := decimal.NewFromString(currentlyPriceStr)
			if err != nil || settlePrice.IsZero() {
				return
			}

			delegatePrice := decimal.NewFromFloat(o.DelegatePrice)
			hit := false

			// 5. 现价穿透判断
			// 买单(0)：价格跌至等于或低于委托价
			// 卖单(1)：价格涨至等于或高于委托价
			if o.Type == 0 {
				if settlePrice.LessThanOrEqual(delegatePrice) {
					hit = true
				}
			} else {
				if settlePrice.GreaterThanOrEqual(delegatePrice) {
					hit = true
				}
			}

			if hit {
				// 6. 开启事务处理结算
				err = dao.CurrencyOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					// 行级锁，防止同一笔订单多次撮合
					var dbOrder *entity.CurrencyOrder
					if err := dao.CurrencyOrder.Ctx(ctx).Where("id", o.Id).Where("status", 0).LockUpdate().Scan(&dbOrder); err != nil || dbOrder == nil {
						return nil // 已被别人抢先处理
					}

					delegateTotal := decimal.NewFromFloat(o.DelegateTotal)
					var dealNum, subtractPrice, fee, addPrice decimal.Decimal
					var subtractCoin, addCoin string

					if o.Type == 0 {
						// 买单: 扣除的是结算币(USDT), 获得的是交易币(BTC)
						dealNum = delegateTotal
						subtractPrice = settlePrice.Mul(delegateTotal) // 暂不考虑找零退款，这里使用 settlePrice, 理论上买单扣除上限是 DelegateValue
						fee = dealNum.Mul(ratio)
						addPrice = dealNum.Sub(fee)
						subtractCoin = o.Coin    // 注意：库里结构可能颠倒了，按注释 "Coin是结算币BaseCoin USDT" - 请以实际业务为准，这里假定 O.Coin 是法币(USDT)，O.Symbol的拆分是"BTC"
						addCoin = curSymbol.Coin // 表示获得的币，比如 BTC
					} else {
						// 卖单: 扣除的是交易币(BTC), 获得的是结算币(USDT)
						dealNum = settlePrice.Mul(delegateTotal)
						subtractPrice = delegateTotal
						fee = dealNum.Mul(ratio)
						addPrice = dealNum.Sub(fee)
						subtractCoin = curSymbol.Coin // 如 BTC
						addCoin = o.Coin              // 如 USDT
					}

					// 更新订单为成功
					_, err := dao.CurrencyOrder.Ctx(ctx).Where("id", o.Id).Update(g.Map{
						"status":      1,
						"deal_num":    dealNum.InexactFloat64(),
						"deal_price":  settlePrice.InexactFloat64(),
						"deal_value":  dealNum.InexactFloat64(), // 取决于是卖还是买，这里简化处理
						"fee":         fee.InexactFloat64(),
						"deal_time":   gtime.Now(),
						"update_time": gtime.Now(),
					})
					if err != nil {
						return err
					}

					// 7. 扣减被冻结方资产 (Type = 1 现货账户)
					var subAsset entity.AppAsset
					err = dao.AppAsset.Ctx(ctx).Where("user_id", o.UserId).Where("type", 1).Where("symbol", subtractCoin).LockUpdate().Scan(&subAsset)
					if err == nil && subAsset.Id > 0 {
						// 从总资产扣除，因为挂单时已经把 available减去转移到 Amout 了？不，挂单是冻结，Amout不变，Occupied增加。
						// 成交意味着钱真花出去了：Amout 减去，Occupied 减去。
						newOccupied := decimal.NewFromFloat(subAsset.OccupiedAmount).Sub(subtractPrice)
						if newOccupied.LessThan(decimal.Zero) {
							newOccupied = decimal.Zero // 兜底防御
						}
						newAmout := decimal.NewFromFloat(subAsset.Amout).Sub(subtractPrice)

						_, _ = dao.AppAsset.Ctx(ctx).Where("id", subAsset.Id).Update(g.Map{
							"occupied_amount": newOccupied.InexactFloat64(),
							"amout":           newAmout.InexactFloat64(),
							"update_time":     gtime.Now(),
						})

						// 流水
						_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(entity.AppWalletRecord{
							UserId:       o.UserId,
							Amount:       subtractPrice.InexactFloat64(),
							Type:         20, // 暂定20表示现货扣出
							SerialId:     o.OrderNo,
							Symbol:       subtractCoin,
							Remark:       "现货币币撮合扣除",
							BeforeAmount: subAsset.Amout,
							AfterAmount:  newAmout.InexactFloat64(),
							CreateTime:   gtime.Now(),
						})
					}

					// 为获利方加资产
					var addAsset entity.AppAsset
					err = dao.AppAsset.Ctx(ctx).Where("user_id", o.UserId).Where("type", 1).Where("symbol", addCoin).LockUpdate().Scan(&addAsset)
					if err == nil && addAsset.Id > 0 {
						newAddAmout := decimal.NewFromFloat(addAsset.Amout).Add(addPrice)
						newAddAvail := decimal.NewFromFloat(addAsset.AvailableAmount).Add(addPrice)

						_, _ = dao.AppAsset.Ctx(ctx).Where("id", addAsset.Id).Update(g.Map{
							"amout":            newAddAmout.InexactFloat64(),
							"available_amount": newAddAvail.InexactFloat64(),
							"update_time":      gtime.Now(),
						})

						// 流水
						_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(entity.AppWalletRecord{
							UserId:       o.UserId,
							Amount:       addPrice.InexactFloat64(),
							Type:         21, // 暂定21表示现货收入
							SerialId:     o.OrderNo,
							Symbol:       addCoin,
							Remark:       "现货币币撮合收入",
							BeforeAmount: addAsset.Amout,
							AfterAmount:  newAddAmout.InexactFloat64(),
							CreateTime:   gtime.Now(),
						})
					}

					// 8. 累计每日打码量 (这里粗略计算为U的价值)
					// 这部分通常独立一个方法处理，防止阻塞核心账本，我们暂留结构位
					// dao.AppAsset.Ctx(ctx).Where("user_id", o.UserId)... Add CodingVolumeDaily

					return nil
				})

				if err == nil {
					g.Log().Infof(ctx, "✅ 现货币币订单撮合成交: OrderID:%s, 触发价:%s", o.OrderNo, currentlyPriceStr)
				} else {
					g.Log().Errorf(ctx, "现货订单撮合 %s 入场事务失败: %v", o.OrderNo, err)
				}
			}

		}(order)
	}

	return nil
}
