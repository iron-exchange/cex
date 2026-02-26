// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DefiOrder is the golang structure of table t_defi_order for DAO operations like Where/Data.
type DefiOrder struct {
	g.Meta         `orm:"table:t_defi_order, do:true"`
	Id             any         // id
	Amount         any         // 收益金额
	TotleAmount    any         // 钱包金额
	Rate           any         // 收益率
	CreateBy       any         // 创建者
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       any         // 更新者
	UpdateTime     *gtime.Time // 更新时间
	Remark         any         // 备注
	SearchValue    any         //
	UserId         any         // 用户id
	AdminParentIds any         // 代理ids
}
