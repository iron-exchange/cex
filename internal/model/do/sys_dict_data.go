// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure of table sys_dict_data for DAO operations like Where/Data.
type SysDictData struct {
	g.Meta     `orm:"table:sys_dict_data, do:true"`
	DictCode   any         // 字典编码
	DictSort   any         // 字典排序
	DictLabel  any         // 字典标签
	DictValue  any         // 字典键值
	DictType   any         // 字典类型
	CssClass   any         // 样式属性（其他样式扩展）
	ListClass  any         // 表格回显样式
	IsDefault  any         // 是否默认（Y是 N否）
	ImgUrl     any         // 图片
	Status     any         // 状态（0正常 1停用）
	CreateBy   any         // 创建者
	CreateTime *gtime.Time // 创建时间
	UpdateBy   any         // 更新者
	UpdateTime *gtime.Time // 更新时间
	Remark     any         // 备注
}
