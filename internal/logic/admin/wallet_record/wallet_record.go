package wallet_record

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminWalletRecord struct{}

func New() *sAdminWalletRecord {
	return &sAdminWalletRecord{}
}

func (s *sAdminWalletRecord) GetWalletRecordList(ctx context.Context, req *v1.GetWalletRecordListReq) (*v1.GetWalletRecordListRes, error) {
	m := dao.AppWalletRecord.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Symbol != "" {
		m = m.Where("symbol", req.Symbol)
	}
	if req.Type != nil {
		m = m.Where("type", *req.Type)
	}

	total, _ := m.Count()
	var list []entity.AppWalletRecord
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	// 为简单起见，这里不再连表查 User 表提取 loginName，真实环境可以关联查询或左连接
	resList := make([]v1.WalletRecordInfo, 0, len(list))
	for _, r := range list {
		resList = append(resList, v1.WalletRecordInfo{
			Id:           r.Id,
			UserId:       r.UserId,
			LoginName:    "", // 需要时使用 dao.AppUser.Ctx(ctx).Fields("login_name").Where("user_id", r.UserId) 获取
			Symbol:       r.Symbol,
			Type:         r.Type,
			Amount:       r.Amount,
			UAmount:      r.UAmount,
			BeforeAmount: r.BeforeAmount,
			AfterAmount:  r.AfterAmount,
			OperateTime:  r.OperateTime.Format("2006-01-02 15:04:05"),
			Remark:       r.Remark,
		})
	}

	return &v1.GetWalletRecordListRes{
		List:  resList,
		Total: total,
	}, nil
}
