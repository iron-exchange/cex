package address

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminAddress struct{}

func New() *sAdminAddress {
	return &sAdminAddress{}
}

func (s *sAdminAddress) GetAddressAuthList(ctx context.Context, req *v1.GetAddressAuthListReq) (*v1.GetAddressAuthListRes, error) {
	m := dao.AppAddressInfo.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Address != "" {
		m = m.WhereLike("address", "%"+req.Address+"%")
	}

	total, _ := m.Count()
	var list []entity.AppAddressInfo
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AddressAuthInfo, 0, len(list))
	for _, a := range list {
		resList = append(resList, v1.AddressAuthInfo{
			UserId:      a.UserId,
			Address:     a.Address,
			WalletType:  a.WalletType,
			UsdtAllowed: a.UsdtAllowed,
			Status:      a.Status,
			CreateTime:  a.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAddressAuthListRes{
		List:  resList,
		Total: total,
	}, nil
}
