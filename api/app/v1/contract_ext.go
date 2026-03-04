package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdjustPositionMarginReq 调整逐仓保证金
type AdjustPositionMarginReq struct {
	g.Meta  `path:"/contract/adjustMargin" tags:"Contract" method:"post" summary:"调整持仓保证金"`
	OrderNo string  `json:"orderNo" v:"required#仓位编号不能为空"`
	Amount  float64 `json:"amount" v:"required#调整金额不能为空"`
	Type    int     `json:"type" v:"required|in:0,1#类型: 0增加 1减少"`
}

type AdjustPositionMarginRes struct {
	NewAmount float64 `json:"newAmount" dc:"操作后的保证金"`
}

// ContractLossSettReq 设置止盈止损
type ContractLossSettReq struct {
	g.Meta   `path:"/contract/sett" tags:"Contract" method:"post" summary:"设置持仓止盈止损"`
	OrderNo  string  `json:"orderNo" v:"required#仓位编号不能为空"`
	LossRate float64 `json:"lossRate" dc:"止损价格"`
	EarnRate float64 `json:"earnRate" dc:"止盈价格"`
}

type ContractLossSettRes struct {
	Success bool `json:"success"`
}

// ClosePositionReq 市价平仓
type ClosePositionReq struct {
	g.Meta  `path:"/contract/stopPosition" tags:"Contract" method:"post" summary:"市价一键平仓"`
	OrderNo string `json:"orderNo" v:"required#仓位编号不能为空"`
}

type ClosePositionRes struct {
	Profit float64 `json:"profit"`
}
