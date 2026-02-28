package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RechargeSubmitReq 会员充值提交请求
type RechargeSubmitReq struct {
	g.Meta  `path:"/recharge/submit" tags:"Funding" method:"post" summary:"会员充值提交"`
	Amount  float64 `json:"amount" v:"required|min:0.000001#请输入有效的充值金额"`
	Type    string  `json:"type" v:"required#请输入充值币种 (如USDT/BTC)"`
	Address string  `json:"address" dc:"转账的链上地址或凭证"`
	TxId    string  `json:"txId" dc:"链上Hash值"`
	Pic     string  `json:"pic" dc:"转账截图URL"`
}

// RechargeSubmitRes 会员充值提交响应
type RechargeSubmitRes struct{}

// WithdrawSubmitReq 用户提现提交请求
type WithdrawSubmitReq struct {
	g.Meta   `path:"/withdraw/submit" tags:"Funding" method:"post" summary:"用户提现提交"`
	Amount   float64 `json:"amount" v:"required|min:1#请输入提现金额，最低额度1"`
	CoinType string  `json:"coinType" v:"required#请输入提现通道类型 (如 ERC20/TRC20)"`
	Pwd      string  `json:"pwd" v:"required#请输入资金交易密码"`
	Address  string  `json:"address" v:"required#请输入提现目标地址"`
	Coin     string  `json:"coin" v:"required#请输入具体的提现币种 (如USDT)"`
}

// WithdrawSubmitRes 用户提现提交响应
type WithdrawSubmitRes struct{}
