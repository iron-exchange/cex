package kyc

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sAdminKyc struct{}

func New() *sAdminKyc {
	return &sAdminKyc{}
}

func (s *sAdminKyc) GetKycList(ctx context.Context, req *v1.GetKycListReq) (*v1.GetKycListRes, error) {
	m := dao.AppUserDetail.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.AuditStatus != nil {
		m = m.Where("audit_status_primary", *req.AuditStatus)
	}

	total, _ := m.Count()
	var list []entity.AppUserDetail
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.KycInfo, 0, len(list))
	for _, k := range list {
		resList = append(resList, v1.KycInfo{
			UserId:              k.UserId,
			RealName:            k.RealName,
			IdCard:              k.IdCard,
			FrontUrl:            k.FrontUrl,
			BackUrl:             k.BackUrl,
			HandelUrl:           k.HandelUrl,
			AuditStatusPrimary:  k.AuditStatusPrimary,
			AuditStatusAdvanced: k.AuditStatusAdvanced,
			CreateTime:          k.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetKycListRes{
		List:  resList,
		Total: total,
	}, nil
}

func (s *sAdminKyc) AuditKyc(ctx context.Context, req *v1.AuditKycReq) error {
	status := 0
	if req.Action == "pass" {
		status = 1
	} else if req.Action == "reject" {
		status = 2
	}

	updateData := map[string]interface{}{}
	if req.AuditLevel == 1 {
		updateData["audit_status_primary"] = status
	} else {
		updateData["audit_status_advanced"] = status
	}

	if req.Remark != "" {
		updateData["remark"] = req.Remark
	}

	_, err := dao.AppUserDetail.Ctx(ctx).Where("user_id", req.UserId).Update(updateData)
	if err != nil {
		return gerror.Wrap(err, "审核操作失败")
	}
	return nil
}
