package ieo

import (
	"context"
	"fmt"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminIeo struct{}

func New() *sAdminIeo {
	return &sAdminIeo{}
}

// GetOwnCoinList 查询 IEO 新币列表
func (s *sAdminIeo) GetOwnCoinList(ctx context.Context, req *v1.GetOwnCoinListReq) (*v1.GetOwnCoinListRes, error) {
	m := dao.OwnCoin.Ctx(ctx)
	if req.Coin != "" {
		m = m.WhereLike("coin", "%"+req.Coin+"%")
	}

	total, _ := m.Count()
	var list []entity.OwnCoin
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.OwnCoinInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.OwnCoinInfo{
			Id:            c.Id,
			Coin:          c.Coin,
			Logo:          c.Logo,
			ReferCoin:     c.ReferCoin,
			ShowSymbol:    c.ShowSymbol,
			Price:         c.Price,
			Proportion:    c.Proportion,
			RaisingAmount: c.RaisingAmount,
			RaisedAmount:  c.RaisedAmount,
			PurchaseLimit: c.PurchaseLimit,
			TotalAmount:   c.TotalAmount,
			Status:        c.Status,
			BeginTime:     c.BeginTime.Format("2006-01-02 15:04:05"),
			EndTime:       c.EndTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetOwnCoinListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetOwnCoinSubscribeOrderList 查询申购打新订单列表
func (s *sAdminIeo) GetOwnCoinSubscribeOrderList(ctx context.Context, req *v1.GetOwnCoinSubscribeOrderListReq) (*v1.GetOwnCoinSubscribeOrderListRes, error) {
	m := dao.OwnCoinSubscribeOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OwnCoin != "" {
		m = m.WhereLike("own_coin", "%"+req.OwnCoin+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.OwnCoinSubscribeOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.OwnCoinSubscribeOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.OwnCoinSubscribeOrderInfo{
			Id:          o.Id,
			SubscribeId: o.SubscribeId,
			UserId:      o.UserId,
			OrderId:     o.OrderId,
			OwnCoin:     o.OwnCoin,
			AmountLimit: o.AmountLimit,
			NumLimit:    o.NumLimit,
			Price:       o.Price,
			Status:      o.Status,
			CreateTime:  o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetOwnCoinSubscribeOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetOwnCoinOrderList 查询认购下单记录
func (s *sAdminIeo) GetOwnCoinOrderList(ctx context.Context, req *v1.GetOwnCoinOrderListReq) (*v1.GetOwnCoinOrderListRes, error) {
	adminId := gconv.Int(ctx.Value("adminId"))

	m := dao.OwnCoinOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderId != "" {
		m = m.Where("order_id", req.OrderId)
	}
	if req.OwnId > 0 {
		m = m.Where("own_id", req.OwnId)
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	// 数据隔离: 非超级管理员只能看自己及其下级的资产
	if adminId != 1 {
		m = m.Where(fmt.Sprintf("FIND_IN_SET(%d, admin_parent_ids)", adminId))
	}

	total, _ := m.Count()
	var list []entity.OwnCoinOrder
	err := m.Page(req.PageNum, req.PageSize).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.OwnCoinOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.OwnCoinOrderInfo{
			Id:             o.Id,
			UserId:         o.UserId,
			OrderId:        o.OrderId,
			OwnId:          o.OwnId,
			OwnCoin:        o.OwnCoin,
			Number:         o.Number,
			Price:          o.Price,
			Amount:         o.Amount,
			Status:         o.Status,
			AdminUserIds:   o.AdminUserIds,
			AdminParentIds: o.AdminParentIds,
			Remark:         o.Remark,
			CreateBy:       o.CreateBy,
			UpdateBy:       o.UpdateBy,
			CreateTime:     o.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:     o.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetOwnCoinOrderListRes{
		Rows:  resList,
		Total: total,
	}, nil
}

// EditOwnCoinOrderPlacing 审批/调整订单 (支持自动退款)
func (s *sAdminIeo) EditOwnCoinOrderPlacing(ctx context.Context, req *v1.EditOwnCoinOrderPlacingReq) error {
	return dao.OwnCoinOrder.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var order entity.OwnCoinOrder
		err := dao.OwnCoinOrder.Ctx(ctx).TX(tx).LockUpdate().Where("id", req.Id).Scan(&order)
		if err != nil || order.Id == 0 {
			return gerror.New("订单不存在")
		}

		beforeNum := order.Number
		currentNum := req.Number

		// 如果调低了数量，需要退还差额
		if currentNum < beforeNum {
			diffNum := beforeNum - currentNum
			refundAmount := float64(diffNum) * order.Price

			// 退回 USDT 到资产账户
			_, err = asset.New().SubAmount(ctx, &v1.SubAmountReq{
				UserId:     order.UserId,
				Symbol:     "usdt",
				Amount:     -refundAmount, // 负数表示增加余额
				RecordType: 52,            // 新币申购退回 (使用 52)
				Remark:     fmt.Sprintf("新币[%s]认购配额调整退回: %v USDT", order.OwnCoin, refundAmount),
			})
			if err != nil {
				return gerror.Wrap(err, "资金退回失败")
			}
		}

		// 更新订单数量与金额
		_, err = dao.OwnCoinOrder.Ctx(ctx).TX(tx).Data(g.Map{
			dao.OwnCoinOrder.Columns().Number:     currentNum,
			dao.OwnCoinOrder.Columns().Amount:     float64(currentNum) * order.Price,
			dao.OwnCoinOrder.Columns().UpdateTime: gtime.Now(),
			dao.OwnCoinOrder.Columns().UpdateBy:   gconv.String(ctx.Value("adminAccount")),
		}).Where("id", req.Id).Update()

		return err
	})
}

// Get 获取单条认购详情
func (s *sAdminIeo) Get(ctx context.Context, id int64) (*v1.OwnCoinOrderInfo, error) {
	var o entity.OwnCoinOrder
	err := dao.OwnCoinOrder.Ctx(ctx).Where("id", id).Scan(&o)
	if err != nil || o.Id == 0 {
		return nil, gerror.New("认购详情不存在")
	}

	return &v1.OwnCoinOrderInfo{
		Id:             o.Id,
		UserId:         o.UserId,
		OrderId:        o.OrderId,
		OwnId:          o.OwnId,
		OwnCoin:        o.OwnCoin,
		Number:         o.Number,
		Price:          o.Price,
		Amount:         o.Amount,
		Status:         o.Status,
		AdminUserIds:   o.AdminUserIds,
		AdminParentIds: o.AdminParentIds,
		Remark:         o.Remark,
		CreateBy:       o.CreateBy,
		UpdateBy:       o.UpdateBy,
		CreateTime:     o.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:     o.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}

// Create 手动补单
func (s *sAdminIeo) Create(ctx context.Context, req *v1.CreateOwnCoinOrderReq) error {
	var own entity.OwnCoin
	err := dao.OwnCoin.Ctx(ctx).Where("id", req.OwnId).Scan(&own)
	if err != nil || own.Id == 0 {
		return gerror.New("对应新币种不存在")
	}

	var user entity.AppUser
	_ = dao.AppUser.Ctx(ctx).Where("user_id", req.UserId).Scan(&user)

	orderId := fmt.Sprintf("M%d%d", req.UserId, gtime.Timestamp())
	_, err = dao.OwnCoinOrder.Ctx(ctx).Data(g.Map{
		dao.OwnCoinOrder.Columns().UserId:         req.UserId,
		dao.OwnCoinOrder.Columns().OrderId:        orderId,
		dao.OwnCoinOrder.Columns().OwnId:          req.OwnId,
		dao.OwnCoinOrder.Columns().OwnCoin:        own.Coin,
		dao.OwnCoinOrder.Columns().Number:         req.Number,
		dao.OwnCoinOrder.Columns().Price:          req.Price,
		dao.OwnCoinOrder.Columns().Amount:         float64(req.Number) * req.Price,
		dao.OwnCoinOrder.Columns().Status:         "1", // 默认已支付/订阅中
		dao.OwnCoinOrder.Columns().AdminUserIds:   user.AppParentIds,
		dao.OwnCoinOrder.Columns().AdminParentIds: user.AdminParentIds,
		dao.OwnCoinOrder.Columns().CreateTime:     gtime.Now(),
		dao.OwnCoinOrder.Columns().CreateBy:       gconv.String(ctx.Value("adminAccount")),
		dao.OwnCoinOrder.Columns().Remark:         req.Remark,
	}).Insert()

	return err
}

// Delete 批量删除认购订单
func (s *sAdminIeo) Delete(ctx context.Context, ids []int64) error {
	_, err := dao.OwnCoinOrder.Ctx(ctx).WhereIn("id", ids).Delete()
	return err
}

// Export 导出认购数据 (返回 List 由 Controller 处理格式)
func (s *sAdminIeo) Export(ctx context.Context, req *v1.ExportOwnCoinOrderReq) ([]v1.OwnCoinOrderInfo, error) {
	m := dao.OwnCoinOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderId != "" {
		m = m.Where("order_id", req.OrderId)
	}
	if req.OwnId > 0 {
		m = m.Where("own_id", req.OwnId)
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	var list []entity.OwnCoinOrder
	_ = m.OrderDesc("id").Scan(&list)

	resList := make([]v1.OwnCoinOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.OwnCoinOrderInfo{
			Id:             o.Id,
			UserId:         o.UserId,
			OrderId:        o.OrderId,
			OwnId:          o.OwnId,
			OwnCoin:        o.OwnCoin,
			Number:         o.Number,
			Price:          o.Price,
			Amount:         o.Amount,
			Status:         o.Status,
			AdminUserIds:   o.AdminUserIds,
			AdminParentIds: o.AdminParentIds,
			CreateTime:     o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return resList, nil
}
