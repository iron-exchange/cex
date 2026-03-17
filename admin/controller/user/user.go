package user

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/user"

	"github.com/gogf/gf/v2/frame/g"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetAppUserList 玩家列表查询
func (c *Controller) GetAppUserList(ctx context.Context, req *v1.GetAppUserListReq) (res *v1.GetAppUserListRes, err error) {
	data, err := user.New().GetAppUserList(ctx, req)
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
	return nil, nil // 返回 nil 绕过中间件包装
}

// FreezeUser 冻结/解冻玩家
func (c *Controller) FreezeUser(ctx context.Context, req *v1.FreezeUserReq) (res *v1.FreezeUserRes, err error) {
	err = user.New().FreezeUser(ctx, req)
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

// UpdateUserAppIds 修改玩家归属代理
func (c *Controller) UpdateUserAppIds(ctx context.Context, req *v1.UpdateUserAppIdsReq) (res *v1.UpdateUserAppIdsRes, err error) {
	err = user.New().UpdateUserAppIds(ctx, req)
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

// UpdateUser 修改玩家基本信息和状态
func (c *Controller) UpdateUser(ctx context.Context, req *v1.UpdateAppUserReq) (res *v1.UpdateAppUserRes, err error) {
	err = user.New().UpdateUser(ctx, req)
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

// SubAmount 人工上下分
func (c *Controller) SubAmount(ctx context.Context, req *v1.SubUserAmountReq) (res *v1.SubUserAmountRes, err error) {
	err = user.New().SubAmount(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// SendBonus 赠送彩金/扣减彩金
func (c *Controller) SendBonus(ctx context.Context, req *v1.SendBonusReq) (res *v1.SendBonusRes, err error) {
	err = user.New().SendBonus(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// AuditUserRealName 审核玩家实名认证(通过/拒绝)
func (c *Controller) AuditUserRealName(ctx context.Context, req *v1.AuditUserRealNameReq) (res *v1.AuditUserRealNameRes, err error) {
	err = user.New().AuditUserRealName(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// ResetUserRealName 重置实名认证(打回原形)
func (c *Controller) ResetUserRealName(ctx context.Context, req *v1.ResetUserRealNameReq) (res *v1.ResetUserRealNameRes, err error) {
	err = user.New().ResetUserRealName(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// UpdateUserRealName 擦除/清理关联的主表实名数据
func (c *Controller) UpdateUserRealName(ctx context.Context, req *v1.UpdateUserRealNameReq) (res *v1.UpdateUserRealNameRes, err error) {
	err = user.New().UpdateUserRealName(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}
