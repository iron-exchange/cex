package task

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"

	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

// CheckContractPosition U本位合约强平监控任务
func (s *sTask) CheckContractPosition(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: U本位合约强平及止盈止损监控")

	// 1. 查询全部持仓中的永续合约订单: TContractPosition `status = 0 (持仓中)`  // 注：Java代码中写的是 ContractOrderStatusEmun.DEAL.getCode() 这里假设0代表持仓
	var positions []*entity.ContractPosition
	err := dao.ContractPosition.Ctx(ctx).Where("status", 0).Scan(&positions)
	if err != nil {
		g.Log().Error(ctx, "查询持仓合约订单失败:", err)
		return err
	}

	for _, position := range positions {
		// 捕捉异常防止单笔错误阻塞全体
		func(p *entity.ContractPosition) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "处理仓位异常 id:%d, err:%v", p.Id, r)
				}
			}()

			symbol := p.Symbol
			redisKey := "CURRENCY_PRICE:" + symbol

			// 2. 从 PG sys_config (替代Redis) 获取对应币种最新市价
			redisV, _ := g.DB().GetValue(ctx, "SELECT config_value FROM sys_config WHERE config_key = ?", redisKey)
			if redisV.IsEmpty() {
				g.Log().Warningf(ctx, "获取不到 %s 的最新价格，跳过", symbol)
				return
			}
			currentlyPriceStr := strings.Trim(redisV.String(), "\"") // 防 JSON 序列化的多余双引号
			currentlyPrice, err := decimal.NewFromString(currentlyPriceStr)
			if err != nil || currentlyPrice.IsZero() {
				g.Log().Warningf(ctx, "价格解析异常 %s, val: %s", symbol, currentlyPriceStr)
				return
			}

			// 强平/止盈止损 闭环清算所需关键变量
			posType := p.Type // 0 买多 1卖空
			closePrice := decimal.NewFromFloat(p.ClosePrice)
			openPrice := decimal.NewFromFloat(p.OpenPrice)
			openNum := decimal.NewFromFloat(p.OpenNum)
			adjustAmount := decimal.NewFromFloat(p.AdjustAmount)
			leverage := decimal.NewFromFloat(p.Leverage)

			// 3. 获取后台合约风控配置 (TContractCoin)
			var tContractCoin *entity.ContractCoin
			_ = dao.ContractCoin.Ctx(ctx).Where("symbol", symbol).Scan(&tContractCoin)

			if tContractCoin != nil && tContractCoin.BaseCoin != "" /* 检查是否查到配置 */ {
				// 检查系统黑手干预标
				profitSignal := checkProfit(tContractCoin)
				if profitSignal == "ok" {
					// 走正常逻辑：若爆仓 (comparePrice) 则强平
					if comparePrice(posType, currentlyPrice, closePrice) {
						p.DealPrice, _ = currentlyPrice.Float64()
						p.SellFee = 0 // 爆仓免手续费，因为本金都已被吃没
						updatePositoinStop(ctx, p, currentlyPrice)
						return // 这笔仓位已经死了，无须走下面的止盈止损
					}
				} else {
					// 开启了外挂干预：只要浮动亏碎把保证金跌穿，不管现价有没有碰到 ClosePrice 都强平
					// 收益率 sub = (目前跌落幅度) / openPrice
					sub := getContractRate(openPrice, currentlyPrice, posType)
					if sub.IsNegative() {
						floatProfit := decimal.NewFromFloat(tContractCoin.FloatProfit)
						profitLoss := decimal.NewFromFloat(tContractCoin.ProfitLoss)
						if floatProfit.IsZero() {
							floatProfit = decimal.NewFromInt(1) // 兜底零除
						}

						// earn = sub * profitLoss * openNum * leverage / floatProfit
						earn := sub.Mul(profitLoss).Mul(openNum).Mul(leverage).Div(floatProfit)

						money := earn.Add(adjustAmount)
						if money.LessThanOrEqual(decimal.Zero) { // 惨遭收割
							p.Earn, _ = earn.Float64()
							p.SellFee = 0
							p.DealPrice, _ = currentlyPrice.Float64()
							updateProfitStop(ctx, p, currentlyPrice)
							return // 嗝屁
						}
					}
				}
			}

			// 4. 止盈止损(TP/SL) 订单簿联动扫描
			var lossRecords []*entity.ContractLoss
			_ = dao.ContractLoss.Ctx(ctx).Where("status", 0).Where("position_id", p.Id).Scan(&lossRecords)

			if len(lossRecords) > 0 {
				for _, loss := range lossRecords {
					earnPrice := decimal.NewFromFloat(loss.EarnPrice)
					losePrice := decimal.NewFromFloat(loss.LosePrice)
					earnDelegatePrice := decimal.NewFromFloat(loss.EarnDelegatePrice)
					loseDelegatePrice := decimal.NewFromFloat(loss.LoseDelegatePrice)

					if loss.DelegateType == 1 { // 市价委托的话以现价为准
						earnDelegatePrice = currentlyPrice
						loseDelegatePrice = currentlyPrice
					}

					var hitTrigger bool
					var triggerPrice decimal.Decimal

					closeFeeRate := decimal.Zero
					if tContractCoin != nil {
						closeFeeRate = decimal.NewFromFloat(tContractCoin.CloseFee)
					}

					if posType == 0 { // 做多
						if earnPrice.IsPositive() && currentlyPrice.GreaterThanOrEqual(earnPrice) {
							triggerPrice = earnDelegatePrice
							hitTrigger = true
						} else if losePrice.IsPositive() && currentlyPrice.LessThanOrEqual(losePrice) {
							triggerPrice = loseDelegatePrice
							hitTrigger = true
						}
					} else if posType == 1 { // 做空
						if earnPrice.IsPositive() && currentlyPrice.LessThanOrEqual(earnPrice) {
							triggerPrice = earnDelegatePrice
							hitTrigger = true
						} else if losePrice.IsPositive() && currentlyPrice.GreaterThanOrEqual(losePrice) {
							triggerPrice = loseDelegatePrice
							hitTrigger = true
						}
					}

					if hitTrigger {
						sellFee := adjustAmount.Mul(closeFeeRate)
						p.SellFee, _ = sellFee.Float64()
						p.DealPrice, _ = triggerPrice.Float64()
						updatePositoinStop(ctx, p, currentlyPrice) // 同样走强平函数
						break                                      // 一手交割完结束当前持仓检查
					}
				}
			}

		}(position)
	}

	return nil
}

