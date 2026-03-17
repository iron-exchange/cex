package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UncCallbackReq 优盾第三方代收付 Webhook 回调结构体
// 原系统映射: @PostMapping(value = "/unc")
type UncCallbackReq struct {
	g.Meta    `path:"/unc" tags:"Recall" method:"post" summary:"优盾支付回调网关"`
	Timestamp string `json:"timestamp" v:"required#时间戳丢失"`
	Nonce     string `json:"nonce"     v:"required#随机数丢失"`
	Sign      string `json:"sign"      v:"required#签名校验失败"`
	Body      string `json:"body"      v:"required#报文主体丢失"`
}

// UdunBody 优盾回调的内层 Body 结构反序列化载体
type UdunBody struct {
	TradeType    int    `json:"tradeType"`    // 业务类型 (1: 充值/代收, 2: 提现/代付)
	Status       int    `json:"status"`       // 交易状态 (0待审核, 1审核通过, 2审核驳回, 3交易成功, 4交易失败)
	BusinessId   string `json:"businessId"`   // 业务 ID (提现订单号)
	Address      string `json:"address"`      // 交易钱包地址
	Amount       string `json:"amount"`       // 交易金额 (最小单位整数型)
	Decimals     int    `json:"decimals"`     // 币种精度
	Fee          string `json:"fee"`          // 链上实际消耗手续费
	TxId         string `json:"txId"`         // 链上哈希 Hash
	MainCoinType string `json:"mainCoinType"` // 优盾币种主 ID
	CoinType     string `json:"coinType"`     // 优盾币种 ID
	Coin         string `json:"coin"`         // 币种名称 (如 USDT)
}

// UncCallbackRes 优盾回调响应
type UncCallbackRes struct{}
