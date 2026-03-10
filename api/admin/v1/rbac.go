package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 部门管理 (Sys Dept) ---------
type AdminSysDeptInfo struct {
	DeptId     int64  `json:"deptId"`
	ParentId   int64  `json:"parentId"`
	DeptName   string `json:"deptName"`
	OrderNum   int    `json:"orderNum"`
	Leader     string `json:"leader"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Status     string `json:"status"` // 0正常 1停用
	CreateTime string `json:"createTime"`
}

type GetAdminSysDeptListReq struct {
	g.Meta   `path:"/rbac/dept/list" tags:"AdminRBAC" method:"get" summary:"获取部门列表"`
	DeptName string `json:"deptName" dc:"部门名称"`
	Status   string `json:"status" dc:"状态"`
}

type GetAdminSysDeptListRes struct {
	List []AdminSysDeptInfo `json:"list"`
}

// --------- 角色管理 (Sys Role) ---------
type AdminSysRoleInfo struct {
	RoleId     int64  `json:"roleId"`
	RoleName   string `json:"roleName"`
	RoleKey    string `json:"roleKey"`
	RoleSort   int    `json:"roleSort"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
	Remark     string `json:"remark"`
}

type GetAdminSysRoleListReq struct {
	g.Meta   `path:"/rbac/role/list" tags:"AdminRBAC" method:"get" summary:"获取角色列表"`
	Page     int    `json:"page" d:"1"`
	Size     int    `json:"size" d:"20"`
	RoleName string `json:"roleName" dc:"角色名称"`
	RoleKey  string `json:"roleKey" dc:"权限字符"`
	Status   string `json:"status" dc:"状态"`
}

type GetAdminSysRoleListRes struct {
	List  []AdminSysRoleInfo `json:"list"`
	Total int                `json:"total"`
}

// --------- 岗位管理 (Sys Post) ---------
type AdminSysPostInfo struct {
	PostId     int64  `json:"postId"`
	PostCode   string `json:"postCode"`
	PostName   string `json:"postName"`
	PostSort   int    `json:"postSort"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
	Remark     string `json:"remark"`
}

type GetAdminSysPostListReq struct {
	g.Meta   `path:"/rbac/post/list" tags:"AdminRBAC" method:"get" summary:"获取岗位列表"`
	Page     int    `json:"page" d:"1"`
	Size     int    `json:"size" d:"20"`
	PostCode string `json:"postCode" dc:"岗位编码"`
	PostName string `json:"postName" dc:"岗位名称"`
	Status   string `json:"status" dc:"状态"`
}

type GetAdminSysPostListRes struct {
	List  []AdminSysPostInfo `json:"list"`
	Total int                `json:"total"`
}

// --------- 用户管理 (Sys User) ---------
type AdminSysUserInfo struct {
	UserId      int64  `json:"userId"`
	DeptId      int64  `json:"deptId"`
	UserName    string `json:"userName"`
	NickName    string `json:"nickName"`
	UserType    string `json:"userType"` // 0普通用户 1组长 2代理
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Sex         string `json:"sex"`
	Status      string `json:"status"`
	LoginIp     string `json:"loginIp"`
	LoginDate   string `json:"loginDate"`
	CreateTime  string `json:"createTime"`
}

type GetAdminSysUserListReq struct {
	g.Meta      `path:"/rbac/user/list" tags:"AdminRBAC" method:"get" summary:"获取后台用户列表"`
	Page        int    `json:"page" d:"1"`
	Size        int    `json:"size" d:"20"`
	UserName    string `json:"userName" dc:"用户账号"`
	Phonenumber string `json:"phonenumber" dc:"手机号码"`
	Status      string `json:"status" dc:"状态"`
	DeptId      int64  `json:"deptId" dc:"部门ID"`
}

type GetAdminSysUserListRes struct {
	List  []AdminSysUserInfo `json:"list"`
	Total int                `json:"total"`
}

// --------- 登录日志 (Sys Logininfor) ---------
type AdminSysLogininforInfo struct {
	InfoId        int64  `json:"infoId"`
	UserName      string `json:"userName"`
	Ipaddr        string `json:"ipaddr"`
	LoginLocation string `json:"loginLocation"`
	Browser       string `json:"browser"`
	Os            string `json:"os"`
	Status        string `json:"status"` // 0成功 1失败
	Msg           string `json:"msg"`
	LoginTime     string `json:"loginTime"`
}

type GetAdminSysLogininforListReq struct {
	g.Meta   `path:"/rbac/logininfor/list" tags:"AdminRBAC" method:"get" summary:"获取登录日志"`
	Page     int    `json:"page" d:"1"`
	Size     int    `json:"size" d:"20"`
	Ipaddr   string `json:"ipaddr" dc:"登录IP地址"`
	UserName string `json:"userName" dc:"用户账号"`
	Status   string `json:"status" dc:"状态"`
}

type GetAdminSysLogininforListRes struct {
	List  []AdminSysLogininforInfo `json:"list"`
	Total int                      `json:"total"`
}

// --------- 操作日志 (Sys Oper Log) ---------
type AdminSysOperLogInfo struct {
	OperId        int64  `json:"operId"`
	Title         string `json:"title"`
	BusinessType  int    `json:"businessType"` // 0其它 1新增 2修改 3删除
	Method        string `json:"method"`
	RequestMethod string `json:"requestMethod"`
	OperName      string `json:"operName"`
	DeptName      string `json:"deptName"`
	OperUrl       string `json:"operUrl"`
	OperIp        string `json:"operIp"`
	OperLocation  string `json:"operLocation"`
	Status        int    `json:"status"` // 0正常 1异常
	OperTime      string `json:"operTime"`
	CostTime      int64  `json:"costTime"`
}

type GetAdminSysOperLogListReq struct {
	g.Meta       `path:"/rbac/operlog/list" tags:"AdminRBAC" method:"get" summary:"获取操作日志"`
	Page         int    `json:"page" d:"1"`
	Size         int    `json:"size" d:"20"`
	Title        string `json:"title" dc:"模块标题"`
	OperName     string `json:"operName" dc:"操作人员"`
	Status       *int   `json:"status" dc:"操作状态"`
	BusinessType *int   `json:"businessType" dc:"业务类型"`
}

type GetAdminSysOperLogListRes struct {
	List  []AdminSysOperLogInfo `json:"list"`
	Total int                   `json:"total"`
}
