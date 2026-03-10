package v1

import "github.com/gogf/gf/v2/frame/g"

type WalletRecordInfo struct {
	Id           int64   `json:"id"`
	UserId       int64   `json:"userId"`
	LoginName    string  `json:"loginName"`
	Symbol       string  `json:"symbol"`
	Type         int     `json:"type"`
	Amount       float64 `json:"amount"`
	UAmount      float64 `json:"uAmount"`
	BeforeAmount float64 `json:"beforeAmount"`
	AfterAmount  float64 `json:"afterAmount"`
	OperateTime  string  `json:"operateTime"`
	Remark       string  `json:"remark"`
}

// GetWalletRecordListReq 账变信息列表查询
type GetWalletRecordListReq struct {
	g.Meta `path:"/walletRecord/list" tags:"AdminReport" method:"get" summary:"获取系统账变流水信息"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	UserId int64  `json:"userId" dc:"根据用户ID查询"`
	Symbol string `json:"symbol" dc:"币种筛选"`
	Type   *int   `json:"type" dc:"账变类型筛选"`
}

type GetWalletRecordListRes struct {
	List  []WalletRecordInfo `json:"list"`
	Total int                `json:"total"`
}
