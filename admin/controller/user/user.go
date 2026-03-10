package user

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/user"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// GetAppUserList 玩家列表查询
func (c *Controller) GetAppUserList(ctx context.Context, req *v1.GetAppUserListReq) (res *v1.GetAppUserListRes, err error) {
	return user.New().GetAppUserList(ctx, req)
}

// FreezeUser 冻结/解冻玩家
func (c *Controller) FreezeUser(ctx context.Context, req *v1.FreezeUserReq) (res *v1.FreezeUserRes, err error) {
	err = user.New().FreezeUser(ctx, req)
	return &v1.FreezeUserRes{}, err
}

// UpdateUserParent 修改代理层级
func (c *Controller) UpdateUserParent(ctx context.Context, req *v1.UpdateUserParentReq) (res *v1.UpdateUserParentRes, err error) {
	err = user.New().UpdateParent(ctx, req)
	return &v1.UpdateUserParentRes{}, err
}
