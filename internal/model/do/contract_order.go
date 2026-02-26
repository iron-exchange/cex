// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractOrder is the golang structure of table t_contract_order for DAO operations like Where/Data.
type ContractOrder struct {
	g.Meta         `orm:"table:t_contract_order, do:true"`
	Id             any         // 主键
	Type           any         // (0 买多 1卖空)
	DelegateType   any         // 委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）
	Status         any         // 状态  0 （等待成交  1 完全成交  3已撤销）
	DelegateTotal  any         // 委托总量
	DelegatePrice  any         // 委托价格
	DealNum        any         // 已成交量
	DealPrice      any         // 成交价
	DelegateValue  any         // 委托价值
	DealValue      any         // 成交价值
	DelegateTime   *gtime.Time // 委托时间
	DealTime       *gtime.Time // 成交时间
	CoinSymbol     any         // 交易币种
	CreateTime     *gtime.Time // 创建时间
	OrderNo        any         // 订单编号
	UserId         any         // 用户id
	UpdateTime     *gtime.Time // 更新时间
	Fee            any         // 手续费
	BaseCoin       any         // 基础币种（USDT）
	Leverage       any         // 杠杆
	Symbol         any         // 交易对
	AdminUserIds   any         // 代理IDS
	AdminParentIds any         //
}
