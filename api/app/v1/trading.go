package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CurrencyOrderSubmitReq 现货挂单 (市价/限价)
type CurrencyOrderSubmitReq struct {
	g.Meta        `path:"/currency/order/submit" tags:"Trading" method:"post" summary:"现货撮合下单"`
	Symbol        string  `json:"symbol" v:"required#交易对不能为空 (例如 BTC/USDT)"`
	Type          int     `json:"type" v:"in:0,1#类型: 0 买入, 1 卖出"`
	DelegateType  int     `json:"delegateType" v:"in:0,1#委托类型: 0 限价, 1 市价"`
	DelegatePrice float64 `json:"delegatePrice" v:"min:0#委托价格不能低于0"`     // 市价时可以为0或者不填
	DelegateTotal float64 `json:"delegateTotal" v:"min:0.00000001#数量太低"` // 限价时为数量，市价买入时为额度
}

type CurrencyOrderSubmitRes struct {
	OrderNo string `json:"orderNo" dc:"系统订单编号"`
}

// SecondContractSubmitReq 期权/秒合约下注
type SecondContractSubmitReq struct {
	g.Meta     `path:"/secondContractOrder/submit" tags:"Trading" method:"post" summary:"秒合约下注"`
	Symbol     string  `json:"symbol" v:"required#交易对不能为空 (例如 BTC/USDT)"`
	BetContent string  `json:"betContent" v:"in:0,1#预测方向: 0 买涨, 1 买跌"`
	BetAmount  float64 `json:"betAmount" v:"required|min:1#下注金额不能少于1"` // 需结合后台配置的最低下注
	Period     int64   `json:"period" v:"required#请选择猜测周期 (例如 30秒, 60秒)"`
}

type SecondContractSubmitRes struct {
	OrderNo string `json:"orderNo" dc:"下注单号"`
}

// ContractOrderSubmitReq 永续合约下单 (U本位)
type ContractOrderSubmitReq struct {
	g.Meta        `path:"/contract/order/buyContractOrder" tags:"Trading" method:"post" summary:"永续合约开仓"`
	Symbol        string  `json:"symbol" v:"required#交易对不能为空 (例如 BTC/USDT)"`
	Type          int     `json:"type" v:"in:0,1#类型: 0 买多, 1 卖空"`
	DelegateType  int     `json:"delegateType" v:"in:0,1#委托类型: 0 限价, 1 市价"`
	Leverage      float64 `json:"leverage" v:"required|min:1#必须提供杠杆倍数"`
	DelegatePrice float64 `json:"delegatePrice" v:"min:0#委托价格"`
	DelegateTotal float64 `json:"delegateTotal" v:"min:0.01#委托保证金数量不能过低"`
}

type ContractOrderSubmitRes struct {
	OrderNo string `json:"orderNo" dc:"系统订单编号"`
}
