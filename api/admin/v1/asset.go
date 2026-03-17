package v1

import "github.com/gogf/gf/v2/frame/g"

// SubAmountReq 人工上下分请求 (或系统通用的资金调账接口入参)
type SubAmountReq struct {
	g.Meta     `path:"/asset/subAmount" tags:"Asset" method:"post" summary:"人工调账/资金扣减"`
	UserId     int64   `json:"userId" v:"required#用户ID不能为空" dc:"目标用户ID"`
	Symbol     string  `json:"symbol" v:"required#币种不能为空" dc:"资产币种，如 USDT, BTC"`
	Amount     float64 `json:"amount" dc:"变动金额 (正数为增加，负数为扣减)"`
	AmountStr  string  `json:"amountStr" dc:"高精度字符串变动金额(可选，防 float64 截断)"`
	UAmount    float64 `json:"uAmount" dc:"折合 USDT 金额 (可选，不传则内部自动根据汇率计算)"`
	RecordType int     `json:"recordType" v:"required#流水类型不能为空" dc:"账变类型 (例如: 1后台充值, 2后台扣款)"`
	Remark     string  `json:"remark" dc:"账变备注"`
}

type SubAmountRes struct {
	RecordId      int64   `json:"recordId" dc:"生成的账变流水记录ID"`
	CurrentAmount float64 `json:"currentAmount" dc:"操作后最新可用余额"`
}

type AppAssetInfo struct {
	CreateBy             interface{} `json:"createBy"`
	CreateTime           string      `json:"createTime"`
	UpdateBy             interface{} `json:"updateBy"`
	UpdateTime           string      `json:"updateTime"`
	Remark               interface{} `json:"remark"`
	UserId               int64       `json:"userId"`
	Adress               *string     `json:"adress"` // 注意拼写
	Symbol               string      `json:"symbol"`
	Amout                float64     `json:"amout"` // 注意拼写
	OccupiedAmount       float64     `json:"occupiedAmount"`
	AvailableAmount      float64     `json:"availableAmount"`
	AvailableAmountDaily float64     `json:"availableAmountDaily"`
	CodingVolumeDaily    float64     `json:"codingVolumeDaily"`
	Type                 int         `json:"type"`
	ExchageAmount        interface{} `json:"exchageAmount"`
	AdminParentIds       interface{} `json:"adminParentIds"`
	Loge                 interface{} `json:"loge"`
}

// GetAppAssetListReq 获取玩家资产列表 (业务)
type GetAppAssetListReq struct {
	g.Meta   `path:"/bussiness/asset/list" tags:"AdminAsset" method:"get" summary:"获取玩家资产列表"`
	PageNum  int    `json:"pageNum" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	UserId   int64  `json:"userId"`
	Adress   string `json:"adress"`
	Symbol   string `json:"symbol"`
	Type     int    `json:"type"`
	Params   struct {
		AmountMin          string `json:"amountMin"`
		AmountMax          string `json:"amountMax"`
		AvailableAmountMin string `json:"availableAmountMin"`
		AvailableAmountMax string `json:"availableAmountMax"`
		OccupiedAmountMin  string `json:"occupiedAmountMin"`
		OccupiedAmountMax  string `json:"occupiedAmountMax"`
		BeginTime          string `json:"beginTime"`
		EndTime            string `json:"endTime"`
	} `json:"params"`
	SearchValue string `json:"searchValue"`
}

type GetAppAssetListRes struct {
	Total int            `json:"total"`
	Rows  []AppAssetInfo `json:"data"`
}
