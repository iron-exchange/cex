package wallet_record

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/wallet_record"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetWalletRecordList 归集订单查询
func (c *Controller) GetWalletRecordList(ctx context.Context, req *v1.GetWalletRecordListReq) (res *v1.GetWalletRecordListRes, err error) {
	out, err := wallet_record.New().GetWalletRecordList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "操作成功",
		"total": out.Total,
		"data":  out.Rows,
	})
	return nil, nil
}

func (c *Controller) GetWalletRecordTypes(ctx context.Context, req *v1.GetWalletRecordTypesReq) (res v1.GetWalletRecordTypesRes, err error) {
	data, err := wallet_record.New().GetWalletRecordTypes(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": data,
	})
	return nil, nil
}

func (c *Controller) GetWalletStatistics(ctx context.Context, req *v1.GetWalletStatisticsReq) (res *v1.GetWalletStatisticsRes, err error) {
	data, err := wallet_record.New().GetWalletStatistics(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": data,
	})
	return nil, nil
}
