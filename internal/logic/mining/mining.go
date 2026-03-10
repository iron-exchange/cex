package mining

import (
	"context"
	"fmt"
	"time"

	v1a "GoCEX/api/admin/v1"
	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

type sMining struct{}

func New() *sMining {
	return &sMining{}
}

// GetFinancialIncome 获取理财汇总数据
func (s *sMining) GetFinancialIncome(ctx context.Context, userId uint64) (res *v1.PersonalIncomeRes, err error) {
	var orders []entity.MineOrder
	err = dao.MineOrder.Ctx(ctx).Where("user_id", userId).Where("status", 0).Scan(&orders)
	if err != nil {
		return nil, err
	}

	var totalIncome, todayIncome float64
	for _, o := range orders {
		totalIncome += o.AccumulaEarn
		// 估算今日收益 (简单逻辑: 金额 * 最小利率 / 100)
		todayIncome += (o.Amount * o.MinOdds) / 100
	}

	res = &v1.PersonalIncomeRes{
		TotalIncome: totalIncome,
		TodayIncome: todayIncome,
		OrderCount:  len(orders),
	}
	return
}

// BuyFinancial 购买理财产品
func (s *sMining) BuyFinancial(ctx context.Context, userId uint64, req *v1.FinancialSubmitReq) (err error) {
	// 1. 查询产品信息
	var product entity.MineFinancial
	err = dao.MineFinancial.Ctx(ctx).Where("id", req.PlanId).Scan(&product)
	if err != nil || product.Id == 0 {
		return gerror.NewCode(codes.Failed, "理财产品不存在")
	}

	// 2. 校验金额
	if req.Money < product.LimitMin || (product.LimitMax > 0 && req.Money > product.LimitMax) {
		return gerror.NewCode(codes.Failed, "申购金额不在允许范围内")
	}

	// 3. 扣减资产 (通过 Asset 模块的 SubAmount)
	_, err = asset.New().SubAmount(ctx, &v1a.SubAmountReq{
		UserId:     gconv.Int64(userId),
		Symbol:     product.Coin,
		Amount:     -req.Money,
		RecordType: 30, // 假设 30 为理财扣款
		Remark:     fmt.Sprintf("购买理财产品: %s", product.Title),
	}, func(ctx context.Context, tx gdb.TX) error {
		// 4. 事务回调: 插入订单
		orderNo := uuid.New().String()
		now := gtime.Now()
		endTime := now.Add(time.Duration(req.Days) * 24 * time.Hour)

		_, err := dao.MineOrder.Ctx(ctx).TX(tx).Data(entity.MineOrder{
			OrderNo:      orderNo,
			UserId:       gconv.Int64(userId),
			PlanId:       gconv.Int64(product.Id),
			PlanTitle:    product.Title,
			Amount:       req.Money,
			OrderAmount:  req.Money,
			Days:         req.Days,
			MinOdds:      product.MinOdds,
			MaxOdds:      product.MaxOdds,
			DefaultOdds:  product.DefaultOdds,
			Coin:         product.Coin,
			Status:       0,
			CreateTime:   now,
			EndTime:      gtime.New(endTime),
			AccumulaEarn: 0,
		}).Insert()
		return err
	})

	return
}

// ShowMining 展示矿机产品与用户订单
func (s *sMining) ShowMining(ctx context.Context, userId uint64) (res *v1.MiningShowRes, err error) {
	res = &v1.MiningShowRes{}

	// 1. 获取有效产品
	var products []entity.MingProduct
	_ = dao.MingProduct.Ctx(ctx).Where("status", 1).Order("sort asc").Scan(&products)
	for _, p := range products {
		res.Products = append(res.Products, v1.MiningProductInfo{
			Id:       gconv.Int64(p.Id),
			Title:    p.Title,
			MinPrice: p.LimitMin,
			MaxPrice: p.LimitMax,
			MinOdds:  p.MinOdds,
			MaxOdds:  p.MaxOdds,
		})
	}

	// 2. 获取用户订单
	var orders []entity.MingOrder
	_ = dao.MingOrder.Ctx(ctx).Where("user_id", userId).Order("create_time desc").Scan(&orders)
	for _, o := range orders {
		res.Orders = append(res.Orders, v1.MiningOrderInfo{
			OrderNo:    o.OrderNo,
			Amount:     o.Amount,
			Days:       o.Days,
			Status:     o.Status,
			CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

// BuyMining 购买矿机
func (s *sMining) BuyMining(ctx context.Context, userId uint64, req *v1.MiningSubmitReq) (err error) {
	// 1. 查询产品信息
	var product entity.MingProduct
	err = dao.MingProduct.Ctx(ctx).Where("id", req.PlanId).Scan(&product)
	if err != nil || product.Id == 0 {
		return gerror.NewCode(codes.Failed, "矿机产品不存在")
	}

	// 2. 扣减资产
	_, err = asset.New().SubAmount(ctx, &v1a.SubAmountReq{
		UserId:     gconv.Int64(userId),
		Symbol:     product.Coin,
		Amount:     -req.Amount,
		RecordType: 31, // 假设 31 为矿机扣款
		Remark:     fmt.Sprintf("购买矿机: %s", product.Title),
	}, func(ctx context.Context, tx gdb.TX) error {
		orderNo := uuid.New().String()
		now := gtime.Now()
		// 矿机天数逻辑可能较复杂，这里采用简单 Days 分隔
		days := gconv.Int(product.Days)
		if days == 0 {
			days = 30
		} // 默认 30 天

		endTime := now.Add(time.Duration(days) * 24 * time.Hour)

		_, err := dao.MingOrder.Ctx(ctx).TX(tx).Data(entity.MingOrder{
			OrderNo:      orderNo,
			UserId:       gconv.Int64(userId),
			PlanId:       gconv.Int64(product.Id),
			PlanTitle:    product.Title,
			Amount:       req.Amount,
			OrderAmount:  req.Amount,
			Days:         days,
			MinOdds:      product.MinOdds,
			MaxOdds:      product.MaxOdds,
			Status:       0,
			CreateTime:   now,
			EndTime:      gtime.New(endTime),
			AccumulaEarn: 0,
		}).Insert()
		return err
	})

	return
}

// Redemption 提前赎回矿机
func (s *sMining) Redemption(ctx context.Context, userId uint64, orderNo string) (err error) {
	// 1. 锁单查询
	var order entity.MingOrder
	err = dao.MingOrder.Ctx(ctx).Where("order_no", orderNo).Where("user_id", userId).Scan(&order)
	if err != nil || order.Id == 0 {
		return gerror.NewCode(codes.Failed, "订单不存在")
	}
	if order.Status != 0 {
		return gerror.NewCode(codes.Failed, "该订单已结算或赎回")
	}

	// 2. 退还本金 (AddAmount)
	// 在原系统中，提前赎回可能扣除手续费或不退收益，此处简化为全退本金
	_, err = asset.New().AddAmount(ctx, &v1a.SubAmountReq{
		UserId:     gconv.Int64(userId),
		Symbol:     "USDT", // 矿机通常退 USDT
		Amount:     order.Amount,
		RecordType: 32, // 矿机赎回
		Remark:     fmt.Sprintf("矿机提前赎回: %s", order.PlanTitle),
	}, func(ctx context.Context, tx gdb.TX) error {
		// 修改订单状态
		_, err := dao.MingOrder.Ctx(ctx).TX(tx).
			Data(g.Map{"status": 1, "update_time": gtime.Now()}).
			Where("id", order.Id).
			Update()
		return err
	})

	return
}

// GetProductList 单独查询可用矿机列表 (对应 /api/mingProduct/list)
func (s *sMining) GetProductList(ctx context.Context, req *v1.MingProductListReq) (res *v1.MingProductListRes, err error) {
	res = &v1.MingProductListRes{}
	var products []entity.MingProduct
	err = dao.MingProduct.Ctx(ctx).Where("status", 1).Order("sort asc").Scan(&products)
	for _, p := range products {
		res.Rows = append(res.Rows, v1.MiningProductInfo{
			Id:       gconv.Int64(p.Id),
			Title:    p.Title,
			MinPrice: p.LimitMin,
			MaxPrice: p.LimitMax,
			MinOdds:  p.MinOdds,
			MaxOdds:  p.MaxOdds,
		})
	}
	return
}

// GetOrderList 单独查询我的订单 (对应 /api/mingOrder/list)
func (s *sMining) GetOrderList(ctx context.Context, userId uint64, req *v1.MingOrderListReq) (res *v1.MingOrderListRes, err error) {
	res = &v1.MingOrderListRes{}
	var orders []entity.MingOrder
	err = dao.MingOrder.Ctx(ctx).Where("user_id", userId).Order("create_time desc").Scan(&orders)
	for _, o := range orders {
		res.Rows = append(res.Rows, v1.MiningOrderInfo{
			OrderNo:    o.OrderNo,
			Amount:     o.Amount,
			Days:       o.Days,
			Status:     o.Status,
			CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

// GetOrderDetail 查询单笔订单详情 (对应 /api/mingOrder/{id})
func (s *sMining) GetOrderDetail(ctx context.Context, userId uint64, req *v1.MingOrderDetailReq) (res *v1.MingOrderDetailRes, err error) {
	var o entity.MingOrder
	err = dao.MingOrder.Ctx(ctx).Where("id", req.Id).Where("user_id", userId).Scan(&o)
	if err != nil || o.Id == 0 {
		return nil, gerror.NewCode(codes.Failed, "订单不存在或无权访问")
	}

	res = &v1.MingOrderDetailRes{
		Data: v1.MiningOrderInfo{
			OrderNo:    o.OrderNo,
			Amount:     o.Amount,
			Days:       o.Days,
			Status:     o.Status,
			CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}
	return res, nil
}