// 检查系统的盈利干预标志
func checkProfit(t *entity.ContractCoin) string {
	res := "ok"
	if t.FloatProfit > 0 {
		res += "earn"
	}
	if t.ProfitLoss > 0 {
		res += "loss"
	}
	return res
}

// comparePrice 正常对比是否爆仓 (0买多, 1卖空)
func comparePrice(posType int, currentlyPrice, closePrice decimal.Decimal) bool {
	if posType == 0 { // 做多，跌到底爆仓
		if currentlyPrice.LessThanOrEqual(closePrice) {
			return true
		}
	} else { // 卖空，涨上天爆仓
		if closePrice.IsZero() {
			return false
		}
		if currentlyPrice.GreaterThanOrEqual(closePrice) {
			return true
		}
	}
	return false
}

// 计算目前合约价格的变动方向比率
func getContractRate(openPrice, currentlyPrice decimal.Decimal, posType int) decimal.Decimal {
	if posType == 0 {
		// 买多：现价 - 入场价
		return currentlyPrice.Sub(openPrice).Div(openPrice)
	}
	// 卖空：入场价 - 现价
	return openPrice.Sub(currentlyPrice).Div(openPrice)
}

// PositionEarn 收益公式
func PositionEarn(openPrice, openNum, dealPrice decimal.Decimal, posType int) decimal.Decimal {
	// Java 中的 ContractComputerUtil.getPositionEarn
	// 买多赚= (卖-买)*数量  卖空赚=(买-卖)*数量
	if posType == 0 {
		return dealPrice.Sub(openPrice).Mul(openNum)
	}
	return openPrice.Sub(dealPrice).Mul(openNum)
}

