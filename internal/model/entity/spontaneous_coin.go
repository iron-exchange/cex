// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpontaneousCoin is the golang structure for table spontaneous_coin.
type SpontaneousCoin struct {
	Id          int64       `json:"id"           orm:"id"           description:"主键ID"`
	Coin        string      `json:"coin"         orm:"coin"         description:"币种"`
	Logo        string      `json:"logo"         orm:"logo"         description:"图标"`
	ReferCoin   string      `json:"refer_coin"   orm:"refer_coin"   description:"参考币种"`
	ReferMarket string      `json:"refer_market" orm:"refer_market" description:"参考币种交易所"`
	ShowSymbol  string      `json:"show_symbol"  orm:"show_symbol"  description:"展示名称"`
	Price       float64     `json:"price"        orm:"price"        description:"初始价格（单位USDT）"`
	Proportion  float64     `json:"proportion"   orm:"proportion"   description:"价格百分比"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:"创建人"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:"更新者"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`
}
