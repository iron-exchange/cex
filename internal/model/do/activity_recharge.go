// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityRecharge is the golang structure of table t_activity_recharge for DAO operations like Where/Data.
type ActivityRecharge struct {
	g.Meta      `orm:"table:t_activity_recharge, do:true"`
	Id          any         // id
	OnOff       any         // 0-关闭 1-开启
	RechargePro any         // 充值返点比例
	MaxRebate   any         // 充值返点最大值
	CreateBy    any         //
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         //
	UpdateTime  *gtime.Time // 更新时间
	SearchValue any         //
}
