// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// KlineSymbol is the golang structure for table kline_symbol.
type KlineSymbol struct {
	Id          int64       `json:"id"           orm:"id"           description:"id"`
	Market      string      `json:"market"       orm:"market"       description:"交易所"`
	Symbol      string      `json:"symbol"       orm:"symbol"       description:"币种简称"`
	Slug        string      `json:"slug"         orm:"slug"         description:"币种名称"`
	Status      int         `json:"status"       orm:"status"       description:"是否开启"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
	Logo        string      `json:"logo"         orm:"logo"         description:""`
	Remark      string      `json:"remark"       orm:"remark"       description:"用户备注"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:""`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:""`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:""`
	ReferMarket string      `json:"refer_market" orm:"refer_market" description:"参考币种交易所"`
	ReferCoin   string      `json:"refer_coin"   orm:"refer_coin"   description:"参考币种"`
	Proportion  float64     `json:"proportion"   orm:"proportion"   description:"价格百分比"`
}