// updatePositoinStop 执行标准结算清仓 (撤毁/盈利下发等)
func updatePositoinStop(ctx context.Context, p *entity.ContractPosition, finalPrice decimal.Decimal) {
	adjustAmount := decimal.NewFromFloat(p.AdjustAmount)
	openNum := decimal.NewFromFloat(p.OpenNum)
	openPrice := decimal.NewFromFloat(p.OpenPrice)
	dealPrice := decimal.NewFromFloat(p.DealPrice) // 刚才刚设进去的平仓价
	sellFee := decimal.NewFromFloat(p.SellFee)

	// 计算真实盈亏
	earn := PositionEarn(openPrice, openNum, dealPrice, p.Type)

	// 更新持仓状态
	p.Status = 1 // 已结算
	p.DealTime = gtime.Now()
	p.DealNum = p.OpenNum
	dealVal := openNum.Mul(dealPrice)
	p.DealValue, _ = dealVal.Float64()

	money := adjustAmount.Add(earn)
	p.Earn, _ = earn.Float64()
	_, _ = dao.ContractPosition.Ctx(ctx).Where("id", p.Id).Update(p)

	// 撤销止盈止损单
	_, _ = dao.ContractLoss.Ctx(ctx).Where("position_id", p.Id).Update(g.Map{"status": 2})

	// 下发资产
	var asset entity.AppAsset
	// TYPE = CONTRACT_ASSETS (平台合约账户), AssetEnum.CONTRACT_ASSETS.getCode() -> 可能是 3
	_ = dao.AppAsset.Ctx(ctx).Where("user_id", p.UserId).Where("type", 3).Where("symbol", "usdt").Scan(&asset)

	amout := decimal.NewFromFloat(asset.Amout)
	add := amout.Add(money).Sub(sellFee)
	if add.IsNegative() {
		add = decimal.Zero
	}
	asset.Amout, _ = add.Float64()
	asset.AvailableAmount, _ = add.Float64()
	_, _ = dao.AppAsset.Ctx(ctx).Where("id", asset.Id).Update(asset)

	// 记录 WalletRecord: CONTRACT_TRANSACTION_CLOSING (枚举值为 62)
	// 原 Java Enum: RecordEnum.CONTRACT_TRANSACTION_CLOSING.getCode() -> 62
	record := entity.AppWalletRecord{
		UserId:       p.UserId,
		Amount:       money.InexactFloat64(),
		Type:         62, // CONTRACT_TRANSACTION_CLOSING
		SerialId:     p.OrderNo,
		Symbol:       "usdt",
		Remark:       "合约交易强平",
		BeforeAmount: amout.InexactFloat64(),
		AfterAmount:  add.InexactFloat64(),
		CreateTime:   gtime.Now(),
	}
	_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(record)

	// PUSH To Socket Stream
	g.Redis().Do(ctx, "XADD", "api-redis-stream.names", "*", "settlement", "2")
}

// updateProfitStop 后台干预收割清仓 (没收剩余)
func updateProfitStop(ctx context.Context, p *entity.ContractPosition, finalPrice decimal.Decimal) {
	openNum := decimal.NewFromFloat(p.OpenNum)
	dealPrice := decimal.NewFromFloat(p.DealPrice)

	p.Status = 1
	p.DealTime = gtime.Now()
	p.DealNum = p.OpenNum
	p.DealValue, _ = openNum.Mul(dealPrice).Float64()

	_, _ = dao.ContractPosition.Ctx(ctx).Where("id", p.Id).Update(p)
	_, _ = dao.ContractLoss.Ctx(ctx).Where("position_id", p.Id).Update(g.Map{"status": 2})

	// 没收本金不加钱, 只给空流水
	var asset entity.AppAsset
	_ = dao.AppAsset.Ctx(ctx).Where("user_id", p.UserId).Where("type", 3).Where("symbol", "usdt").Scan(&asset)

	record := entity.AppWalletRecord{
		UserId:       p.UserId,
		Amount:       0,
		Type:         62, // CONTRACT_TRANSACTION_CLOSING
		SerialId:     p.OrderNo,
		Symbol:       "usdt",
		Remark:       "合约交易干预强平",
		BeforeAmount: asset.Amout,
		AfterAmount:  asset.Amout,
		CreateTime:   gtime.Now(),
	}
	_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(record)

	g.Redis().Do(ctx, "XADD", "api-redis-stream.names", "*", "settlement", "2")
}
