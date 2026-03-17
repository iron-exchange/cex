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

type AdminLogoutReq struct {
	g.Meta `path:"/logout" tags:"AdminAuth" method:"post" summary:"管理员登出"`
}

type AdminLogoutRes struct{}

type AdminGetInfoReq struct {
	g.Meta `path:"/getInfo" tags:"AdminAuth" method:"get" summary:"获取当前管理员信息与权限"`
}

type AdminGetInfoRes struct {
	User        AdminInfoUser `json:"user"`
	Roles       []string      `json:"roles"`
	Permissions []string      `json:"permissions"`
}

type AdminInfoUser struct {
	CreateBy    interface{}     `json:"createBy"`
	CreateTime  string          `json:"createTime"`
	UpdateBy    interface{}     `json:"updateBy"`
	UpdateTime  interface{}     `json:"updateTime"`
	Remark      interface{}     `json:"remark"`
	UserId      int64           `json:"userId"`
	DeptId      int64           `json:"deptId"`
	UserName    string          `json:"userName"`
	UserType    string          `json:"userType"`
	NickName    string          `json:"nickName"`
	Email       string          `json:"email"`
	Phonenumber string          `json:"phonenumber"`
	Sex         string          `json:"sex"`
	Avatar      string          `json:"avatar"`
	Password    string          `json:"password"`
	Status      string          `json:"status"`
	DelFlag     string          `json:"delFlag"`
	LoginIp     string          `json:"loginIp"`
	LoginDate   string          `json:"loginDate"`
	ParentId    interface{}     `json:"parentId"`
	Dept        AdminInfoDept   `json:"dept"`
	Roles       []AdminInfoRole `json:"roles"`
	RoleIds     interface{}     `json:"roleIds"`
	PostIds     interface{}     `json:"postIds"`
	RoleId      interface{}     `json:"roleId"`
	GoogleKey   string          `json:"googleKey"`
	Admin       bool            `json:"admin"`
}

type AdminInfoDept struct {
	CreateBy   interface{}     `json:"createBy"`
	CreateTime interface{}     `json:"createTime"`
	UpdateBy   interface{}     `json:"updateBy"`
	UpdateTime interface{}     `json:"updateTime"`
	Remark     interface{}     `json:"remark"`
	DeptId     int64           `json:"deptId"`
	ParentId   interface{}     `json:"parentId"`
	Ancestors  string          `json:"ancestors"`
	DeptName   string          `json:"deptName"`
	OrderNum   int             `json:"orderNum"`
	Leader     string          `json:"leader"`
	Phone      interface{}     `json:"phone"`
	Email      interface{}     `json:"email"`
	Status     string          `json:"status"`
	DelFlag    interface{}     `json:"delFlag"`
	ParentName interface{}     `json:"parentName"`
	Children   []AdminInfoDept `json:"children"`
}

type AdminInfoRole struct {
	CreateBy          interface{} `json:"createBy"`
	CreateTime        interface{} `json:"createTime"`
	UpdateBy          interface{} `json:"updateBy"`
	UpdateTime        interface{} `json:"updateTime"`
	Remark            interface{} `json:"remark"`
	RoleId            int64       `json:"roleId"`
	RoleName          string      `json:"roleName"`
	RoleKey           string      `json:"roleKey"`
	RoleSort          int         `json:"roleSort"`
	DataScope         string      `json:"dataScope"`
	MenuCheckStrictly bool        `json:"menuCheckStrictly"`
	DeptCheckStrictly bool        `json:"deptCheckStrictly"`
	Status            string      `json:"status"`
	DelFlag           interface{} `json:"delFlag"`
	Flag              bool        `json:"flag"`
	MenuIds           interface{} `json:"menuIds"`
	DeptIds           interface{} `json:"deptIds"`
	Permissions       interface{} `json:"permissions"`
	Admin             bool        `json:"admin"`
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

type AdminGetRoutersRes []AdminRouterInfo

// AdminGetAllSettingReq 获取全站配置参数
type AdminGetAllSettingReq struct {
	g.Meta `path:"/common/getAllSetting" tags:"AdminAuth" method:"post,get" summary:"获取全站配置参数"`
}

// 对应原始 Java 返回的 data: { "MARKET_URL": { "url": false, ... }, "SOME_KEY": "string_val" }
type AdminGetAllSettingRes map[string]interface{}

// AdminCaptchaImageReq 生成图形验证码
type AdminCaptchaImageReq struct {
	g.Meta `path:"/captchaImage" tags:"AdminAuth" method:"get" summary:"生成验证码"`
}

type AdminCaptchaImageRes struct {
	CaptchaEnabled bool   `json:"captchaEnabled" dc:"是否开启验证码"`
	Uuid           string `json:"uuid" dc:"验证码唯一ID"`
	Img            string `json:"img" dc:"Base64 验证码图片"`
}
