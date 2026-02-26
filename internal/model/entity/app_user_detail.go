// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUserDetail is the golang structure for table app_user_detail.
type AppUserDetail struct {
	Id                  int         `json:"id"                    orm:"id"                    description:""`
	UserId              int         `json:"user_id"               orm:"user_id"               description:""`
	RealName            string      `json:"real_name"             orm:"real_name"             description:"真实姓名"`
	IdCard              string      `json:"id_card"               orm:"id_card"               description:"身份证号码"`
	FrontUrl            string      `json:"front_url"             orm:"front_url"             description:"身份证正面照片"`
	Country             string      `json:"country"               orm:"country"               description:"国际"`
	CardType            string      `json:"card_type"             orm:"card_type"             description:""`
	HandelUrl           string      `json:"handel_url"            orm:"handel_url"            description:"手持身份证照片"`
	BackUrl             string      `json:"back_url"              orm:"back_url"              description:"身份证反面照片"`
	UserTardPwd         string      `json:"user_tard_pwd"         orm:"user_tard_pwd"         description:"用户交易密码"`
	CreateBy            string      `json:"create_by"             orm:"create_by"             description:""`
	CreateTime          *gtime.Time `json:"create_time"           orm:"create_time"           description:""`
	UpdateBy            string      `json:"update_by"             orm:"update_by"             description:""`
	UpdateTime          *gtime.Time `json:"update_time"           orm:"update_time"           description:""`
	Remark              string      `json:"remark"                orm:"remark"                description:""`
	SearchValue         string      `json:"search_value"          orm:"search_value"          description:""`
	AuditStatusPrimary  int         `json:"audit_status_primary"  orm:"audit_status_primary"  description:"初级验证状态"`
	AuditStatusAdvanced int         `json:"audit_status_advanced" orm:"audit_status_advanced" description:"高级验证状态"`
	Credits             int         `json:"credits"               orm:"credits"               description:"信用分"`
	UserRechargeAddress string      `json:"user_recharge_address" orm:"user_recharge_address" description:"用户充值地址"`
	WinNum              int         `json:"win_num"               orm:"win_num"               description:"连赢场次"`
	LoseNum             int         `json:"lose_num"              orm:"lose_num"              description:"连输场次"`
	TradeFlag           string      `json:"trade_flag"            orm:"trade_flag"            description:"交易是否被限制 1 为限制"`
	AmountFlag          string      `json:"amount_flag"           orm:"amount_flag"           description:"金额是否被限制 1 为限制"`
	PushMessage         string      `json:"push_message"          orm:"push_message"          description:"金额限制提示语"`
	TradeMessage        string      `json:"trade_message"         orm:"trade_message"         description:"交易限制提示语"`
	OperateTime         *gtime.Time `json:"operate_time"          orm:"operate_time"          description:"实名认证时间"`
}
