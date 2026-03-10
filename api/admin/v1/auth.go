package v1

import "github.com/gogf/gf/v2/frame/g"

type AdminLoginReq struct {
	g.Meta   `path:"/login" tags:"AdminAuth" method:"post" summary:"管理员登录"`
	Username string `json:"username" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
	Code     string `json:"code" dc:"验证码"`
	Uuid     string `json:"uuid" dc:"验证码唯一标识"`
}

type AdminLoginRes struct {
	Token string `json:"token"`
}

type AdminGetInfoReq struct {
	g.Meta `path:"/getInfo" tags:"AdminAuth" method:"get" summary:"获取当前管理员信息与权限"`
}

type AdminGetInfoRes struct {
	User        g.Map    `json:"user"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type AdminGetRoutersReq struct {
	g.Meta `path:"/getRouters" tags:"AdminAuth" method:"get" summary:"获取动态路由菜单"`
}

type RouterMeta struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	NoCache bool   `json:"noCache"`
	Link    string `json:"link"`
}

type AdminRouterInfo struct {
	Name       string            `json:"name"`
	Path       string            `json:"path"`
	Hidden     bool              `json:"hidden"`
	Redirect   string            `json:"redirect"`
	Component  string            `json:"component"`
	AlwaysShow bool              `json:"alwaysShow"`
	Meta       RouterMeta        `json:"meta"`
	Children   []AdminRouterInfo `json:"children"`
}

type AdminGetRoutersRes struct {
	Data []AdminRouterInfo `json:"data"` // Default array wrapped in data
}
