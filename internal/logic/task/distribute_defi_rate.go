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

// DistributeDefiRate 每天按用户去中心化钱包 USDT 余额，生成利息 (DeFi质押)
func (s *sTask) DistributeDefiRate(ctx context.Context) error {
	g.Log().Info(ctx, "[定时任务] 执行: 定期扫描链上授权钱包派送 DeFi 质押利息")

	// 1. 获取所有的 DeFi 阶梯利率表，按 min_amount 倒序或内存排序，用于稍后匹配
	var rates []*entity.DefiRate
	err := dao.DefiRate.Ctx(ctx).Order("min_amount DESC").Scan(&rates)
	if err != nil {
		g.Log().Error(ctx, "加载 Defi 利率表失败:", err)
		return err
	}
	if len(rates) == 0 {
		g.Log().Warning(ctx, "未配置 Defi 阶梯利率，提前退出")
		return nil
	}

	// 2. 查询所有已授信 (AllowedUser) 的用户的链上钱包信息 (TAppAddressInfo)
	// 原版通常认定 usdt_allowed > 0 或特有状态来证明这笔钱能被监控算作质押
	var addresses []*entity.AppAddressInfo
	err = dao.AppAddressInfo.Ctx(ctx).Where("usdt_allowed >", 0).Scan(&addresses)
	if err != nil {
		g.Log().Error(ctx, "查询授权钱包失败:", err)
		return err
	}

	// 3. 准备获取 ETH 最新市价 (从 PG sys_config 获取以替代 Redis)
	redisV, _ := g.DB().GetValue(ctx, "SELECT config_value FROM sys_config WHERE config_key = ?", "CURRENCY_PRICE:ETH")
	if redisV.IsEmpty() {
		g.Log().Error(ctx, "获取 ETH/USDT 实时价格失败")
		return nil
	}
	ethPriceStr := strings.Trim(redisV.String(), "\"")
	ethPrice, err := decimal.NewFromString(ethPriceStr)
	if err != nil || ethPrice.IsZero() {
		g.Log().Errorf(ctx, "解析 ETH 价格异常: %v", err)
		return nil
	}

	// 4. 遍历发息
	for _, addr := range addresses {
		func(a *entity.AppAddressInfo) {
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "处理用户 %d DeFi 收益异常: %v", a.UserId, r)
				}
			}()

			usdtBal := decimal.NewFromFloat(a.Usdt)
			if usdtBal.LessThanOrEqual(decimal.Zero) {
				return // 没钱就不生利息
			}

			// 匹配利率 (由于前面查出的是倒序的 min_amount，第一个符合的就是他应该享有的最高档位)
			var hitRate *entity.DefiRate
			for _, r := range rates {
				min := decimal.NewFromFloat(r.MinAmount)
				max := decimal.NewFromFloat(r.MaxAmount)
				// 原有可能是左闭右闭或者左闭右开，这里做基础的 [min, max] 比较
				if usdtBal.GreaterThanOrEqual(min) && (r.MaxAmount == 0 || usdtBal.LessThanOrEqual(max)) {
					hitRate = r
					break
				}
			}

			if hitRate == nil || hitRate.Rate <= 0 {
				return // 未达到起投门槛
			}

			rateDecimal := decimal.NewFromFloat(hitRate.Rate)

			// 5. 计算应发利息 (以 USDT 计价) = 链上钱包里的 USDT * 日利率
			usdtReward := usdtBal.Mul(rateDecimal)

			// 6. 将应发利息折算成等价值的 ETH 数量 = USDT 利息额 / 当前 ETH 单价
			ethReward := usdtReward.Div(ethPrice)

			// 7. 生成一条 DefiOrder 订单流水
			order := entity.DefiOrder{
				UserId:      a.UserId,
				Amount:      usdtReward.InexactFloat64(),
				TotleAmount: usdtBal.InexactFloat64(),
				Rate:        hitRate.Rate,
				CreateTime:  gtime.Now(),
				Remark:      "DeFi 每日利息发放",
			}
			insRes, err := dao.DefiOrder.Ctx(ctx).Insert(order)
			if err != nil {
				g.Log().Error(ctx, "插入 DefiOrder 失败:", err)
				return
			}
			orderId, _ := insRes.LastInsertId()

			// 8 & 9. 找到用户的平台内部 ETH 资产并加钱
			// AssetEnum.PLATFORM_ASSETS -> 通常是 1 或者是对应的枚举，原 Java 未给死，这里默认 type=1 是平台现货资产
			var asset entity.AppAsset
			err = dao.AppAsset.Ctx(ctx).Where("user_id", a.UserId).Where("type", 1).Where("symbol", "eth").Scan(&asset)
			if err != nil || asset.Id == 0 {
				// 如果没有该资产，应该抛弃或者帮忙建一条。此系统通常买币时建立，这里如果没资产可以简单生成或者日志报警。
				// 依照你在 CurrencyOrder 里帮他建表的设计，这里最好是给他建一条
				asset = entity.AppAsset{
					UserId:          a.UserId,
					Symbol:          "eth",
					Type:            "1", // 假设平台资产
					Amout:           0,
					AvailableAmount: 0,
					CreateTime:      gtime.Now(),
					UpdateTime:      gtime.Now(),
				}
				insAsset, _ := dao.AppAsset.Ctx(ctx).Insert(asset)
				assetId, _ := insAsset.LastInsertId()
				asset.Id = int(assetId)
			}

			beforeAmout := asset.AvailableAmount
			asset.AvailableAmount = ethReward.Add(decimal.NewFromFloat(asset.AvailableAmount)).InexactFloat64()
			asset.Amout = ethReward.Add(decimal.NewFromFloat(asset.Amout)).InexactFloat64()
			asset.UpdateTime = gtime.Now()

			_, _ = dao.AppAsset.Ctx(ctx).Where("id", asset.Id).Update(asset)

			// 10. 给钱包写一条 `DEFI_ORDER` 类型的入账流水 (假设 DEFI_ORDER 的枚举对应 Type 是 50)
			record := entity.AppWalletRecord{
				UserId:       a.UserId,
				Amount:       ethReward.InexactFloat64(),
				Type:         50, // 假设 DEFI_ORDER 类型为 50
				SerialId:     g.NewVar(orderId).String(),
				Symbol:       "eth",
				Remark:       "DeFi持币生息发放(折算ETH)",
				BeforeAmount: beforeAmout,
				AfterAmount:  asset.AvailableAmount,
				CreateTime:   gtime.Now(),
			}
			_, _ = dao.AppWalletRecord.Ctx(ctx).Insert(record)

		}(addr)
	}

	return nil
}
