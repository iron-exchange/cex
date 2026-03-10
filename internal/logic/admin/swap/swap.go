package swap

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminSwap struct{}

func New() *sAdminSwap {
	return &sAdminSwap{}
}

// GetSymbolManageList 获取闪兑币种配置表
func (s *sAdminSwap) GetSymbolManageList(ctx context.Context, req *v1.GetAdminSymbolManageListReq) (*v1.GetAdminSymbolManageListRes, error) {
	m := dao.SymbolManage.Ctx(ctx)
	if req.Symbol != "" {
		m = m.WhereLike("symbol", "%"+req.Symbol+"%")
	}
	m = m.Where("del_flag", "0") // 0正常 2删除

	total, _ := m.Count()
	var list []entity.SymbolManage
	err := m.Page(req.Page, req.Size).OrderDesc("sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSymbolManageInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.AdminSymbolManageInfo{
			Id:           c.Id,
			Symbol:       c.Symbol,
			MinChargeNum: c.MinChargeNum,
			MaxChargeNum: c.MaxChargeNum,
			Commission:   c.Commission,
			Sort:         c.Sort,
			Enable:       c.Enable,
			Logo:         c.Logo,
			Market:       c.Market,
			CreateTime:   c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminSymbolManageListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetExchangeCoinRecordList 获取闪兑订单记录
func (s *sAdminSwap) GetExchangeCoinRecordList(ctx context.Context, req *v1.GetAdminExchangeCoinRecordListReq) (*v1.GetAdminExchangeCoinRecordListRes, error) {
	m := dao.ExchangeCoinRecord.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}

	total, _ := m.Count()
	var list []entity.ExchangeCoinRecord
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminExchangeCoinRecordInfo, 0, len(list))
	for _, r := range list {
		resList = append(resList, v1.AdminExchangeCoinRecordInfo{
			Id:         r.Id,
			UserId:     r.UserId,
			Username:   r.Username,
			FromCoin:   r.FromCoin,
			ToCoin:     r.ToCoin,
			Amount:     r.Amount,
			ThirdRate:  r.ThirdRate,
			SystemRate: r.SystemRate,
			Status:     r.Status,
			CreateTime: r.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminExchangeCoinRecordListRes{
		List:  resList,
		Total: total,
	}, nil
}
