package owncoin

import (
	"context"
	"fmt"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// GetOwnCoinDetail 获取自发币详情
func (s *sAppOwnCoin) GetOwnCoinDetail(ctx context.Context, userId int64, req *v1.GetOwnCoinDetailReq) (*v1.GetOwnCoinDetailRes, error) {
	var info entity.OwnCoin
	err := dao.OwnCoin.Ctx(ctx).Where("id", req.OwnId).Scan(&info)
	if err != nil || info.Id == 0 {
		return nil, gerror.New("当前发行币不存在")
	}

	// userId 已作为参数传入

	// 统计该用户当前已申购的总额
	purchasedAmt := 0.0
	if userId > 0 {
		var sum struct {
			SumAmt float64 `orm:"sumAmt"`
		}
		_ = dao.OwnCoinOrder.Ctx(ctx).
			Fields("COALESCE(SUM(amount), 0) sumAmt").
			Where("user_id", userId).
			Where("own_id", req.OwnId).
			Scan(&sum)
		purchasedAmt = sum.SumAmt
	}

	appInfo := v1.OwnCoinAppInfo{
		Id:            info.Id,
		Coin:          info.Coin,
		Logo:          info.Logo,
		ReferCoin:     info.ReferCoin,
		ShowSymbol:    info.ShowSymbol,
		Price:         info.Price,
		Proportion:    info.Proportion,
		RaisingAmount: info.RaisingAmount,
		RaisedAmount:  info.RaisedAmount,
		PurchaseLimit: info.PurchaseLimit,
		TotalAmount:   info.TotalAmount,
		Status:        info.Status,
		BeginTime:     info.BeginTime.String(),
		EndTime:       info.EndTime.String(),
		Introduce:     info.Introduce,
	}

	res := &v1.GetOwnCoinDetailRes{}
	res.Data.CoinInfo = appInfo
	res.Data.PurchasedAmt = purchasedAmt
	res.Data.RemainingAmt = info.TotalAmount - info.RaisedAmount
	if res.Data.RemainingAmt < 0 {
		res.Data.RemainingAmt = 0
	}
	res.Data.UserRemaining = info.PurchaseLimit // 假定 PurchaseLimit 表示最多申购几次? 还是数量? 这里简单塞回去

	return res, nil
}

// SubscribeOwnCoin 用户认购自发币
func (s *sAppOwnCoin) SubscribeOwnCoin(ctx context.Context, userId int64, req *v1.SubscribeOwnCoinReq) (*v1.SubscribeOwnCoinRes, error) {
	if userId == 0 {
		return nil, gerror.New("请先登录")
	}

	if req.PayAmount <= 0 {
		return nil, gerror.New("支付金额必须大于0")
	}

	return nil, dao.OwnCoin.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 锁表查币种状态
		var info entity.OwnCoin
		err := dao.OwnCoin.Ctx(ctx).Where("id", req.OwnId).LockUpdate().Scan(&info)
		if err != nil || info.Id == 0 {
			return gerror.New("非法的币种ID")
		}

		if info.Status != 2 {
			return gerror.New("当前不在筹集中(申购)状态")
		}
		if info.ReferCoin != req.PayCoin {
			return gerror.Newf("支付币种不符，仅支持 %s", info.ReferCoin)
		}

		// 2. 扣减用户参考币余额 (例如 USDT)
		// Assuming you have an Asset table
		var userAsset entity.AppAsset
		err = dao.AppAsset.Ctx(ctx).Where("user_id", userId).Where("currency_name", req.PayCoin).LockUpdate().Scan(&userAsset)
		if err != nil || userAsset.AvailableAmount < req.PayAmount {
			return gerror.New("可用余额不足")
		}

		// 扣减余额
		_, err = dao.AppAsset.Ctx(ctx).Where("id", userAsset.Id).Decrement("available_amount", req.PayAmount)
		if err != nil {
			return err
		}

		// 计算获得的数量
		number := int(req.PayAmount / info.Price)

		// 4. 更新已募集总数 (金额)
		_, err = dao.OwnCoin.Ctx(ctx).Where("id", info.Id).Increment("raised_amount", req.PayAmount)
		if err != nil {
			return err
		}

		// 5. 插入订单表
		_, err = dao.OwnCoinOrder.Ctx(ctx).Data(entity.OwnCoinOrder{
			OwnId:   info.Id,
			OwnCoin: info.Coin,
			UserId:  userId,
			Amount:  req.PayAmount, // 消耗的参考币总额(PayAmount)
			Number:  number,        // 买到的新币数量
			Price:   info.Price,
			Status:  "1", // 1=待发放 (发红后变为2, 或者失败变为3等)
			OrderId: fmt.Sprintf("IEOX%dX%d", userId, number),
		}).Insert()

		return err
	})
}
