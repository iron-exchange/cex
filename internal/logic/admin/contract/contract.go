package contract

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminContract struct{}

func New() *sAdminContract {
	return &sAdminContract{}
}

// GetContractCoinList 获取U本位合约币种配置
func (s *sAdminContract) GetContractCoinList(ctx context.Context, req *v1.GetContractCoinListReq) (*v1.GetContractCoinListRes, error) {
	m := dao.ContractCoin.Ctx(ctx)
	if req.Symbol != "" {
		m = m.WhereLike("symbol", "%"+req.Symbol+"%")
	}

	total, _ := m.Count()
	var list []entity.ContractCoin
	err := m.Page(req.Page, req.Size).OrderDesc("sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.ContractCoinInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.ContractCoinInfo{
			Id:           c.Id,
			Symbol:       c.Symbol,
			Coin:         c.Coin,
			BaseCoin:     c.BaseCoin,
			ShareNumber:  c.ShareNumber,
			Leverage:     c.Leverage,
			Enable:       c.Enable,
			Exchangeable: c.Exchangeable,
			OpenFee:      c.OpenFee,
			CloseFee:     c.CloseFee,
			UsdtRate:     c.UsdtRate,
			Visible:      c.Visible,
			MinShare:     c.MinShare,
			MaxShare:     c.MaxShare,
			CreateTime:   c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetContractCoinListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetContractOrderList 获取合约历史委托订单
func (s *sAdminContract) GetContractOrderList(ctx context.Context, req *v1.GetContractOrderListReq) (*v1.GetContractOrderListRes, error) {
	m := dao.ContractOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderNo != "" {
		m = m.WhereLike("order_no", "%"+req.OrderNo+"%")
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.ContractOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.ContractOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.ContractOrderInfo{
			Id:            o.Id,
			OrderNo:       o.OrderNo,
			UserId:        o.UserId,
			Symbol:        o.Symbol,
			Type:          o.Type,
			DelegateType:  o.DelegateType,
			Status:        o.Status,
			DelegateTotal: o.DelegateTotal,
			DelegatePrice: o.DelegatePrice,
			DealNum:       o.DealNum,
			DealPrice:     o.DealPrice,
			Leverage:      o.Leverage,
			Fee:           o.Fee,
			CreateTime:    o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetContractOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetContractPositionList 获取当前未平仓头寸仓位
func (s *sAdminContract) GetContractPositionList(ctx context.Context, req *v1.GetContractPositionListReq) (*v1.GetContractPositionListRes, error) {
	m := dao.ContractPosition.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderNo != "" {
		m = m.WhereLike("order_no", "%"+req.OrderNo+"%")
	}
	if req.Symbol != "" {
		m = m.Where("symbol", req.Symbol)
	}

	total, _ := m.Count()
	var list []entity.ContractPosition
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.ContractPositionInfo, 0, len(list))
	for _, p := range list {
		resList = append(resList, v1.ContractPositionInfo{
			Id:         p.Id,
			OrderNo:    p.OrderNo,
			UserId:     p.UserId,
			Symbol:     p.Symbol,
			Type:       p.Type,
			Status:     p.Status,
			Amount:     p.Amount,
			OpenNum:    p.OpenNum,
			OpenPrice:  p.OpenPrice,
			ClosePrice: p.ClosePrice,
			Leverage:   p.Leverage,
			Earn:       p.Earn,
			OpenFee:    p.OpenFee,
			CreateTime: p.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetContractPositionListRes{
		List:  resList,
		Total: total,
	}, nil
}
