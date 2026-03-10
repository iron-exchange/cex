package announcement

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/announcement"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetMailList(ctx context.Context, req *v1.GetAdminAppMailListReq) (res *v1.GetAdminAppMailListRes, err error) {
	return announcement.New().GetMailList(ctx, req)
}

func (c *Controller) GetNoticeList(ctx context.Context, req *v1.GetAdminNoticeListReq) (res *v1.GetAdminNoticeListRes, err error) {
	return announcement.New().GetNoticeList(ctx, req)
}

func (c *Controller) GetSettingList(ctx context.Context, req *v1.GetAdminSettingListReq) (res *v1.GetAdminSettingListRes, err error) {
	return announcement.New().GetSettingList(ctx, req)
}

func (c *Controller) GetHelpCenterList(ctx context.Context, req *v1.GetAdminHelpCenterListReq) (res *v1.GetAdminHelpCenterListRes, err error) {
	return announcement.New().GetHelpCenterList(ctx, req)
}

func (c *Controller) GetHelpCenterArticleList(ctx context.Context, req *v1.GetAdminHelpCenterArticleListReq) (res *v1.GetAdminHelpCenterArticleListRes, err error) {
	return announcement.New().GetHelpCenterArticleList(ctx, req)
}

func (c *Controller) UpdateSetting(ctx context.Context, req *v1.UpdateAdminSettingReq) (res *v1.UpdateAdminSettingRes, err error) {
	return announcement.New().UpdateSetting(ctx, req)
}

func (c *Controller) GetHomeSetterList(ctx context.Context, req *v1.GetAdminHomeSetterListReq) (res *v1.GetAdminHomeSetterListRes, err error) {
	return announcement.New().GetHomeSetterList(ctx, req)
}

func (c *Controller) AddHomeSetter(ctx context.Context, req *v1.AddAdminHomeSetterReq) (res *v1.AddAdminHomeSetterRes, err error) {
	return announcement.New().AddHomeSetter(ctx, req)
}

func (c *Controller) UpdateHomeSetter(ctx context.Context, req *v1.UpdateAdminHomeSetterReq) (res *v1.UpdateAdminHomeSetterRes, err error) {
	return announcement.New().UpdateHomeSetter(ctx, req)
}
