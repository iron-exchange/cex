// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractLoss is the golang structure of table t_contract_loss for DAO operations like Where/Data.
type ContractLoss struct {
	g.Meta            `orm:"table:t_contract_loss, do:true"`
	Id                any         // 主键
	DelegateType      any         // 委托类型（0 限价 1 市价）
	Status            any         // 状态  0  正常 1 删除  2 撤销
	PositionId        any         // 仓位ID
	UserId            any         // 用户id
	EarnPrice         any         // 止盈触发价
	LosePrice         any         // 止损触发价
	CreateTime        *gtime.Time // 创建时间
	EarnDelegatePrice any         // 止盈委托价
	LoseDelegatePrice any         // 止损委托价
	EarnNumber        any         // 止盈数量
	LoseNumber        any         // 止损数量
	LossType          any         // 0 止盈    1止损
	UpdateTime        *gtime.Time // 更新时间
	Type              any         //
	Leverage          any         //
	Symbol            any         //
}
