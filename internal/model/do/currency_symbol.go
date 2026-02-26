// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CurrencySymbol is the golang structure of table t_currency_symbol for DAO operations like Where/Data.
type CurrencySymbol struct {
	g.Meta        `orm:"table:t_currency_symbol, do:true"`
	Id            any         // 主键id
	Symbol        any         // 交易对
	ShowSymbol    any         // 展示交易对
	Coin          any         // 交易币种
	BaseCoin      any         // 结算币种
	FeeRate       any         // 手续费率
	CoinPrecision any         // 交易币种精度
	BasePrecision any         // 结算币种精度
	SellMin       any         // 最低卖单价
	BuyMax        any         // 最高买单价
	OrderMin      any         // 最小下单量
	OrderMax      any         // 最大下单量
	Enable        any         // 启用禁用  1=启用 2=禁用
	IsShow        any         // 前端是否显示 1=显示  2=隐藏
	IsDeal        any         // 是否可交易 1=是 2=否
	MarketBuy     any         // 市价买 1=可以 2=不可以
	MarketSell    any         // 市价卖 1=可以 2=不可以
	LimitedBuy    any         // 限价买 1=可以 2=不可以
	LimitedSell   any         // 限价卖 1=可以 2=不可以
	Logo          any         // 图标
	Market        any         // 交易所
	CreateBy      any         //
	CreateTime    *gtime.Time //
	UpdateBy      any         //
	UpdateTime    *gtime.Time //
	SearchValue   any         //
	Remark        any         //
	MinSell       any         // 最低卖出量
}
