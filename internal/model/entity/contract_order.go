// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractOrder is the golang structure for table contract_order.
type ContractOrder struct {
	Id             int64       `json:"id"               orm:"id"               description:"主键"`
	Type           int         `json:"type"             orm:"type"             description:"(0 买多 1卖空)"`
	DelegateType   int         `json:"delegate_type"    orm:"delegate_type"    description:"委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）"`
	Status         int         `json:"status"           orm:"status"           description:"状态  0 （等待成交  1 完全成交  3已撤销）"`
	DelegateTotal  float64     `json:"delegate_total"   orm:"delegate_total"   description:"委托总量"`
	DelegatePrice  float64     `json:"delegate_price"   orm:"delegate_price"   description:"委托价格"`
	DealNum        float64     `json:"deal_num"         orm:"deal_num"         description:"已成交量"`
	DealPrice      float64     `json:"deal_price"       orm:"deal_price"       description:"成交价"`
	DelegateValue  float64     `json:"delegate_value"   orm:"delegate_value"   description:"委托价值"`
	DealValue      float64     `json:"deal_value"       orm:"deal_value"       description:"成交价值"`
	DelegateTime   *gtime.Time `json:"delegate_time"    orm:"delegate_time"    description:"委托时间"`
	DealTime       *gtime.Time `json:"deal_time"        orm:"deal_time"        description:"成交时间"`
	CoinSymbol     string      `json:"coin_symbol"      orm:"coin_symbol"      description:"交易币种"`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	OrderNo        string      `json:"order_no"         orm:"order_no"         description:"订单编号"`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"用户id"`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	Fee            float64     `json:"fee"              orm:"fee"              description:"手续费"`
	BaseCoin       string      `json:"base_coin"        orm:"base_coin"        description:"基础币种（USDT）"`
	Leverage       float64     `json:"leverage"         orm:"leverage"         description:"杠杆"`
	Symbol         string      `json:"symbol"           orm:"symbol"           description:"交易对"`
	AdminUserIds   string      `json:"admin_user_ids"   orm:"admin_user_ids"   description:"代理IDS"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:""`
}
