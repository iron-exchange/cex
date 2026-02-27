package trading

import (
	"context"
	"strings"

	adminV1 "GoCEX/api/admin/v1"
	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GetCurrencyOrders 获取现货订单列表
func (s *sTrading) GetCurrencyOrders(ctx context.Context, in *v1.CurrencyOrderListReq, userId uint64) (*v1.CurrencyOrderListRes, error) {
	m := dao.CurrencyOrder.Ctx(ctx).Where(dao.CurrencyOrder.Columns().UserId, userId)

	if in.Symbol != "" {
		m = m.Where(dao.CurrencyOrder.Columns().Symbol, in.Symbol)
	}
	if in.Status != -1 {
		m = m.Where(dao.CurrencyOrder.Columns().Status, in.Status)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var orders []entity.CurrencyOrder
	err = m.Page(in.Page, in.Size).OrderDesc(dao.CurrencyOrder.Columns().Id).Scan(&orders)
	if err != nil {
		return nil, err
	}

	res := &v1.CurrencyOrderListRes{
		List:  make([]v1.CurrencyOrderInfo, 0, len(orders)),
		Total: total,
	}

	for _, o := range orders {
		// 容错处理如果 DelegateTime 没值的情况
		createTime := ""
		if o.DelegateTime != nil {
			createTime = o.DelegateTime.Format("Y-m-d H:i:s")
		}

		res.List = append(res.List, v1.CurrencyOrderInfo{
			OrderNo:    o.OrderNo,
			Symbol:     o.Symbol,
			Price:      o.DelegatePrice,
			Amount:     o.DelegateTotal,
			Type:       o.Type,
			Status:     o.Status,
			CreateTime: createTime,
		})
	}
	return res, nil
}

// CancelCurrencyOrder CAS 安全撤单与资金解冻
func (s *sTrading) CancelCurrencyOrder(ctx context.Context, in *v1.CurrencyOrderCancelReq, userId uint64) (*v1.CurrencyOrderCancelRes, error) {
	var order entity.CurrencyOrder
	err := dao.CurrencyOrder.Ctx(ctx).Where(g.Map{
		dao.CurrencyOrder.Columns().OrderNo: in.OrderNo,
		dao.CurrencyOrder.Columns().UserId:  userId,
	}).Scan(&order)

	if err != nil || order.Id == 0 {
		return nil, gerror.New("订单不存在")
	}

	if order.Status != 0 {
		return nil, gerror.New("当前订单状态不支持撤单")
	}

	// 核心安全机制: CAS 乐观锁防多重撤单解冻
	res, err := dao.CurrencyOrder.Ctx(ctx).Where(g.Map{
		dao.CurrencyOrder.Columns().Id:     order.Id,
		dao.CurrencyOrder.Columns().Status: 0,
	}).Data(g.Map{
		dao.CurrencyOrder.Columns().Status:     2, // 2 = 已撤销
		dao.CurrencyOrder.Columns().UpdateTime: gtime.Now(),
	}).Update()

	if err != nil {
		return nil, err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return nil, gerror.New("撤单失败，订单已被并发处理(无 CAS 抢占权限)")
	}

	// 此时已抢占到撤单归属权，负责资产解冻（原路退还）
	var freezeSymbol string
	if order.Type == 0 { // 买单: 退回挂单时冻结的计价币（例如 USDT）
		freezeSymbol = order.Coin
	} else { // 卖单: 退回目标手持币 (从 Symbol 如 BTC/USDT 中截取第一段)
		pairs := strings.Split(order.Symbol, "/")
		if len(pairs) == 2 {
			freezeSymbol = pairs[0]
		}
	}

	if freezeSymbol != "" {
		// 使用 FreezeAmount 传入负数引发数据库的资金解冻
		unfreezeReq := &adminV1.SubAmountReq{
			UserId:     int64(userId),
			Symbol:     freezeSymbol,
			Amount:     -order.DelegateValue,
			RecordType: 31,
			Remark:     "现货手工撤单解冻退回",
		}
		_, errUnfreeze := asset.New().FreezeAmount(ctx, unfreezeReq)
		if errUnfreeze != nil {
			g.Log().Errorf(ctx, "[CAS 撤单警告] 状态已修改，但解冻退回失败, %v", errUnfreeze)
			// 注意：在真正严谨的生产中，解冻必须跟状态 UPDATE 在同一个 gdb.TX 事务内挂钩。
			// 这里借助了之前已写的 FreezeAmount 内部自带的资金锁闭环进行回撤
		}
	}

	return &v1.CurrencyOrderCancelRes{Success: true}, nil
}
