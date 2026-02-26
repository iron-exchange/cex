// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractPosition is the golang structure of table t_contract_position for DAO operations like Where/Data.
type ContractPosition struct {
	g.Meta           `orm:"table:t_contract_position, do:true"`
	Id               any         // 主键
	Type             any         // (0 买多 1卖空)
	DelegateType     any         // 委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）
	Status           any         // 状态  0 （等待成交  1 完全成交
	Amount           any         // 保证金
	OpenNum          any         // 持仓数量
	OpenPrice        any         // 开仓均价
	ClosePrice       any         // 预计强平价
	OrderNo          any         // 仓位编号
	UserId           any         // 用户id
	OpenFee          any         // 开仓手续费
	Leverage         any         // 杠杆
	Symbol           any         // 交易对
	CreateTime       *gtime.Time // 创建时间
	AdjustAmount     any         // 调整保证金
	Earn             any         // 收益
	DealPrice        any         // 成交价
	DealNum          any         // 成交量
	DealTime         *gtime.Time // 成交时间
	SellFee          any         // 卖出手续费
	RemainMargin     any         // 剩余保证金
	AssetFee         any         // 周期手续费
	EntrustmentValue any         //
	DealValue        any         //
	UpdateTime       *gtime.Time //
	AdminParentIds   any         // 代理IDS
	AuditStatus      any         // 审核
	DeliveryDays     any         // 交割时间
	MinMargin        any         // 最小保证金
	LossRate         any         // 止损率
	EarnRate         any         // 止盈率
	SubTime          *gtime.Time // 提交时间
}
