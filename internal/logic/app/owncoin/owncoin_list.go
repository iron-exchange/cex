package owncoin

import (
	"context"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAppOwnCoin struct{}

func New() *sAppOwnCoin {
	return &sAppOwnCoin{}
}

// GetOwnCoinList 获取正在发行或预热的自发币列表
func (s *sAppOwnCoin) GetOwnCoinList(ctx context.Context, req *v1.GetOwnCoinListReq) (*v1.GetOwnCoinListRes, error) {
	m := dao.OwnCoin.Ctx(ctx)
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.OwnCoin
	err := m.Page(req.Page, req.Size).OrderDesc("id").Scan(&list)

	resList := make([]v1.OwnCoinAppInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.OwnCoinAppInfo{
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
			BeginTime:     c.BeginTime.String(),
			EndTime:       c.EndTime.String(),
			Introduce:     c.Introduce,
		})
	}

	return &v1.GetOwnCoinListRes{
		Total: total,
		Rows:  resList,
	}, err
}
