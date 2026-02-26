// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Withdraw is the golang structure for table withdraw.
type Withdraw struct {
	Id                int         `json:"id"                  orm:"id"                  description:"卡ID"`
	CreateBy          string      `json:"create_by"           orm:"create_by"           description:""`
	CreateTime        *gtime.Time `json:"create_time"         orm:"create_time"         description:"创建时间"`
	UpdateBy          string      `json:"update_by"           orm:"update_by"           description:""`
	UpdateTime        *gtime.Time `json:"update_time"         orm:"update_time"         description:"更新时间"`
	Remark            string      `json:"remark"              orm:"remark"              description:""`
	UserId            int         `json:"user_id"             orm:"user_id"             description:"用户"`
	Username          string      `json:"username"            orm:"username"            description:"用户名"`
	Address           string      `json:"address"             orm:"address"             description:"用户名"`
	Amount            float64     `json:"amount"              orm:"amount"              description:"提现金额"`
	Status            int         `json:"status"              orm:"status"              description:"0审核中1成功2失败"`
	SerialId          string      `json:"serial_id"           orm:"serial_id"           description:""`
	SearchValue       string      `json:"search_value"        orm:"search_value"        description:""`
	FromAddr          string      `json:"from_addr"           orm:"from_addr"           description:"用户名"`
	Type              string      `json:"type"                orm:"type"                description:"0审核中1成功2失败"`
	Coin              string      `json:"coin"                orm:"coin"                description:"用户名"`
	Ratio             float64     `json:"ratio"               orm:"ratio"               description:""`
	Fee               float64     `json:"fee"                 orm:"fee"                 description:"手续费"`
	WithdrawId        string      `json:"withdraw_id"         orm:"withdraw_id"         description:"用户名"`
	Host              string      `json:"host"                orm:"host"                description:"Host"`
	RealAmount        float64     `json:"real_amount"         orm:"real_amount"         description:"实际金额"`
	ToAdress          string      `json:"to_adress"           orm:"to_adress"           description:"收款地址"`
	AdminParentIds    string      `json:"admin_parent_ids"    orm:"admin_parent_ids"    description:"后台用户id"`
	NoticeFlag        int         `json:"notice_flag"         orm:"notice_flag"         description:"通知字段 0未通知 1通知了"`
	WithDrawRemark    string      `json:"with_draw_remark"    orm:"with_draw_remark"    description:"提现说明"`
	BankName          string      `json:"bank_name"           orm:"bank_name"           description:"银行名称"`
	BankUserName      string      `json:"bank_user_name"      orm:"bank_user_name"      description:"银行收款人名称"`
	BankBranch        string      `json:"bank_branch"         orm:"bank_branch"         description:""`
	AdminUserIds      string      `json:"admin_user_ids"      orm:"admin_user_ids"      description:""`
	OperateTime       *gtime.Time `json:"operate_time"        orm:"operate_time"        description:"操作时间"`
	FixedFee          float64     `json:"fixed_fee"           orm:"fixed_fee"           description:"固定手续费"`
	OrderType         string      `json:"order_type"          orm:"order_type"          description:"订单类型 1/null 提现  2=彩金扣减"`
	ExchangeRate      float64     `json:"exchange_rate"       orm:"exchange_rate"       description:"汇率"`
	ReceiptAmount     float64     `json:"receipt_amount"      orm:"receipt_amount"      description:"应到账金额"`
	ReceiptRealAmount float64     `json:"receipt_real_amount" orm:"receipt_real_amount" description:"实际到账金额"`
	ReceiptCoin       string      `json:"receipt_coin"        orm:"receipt_coin"        description:"到账币种"`
}
