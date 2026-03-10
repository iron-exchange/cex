package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginLogInfo struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"userId"`
	Ipaddr        string `json:"ipaddr"`
	LoginLocation string `json:"loginLocation"`
	Browser       string `json:"browser"`
	Os            string `json:"os"`
	Status        string `json:"status"`
	Msg           string `json:"msg"`
	LoginTime     string `json:"loginTime"`
}

// GetLoginLogListReq 获取登录日志列表
type GetLoginLogListReq struct {
	g.Meta `path:"/loginLog/list" tags:"AdminLog" method:"get" summary:"获取玩家登录日志"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	UserId int64  `json:"userId" dc:"按用户ID搜索"`
	Ipaddr string `json:"ipaddr" dc:"按IP地址搜索"`
}

type GetLoginLogListRes struct {
	List  []LoginLogInfo `json:"list"`
	Total int            `json:"total"`
}
