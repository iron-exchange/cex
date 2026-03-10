package second_contract

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminSecondContract struct{}

func New() *sAdminSecondContract {
	return &sAdminSecondContract{}
}

// GetSecondCoinConfigList 查询币种配置列表
func (s *sAdminSecondContract) GetSecondCoinConfigList(ctx context.Context, req *v1.GetSecondCoinConfigListReq) (*v1.GetSecondCoinConfigListRes, error) {
	m := dao.SecondCoinConfig.Ctx(ctx)
	if req.Symbol != "" {
		m = m.WhereLike("symbol", "%"+req.Symbol+"%")
	}

	total, _ := m.Count()
	var list []entity.SecondCoinConfig
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.SecondCoinConfigInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.SecondCoinConfigInfo{
			Id:         c.Id,
			Symbol:     c.Symbol,
			Market:     c.Market,
			Status:     c.Status,
			ShowFlag:   c.ShowFlag,
			Coin:       c.Coin,
			Sort:       c.Sort,
			Logo:       c.Logo,
			BaseCoin:   c.BaseCoin,
			ShowSymbol: c.ShowSymbol,
			Type:       c.Type,
			CreateTime: c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetSecondCoinConfigListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetSecondContractOrderList 查询秒合约订单列表
func (s *sAdminSecondContract) GetSecondContractOrderList(ctx context.Context, req *v1.GetSecondContractOrderListReq) (*v1.GetSecondContractOrderListRes, error) {
	m := dao.SecondContractOrder.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.OrderNo != "" {
		m = m.WhereLike("order_no", "%"+req.OrderNo+"%")
	}
	if req.Symbol != "" {
		m = m.WhereLike("symbol", "%"+req.Symbol+"%")
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.SecondContractOrder
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.SecondContractOrderInfo, 0, len(list))
	for _, o := range list {
		resList = append(resList, v1.SecondContractOrderInfo{
			Id:                 o.Id,
			OrderNo:            o.OrderNo,
			Symbol:             o.Symbol,
			Type:               o.Type,
			UserId:             o.UserId,
			UserAddress:        o.UserAddress,
			BetContent:         o.BetContent,
			OpenResult:         o.OpenResult,
			Status:             o.Status,
			RateFlag:           o.RateFlag,
			BetAmount:          o.BetAmount,
			RewardAmount:       o.RewardAmount,
			CompensationAmount: o.CompensationAmount,
			OpenPrice:          o.OpenPrice,
			ClosePrice:         o.ClosePrice,
			CoinSymbol:         o.CoinSymbol,
			BaseSymbol:         o.BaseSymbol,
			Sign:               o.Sign,
			ManualIntervention: o.ManualIntervention,
			Rate:               o.Rate,
			CreateTime:         o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetSecondContractOrderListRes{
		List:  resList,
		Total: total,
	}, nil
}
