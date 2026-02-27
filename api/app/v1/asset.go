package v1

import "github.com/gogf/gf/v2/frame/g"

// AssetListReq 获取用户资产大盘
type AssetListReq struct {
	g.Meta `path:"/asset/list" tags:"Asset" method:"get" summary:"获取用户资产大盘"`
}

type AssetInfo struct {
	Symbol          string `json:"symbol" dc:"币种"`
	AvailableAmount string `json:"availableAmount" dc:"可用金额"`
	FrozenAmount    string `json:"frozenAmount" dc:"冻结金额"`
	TotalAmount     string `json:"totalAmount" dc:"总金额"`
	UsdtValuation   string `json:"usdtValuation" dc:"折合 USDT 价值"`
}

type AssetListRes struct {
	TotalUsdtValuation string      `json:"totalUsdtValuation" dc:"总资产折合 USDT"`
	List               []AssetInfo `json:"list" dc:"币种资产明细"`
}

// WalletRecordReq 获取财务流水
type WalletRecordReq struct {
	g.Meta `path:"/wallet/records" tags:"Asset" method:"get" summary:"获取财务流水"`
	Page   int `json:"page" d:"1" dc:"页码"`
	Size   int `json:"size" d:"20" dc:"每页数量"`
	Type   int `json:"type" d:"0" dc:"流水类型过滤，0为全部"`
}

type WalletRecordInfo struct {
	Id           int64  `json:"id"`
	Symbol       string `json:"symbol" dc:"币种"`
	Amount       string `json:"amount" dc:"变动金额"`
	BeforeAmount string `json:"beforeAmount" dc:"变动前"`
	AfterAmount  string `json:"afterAmount" dc:"变动后"`
	Type         int    `json:"type" dc:"类型"`
	Remark       string `json:"remark" dc:"备注"`
	CreateTime   string `json:"createTime" dc:"创建时间"`
}

type WalletRecordRes struct {
	List  []WalletRecordInfo `json:"list" dc:"流水列表"`
	Total int                `json:"total" dc:"总条数"`
}
