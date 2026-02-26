// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Activity is the golang structure for table activity.
type Activity struct {
	Id           int         `json:"id"            orm:"id"            description:"卡ID"`
	CreateBy     string      `json:"create_by"     orm:"create_by"     description:""`
	CreateTime   *gtime.Time `json:"create_time"   orm:"create_time"   description:"创建时间"`
	UpdateBy     string      `json:"update_by"     orm:"update_by"     description:""`
	UpdateTime   *gtime.Time `json:"update_time"   orm:"update_time"   description:"更新时间"`
	SearchValue  string      `json:"search_value"  orm:"search_value"  description:""`
	Status       int         `json:"status"        orm:"status"        description:"状态0编辑1上线"`
	Remark       string      `json:"remark"        orm:"remark"        description:""`
	DisplayUrl   string      `json:"display_url"   orm:"display_url"   description:""`
	DetailUrl    string      `json:"detail_url"    orm:"detail_url"    description:"详情页"`
	DetailStatus int         `json:"detail_status" orm:"detail_status" description:""`
	LoadingUrl   string      `json:"loading_url"   orm:"loading_url"   description:""`
	Tags         string      `json:"tags"          orm:"tags"          description:""`
	Type         int         `json:"type"          orm:"type"          description:"1广告2活动"`
	Source       int         `json:"source"        orm:"source"        description:"展示端,1:PC,2:H5"`
	LanguageId   int         `json:"language_id"   orm:"language_id"   description:"zh:1,cht:2,en:3,pt:4,sa:5,ko:6,ja:7,es:8,th:9,ms:10,id:11,fr:12,ru:13"`
	Agent        int         `json:"agent"         orm:"agent"         description:"0代理无关，1相关"`
	JumpType     int         `json:"jump_type"     orm:"jump_type"     description:"跳转类型 0 内链 1外链"`
	SortNum      int         `json:"sort_num"      orm:"sort_num"      description:"排序"`
}
