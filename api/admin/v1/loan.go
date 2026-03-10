package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 借贷产品 (Load Product) ---------
type AdminLoadProductInfo struct {
	Id         int64   `json:"id"`
	AmountMin  float64 `json:"amountMin"`
	AmountMax  float64 `json:"amountMax"`
	CycleType  int     `json:"cycleType"` // 0-7天 1-14天 2-30天
	RepayType  int     `json:"repayType"`
	Status     int     `json:"status"` // 0未开启 1已开启
	Odds       float64 `json:"odds"`
	RepayOrg   string  `json:"repayOrg"`
	IsFreeze   string  `json:"isFreeze"` // 1正常 2冻结
	CreateTime string  `json:"createTime"`
}

type GetAdminLoadProductListReq struct {
	g.Meta `path:"/loan/product/list" tags:"AdminLoan" method:"get" summary:"获取借贷产品列表"`
	Page   int  `json:"page" d:"1"`
	Size   int  `json:"size" d:"20"`
	Status *int `json:"status" dc:"状态 0/1"`
}

type GetAdminLoadProductListRes struct {
	List  []AdminLoadProductInfo `json:"list"`
	Total int                    `json:"total"`
}

// --------- 借贷订单 (Load Order) ---------
type AdminLoadOrderInfo struct {
	Id             int64   `json:"id"`
	OrderNo        string  `json:"orderNo"`
	UserId         int64   `json:"userId"`
	ProId          int64   `json:"proId"`
	Amount         float64 `json:"amount"`
	Rate           float64 `json:"rate"`
	Interest       float64 `json:"interest"`
	Status         int     `json:"status"` // 0=待审核 1=审核通过  2=审核拒绝  3=已结清  4=已逾期
	CycleType      int     `json:"cycleType"`
	FinalRepayTime string  `json:"finalRepayTime"` // 最后还款日
	DisburseTime   string  `json:"disburseTime"`   // 放款日期
	ReturnTime     string  `json:"returnTime"`     // 还款日期
	DisburseAmount float64 `json:"disburseAmount"` // 审批金额
	CreateTime     string  `json:"createTime"`
}

type GetAdminLoadOrderListReq struct {
	g.Meta  `path:"/loan/order/list" tags:"AdminLoan" method:"get" summary:"获取用户借贷订单记录"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OrderNo string `json:"orderNo" dc:"订单编号"`
	Status  *int   `json:"status" dc:"状态"`
}

type GetAdminLoadOrderListRes struct {
	List  []AdminLoadOrderInfo `json:"list"`
	Total int                  `json:"total"`
}
