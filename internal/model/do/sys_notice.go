// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNotice is the golang structure of table sys_notice for DAO operations like Where/Data.
type SysNotice struct {
	g.Meta        `orm:"table:sys_notice, do:true"`
	NoticeId      any         // 公告ID
	NoticeTitle   any         // 公告标题
	NoticeType    any         // 公告类型（1通知 2公告）
	NoticeContent any         // 公告内容
	Status        any         // 公告状态（0正常 1关闭）
	CreateBy      any         // 创建者
	CreateTime    *gtime.Time // 创建时间
	UpdateBy      any         // 更新者
	UpdateTime    *gtime.Time // 更新时间
	Remark        any         // 备注
}
