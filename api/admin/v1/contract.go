package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 币种配置 (Contract Coin Config) ---------
type ContractCoinInfo struct {
	Id           int64   `json:"id"`
	Symbol       string  `json:"symbol"`
	Coin         string  `json:"coin"`
	BaseCoin     string  `json:"baseCoin"`
	ShareNumber  float64 `json:"shareNumber"`
	Leverage     string  `json:"leverage"`
	Enable       int     `json:"enable"` // 0 启用 1 禁止
	Exchangeable int     `json:"exchangeable"`
	OpenFee      float64 `json:"openFee"`
	CloseFee     float64 `json:"closeFee"`
	UsdtRate     float64 `json:"usdtRate"`
	Visible      int     `json:"visible"` // 0 启用 1 禁止
	MinShare     float64 `json:"minShare"`
	MaxShare     float64 `json:"maxShare"`
	CreateTime   string  `json:"createTime"`
}

type GetContractCoinListReq struct {
	g.Meta `path:"/contract/coin/list" tags:"AdminContract" method:"get" summary:"获取U本位合约币种配置"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Symbol string `json:"symbol" dc:"交易对名称"`
}

type GetContractCoinListRes struct {
	List  []ContractCoinInfo `json:"list"`
	Total int                `json:"total"`
}

// --------- 委托列表 (Contract Orders) ---------
type ContractOrderInfo struct {
	Id            int64   `json:"id"`
	OrderNo       string  `json:"orderNo"`
	UserId        int64   `json:"userId"`
	Symbol        string  `json:"symbol"`
	Type          int     `json:"type"`         // 0 买多 1卖空
	DelegateType  int     `json:"delegateType"` // 0 限价 1 市价
	Status        int     `json:"status"`       // 0 等待成交 1 完全成交 3 已撤销
	DelegateTotal float64 `json:"delegateTotal"`
	DelegatePrice float64 `json:"delegatePrice"`
	DealNum       float64 `json:"dealNum"`
	DealPrice     float64 `json:"dealPrice"`
	Leverage      float64 `json:"leverage"`
	Fee           float64 `json:"fee"`
	CreateTime    string  `json:"createTime"`
}

type GetContractOrderListReq struct {
	g.Meta  `path:"/contract/order/list" tags:"AdminContract" method:"get" summary:"获取U本位委托订单"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetContractOrderListRes struct {
	List  []ContractOrderInfo `json:"list"`
	Total int                 `json:"total"`
}

// --------- 持仓列表 (Contract Positions) ---------
type ContractPositionInfo struct {
	Id         int64   `json:"id"`
	OrderNo    string  `json:"orderNo"`
	UserId     int64   `json:"userId"`
	Symbol     string  `json:"symbol"`
	Type       int     `json:"type"` // 0 买多 1卖空
	Status     int     `json:"status"`
	Amount     float64 `json:"amount"` // 保证金
	OpenNum    float64 `json:"openNum"`
	OpenPrice  float64 `json:"openPrice"`
	ClosePrice float64 `json:"closePrice"`
	Leverage   float64 `json:"leverage"`
	Earn       float64 `json:"earn"`
	OpenFee    float64 `json:"openFee"`
	CreateTime string  `json:"createTime"`
}

type GetContractPositionListReq struct {
	g.Meta  `path:"/contract/position/list" tags:"AdminContract" method:"get" summary:"获取U本位持仓记录"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"仓位编号"`
	Symbol  string `json:"symbol" dc:"交易对"`
}

type GetContractPositionListRes struct {
	List  []ContractPositionInfo `json:"list"`
	Total int                    `json:"total"`
}
