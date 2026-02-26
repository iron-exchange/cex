// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta      `orm:"table:sys_config, do:true"`
	ConfigId    any         // 参数主键
	ConfigName  any         // 参数名称
	ConfigKey   any         // 参数键名
	ConfigValue any         // 参数键值
	ConfigType  any         // 系统内置（Y是 N否）
	CreateBy    any         // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
}
