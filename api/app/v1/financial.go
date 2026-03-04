package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FinancialSubmitReq 理财产品购买
type FinancialSubmitReq struct {
	g.Meta `path:"/financial/submit" tags:"Financial" method:"post" summary:"理财产品购买"`
	PlanId int64   `json:"planId" v:"required#产品ID不能为空"`
	Money  float64 `json:"money" v:"required|min:1#申购金额无效"`
	Days   int     `json:"days" v:"required#周期不能为空"`
}

type FinancialSubmitRes struct {
	OrderNo string `json:"orderNo"`
}

// PersonalIncomeReq 个人收益类加详情
type PersonalIncomeReq struct {
	g.Meta `path:"/financial/personalIncome" tags:"Financial" method:"post" summary:"个人收益详情汇总"`
}

type PersonalIncomeRes struct {
	TotalIncome float64 `json:"totalIncome" dc:"累计收益"`
	TodayIncome float64 `json:"todayIncome" dc:"今日预计收益"`
	OrderCount  int     `json:"orderCount" dc:"在持订单数"`
}
