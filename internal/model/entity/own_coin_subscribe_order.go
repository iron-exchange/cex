// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OwnCoinSubscribeOrder is the golang structure for table own_coin_subscribe_order.
type OwnCoinSubscribeOrder struct {
	Id          int64       `json:"id"           orm:"id"           description:"主键"`
	SubscribeId string      `json:"subscribe_id" orm:"subscribe_id" description:"订阅单号ID"`
	UserId      int64       `json:"user_id"      orm:"user_id"      description:"用户ID"`
	OrderId     string      `json:"order_id"     orm:"order_id"     description:"订单ID"`
	OwnId       int64       `json:"own_id"       orm:"own_id"       description:"申购币种ID"`
	OwnCoin     string      `json:"own_coin"     orm:"own_coin"     description:"申购币种"`
	AmountLimit float64     `json:"amount_limit" orm:"amount_limit" description:"申购额（usdt）"`
	NumLimit    int         `json:"num_limit"    orm:"num_limit"    description:"申购数量上限"`
	Price       float64     `json:"price"        orm:"price"        description:"申购价"`
	Status      string      `json:"status"       orm:"status"       description:"状态，1订阅中、2订阅成功、3成功消息推送完成"`
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:""`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:""`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:""`
}
