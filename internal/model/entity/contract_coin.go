// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContractCoin is the golang structure for table contract_coin.
type ContractCoin struct {
	Id               int64       `json:"id"                 orm:"id"                 description:""`
	Symbol           string      `json:"symbol"             orm:"symbol"             description:"交易对"`
	Coin             string      `json:"coin"               orm:"coin"               description:"币种"`
	BaseCoin         string      `json:"base_coin"          orm:"base_coin"          description:"基础币种"`
	ShareNumber      float64     `json:"share_number"       orm:"share_number"       description:"合约面值（1手多少 如 1手=0.01BTC）"`
	Leverage         string      `json:"leverage"           orm:"leverage"           description:"杠杆倍数"`
	Enable           int         `json:"enable"             orm:"enable"             description:"0 启用  1 禁止"`
	Visible          int         `json:"visible"            orm:"visible"            description:"前端显示0启用 1 禁止"`
	Exchangeable     int         `json:"exchangeable"       orm:"exchangeable"       description:"是否可交易（0 可以 1 禁止）"`
	EnableOpenSell   int         `json:"enable_open_sell"   orm:"enable_open_sell"   description:"开空  （0  是  1 否）"`
	EnableOpenBuy    int         `json:"enable_open_buy"    orm:"enable_open_buy"    description:"开多  （0  是  1 否）"`
	EnableMarketSell int         `json:"enable_market_sell" orm:"enable_market_sell" description:"市价开空（0 是 1否）"`
	EnableMarketBuy  int         `json:"enable_market_buy"  orm:"enable_market_buy"  description:"市价开多（0 是 1否）"`
	OpenFee          float64     `json:"open_fee"           orm:"open_fee"           description:"开仓手续费"`
	CloseFee         float64     `json:"close_fee"          orm:"close_fee"          description:"平仓手续费"`
	UsdtRate         float64     `json:"usdt_rate"          orm:"usdt_rate"          description:"资金费率"`
	IntervalHour     float64     `json:"interval_hour"      orm:"interval_hour"      description:"资金周期"`
	CoinScale        float64     `json:"coin_scale"         orm:"coin_scale"         description:"币种小数精度"`
	BaseScale        float64     `json:"base_scale"         orm:"base_scale"         description:"基础币小数精度"`
	MinShare         float64     `json:"min_share"          orm:"min_share"          description:"最小数（以手为单位 ）"`
	MaxShare         float64     `json:"max_share"          orm:"max_share"          description:"最大数（以手为单位 ）"`
	TotalProfit      float64     `json:"total_profit"       orm:"total_profit"       description:"平台收益"`
	Sort             int         `json:"sort"               orm:"sort"               description:"排序字段"`
	CreateTime       *gtime.Time `json:"create_time"        orm:"create_time"        description:""`
	UpdateTime       *gtime.Time `json:"update_time"        orm:"update_time"        description:""`
	ShowSymbol       string      `json:"show_symbol"        orm:"show_symbol"        description:"显示币种"`
	Logo             string      `json:"logo"               orm:"logo"               description:""`
	Market           string      `json:"market"             orm:"market"             description:""`
	DeliveryDays     int         `json:"delivery_days"      orm:"delivery_days"      description:"交割时间"`
	MinMargin        float64     `json:"min_margin"         orm:"min_margin"         description:"最小保证金"`
	LossRate         float64     `json:"loss_rate"          orm:"loss_rate"          description:"止损率"`
	EarnRate         float64     `json:"earn_rate"          orm:"earn_rate"          description:"止盈率"`
	FloatProfit      float64     `json:"float_profit"       orm:"float_profit"       description:"浮动盈利点"`
	ProfitLoss       float64     `json:"profit_loss"        orm:"profit_loss"        description:"浮动盈亏"`
}
