package collection

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/collection"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetCollectionOrderList 归集订单查询
func (c *Controller) GetCollectionOrderList(ctx context.Context, req *v1.GetCollectionOrderListReq) (res *v1.GetCollectionOrderListRes, err error) {
	data, err := collection.New().GetCollectionOrderList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "操作成功",
		"total": data.Total,
		"data":  data.Rows,
	})
	return nil, nil // 返回 nil 绕过中间件包装
}
