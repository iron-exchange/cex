package kyc

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sAdminKyc struct{}

func New() *sAdminKyc {
	return &sAdminKyc{}
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
