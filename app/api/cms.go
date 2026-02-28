package v1

import "github.com/gogf/gf/v2/frame/g"

// GetAllNoticeListReq 获取首页轮播与公告
type GetAllNoticeListReq struct {
	g.Meta `path:"/notice/getAllNoticeList" tags:"CMS" method:"get,post" summary:"获取轮播与公告"`
}

type NoticeInfo struct {
	NoticeId      int    `json:"noticeId" dc:"公告ID"`
	NoticeTitle   string `json:"noticeTitle" dc:"标题"`
	NoticeType    string `json:"noticeType" dc:"类型 1公告 2活动 3轮播"`
	NoticeContent string `json:"noticeContent" dc:"内容"`
	Cover         string `json:"cover" dc:"封面图"`
	CreateTime    string `json:"createTime" dc:"创建时间"`
}

type GetAllNoticeListRes struct {
	List []NoticeInfo `json:"list" dc:"公告列表"`
}

// GetHelpCenterListReq 获取帮助中心与教学指南
type GetHelpCenterListReq struct {
	g.Meta `path:"/helpcenter/list" tags:"CMS" method:"get,post" summary:"获取帮助中心指南"`
}

type HelpCenterInfo struct {
	Id         int64  `json:"id" dc:"帮助ID"`
	Title      string `json:"title" dc:"标题"`
	Language   string `json:"language" dc:"语言"`
	CreateTime string `json:"createTime" dc:"创建时间"`
}

type GetHelpCenterListRes struct {
	List []HelpCenterInfo `json:"list" dc:"帮助中心列表"`
}

// GetUserMailReq 获取个人站内信箱
type GetUserMailReq struct {
	g.Meta `path:"/mail/listByUserId" tags:"CMS" method:"get,post" summary:"获取个人站内信"`
	Page   int `json:"page" d:"1" dc:"页码"`
	Size   int `json:"size" d:"20" dc:"每页条数"`
}

type MailInfo struct {
	Id         int64  `json:"id"`
	Title      string `json:"title" dc:"标题"`
	Content    string `json:"content" dc:"内容"`
	Type       string `json:"type" dc:"1普通消息 2全站消息"`
	Status     int    `json:"status" dc:"0未读 1已读"`
	CreateTime string `json:"createTime" dc:"发送时间"`
}

type GetUserMailRes struct {
	List  []MailInfo `json:"list" dc:"站内信列表"`
	Total int        `json:"total" dc:"总数"`
}
