package cms

import (
	"context"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/cms"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/util/gconv"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetAllNoticeList 首页获取公告和轮播
func (c *Controller) GetAllNoticeList(ctx context.Context, req *v1.GetAllNoticeListReq) (res *v1.GetAllNoticeListRes, err error) {
	return cms.New().GetAllNoticeList(ctx)
}

// GetHelpCenterList 获取帮助中心列表
func (c *Controller) GetHelpCenterList(ctx context.Context, req *v1.GetHelpCenterListReq) (res *v1.GetHelpCenterListRes, err error) {
	return cms.New().GetHelpCenterList(ctx)
}

// GetUserMail 获取个人站内信 (必须带 Token 鉴权)
func (c *Controller) GetUserMail(ctx context.Context, req *v1.GetUserMailReq) (res *v1.GetUserMailRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return cms.New().GetUserMail(ctx, userId, req)
}
