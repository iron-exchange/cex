// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost struct {
	PostId     int64       `json:"post_id"     orm:"post_id"     description:"岗位ID"`
	PostCode   string      `json:"post_code"   orm:"post_code"   description:"岗位编码"`
	PostName   string      `json:"post_name"   orm:"post_name"   description:"岗位名称"`
	PostSort   int         `json:"post_sort"   orm:"post_sort"   description:"显示顺序"`
	Status     string      `json:"status"      orm:"status"      description:"状态（0正常 1停用）"`
	CreateBy   string      `json:"create_by"   orm:"create_by"   description:"创建者"`
	CreateTime *gtime.Time `json:"create_time" orm:"create_time" description:"创建时间"`
	UpdateBy   string      `json:"update_by"   orm:"update_by"   description:"更新者"`
	UpdateTime *gtime.Time `json:"update_time" orm:"update_time" description:"更新时间"`
	Remark     string      `json:"remark"      orm:"remark"      description:"备注"`
}
