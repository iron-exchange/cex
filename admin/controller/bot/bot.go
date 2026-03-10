package bot

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/bot"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetBotModelList(ctx context.Context, req *v1.GetAdminBotKlineModelListReq) (res *v1.GetAdminBotKlineModelListRes, err error) {
	return bot.New().GetBotModelList(ctx, req)
}

func (c *Controller) GetBotModelDataList(ctx context.Context, req *v1.GetAdminBotKlineModelDataListReq) (res *v1.GetAdminBotKlineModelDataListRes, err error) {
	return bot.New().GetBotModelDataList(ctx, req)
}
