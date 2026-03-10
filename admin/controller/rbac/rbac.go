package rbac

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/rbac"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// --------- 部门管理 ---------
func (c *Controller) GetDeptList(ctx context.Context, req *v1.GetAdminSysDeptListReq) (res *v1.GetAdminSysDeptListRes, err error) {
	return rbac.New().GetDeptList(ctx, req)
}

// --------- 角色管理 ---------
func (c *Controller) GetRoleList(ctx context.Context, req *v1.GetAdminSysRoleListReq) (res *v1.GetAdminSysRoleListRes, err error) {
	return rbac.New().GetRoleList(ctx, req)
}

// --------- 岗位管理 ---------
func (c *Controller) GetPostList(ctx context.Context, req *v1.GetAdminSysPostListReq) (res *v1.GetAdminSysPostListRes, err error) {
	return rbac.New().GetPostList(ctx, req)
}

// --------- 用户管理 ---------
func (c *Controller) GetUserList(ctx context.Context, req *v1.GetAdminSysUserListReq) (res *v1.GetAdminSysUserListRes, err error) {
	return rbac.New().GetUserList(ctx, req)
}

// --------- 登录日志 ---------
func (c *Controller) GetLogininforList(ctx context.Context, req *v1.GetAdminSysLogininforListReq) (res *v1.GetAdminSysLogininforListRes, err error) {
	return rbac.New().GetLogininforList(ctx, req)
}

// --------- 操作日志 ---------
func (c *Controller) GetOperLogList(ctx context.Context, req *v1.GetAdminSysOperLogListReq) (res *v1.GetAdminSysOperLogListRes, err error) {
	return rbac.New().GetOperLogList(ctx, req)
}
