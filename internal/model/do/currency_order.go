// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CurrencyOrder is the golang structure of table t_currency_order for DAO operations like Where/Data.
type CurrencyOrder struct {
	g.Meta         `orm:"table:t_currency_order, do:true"`
	Id             any         // 主键
	UserId         any         // 用户id
	Type           any         // (0 买入 1卖出)
	DelegateType   any         // 委托类型（0 限价 1 市价 2 止盈止损  3 计划委托）
	Status         any         // 状态  0 （等待成交  1 完全成交  3已撤销）
	OrderNo        any         // 订单编号
	Symbol         any         // 交易币种
	Coin           any         // 结算币种
	DelegateTotal  any         // 委托总量
	DelegatePrice  any         // 委托价格
	DelegateValue  any         // 委托价值
	DelegateTime   *gtime.Time // 委托时间
	DealNum        any         // 成交总量
	DealPrice      any         // 成交价格
	DealValue      any         // 成交价值
	DealTime       *gtime.Time // 成交时间
	Fee            any         // 手续费
	AdminParentIds any         // 后台代理ids
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
	SearchValue    any         //
	CreateBy       any         //
	UpdateBy       any         //
	Remark         any         //
}
