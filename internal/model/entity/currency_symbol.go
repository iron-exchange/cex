// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CurrencySymbol is the golang structure for table currency_symbol.
type CurrencySymbol struct {
	Id            int64       `json:"id"             orm:"id"             description:"主键id"`
	Symbol        string      `json:"symbol"         orm:"symbol"         description:"交易对"`
	ShowSymbol    string      `json:"show_symbol"    orm:"show_symbol"    description:"展示交易对"`
	Coin          string      `json:"coin"           orm:"coin"           description:"交易币种"`
	BaseCoin      string      `json:"base_coin"      orm:"base_coin"      description:"结算币种"`
	FeeRate       float64     `json:"fee_rate"       orm:"fee_rate"       description:"手续费率"`
	CoinPrecision int         `json:"coin_precision" orm:"coin_precision" description:"交易币种精度"`
	BasePrecision int         `json:"base_precision" orm:"base_precision" description:"结算币种精度"`
	SellMin       float64     `json:"sell_min"       orm:"sell_min"       description:"最低卖单价"`
	BuyMax        float64     `json:"buy_max"        orm:"buy_max"        description:"最高买单价"`
	OrderMin      float64     `json:"order_min"      orm:"order_min"      description:"最小下单量"`
	OrderMax      float64     `json:"order_max"      orm:"order_max"      description:"最大下单量"`
	Enable        string      `json:"enable"         orm:"enable"         description:"启用禁用  1=启用 2=禁用"`
	IsShow        string      `json:"is_show"        orm:"is_show"        description:"前端是否显示 1=显示  2=隐藏"`
	IsDeal        string      `json:"is_deal"        orm:"is_deal"        description:"是否可交易 1=是 2=否"`
	MarketBuy     string      `json:"market_buy"     orm:"market_buy"     description:"市价买 1=可以 2=不可以"`
	MarketSell    string      `json:"market_sell"    orm:"market_sell"    description:"市价卖 1=可以 2=不可以"`
	LimitedBuy    string      `json:"limited_buy"    orm:"limited_buy"    description:"限价买 1=可以 2=不可以"`
	LimitedSell   string      `json:"limited_sell"   orm:"limited_sell"   description:"限价卖 1=可以 2=不可以"`
	Logo          string      `json:"logo"           orm:"logo"           description:"图标"`
	Market        string      `json:"market"         orm:"market"         description:"交易所"`
	CreateBy      string      `json:"create_by"      orm:"create_by"      description:""`
	CreateTime    *gtime.Time `json:"create_time"    orm:"create_time"    description:""`
	UpdateBy      string      `json:"update_by"      orm:"update_by"      description:""`
	UpdateTime    *gtime.Time `json:"update_time"    orm:"update_time"    description:""`
	SearchValue   string      `json:"search_value"   orm:"search_value"   description:""`
	Remark        string      `json:"remark"         orm:"remark"         description:""`
	MinSell       float64     `json:"min_sell"       orm:"min_sell"       description:"最低卖出量"`
}
