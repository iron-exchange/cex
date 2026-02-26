// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OwnCoinOrder is the golang structure for table own_coin_order.
type OwnCoinOrder struct {
	Id             int64       `json:"id"               orm:"id"               description:"主键"`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"用户ID"`
	OrderId        string      `json:"order_id"         orm:"order_id"         description:"订单ID"`
	OwnId          int64       `json:"own_id"           orm:"own_id"           description:"申购币种ID"`
	OwnCoin        string      `json:"own_coin"         orm:"own_coin"         description:"申购币种"`
	Amount         float64     `json:"amount"           orm:"amount"           description:"申购额（usdt）"`
	Number         int         `json:"number"           orm:"number"           description:"申购数量"`
	Price          float64     `json:"price"            orm:"price"            description:"申购价"`
	Status         string      `json:"status"           orm:"status"           description:"状态"`
	AdminUserIds   string      `json:"admin_user_ids"   orm:"admin_user_ids"   description:"上级用户IDS"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:"上级后台用户IDS"`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:""`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:""`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:""`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:""`
	Remark         string      `json:"remark"           orm:"remark"           description:""`
}
