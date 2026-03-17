package auth

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/auth"
	"GoCEX/internal/logic/common"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Login(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	return auth.New().Login(ctx, req)
}

func (c *Controller) GetInfo(ctx context.Context, req *v1.AdminGetInfoReq) (res *v1.AdminGetInfoRes, err error) {
	out, err := auth.New().GetInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": g.Map{
			"permissions": out.Permissions,
			"roles":       out.Roles,
			"user":        out.User,
		},
	})
	return nil, nil
}

func (c *Controller) GetRouters(ctx context.Context, req *v1.AdminGetRoutersReq) (res *v1.AdminGetRoutersRes, err error) {
	return auth.New().GetRouters(ctx, req)
}

func (c *Controller) CaptchaImage(ctx context.Context, req *v1.AdminCaptchaImageReq) (res *v1.AdminCaptchaImageRes, err error) {
	// 复用 common 已写好的验证码逻辑，做类型转换
	out, err := common.New().CaptchaImage(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.AdminCaptchaImageRes{
		CaptchaEnabled: out.CaptchaEnabled,
		Uuid:           out.Uuid,
		Img:            out.Img,
	}, nil
}

func (c *Controller) GetAllSetting(ctx context.Context, req *v1.AdminGetAllSettingReq) (res v1.AdminGetAllSettingRes, err error) {
	out, err := common.New().GetAllSetting(ctx)
	if err != nil {
		return nil, err
	}

	result := make(v1.AdminGetAllSettingRes)

	// Admin 接口特殊要求：只返回 MARKET_URL 这一个 Key 的结果
	if val, ok := out["MARKET_URL"]; ok {
		result["MARKET_URL"] = val
	}

	return result, nil
}
