package funding

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminFunding struct{}

func New() *sAdminFunding {
	return &sAdminFunding{}
}

// GetRechargeList 充值列表查询
func (s *sAdminFunding) GetRechargeList(ctx context.Context, req *v1.GetRechargeListReq) (*v1.GetRechargeListRes, error) {
	m := dao.AppRecharge.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderNo != "" {
		m = m.WhereLike("serial_id", "%"+req.OrderNo+"%")
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.AppRecharge
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.RechargeInfo, 0, len(list))
	for _, r := range list {
		resList = append(resList, v1.RechargeInfo{
			Id:         r.Id,
			UserId:     r.UserId,
			LoginName:  r.Username,
			OrderNo:    r.SerialId,
			Coin:       r.Coin,
			Amount:     r.Amount,
			Address:    r.Address,
			TxHash:     r.TxId,
			Status:     r.Status,
			CreateTime: r.CreateTime.Format("2006-01-02 15:04:05"),
			Remark:     r.Remark,
		})
	}

	return &v1.GetRechargeListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetWithdrawList 提现列表查询
func (s *sAdminFunding) GetWithdrawList(ctx context.Context, req *v1.GetWithdrawListReq) (*v1.GetWithdrawListRes, error) {
	m := dao.Withdraw.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderNo != "" {
		m = m.WhereLike("serial_id", "%"+req.OrderNo+"%")
	}
	// Withdraw Status definition might vary, mapping exactly to DB
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.Withdraw
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.WithdrawInfo, 0, len(list))
	for _, w := range list {
		resList = append(resList, v1.WithdrawInfo{
			Id:         w.Id,
			UserId:     w.UserId,
			LoginName:  w.Username,
			OrderNo:    w.SerialId,
			Coin:       w.Coin,
			Amount:     w.Amount,
			Address:    w.Address,
			Status:     w.Status,
			CreateTime: w.CreateTime.Format("2006-01-02 15:04:05"),
			Remark:     w.Remark,
		})
	}

	return &v1.GetWithdrawListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetRechargeChannelList 充值通道配置
func (s *sAdminFunding) GetRechargeChannelList(ctx context.Context, req *v1.GetRechargeChannelListReq) (*v1.GetRechargeChannelListRes, error) {
	m := dao.SymbolManage.Ctx(ctx)
	if req.Symbol != "" {
		m = m.WhereLike("symbol", "%"+req.Symbol+"%")
	}

	total, _ := m.Count()
	var list []entity.SymbolManage
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.RechargeChannelInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.RechargeChannelInfo{
			Id:         c.Id,
			Symbol:     c.Symbol,
			Enable:     c.Enable,
			Type:       "Digital", // Custom mapped
			CreateTime: c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetRechargeChannelListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetWithdrawChannelList 提现通道配置
func (s *sAdminFunding) GetWithdrawChannelList(ctx context.Context, req *v1.GetWithdrawChannelListReq) (*v1.GetWithdrawChannelListRes, error) {
	m := dao.CurrencySymbol.Ctx(ctx)
	if req.Coin != "" {
		m = m.WhereLike("coin", "%"+req.Coin+"%")
	}

	total, _ := m.Count()
	var list []entity.CurrencySymbol
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.WithdrawChannelInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.WithdrawChannelInfo{
			Id:         c.Id,
			Coin:       c.Coin,
			Enable:     c.Enable,
			FeeRate:    "Fixed Rate", // Format properly based on DB struct if needed, ignoring detailed decimals for simple CRUD mapping here
			MinLimit:   "0",
			MaxLimit:   "1000000",
			CreateTime: c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetWithdrawChannelListRes{
		List:  resList,
		Total: total,
	}, nil
}
