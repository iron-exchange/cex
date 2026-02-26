// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SecondCoinConfig is the golang structure for table second_coin_config.
type SecondCoinConfig struct {
	Id          int64       `json:"id"           orm:"id"           description:"id"`
	Symbol      string      `json:"symbol"       orm:"symbol"       description:"合约交易对"`
	Market      string      `json:"market"       orm:"market"       description:"所属交易所"`
	Status      int         `json:"status"       orm:"status"       description:"是否启用 2关闭 1启用"`
	ShowFlag    int         `json:"show_flag"    orm:"show_flag"    description:"是否展示 2不展示 1展示"`
	Coin        string      `json:"coin"         orm:"coin"         description:"币种"`
	Sort        int         `json:"sort"         orm:"sort"         description:"排序"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:"创建人"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
	Logo        string      `json:"logo"         orm:"logo"         description:"图标"`
	BaseCoin    string      `json:"base_coin"    orm:"base_coin"    description:"结算币种"`
	ShowSymbol  string      `json:"show_symbol"  orm:"show_symbol"  description:"展示币种"`
	Type        int         `json:"type"         orm:"type"         description:"币种类型 1 外汇  2虚拟币"`
}
