// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure of table sys_post for DAO operations like Where/Data.
type SysPost struct {
	g.Meta     `orm:"table:sys_post, do:true"`
	PostId     any         // 岗位ID
	PostCode   any         // 岗位编码
	PostName   any         // 岗位名称
	PostSort   any         // 显示顺序
	Status     any         // 状态（0正常 1停用）
	CreateBy   any         // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   any         // 更新者
	UpdateTime *gtime.Time // 更新时间
	Remark     any         // 备注
}
