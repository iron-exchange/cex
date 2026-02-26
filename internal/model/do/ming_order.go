// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MingOrder is the golang structure of table t_ming_order for DAO operations like Where/Data.
type MingOrder struct {
	g.Meta          `orm:"table:t_ming_order, do:true"`
	Id              any         //
	Amount          any         // 投资金额（分）
	Days            any         // 投资期限（天）
	Status          any         // 0 收益  1 结算
	PlanId          any         // 项目id
	PlanTitle       any         // 项目名称
	OrderNo         any         // 订单编号
	CreateTime      *gtime.Time // 投资时间
	EndTime         *gtime.Time // 到期时间
	SettleTime      *gtime.Time // 结算时间
	AccumulaEarn    any         // 累计收益
	UpdateTime      *gtime.Time //
	MinOdds         any         // 最小利率
	MaxOdds         any         // 最大利率
	AdminUserIds    any         // 后台用户id
	CollectionOrder any         //
	UserId          any         //
	OrderAmount     any         //
}
