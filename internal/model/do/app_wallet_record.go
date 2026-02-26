// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppWalletRecord is the golang structure of table t_app_wallet_record for DAO operations like Where/Data.
type AppWalletRecord struct {
	g.Meta         `orm:"table:t_app_wallet_record, do:true"`
	Id             any         // 卡ID
	Amount         any         // 余额
	UAmount        any         // 换算U金额
	CreateBy       any         //
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       any         //
	UpdateTime     *gtime.Time // 更新时间
	Remark         any         //
	UserId         any         // 用户id
	SearchValue    any         //
	BeforeAmount   any         // 前值
	AfterAmount    any         // 后值
	SerialId       any         //
	Type           any         // 余额
	Symbol         any         // 币种
	AdminParentIds any         // 代理ID
	OperateTime    *gtime.Time // 操作时间
}
