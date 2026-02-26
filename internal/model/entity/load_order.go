// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LoadOrder is the golang structure for table load_order.
type LoadOrder struct {
	Id             int64       `json:"id"               orm:"id"               description:"主键"`
	ProId          int64       `json:"pro_id"           orm:"pro_id"           description:"贷款商品表id"`
	UserId         int64       `json:"user_id"          orm:"user_id"          description:"用户id"`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"购买时间"`
	Amount         float64     `json:"amount"           orm:"amount"           description:"贷款金额"`
	Rate           float64     `json:"rate"             orm:"rate"             description:"贷款利率"`
	Interest       float64     `json:"interest"         orm:"interest"         description:"利息"`
	Status         int         `json:"status"           orm:"status"           description:"0=待审核 1=审核通过  2=审核拒绝  3=已结清  4=已逾期"`
	FinalRepayTime *gtime.Time `json:"final_repay_time" orm:"final_repay_time" description:"最后还款日"`
	DisburseTime   *gtime.Time `json:"disburse_time"    orm:"disburse_time"    description:"放款日期"`
	ReturnTime     *gtime.Time `json:"return_time"      orm:"return_time"      description:"还款日期"`
	DisburseAmount float64     `json:"disburse_amount"  orm:"disburse_amount"  description:"审批金额"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:"后台代理ids"`
	CardUrl        string      `json:"card_url"         orm:"card_url"         description:"手持身份证"`
	CardBackUrl    string      `json:"card_back_url"    orm:"card_back_url"    description:"身份证正面"`
	CapitalUrl     string      `json:"capital_url"      orm:"capital_url"      description:"身份证反面"`
	LicenseUrl     string      `json:"license_url"      orm:"license_url"      description:""`
	OrderNo        string      `json:"order_no"         orm:"order_no"         description:""`
	CycleType      int         `json:"cycle_type"       orm:"cycle_type"       description:"还款周期  多少天"`
	Remark         string      `json:"remark"           orm:"remark"           description:"用户备注"`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:""`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:""`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
}
