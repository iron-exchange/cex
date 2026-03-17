package asset

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/asset"

	"github.com/gogf/gf/v2/frame/g"
)

// Controller admin 端的资产接口 (供运营人员调整资金)
type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SubAmount 后台人工上下分 / 强制资金划转
func (c *Controller) SubAmount(ctx context.Context, req *v1.SubAmountReq) (res *v1.SubAmountRes, err error) {
	return asset.New().SubAmount(ctx, req)
}

// GetAppAssetList 获取玩家资产列表
func (c *Controller) GetAppAssetList(ctx context.Context, req *v1.GetAppAssetListReq) (res *v1.GetAppAssetListRes, err error) {
	data, err := asset.New().GetAppAssetList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "查询成功",
		"total": data.Total,
		"data":  data.Rows,
	})
	return nil, nil
}
