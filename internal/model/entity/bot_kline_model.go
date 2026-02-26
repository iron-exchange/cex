// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BotKlineModel is the golang structure for table bot_kline_model.
type BotKlineModel struct {
	Id            int64       `json:"id"              orm:"id"              description:"id"`
	Decline       int         `json:"decline"         orm:"decline"         description:"最大跌幅"`
	Granularity   int         `json:"granularity"     orm:"granularity"     description:"控制粒度"`
	Increase      int         `json:"increase"        orm:"increase"        description:"最大涨幅"`
	Model         int         `json:"model"           orm:"model"           description:"控盘策略"`
	PricePencent  int         `json:"price_pencent"   orm:"price_pencent"   description:"浮动比例"`
	Symbol        string      `json:"symbol"          orm:"symbol"          description:"交易对"`
	CreateBy      string      `json:"create_by"       orm:"create_by"       description:"创建人"`
	CreateTime    *gtime.Time `json:"create_time"     orm:"create_time"     description:"创建时间"`
	UpdateBy      string      `json:"update_by"       orm:"update_by"       description:"修改人"`
	UpdateTime    *gtime.Time `json:"update_time"     orm:"update_time"     description:"更新时间"`
	SearchValue   string      `json:"search_value"    orm:"search_value"    description:"值"`
	BeginTime     *gtime.Time `json:"begin_time"      orm:"begin_time"      description:"开始时间"`
	EndTime       *gtime.Time `json:"end_time"        orm:"end_time"        description:"结束时间"`
	LineChartData string      `json:"line_chart_data" orm:"line_chart_data" description:""`
	Remark        string      `json:"remark"          orm:"remark"          description:""`
	ConPrice      float64     `json:"con_price"       orm:"con_price"       description:""`
}
