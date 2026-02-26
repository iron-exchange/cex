// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Markets is the golang structure of table t_markets for DAO operations like Where/Data.
type Markets struct {
	g.Meta      `orm:"table:t_markets, do:true"`
	Id          any //
	Slug        any // 交易所名称(ID)
	Fullname    any // 交易所全称
	WebsiteUrl  any // 交易所官网链接
	Status      any // 状态: [enable, disable]. disable为停止更新数据
	Kline       any // 是否接入K线数据
	Spot        any // 是否支持现货
	Futures     any // 是否支持期货
	SearchValue any //
}
