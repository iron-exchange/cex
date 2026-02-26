package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RegisterReq 注册请求结构体
type RegisterReq struct {
	g.Meta     `path:"/user/register" tags:"User" method:"post" summary:"用户注册"`
	Username   string `json:"username" v:"required|length:4,30#请输入账号|账号长度应当在4到30个字符之间" dc:"用户名"`
	Password   string `json:"password" v:"required|length:6,30#请输入密码|密码长度应当在6到30字符之间" dc:"密码"`
	InviteCode string `json:"inviteCode" dc:"邀请码/推广码"` // 通过推荐链接进来时附带
}

// RegisterRes 注册返回结构体
type RegisterRes struct {
	UserId   uint64 `json:"userId" dc:"新创建的用户ID"`
	Username string `json:"username" dc:"用户名"`
	Avatar   string `json:"avatar" dc:"用户头像"`
}

// LoginReq 登录请求结构体
type LoginReq struct {
	g.Meta   `path:"/user/login" tags:"User" method:"post" summary:"用户登录"`
	Username string `json:"username" v:"required#请输入账号" dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
}

// LoginRes 登录返回结构体
type LoginRes struct {
	Token    string `json:"token" dc:"JWT 授权 Token"`
	Expire   string `json:"expire" dc:"Token 过期时间"`
	UserId   uint64 `json:"userId" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
}
