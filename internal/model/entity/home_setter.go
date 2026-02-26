// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HomeSetter is the golang structure for table home_setter.
type HomeSetter struct {
	Id           int64       `json:"id"            orm:"id"            description:""`
	Title        string      `json:"title"         orm:"title"         description:"标题"`
	Author       string      `json:"author"        orm:"author"        description:"作者"`
	Content      string      `json:"content"       orm:"content"       description:"内容"`
	CreateTime   *gtime.Time `json:"create_time"   orm:"create_time"   description:"创建时间"`
	ImgUrl       string      `json:"img_url"       orm:"img_url"       description:"图片地址"`
	Sort         int         `json:"sort"          orm:"sort"          description:"排序"`
	IsShow       int         `json:"is_show"       orm:"is_show"       description:"是否展示 0展示  2不展示"`
	LanguageName string      `json:"language_name" orm:"language_name" description:"语言"`
	LikesNum     int         `json:"likes_num"     orm:"likes_num"     description:"点赞数"`
	HomeType     int         `json:"home_type"     orm:"home_type"     description:"类型（0 首页文本  1 问题列表）"`
	ModelType    int         `json:"model_type"    orm:"model_type"    description:"功能（0=首页  1=defi挖矿 2=助力贷）"`
	SearchValue  string      `json:"search_value"  orm:"search_value"  description:""`
}
