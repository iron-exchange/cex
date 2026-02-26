// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpontaneousCoin is the golang structure of table t_spontaneous_coin for DAO operations like Where/Data.
type SpontaneousCoin struct {
	g.Meta      `orm:"table:t_spontaneous_coin, do:true"`
	Id          any         // 主键ID
	Coin        any         // 币种
	Logo        any         // 图标
	ReferCoin   any         // 参考币种
	ReferMarket any         // 参考币种交易所
	ShowSymbol  any         // 展示名称
	Price       any         // 初始价格（单位USDT）
	Proportion  any         // 价格百分比
	CreateBy    any         // 创建人
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
}
