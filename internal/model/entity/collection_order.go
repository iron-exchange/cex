// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectionOrder is the golang structure for table collection_order.
type CollectionOrder struct {
	Id          int         `json:"id"           orm:"id"           description:"主键ID"`
	OrderId     string      `json:"order_id"     orm:"order_id"     description:"订单号"`
	UserId      int64       `json:"user_id"      orm:"user_id"      description:"用户ID"`
	Address     string      `json:"address"      orm:"address"      description:"归集地址"`
	Chain       string      `json:"chain"        orm:"chain"        description:"地址类型"`
	Hash        string      `json:"hash"         orm:"hash"         description:"hash"`
	Coin        string      `json:"coin"         orm:"coin"         description:"币种"`
	Amount      float64     `json:"amount"       orm:"amount"       description:"归集金额"`
	Status      string      `json:"status"       orm:"status"       description:"1  进行中   2 归集成功  3 归集失败"`
	ClientName  string      `json:"client_name"  orm:"client_name"  description:"客户端名称"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:"创建人"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"修改时间"`
	UpdateBy    *gtime.Time `json:"update_by"    orm:"update_by"    description:"修改人"`
	Remark      string      `json:"remark"       orm:"remark"       description:"备注"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
}
