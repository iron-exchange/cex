package v1

import "github.com/gogf/gf/v2/frame/g"

// SubAmountReq 人工上下分请求 (或系统通用的资金调账接口入参)
type SubAmountReq struct {
	g.Meta     `path:"/asset/subAmount" tags:"Asset" method:"post" summary:"人工调账/资金扣减"`
	UserId     int64   `json:"userId" v:"required#用户ID不能为空" dc:"目标用户ID"`
	Symbol     string  `json:"symbol" v:"required#币种不能为空" dc:"资产币种，如 USDT, BTC"`
	Amount     float64 `json:"amount" dc:"变动金额 (正数为增加，负数为扣减)"`
	AmountStr  string  `json:"amountStr" dc:"高精度字符串变动金额(可选，防 float64 截断)"`
	RecordType int     `json:"recordType" v:"required#流水类型不能为空" dc:"账变类型 (例如: 1后台充值, 2后台扣款)"`
	Remark     string  `json:"remark" dc:"账变备注"`
}

type SubAmountRes struct {
	RecordId      int64   `json:"recordId" dc:"生成的账变流水记录ID"`
	CurrentAmount float64 `json:"currentAmount" dc:"操作后最新可用余额"`
}

type AssetInfo struct {
	UserId       int64   `json:"userId"`
	Symbol       string  `json:"symbol"`
	Available    float64 `json:"available"`
	FreezeAmount float64 `json:"freezeAmount"`
	UpdateTime   string  `json:"updateTime"`
}

// GetAssetListReq 获取玩家资产列表
type GetAssetListReq struct {
	g.Meta `path:"/asset/list" tags:"AdminAsset" method:"get" summary:"获取玩家资产列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	UserId int64  `json:"userId" dc:"按用户ID搜索"`
	Symbol string `json:"symbol" dc:"按币种搜索"`
}

type GetAssetListRes struct {
	List  []AssetInfo `json:"list"`
	Total int         `json:"total"`
}
