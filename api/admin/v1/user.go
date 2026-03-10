package v1

import "github.com/gogf/gf/v2/frame/g"

type AppUserInfo struct {
	UserId         int64   `json:"userId"`
	LoginName      string  `json:"loginName"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
	Address        string  `json:"address"`
	Status         int     `json:"status"`
	Buff           int     `json:"buff"`
	TotalAmount    float64 `json:"totalAmount"`
	AdminParentIds string  `json:"adminParentIds"`
	CreateTime     string  `json:"createTime"`
}

// GetAppUserListReq 获取玩家用户列表
type GetAppUserListReq struct {
	g.Meta         `path:"/user/list" tags:"AdminUser" method:"get" summary:"获取玩家用户列表"`
	Page           int    `json:"page" d:"1"`
	Size           int    `json:"size" d:"20"`
	LoginName      string `json:"loginName" dc:"按照登录名搜索"`
	Phone          string `json:"phone" dc:"按照手机号搜索"`
	Address        string `json:"address" dc:"按照地址搜索"`
	AdminParentIds string `json:"adminParentIds" dc:"按照代理线过滤"`
}

type GetAppUserListRes struct {
	List  []AppUserInfo `json:"list"`
	Total int           `json:"total"`
}

// FreezeUserReq 冻结/解冻用户
type FreezeUserReq struct {
	g.Meta `path:"/user/freeze" tags:"AdminUser" method:"post" summary:"冻结或解冻玩家"`
	UserId int64 `json:"userId" v:"required#用户ID不能为空"`
	Status int   `json:"status" v:"required|in:0,1#状态只能为 0(正常) 或 1(冻结)"`
}

type FreezeUserRes struct{}

// UpdateUserParentReq 修改用户代理层级
type UpdateUserParentReq struct {
	g.Meta         `path:"/user/updateParent" tags:"AdminUser" method:"post" summary:"修改用户代理线"`
	UserId         int64  `json:"userId" v:"required#用户ID不能为空"`
	AdminParentIds string `json:"adminParentIds" v:"required#代理上级ID不能为空"`
}

type UpdateUserParentRes struct{}
