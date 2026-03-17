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

// RechargeListReq 获取充值列表请求
type RechargeListReq struct {
	g.Meta   `path:"/recharge/list" tags:"Funding" method:"get" summary:"获取充值列表"`
	PageNum  int    `json:"pageNum" v:"min:1#页码最小为1" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" v:"min:1|max:100#每页数量1-100" dc:"每页数量" d:"10"`
	Coin     string `json:"coin" dc:"币种筛选"`
}

// RechargeListRes 获取充值列表响应
type RechargeListRes struct {
	Total int            `json:"total" dc:"总条数"`
	Rows  []RechargeItem `json:"rows" dc:"数据列表"`
}

// RechargeItem 充值记录单项
type RechargeItem struct {
	Id         int64   `json:"id" dc:"ID"`
	SerialId   string  `json:"serialId" dc:"订单号"`
	Amount     float64 `json:"amount" dc:"申请数量"`
	RealAmount float64 `json:"realAmount" dc:"到账数量"`
	Coin       string  `json:"coin" dc:"币种"`
	Type       string  `json:"type" dc:"通道类型"`
	Status     int     `json:"status" dc:"状态: 0审核中, 1通过, 2拒绝"`
	Address    string  `json:"address" dc:"充值地址"`
	TxId       string  `json:"txId" dc:"链上Hash"`
	Remark     string  `json:"remark" dc:"理由/备注"`
	CreateTime string  `json:"createTime" dc:"创建时间"`
}

// WithdrawListReq 获取提现列表请求
type WithdrawListReq struct {
	g.Meta   `path:"/withdraw/list" tags:"Funding" method:"get" summary:"获取提现列表"`
	PageNum  int    `json:"pageNum" v:"min:1#页码最小为1" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" v:"min:1|max:100#每页数量1-100" dc:"每页数量" d:"10"`
	Coin     string `json:"coin" dc:"币种筛选"`
}

// WithdrawListRes 获取提现列表响应
type WithdrawListRes struct {
	Total int            `json:"total" dc:"总条数"`
	Rows  []WithdrawItem `json:"rows" dc:"数据列表"`
}

// WithdrawItem 提现记录单项
type WithdrawItem struct {
	Id         int     `json:"id" dc:"ID"`
	SerialId   string  `json:"serialId" dc:"订单号"`
	Amount     float64 `json:"amount" dc:"申请数量"`
	Type       string  `json:"type" dc:"通道类型"`
	Coin       string  `json:"coin" dc:"币种"`
	Fee        float64 `json:"fee" dc:"手续费"`
	RealAmount float64 `json:"realAmount" dc:"到账数量"`
	Address    string  `json:"address" dc:"提现地址"`
	Status     int     `json:"status" dc:"状态: 0审核中, 1成功, 2失败"`
	Remark     string  `json:"remark" dc:"理由/备注"`
	CreateTime string  `json:"createTime" dc:"创建时间"`
}
