package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 一对一站内信 (App Mail) ---------
type AdminAppMailInfo struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Type       string `json:"type"`   // 1普通 2全站
	Status     int    `json:"status"` // 0未读 1已读
	CreateTime string `json:"createTime"`
}

type GetAdminAppMailListReq struct {
	g.Meta `path:"/announcement/mail/list" tags:"AdminAnnouncement" method:"get" summary:"获取站内信列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	UserId int64  `json:"userId" dc:"用户ID"`
	Title  string `json:"title" dc:"标题"`
}

type GetAdminAppMailListRes struct {
	List  []AdminAppMailInfo `json:"list"`
	Total int                `json:"total"`
}

// --------- 通知公告 (Notice) ---------
type AdminNoticeInfo struct {
	NoticeId      int    `json:"noticeId"`
	NoticeTitle   string `json:"noticeTitle"`
	NoticeType    string `json:"noticeType"` // 1公告信息 2活动公告 3首页滚动
	NoticeContent string `json:"noticeContent"`
	Cover         string `json:"cover"`
	ViewNum       int    `json:"viewNum"`
	LanguageId    string `json:"languageId"`
	Status        string `json:"status"` // 0正常 1关闭
	Sort          int    `json:"sort"`
	CreateTime    string `json:"createTime"`
}

type GetAdminNoticeListReq struct {
	g.Meta      `path:"/announcement/notice/list" tags:"AdminAnnouncement" method:"get" summary:"获取通知公告列表"`
	Page        int    `json:"page" d:"1"`
	Size        int    `json:"size" d:"20"`
	NoticeTitle string `json:"noticeTitle" dc:"标题"`
	NoticeType  string `json:"noticeType" dc:"类型"`
}

type GetAdminNoticeListRes struct {
	List  []AdminNoticeInfo `json:"list"`
	Total int               `json:"total"`
}

// --------- 前台文本配置 - 基础配置 (Setting) ---------
type AdminSettingInfo struct {
	Id           string `json:"id"`
	SettingValue string `json:"settingValue"`
	CreateTime   string `json:"createTime"`
}

type GetAdminSettingListReq struct {
	g.Meta `path:"/announcement/setting/list" tags:"AdminAnnouncement" method:"get" summary:"获取前台文本配置"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Id     string `json:"id" dc:"配置键名"`
}

type GetAdminSettingListRes struct {
	List  []AdminSettingInfo `json:"list"`
	Total int                `json:"total"`
}

type UpdateAdminSettingReq struct {
	g.Meta       `path:"/announcement/setting/update" tags:"AdminAnnouncement" method:"post" summary:"修改前台基础配置"`
	Id           string `json:"id" v:"required#配置ID不能为空"`
	SettingValue string `json:"settingValue" v:"required#配置内容不能为空"`
}
type UpdateAdminSettingRes struct{}

// --------- 前台文本配置 - 规则说明 (Home Setter) ---------
type AdminHomeSetterInfo struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	ImgUrl       string `json:"imgUrl"`
	Sort         int    `json:"sort"`
	IsShow       int    `json:"isShow"` // 0展示 2不展示
	LanguageName string `json:"languageName"`
	HomeType     int    `json:"homeType"`  // 0首页文本 1问题列表
	ModelType    int    `json:"modelType"` // 0首页 1defi挖矿 2助力贷
	CreateTime   string `json:"createTime"`
}

type GetAdminHomeSetterListReq struct {
	g.Meta       `path:"/announcement/rule/list" tags:"AdminAnnouncement" method:"get" summary:"获取规则说明配置"`
	Page         int    `json:"page" d:"1"`
	Size         int    `json:"size" d:"20"`
	Title        string `json:"title" dc:"标题"`
	LanguageName string `json:"languageName" dc:"语言"`
	ModelType    *int   `json:"modelType" dc:"模块类型"`
}
type GetAdminHomeSetterListRes struct {
	List  []AdminHomeSetterInfo `json:"list"`
	Total int                   `json:"total"`
}

type AddAdminHomeSetterReq struct {
	g.Meta       `path:"/announcement/rule/add" tags:"AdminAnnouncement" method:"post" summary:"添加规则说明"`
	Title        string `json:"title" v:"required#标题不能为空"`
	Content      string `json:"content" v:"required#内容不能为空"`
	LanguageName string `json:"languageName"`
	HomeType     int    `json:"homeType" d:"0"`
	ModelType    int    `json:"modelType" d:"0"`
	Sort         int    `json:"sort" d:"0"`
	IsShow       int    `json:"isShow" d:"0"`
}
type AddAdminHomeSetterRes struct{}

type UpdateAdminHomeSetterReq struct {
	g.Meta       `path:"/announcement/rule/update" tags:"AdminAnnouncement" method:"post" summary:"修改规则说明"`
	Id           int64  `json:"id" v:"required#ID不能为空"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	LanguageName string `json:"languageName"`
	Sort         int    `json:"sort"`
	IsShow       int    `json:"isShow"`
}
type UpdateAdminHomeSetterRes struct{}

// --------- 帮助中心 (Help Center) ---------
type AdminHelpCenterInfo struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Language   string `json:"language"`
	Enable     string `json:"enable"` // 1启用 2禁用
	ShowSymbol string `json:"showSymbol"`
	CreateTime string `json:"createTime"`
}

type GetAdminHelpCenterListReq struct {
	g.Meta `path:"/announcement/help_center/list" tags:"AdminAnnouncement" method:"get" summary:"获取帮助中心分类配置"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Title  string `json:"title" dc:"标题"`
}

type GetAdminHelpCenterListRes struct {
	List  []AdminHelpCenterInfo `json:"list"`
	Total int                   `json:"total"`
}

// 帮助中心详情记录 (Help Center Info)
type AdminHelpCenterArticleInfo struct {
	Id           float64 `json:"id"`
	HelpCenterId int64   `json:"helpCenterId"` // 关联分类ID
	Question     string  `json:"question"`
	Content      string  `json:"content"`
	Language     string  `json:"language"`
	Enable       string  `json:"enable"`
	CreateTime   string  `json:"createTime"`
}

type GetAdminHelpCenterArticleListReq struct {
	g.Meta       `path:"/announcement/help_center/article/list" tags:"AdminAnnouncement" method:"get" summary:"获取帮助中心文章列表"`
	Page         int    `json:"page" d:"1"`
	Size         int    `json:"size" d:"20"`
	HelpCenterId int64  `json:"helpCenterId" dc:"栏目ID"`
	Question     string `json:"question" dc:"问题标题"`
}

type GetAdminHelpCenterArticleListRes struct {
	List  []AdminHelpCenterArticleInfo `json:"list"`
	Total int                          `json:"total"`
}
