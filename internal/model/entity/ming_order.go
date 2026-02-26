// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MingOrder is the golang structure for table ming_order.
type MingOrder struct {
	Id              int         `json:"id"               orm:"id"               description:""`
	Amount          float64     `json:"amount"           orm:"amount"           description:"投资金额（分）"`
	Days            int         `json:"days"             orm:"days"             description:"投资期限（天）"`
	Status          int         `json:"status"           orm:"status"           description:"0 收益  1 结算"`
	PlanId          int64       `json:"plan_id"          orm:"plan_id"          description:"项目id"`
	PlanTitle       string      `json:"plan_title"       orm:"plan_title"       description:"项目名称"`
	OrderNo         string      `json:"order_no"         orm:"order_no"         description:"订单编号"`
	CreateTime      *gtime.Time `json:"create_time"      orm:"create_time"      description:"投资时间"`
	EndTime         *gtime.Time `json:"end_time"         orm:"end_time"         description:"到期时间"`
	SettleTime      *gtime.Time `json:"settle_time"      orm:"settle_time"      description:"结算时间"`
	AccumulaEarn    float64     `json:"accumula_earn"    orm:"accumula_earn"    description:"累计收益"`
	UpdateTime      *gtime.Time `json:"update_time"      orm:"update_time"      description:""`
	MinOdds         float64     `json:"min_odds"         orm:"min_odds"         description:"最小利率"`
	MaxOdds         float64     `json:"max_odds"         orm:"max_odds"         description:"最大利率"`
	AdminUserIds    string      `json:"admin_user_ids"   orm:"admin_user_ids"   description:"后台用户id"`
	CollectionOrder string      `json:"collection_order" orm:"collection_order" description:""`
	UserId          int64       `json:"user_id"          orm:"user_id"          description:""`
	OrderAmount     float64     `json:"order_amount"     orm:"order_amount"     description:""`
}
