// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppMail is the golang structure for table app_mail.
type AppMail struct {
	Id          int64       `json:"id"           orm:"id"           description:""`
	UserId      int64       `json:"user_id"      orm:"user_id"      description:""`
	Title       string      `json:"title"        orm:"title"        description:"标题"`
	Content     string      `json:"content"      orm:"content"      description:"内容"`
	Type        string      `json:"type"         orm:"type"         description:"消息类型 1=普通消息 2=全站消息"`
	Status      int         `json:"status"       orm:"status"       description:"状态（0 未读 1已读）"`
	OpertorId   string      `json:"opertor_id"   orm:"opertor_id"   description:"操作人"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:""`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:""`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
	DelFlag     string      `json:"del_flag"     orm:"del_flag"     description:"0正常 2删除"`
}
