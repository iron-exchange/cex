// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BotKlineModel is the golang structure of table t_bot_kline_model for DAO operations like Where/Data.
type BotKlineModel struct {
	g.Meta        `orm:"table:t_bot_kline_model, do:true"`
	Id            any         // id
	Decline       any         // 最大跌幅
	Granularity   any         // 控制粒度
	Increase      any         // 最大涨幅
	Model         any         // 控盘策略
	PricePencent  any         // 浮动比例
	Symbol        any         // 交易对
	CreateBy      any         // 创建人
	CreateTime    *gtime.Time // 创建时间
	UpdateBy      any         // 修改人
	UpdateTime    *gtime.Time // 更新时间
	SearchValue   any         // 值
	BeginTime     *gtime.Time // 开始时间
	EndTime       *gtime.Time // 结束时间
	LineChartData any         //
	Remark        any         //
	ConPrice      any         //
}
