package ieo

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/ieo"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetOwnCoinList 币种列表
func (c *Controller) GetOwnCoinList(ctx context.Context, req *v1.GetOwnCoinListReq) (res *v1.GetOwnCoinListRes, err error) {
	return ieo.New().GetOwnCoinList(ctx, req)
}

// GetOwnCoinSubscribeOrderList 申购订单查询
func (c *Controller) GetOwnCoinSubscribeOrderList(ctx context.Context, req *v1.GetOwnCoinSubscribeOrderListReq) (res *v1.GetOwnCoinSubscribeOrderListRes, err error) {
	return ieo.New().GetOwnCoinSubscribeOrderList(ctx, req)
}

// GetOwnCoinOrderList 认购下单查询
func (c *Controller) GetOwnCoinOrderList(ctx context.Context, req *v1.GetOwnCoinOrderListReq) (res *v1.GetOwnCoinOrderListRes, err error) {
	return ieo.New().GetOwnCoinOrderList(ctx, req)
}

// EditPlacing 审批/调整订单
func (c *Controller) EditPlacing(ctx context.Context, req *v1.EditOwnCoinOrderPlacingReq) (res *g.Var, err error) {
	err = ieo.New().EditOwnCoinOrderPlacing(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// Get 获取单条详情
func (c *Controller) Get(ctx context.Context, req *v1.GetOwnCoinOrderReq) (res *g.Var, err error) {
	data, err := ieo.New().Get(ctx, req.Id)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功", "data": data})
	return nil, nil
}

// Create 手动补单
func (c *Controller) Create(ctx context.Context, req *v1.CreateOwnCoinOrderReq) (res *g.Var, err error) {
	err = ieo.New().Create(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// Delete 批量删除认购订单
func (c *Controller) Delete(ctx context.Context, req *v1.DeleteOwnCoinOrderReq) (res *g.Var, err error) {
	err = ieo.New().Delete(ctx, req.Ids)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// Export 导出认购数据
func (c *Controller) Export(ctx context.Context, req *v1.ExportOwnCoinOrderReq) (res *g.Var, err error) {
	data, err := ieo.New().Export(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功", "data": data})
	return nil, nil
}
