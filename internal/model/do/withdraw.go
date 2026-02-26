// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Withdraw is the golang structure of table t_withdraw for DAO operations like Where/Data.
type Withdraw struct {
	g.Meta            `orm:"table:t_withdraw, do:true"`
	Id                any         // 卡ID
	CreateBy          any         //
	CreateTime        *gtime.Time // 创建时间
	UpdateBy          any         //
	UpdateTime        *gtime.Time // 更新时间
	Remark            any         //
	UserId            any         // 用户
	Username          any         // 用户名
	Address           any         // 用户名
	Amount            any         // 提现金额
	Status            any         // 0审核中1成功2失败
	SerialId          any         //
	SearchValue       any         //
	FromAddr          any         // 用户名
	Type              any         // 0审核中1成功2失败
	Coin              any         // 用户名
	Ratio             any         //
	Fee               any         // 手续费
	WithdrawId        any         // 用户名
	Host              any         // Host
	RealAmount        any         // 实际金额
	ToAdress          any         // 收款地址
	AdminParentIds    any         // 后台用户id
	NoticeFlag        any         // 通知字段 0未通知 1通知了
	WithDrawRemark    any         // 提现说明
	BankName          any         // 银行名称
	BankUserName      any         // 银行收款人名称
	BankBranch        any         //
	AdminUserIds      any         //
	OperateTime       *gtime.Time // 操作时间
	FixedFee          any         // 固定手续费
	OrderType         any         // 订单类型 1/null 提现  2=彩金扣减
	ExchangeRate      any         // 汇率
	ReceiptAmount     any         // 应到账金额
	ReceiptRealAmount any         // 实际到账金额
	ReceiptCoin       any         // 到账币种
}
