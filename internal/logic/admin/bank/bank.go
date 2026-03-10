package bank

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminBank struct{}

func New() *sAdminBank {
	return &sAdminBank{}
}

func (s *sAdminBank) GetBankList(ctx context.Context, req *v1.GetBankListReq) (*v1.GetBankListRes, error) {
	m := dao.UserBank.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.CardNumber != "" {
		m = m.WhereLike("card_number", "%"+req.CardNumber+"%")
	}

	total, _ := m.Count()
	var list []entity.UserBank
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.BankInfo, 0, len(list))
	for _, b := range list {
		resList = append(resList, v1.BankInfo{
			Id:         b.Id,
			UserId:     b.UserId,
			BankName:   b.BankName,
			BankBranch: b.BankBranch,
			CardNumber: b.CardNumber,
			RealName:   b.UserName,
			CreateTime: b.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetBankListRes{
		List:  resList,
		Total: total,
	}, nil
}
