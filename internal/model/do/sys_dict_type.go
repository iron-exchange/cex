// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure of table sys_dict_type for DAO operations like Where/Data.
type SysDictType struct {
	g.Meta     `orm:"table:sys_dict_type, do:true"`
	DictId     any         // 字典主键
	DictName   any         // 字典名称
	DictType   any         // 字典类型
	Status     any         // 状态（0正常 1停用）
	CreateBy   any         // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   any         // 更新者
	UpdateTime *gtime.Time // 更新时间
	Remark     any         // 备注
}
