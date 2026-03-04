package task

import (
	"context"
	"strings"

	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// SettleContractOrder U本位限价/市价委托单入场结算任务 (对应原 ContractOrderTask)
func (s *sTask) SettleContractOrder(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: U本位委托订单入场检测")

	// 1. 获取所有等待成交的委托单 (status = 0)
	var orders []*entity.ContractOrder
	err := dao.ContractOrder.Ctx(ctx).Where("status", 0).Scan(&orders)
	if err != nil {
		g.Log().Error(ctx, "查询等待成交的委托单失败:", err)
		return err
	}

	for _, order := range orders {
		func(o *entity.ContractOrder) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "处理委托单异常 id:%d, err:%v", o.Id, r)
				}
			}()

			symbol := o.Symbol
			redisKey := "CURRENCY_PRICE:" + symbol

			// 2. 从 PG sys_config (替代Redis) 获取最新报价
			redisV, _ := g.DB().GetValue(ctx, "SELECT config_value FROM sys_config WHERE config_key = ?", redisKey)
			if redisV.IsEmpty() {
				return // 没价格暂时跳过
			}
			currentlyPriceStr := strings.Trim(redisV.String(), "\"")
			currentlyPrice, err := decimal.NewFromString(currentlyPriceStr)
			if err != nil || currentlyPrice.IsZero() {
				return
			}

			delegatePrice := decimal.NewFromFloat(o.DelegatePrice)
			hit := false

			// 3. 价格穿透判断
			// Type 0买多：现价 <= 委托价 触发买入
			// Type 1卖空：现价 >= 委托价 触发卖出
			if o.Type == 0 {
				if currentlyPrice.LessThanOrEqual(delegatePrice) {
					hit = true
				}
			} else if o.Type == 1 {
				if currentlyPrice.GreaterThanOrEqual(delegatePrice) {
					hit = true
				}
			}

			// 如果市价单(DelegateType==1)，理论上挂单瞬间引擎已处理。
			// 若由于极速并发落入定时器，也应视为 hit 且价格以市价为准，但在目前架构，我们先专注限价单逻辑
			if hit {
				// 获取风控配置
				var coinConfig *entity.ContractCoin
				_ = dao.ContractCoin.Ctx(ctx).Where("symbol", symbol).Scan(&coinConfig)

				// 启动事务保证单子入场原子性
				err := dao.ContractOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					// 加锁订单行防止重复处理
					var dbOrder *entity.ContractOrder
					if err := dao.ContractOrder.Ctx(ctx).Where("id", o.Id).Where("status", 0).LockUpdate().Scan(&dbOrder); err != nil || dbOrder == nil {
						return nil // 已经被吃单
					}

					// 计算保证金： DelegateValue / Leverage
					delValue := decimal.NewFromFloat(o.DelegateValue)
					leverage := decimal.NewFromFloat(o.Leverage)
					margin := delValue.Div(leverage)
					fee := decimal.NewFromFloat(o.Fee)

					// 计算强平价 (粗略计算公式，Java中通常为 OpenPrice*(1 ± 1/Leverage))
					closePrice := decimal.Zero
					if o.Type == 0 { // 买多：跌破 OpenPrice * (1 - 1/leverage) 强平
						closePrice = delegatePrice.Mul(decimal.NewFromInt(1).Sub(decimal.NewFromInt(1).Div(leverage)))
					} else { // 卖空：涨破 OpenPrice * (1 + 1/leverage) 强平
						closePrice = delegatePrice.Mul(decimal.NewFromInt(1).Add(decimal.NewFromInt(1).Div(leverage)))
					}

					// 创建持仓 ContractPosition
					pos := entity.ContractPosition{
						Type:         o.Type,
						DelegateType: o.DelegateType,
						Status:       0, // 0 持仓中
						Amount:       margin.InexactFloat64(),
						OpenNum:      o.DelegateTotal,
						OpenPrice:    o.DelegatePrice,
						ClosePrice:   closePrice.InexactFloat64(),
						OrderNo:      o.OrderNo,
						UserId:       o.UserId,
						OpenFee:      fee.InexactFloat64(),
						Leverage:     o.Leverage,
						Symbol:       o.Symbol,
						CreateTime:   gtime.Now(),
						AdjustAmount: margin.InexactFloat64(),
					}
					// 继承后台风控参数
					if coinConfig != nil {
						pos.MinMargin = coinConfig.MinMargin
						pos.LossRate = coinConfig.LossRate
						pos.EarnRate = coinConfig.EarnRate
					}

					_, err := dao.ContractPosition.Ctx(ctx).Insert(pos)
					if err != nil {
						return err
					}

					// 更新订单状态 -> 1 (完全成交)
					_, err = dao.ContractOrder.Ctx(ctx).Where("id", o.Id).Update(g.Map{
						"status":      1,
						"deal_num":    o.DelegateTotal,
						"deal_price":  o.DelegatePrice,
						"deal_value":  o.DelegateValue,
						"deal_time":   gtime.Now(),
						"update_time": gtime.Now(),
					})
					if err != nil {
						return err
					}

					// 扣除 AppAsset 冻结款
					// 当用户下限价单时，AvailableAmount 已经预先扣减了 (Margin + Fee)，但 Amout 没变 (这就是冻结的含义)
					// 当订单真实成交落入仓位时，这部分冻结的钱正式转化为了仓位和手续费
					// 此时我们需要从总资产 Amout 中剔除这批款项，以保证双重记账的严密性
					var asset entity.AppAsset
					err = dao.AppAsset.Ctx(ctx).Where("user_id", o.UserId).Where("type", 3).Where("symbol", o.BaseCoin).LockUpdate().Scan(&asset)
					if err == nil && asset.Id > 0 {
						toDeduct := margin.Add(fee)
						asset.Amout = decimal.NewFromFloat(asset.Amout).Sub(toDeduct).InexactFloat64()
						_, _ = dao.AppAsset.Ctx(ctx).Where("id", asset.Id).Update(g.Map{"amout": asset.Amout, "update_time": gtime.Now()})

						// 流水类型: 假设 61 代表 CONTRACT_TRANSACTION_OPENING
						record := entity.AppWalletRecord{
							UserId:       o.UserId,
							Amount:       toDeduct.InexactFloat64(),
							Type:         61,
							SerialId:     o.OrderNo,
							Symbol:       o.BaseCoin,
							Remark:       "合约限价单吃单入盘 (转换保证金)",
							BeforeAmount: asset.Amout + toDeduct.InexactFloat64(),
							AfterAmount:  asset.Amout,
							CreateTime:   gtime.Now(),
						}
						_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(record)
					}

					return nil
				})

				if err == nil {
					g.Log().Infof(ctx, "✅ 委托单入场成功: OrderID:%s, 触发价:%s", o.OrderNo, currentlyPriceStr)
					// 发送 settlement: 2 的消息刷新 K 线前端
					g.Redis().Do(ctx, "XADD", "api-redis-stream.names", "*", "settlement", "2")
				} else {
					g.Log().Errorf(ctx, "委托单 %s 入场事务失败: %v", o.OrderNo, err)
				}
			}
		}(order)
	}

	return nil
}
