// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DefiActivity is the golang structure of table t_defi_activity for DAO operations like Where/Data.
type DefiActivity struct {
	g.Meta      `orm:"table:t_defi_activity, do:true"`
	Id          any         // id
	TotleAmount any         // 需要金额
	UserId      any         // 用户id
	EndTime     *gtime.Time // 结束时间
	Amount      any         // 奖励金额
	Type        any         // 0-usdt 1-eth
	Status      any         // 0未领取 1已读 2已领取
	CreateBy    any         // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
	SearchValue any         //
}
