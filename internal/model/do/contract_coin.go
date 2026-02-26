// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractCoin is the golang structure of table t_contract_coin for DAO operations like Where/Data.
type ContractCoin struct {
	g.Meta           `orm:"table:t_contract_coin, do:true"`
	Id               any         //
	Symbol           any         // 交易对
	Coin             any         // 币种
	BaseCoin         any         // 基础币种
	ShareNumber      any         // 合约面值（1手多少 如 1手=0.01BTC）
	Leverage         any         // 杠杆倍数
	Enable           any         // 0 启用  1 禁止
	Visible          any         // 前端显示0启用 1 禁止
	Exchangeable     any         // 是否可交易（0 可以 1 禁止）
	EnableOpenSell   any         // 开空  （0  是  1 否）
	EnableOpenBuy    any         // 开多  （0  是  1 否）
	EnableMarketSell any         // 市价开空（0 是 1否）
	EnableMarketBuy  any         // 市价开多（0 是 1否）
	OpenFee          any         // 开仓手续费
	CloseFee         any         // 平仓手续费
	UsdtRate         any         // 资金费率
	IntervalHour     any         // 资金周期
	CoinScale        any         // 币种小数精度
	BaseScale        any         // 基础币小数精度
	MinShare         any         // 最小数（以手为单位 ）
	MaxShare         any         // 最大数（以手为单位 ）
	TotalProfit      any         // 平台收益
	Sort             any         // 排序字段
	CreateTime       *gtime.Time //
	UpdateTime       *gtime.Time //
	ShowSymbol       any         // 显示币种
	Logo             any         //
	Market           any         //
	DeliveryDays     any         // 交割时间
	MinMargin        any         // 最小保证金
	LossRate         any         // 止损率
	EarnRate         any         // 止盈率
	FloatProfit      any         // 浮动盈利点
	ProfitLoss       any         // 浮动盈亏
}
