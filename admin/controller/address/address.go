package address

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/address"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetAddressAuthList 授权地址列表查询
func (c *Controller) GetAddressAuthList(ctx context.Context, req *v1.GetAddressAuthListReq) (res *v1.GetAddressAuthListRes, err error) {
	data, err := address.New().GetAddressAuthList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": data.List,
	})
	return nil, nil
}

// GetAddressInfoList 钱包地址授权详情列表
func (c *Controller) GetAddressInfoList(ctx context.Context, req *v1.GetAddressInfoListReq) (res *v1.GetAddressInfoListRes, err error) {
	data, err := address.New().GetAddressInfoList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "操作成功",
		"total": data.Total,
		"data":  data.Rows,
	})
	return nil, nil // 绕过全局包装
}

// RefreshAddressInfo 强制手动刷新某个钱包地址的链上数据
func (c *Controller) RefreshAddressInfo(ctx context.Context, req *v1.RefreshAddressInfoReq) (res *v1.RefreshAddressInfoRes, err error) {
	_, err = address.New().RefreshAddressInfo(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
	})
	return nil, nil
}

// Collection 手动发起资产归集
func (c *Controller) Collection(ctx context.Context, req *v1.AddressCollectionReq) (res *v1.AddressCollectionRes, err error) {
	_, err = address.New().Collection(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
	})
	return nil, nil
}

// GetAddressInfo 获取单条地址详情
func (c *Controller) GetAddressInfo(ctx context.Context, req *v1.GetAddressInfoReq) (res *v1.GetAddressInfoRes, err error) {
	data, err := address.New().GetAddressInfo(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": data.AppAddressInfoDetail,
	})
	return nil, nil
}

// UpdateAddressInfo 修改单条地址详情
func (c *Controller) UpdateAddressInfo(ctx context.Context, req *v1.UpdateAddressInfoReq) (res *v1.UpdateAddressInfoRes, err error) {
	err = address.New().UpdateAddressInfo(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code": 500,
			"msg":  err.Error(),
		})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
	})
	return nil, nil
}
