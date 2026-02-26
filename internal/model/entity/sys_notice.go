// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNotice is the golang structure for table sys_notice.
type SysNotice struct {
	NoticeId      int         `json:"notice_id"      orm:"notice_id"      description:"公告ID"`
	NoticeTitle   string      `json:"notice_title"   orm:"notice_title"   description:"公告标题"`
	NoticeType    string      `json:"notice_type"    orm:"notice_type"    description:"公告类型（1通知 2公告）"`
	NoticeContent string      `json:"notice_content" orm:"notice_content" description:"公告内容"`
	Status        string      `json:"status"         orm:"status"         description:"公告状态（0正常 1关闭）"`
	CreateBy      string      `json:"create_by"      orm:"create_by"      description:"创建者"`
	CreateTime    *gtime.Time `json:"create_time"    orm:"create_time"    description:"创建时间"`
	UpdateBy      string      `json:"update_by"      orm:"update_by"      description:"更新者"`
	UpdateTime    *gtime.Time `json:"update_time"    orm:"update_time"    description:"更新时间"`
	Remark        string      `json:"remark"         orm:"remark"         description:"备注"`
}
