// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DefiRate is the golang structure of table t_defi_rate for DAO operations like Where/Data.
type DefiRate struct {
	g.Meta      `orm:"table:t_defi_rate, do:true"`
	Id          any         // id
	MinAmount   any         // 最小金额
	MaxAmount   any         // 最大金额
	Rate        any         // 利率
	CreateBy    any         // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
	SearchValue any         //
}
