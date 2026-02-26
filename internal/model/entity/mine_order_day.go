// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MineOrderDay is the golang structure for table mine_order_day.
type MineOrderDay struct {
	Id          int64       `json:"id"           orm:"id"           description:""`
	Amount      float64     `json:"amount"       orm:"amount"       description:"投资金额（分）"`
	Odds        float64     `json:"odds"         orm:"odds"         description:"当日利率"`
	Earn        float64     `json:"earn"         orm:"earn"         description:"收益"`
	PlanId      int64       `json:"plan_id"      orm:"plan_id"      description:"项目id"`
	OrderNo     string      `json:"order_no"     orm:"order_no"     description:"订单编号"`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"时间"`
	Address     string      `json:"address"      orm:"address"      description:"地址"`
	Type        int         `json:"type"         orm:"type"         description:"0 质押挖矿 1 非质押挖矿"`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:""`
	Status      int         `json:"status"       orm:"status"       description:"1 待结算  2  结算"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:""`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:""`
	Remark      string      `json:"remark"       orm:"remark"       description:""`
}
