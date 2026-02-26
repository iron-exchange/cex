// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LoadProduct is the golang structure for table load_product.
type LoadProduct struct {
	Id          int64       `json:"id"           orm:"id"           description:"主键"`
	AmountMin   float64     `json:"amount_min"   orm:"amount_min"   description:"贷款最小额度"`
	AmountMax   float64     `json:"amount_max"   orm:"amount_max"   description:"贷款最大额度"`
	CycleType   int         `json:"cycle_type"   orm:"cycle_type"   description:"周期类型  0-7天 1-14天 2-30天 ,,,,待补充"`
	RepayType   int         `json:"repay_type"   orm:"repay_type"   description:"还款类型 0-到期一次换本息...待补充"`
	Status      int         `json:"status"       orm:"status"       description:"状态 0 未开启 1已开启"`
	CreateBy    string      `json:"create_by"    orm:"create_by"    description:""`
	CreateTime  *gtime.Time `json:"create_time"  orm:"create_time"  description:"创建时间"`
	UpdateBy    string      `json:"update_by"    orm:"update_by"    description:""`
	UpdateTime  *gtime.Time `json:"update_time"  orm:"update_time"  description:"更新时间"`
	Remark      string      `json:"remark"       orm:"remark"       description:"用户备注"`
	SearchValue string      `json:"search_value" orm:"search_value" description:""`
	Odds        float64     `json:"odds"         orm:"odds"         description:"日利率（%）"`
	RepayOrg    string      `json:"repay_org"    orm:"repay_org"    description:"还款机构"`
	IsFreeze    string      `json:"is_freeze"    orm:"is_freeze"    description:"是否冻结  1=正常 2=冻结"`
}
