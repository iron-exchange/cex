package v1

import "github.com/gogf/gf/v2/frame/g"

type BankInfo struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"userId"`
	BankName   string `json:"bankName"`
	BankBranch string `json:"bankBranch"`
	CardNumber string `json:"cardNumber"`
	RealName   string `json:"realName"`
	CreateTime string `json:"createTime"`
}

// GetBankListReq 获取玩家银行卡列表
type GetBankListReq struct {
	g.Meta     `path:"/bank/list" tags:"AdminBank" method:"get" summary:"获取玩家银行卡列表"`
	Page       int    `json:"page" d:"1"`
	Size       int    `json:"size" d:"20"`
	UserId     int    `json:"userId" dc:"按用户ID搜索"`
	CardNumber string `json:"cardNumber" dc:"按卡号搜索"`
}

type GetBankListRes struct {
	List  []BankInfo `json:"list"`
	Total int        `json:"total"`
}
