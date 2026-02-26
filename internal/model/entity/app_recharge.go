// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppRecharge is the golang structure for table app_recharge.
type AppRecharge struct {
	Id             int64       `json:"id"               orm:"id"               description:"卡ID"`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:""`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:""`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	Remark         string      `json:"remark"           orm:"remark"           description:""`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"所有者ID"`
	Username       string      `json:"username"         orm:"username"         description:"用户名"`
	Amount         float64     `json:"amount"           orm:"amount"           description:"充值金额"`
	Bonus          int         `json:"bonus"            orm:"bonus"            description:""`
	Status         int         `json:"status"           orm:"status"           description:"状态"`
	SerialId       string      `json:"serial_id"        orm:"serial_id"        description:"订单号"`
	TxId           string      `json:"tx_id"            orm:"tx_id"            description:"第三方支付订单号"`
	Type           string      `json:"type"             orm:"type"             description:"类型"`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
	Address        string      `json:"address"          orm:"address"          description:"充值地址"`
	Tree           string      `json:"tree"             orm:"tree"             description:""`
	Coin           string      `json:"coin"             orm:"coin"             description:"币总"`
	ToAddress      string      `json:"to_address"       orm:"to_address"       description:"入款地址"`
	BlockTime      *gtime.Time `json:"block_time"       orm:"block_time"       description:"区块时间"`
	Host           string      `json:"host"             orm:"host"             description:""`
	RealAmount     float64     `json:"real_amount"      orm:"real_amount"      description:"实际到账金额"`
	FileName       string      `json:"file_name"        orm:"file_name"        description:"充值凭证"`
	RechargeRemark string      `json:"recharge_remark"  orm:"recharge_remark"  description:""`
	NoticeFlag     int         `json:"notice_flag"      orm:"notice_flag"      description:"通知字段 0未通知 1通知了"`
	AppParentIds   string      `json:"app_parent_ids"   orm:"app_parent_ids"   description:"app代理ids"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:"后台代理ids"`
	OperateTime    *gtime.Time `json:"operate_time"     orm:"operate_time"     description:"操作时间"`
	OrderType      string      `json:"order_type"       orm:"order_type"       description:"订单类型 1/null=充值  2=彩金赠送"`
}
