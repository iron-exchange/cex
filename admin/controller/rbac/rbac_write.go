package rbac

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/rbac"
)

// --------- 菜单管理 (Sys Menu Write) ---------
func (c *Controller) AddMenu(ctx context.Context, req *v1.AddAdminSysMenuReq) (res *v1.AddAdminSysMenuRes, err error) {
	return rbac.New().AddMenu(ctx, req)
}
func (c *Controller) EditMenu(ctx context.Context, req *v1.EditAdminSysMenuReq) (res *v1.EditAdminSysMenuRes, err error) {
	return rbac.New().EditMenu(ctx, req)
}
func (c *Controller) DeleteMenu(ctx context.Context, req *v1.DeleteAdminSysMenuReq) (res *v1.DeleteAdminSysMenuRes, err error) {
	return rbac.New().DeleteMenu(ctx, req)
}

// --------- 角色管理 (Sys Role Write) ---------
func (c *Controller) AddRole(ctx context.Context, req *v1.AddAdminSysRoleReq) (res *v1.AddAdminSysRoleRes, err error) {
	return rbac.New().AddRole(ctx, req)
}
func (c *Controller) EditRole(ctx context.Context, req *v1.EditAdminSysRoleReq) (res *v1.EditAdminSysRoleRes, err error) {
	return rbac.New().EditRole(ctx, req)
}
func (c *Controller) DeleteRole(ctx context.Context, req *v1.DeleteAdminSysRoleReq) (res *v1.DeleteAdminSysRoleRes, err error) {
	return rbac.New().DeleteRole(ctx, req)
}

// --------- 用户管理 (Sys User Write) ---------
func (c *Controller) AddUser(ctx context.Context, req *v1.AddAdminSysUserReq) (res *v1.AddAdminSysUserRes, err error) {
	return rbac.New().AddUser(ctx, req)
}
func (c *Controller) EditUser(ctx context.Context, req *v1.EditAdminSysUserReq) (res *v1.EditAdminSysUserRes, err error) {
	return rbac.New().EditUser(ctx, req)
}
func (c *Controller) DeleteUser(ctx context.Context, req *v1.DeleteAdminSysUserReq) (res *v1.DeleteAdminSysUserRes, err error) {
	return rbac.New().DeleteUser(ctx, req)
}
func (c *Controller) ResetUserPwd(ctx context.Context, req *v1.ResetAdminSysUserPwdReq) (res *v1.ResetAdminSysUserPwdRes, err error) {
	return rbac.New().ResetUserPwd(ctx, req)
}
