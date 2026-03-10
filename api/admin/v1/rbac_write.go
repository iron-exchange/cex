package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 菜单管理 (Sys Menu Write) ---------
type AddAdminSysMenuReq struct {
	g.Meta    `path:"/rbac/menu" tags:"AdminRBAC" method:"post" summary:"新增菜单"`
	ParentId  int64  `json:"parentId"`
	MenuName  string `json:"menuName" v:"required"`
	OrderNum  int    `json:"orderNum"`
	Path      string `json:"path"`
	Component string `json:"component"`
	IsFrame   int    `json:"isFrame"`  // 1是 0否
	IsCache   int    `json:"isCache"`  // 0缓存 1不缓存
	MenuType  string `json:"menuType"` // M目录 C菜单 F按钮
	Visible   string `json:"visible"`  // 0显示 1隐藏
	Status    string `json:"status"`   // 0正常 1停用
	Perms     string `json:"perms"`
	Icon      string `json:"icon"`
}
type AddAdminSysMenuRes struct{}

type EditAdminSysMenuReq struct {
	g.Meta    `path:"/rbac/menu" tags:"AdminRBAC" method:"put" summary:"修改菜单"`
	MenuId    int64  `json:"menuId" v:"required"`
	ParentId  int64  `json:"parentId"`
	MenuName  string `json:"menuName" v:"required"`
	OrderNum  int    `json:"orderNum"`
	Path      string `json:"path"`
	Component string `json:"component"`
	IsFrame   int    `json:"isFrame"`
	IsCache   int    `json:"isCache"`
	MenuType  string `json:"menuType"`
	Visible   string `json:"visible"`
	Status    string `json:"status"`
	Perms     string `json:"perms"`
	Icon      string `json:"icon"`
}
type EditAdminSysMenuRes struct{}

type DeleteAdminSysMenuReq struct {
	g.Meta `path:"/rbac/menu/{menuId}" tags:"AdminRBAC" method:"delete" summary:"删除菜单"`
	MenuId int64 `json:"menuId" in:"path"`
}
type DeleteAdminSysMenuRes struct{}

// --------- 角色管理 (Sys Role Write) ---------
type AddAdminSysRoleReq struct {
	g.Meta   `path:"/rbac/role" tags:"AdminRBAC" method:"post" summary:"新增角色"`
	RoleName string  `json:"roleName" v:"required"`
	RoleKey  string  `json:"roleKey" v:"required"`
	RoleSort int     `json:"roleSort" v:"required"`
	Status   string  `json:"status"`
	Remark   string  `json:"remark"`
	MenuIds  []int64 `json:"menuIds" dc:"关联菜单组"`
}
type AddAdminSysRoleRes struct{}

type EditAdminSysRoleReq struct {
	g.Meta   `path:"/rbac/role" tags:"AdminRBAC" method:"put" summary:"修改角色"`
	RoleId   int64   `json:"roleId" v:"required"`
	RoleName string  `json:"roleName" v:"required"`
	RoleKey  string  `json:"roleKey" v:"required"`
	RoleSort int     `json:"roleSort" v:"required"`
	Status   string  `json:"status"`
	Remark   string  `json:"remark"`
	MenuIds  []int64 `json:"menuIds"`
}
type EditAdminSysRoleRes struct{}

type DeleteAdminSysRoleReq struct {
	g.Meta  `path:"/rbac/role/{roleIds}" tags:"AdminRBAC" method:"delete" summary:"删除角色"`
	RoleIds string `json:"roleIds" in:"path" dc:"多个用逗号隔开"`
}
type DeleteAdminSysRoleRes struct{}

// --------- 用户管理 (Sys User Write) ---------
type AddAdminSysUserReq struct {
	g.Meta      `path:"/rbac/user" tags:"AdminRBAC" method:"post" summary:"新增后台用户"`
	DeptId      int64   `json:"deptId"`
	UserName    string  `json:"userName" v:"required"`
	NickName    string  `json:"nickName" v:"required"`
	Password    string  `json:"password" v:"required"`
	Phonenumber string  `json:"phonenumber"`
	Email       string  `json:"email"`
	Sex         string  `json:"sex"`
	Status      string  `json:"status"`
	RoleIds     []int64 `json:"roleIds" dc:"关联角色组"`
}
type AddAdminSysUserRes struct{}

type EditAdminSysUserReq struct {
	g.Meta      `path:"/rbac/user" tags:"AdminRBAC" method:"put" summary:"修改后台用户"`
	UserId      int64   `json:"userId" v:"required"`
	DeptId      int64   `json:"deptId"`
	UserName    string  `json:"userName" v:"required"`
	NickName    string  `json:"nickName" v:"required"`
	Phonenumber string  `json:"phonenumber"`
	Email       string  `json:"email"`
	Sex         string  `json:"sex"`
	Status      string  `json:"status"`
	RoleIds     []int64 `json:"roleIds"`
}
type EditAdminSysUserRes struct{}

type DeleteAdminSysUserReq struct {
	g.Meta  `path:"/rbac/user/{userIds}" tags:"AdminRBAC" method:"delete" summary:"删除后台用户"`
	UserIds string `json:"userIds" in:"path" dc:"多个用逗号隔开"`
}
type DeleteAdminSysUserRes struct{}

type ResetAdminSysUserPwdReq struct {
	g.Meta   `path:"/rbac/user/resetPwd" tags:"AdminRBAC" method:"put" summary:"重置用户密码"`
	UserId   int64  `json:"userId" v:"required"`
	Password string `json:"password" v:"required"`
}
type ResetAdminSysUserPwdRes struct{}
