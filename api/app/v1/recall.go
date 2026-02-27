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
	Address string `json:"address"` // 收款地址
	Amount  string `json:"amount"`  // 交易金额 (由于可能极小值，优盾默认推 String)
	Coin    string `json:"coin"`    // 链上手续费币种标记或主币标记
	TxId    string `json:"txId"`    // 链上哈希 Hash
	Type    string `json:"type"`    // 充值/提现枚举
	Status  int    `json:"status"`  // 优盾订单状态
}

// UncCallbackRes 优盾回调响应
type UncCallbackRes struct{}
