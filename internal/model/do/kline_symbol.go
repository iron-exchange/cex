// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// KlineSymbol is the golang structure of table t_kline_symbol for DAO operations like Where/Data.
type KlineSymbol struct {
	g.Meta      `orm:"table:t_kline_symbol, do:true"`
	Id          any         // id
	Market      any         // 交易所
	Symbol      any         // 币种简称
	Slug        any         // 币种名称
	Status      any         // 是否开启
	SearchValue any         //
	Logo        any         //
	Remark      any         // 用户备注
	CreateBy    any         //
	UpdateBy    any         //
	UpdateTime  *gtime.Time // 更新时间
	CreateTime  *gtime.Time //
	ReferMarket any         // 参考币种交易所
	ReferCoin   any         // 参考币种
	Proportion  any         // 价格百分比
}
