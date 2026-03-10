package kyc

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/kyc"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetKycList 玩家实名认证列表
func (c *Controller) GetKycList(ctx context.Context, req *v1.GetKycListReq) (res *v1.GetKycListRes, err error) {
	return kyc.New().GetKycList(ctx, req)
}

// AuditKyc 审核玩家实名认证
func (c *Controller) AuditKyc(ctx context.Context, req *v1.AuditKycReq) (res *v1.AuditKycRes, err error) {
	err = kyc.New().AuditKyc(ctx, req)
	return &v1.AuditKycRes{}, err
}
