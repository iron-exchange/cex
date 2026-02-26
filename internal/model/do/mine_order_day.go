// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MineOrderDay is the golang structure of table t_mine_order_day for DAO operations like Where/Data.
type MineOrderDay struct {
	g.Meta      `orm:"table:t_mine_order_day, do:true"`
	Id          any         //
	Amount      any         // 投资金额（分）
	Odds        any         // 当日利率
	Earn        any         // 收益
	PlanId      any         // 项目id
	OrderNo     any         // 订单编号
	CreateTime  *gtime.Time // 时间
	Address     any         // 地址
	Type        any         // 0 质押挖矿 1 非质押挖矿
	UpdateTime  *gtime.Time //
	Status      any         // 1 待结算  2  结算
	SearchValue any         //
	UpdateBy    any         //
	CreateBy    any         //
	Remark      any         //
}
