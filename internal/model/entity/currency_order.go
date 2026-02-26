// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CurrencyOrder is the golang structure for table currency_order.
type CurrencyOrder struct {
	Id             int64       `json:"id"               orm:"id"               description:"主键"`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"用户id"`
	Type           int         `json:"type"             orm:"type"             description:"(0 买入 1卖出)"`
	DelegateType   int         `json:"delegate_type"    orm:"delegate_type"    description:"委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）"`
	Status         int         `json:"status"           orm:"status"           description:"状态  0 （等待成交  1 完全成交  3已撤销）"`
	OrderNo        string      `json:"order_no"         orm:"order_no"         description:"订单编号"`
	Symbol         string      `json:"symbol"           orm:"symbol"           description:"交易币种"`
	Coin           string      `json:"coin"             orm:"coin"             description:"结算币种"`
	DelegateTotal  float64     `json:"delegate_total"   orm:"delegate_total"   description:"委托总量"`
	DelegatePrice  float64     `json:"delegate_price"   orm:"delegate_price"   description:"委托价格"`
	DelegateValue  float64     `json:"delegate_value"   orm:"delegate_value"   description:"委托价值"`
	DelegateTime   *gtime.Time `json:"delegate_time"    orm:"delegate_time"    description:"委托时间"`
	DealNum        float64     `json:"deal_num"         orm:"deal_num"         description:"成交总量"`
	DealPrice      float64     `json:"deal_price"       orm:"deal_price"       description:"成交价格"`
	DealValue      float64     `json:"deal_value"       orm:"deal_value"       description:"成交价值"`
	DealTime       *gtime.Time `json:"deal_time"        orm:"deal_time"        description:"成交时间"`
	Fee            float64     `json:"fee"              orm:"fee"              description:"手续费"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:"后台代理ids"`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:""`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:""`
	Remark         string      `json:"remark"           orm:"remark"           description:""`
}
