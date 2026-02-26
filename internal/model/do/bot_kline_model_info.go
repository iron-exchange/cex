// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BotKlineModelInfo is the golang structure of table t_bot_kline_model_info for DAO operations like Where/Data.
type BotKlineModelInfo struct {
	g.Meta   `orm:"table:t_bot_kline_model_info, do:true"`
	Id       any // id
	ModelId  any // t_bot_kline_model 的主键
	DateTime any // 时间戳
	Open     any // 开盘价
	Close    any // 封盘价
	High     any // 最高价
	Low      any // 最低价
	X        any // x轴
	Y        any // y轴
}
