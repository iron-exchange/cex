package v1

import "github.com/gogf/gf/v2/frame/g"

type WalletRecordInfo struct {
	SearchValue    interface{} `json:"searchValue"`
	CreateBy       string      `json:"createBy"`
	CreateTime     string      `json:"createTime"`
	UpdateBy       interface{} `json:"updateBy"`
	UpdateTime     interface{} `json:"updateTime"`
	Remark         string      `json:"remark"`
	Id             int64       `json:"id"`
	UserId         int64       `json:"userId"`
	IsTest         string      `json:"isTest"`
	Amount         float64     `json:"amount"`
	BeforeAmount   float64     `json:"beforeAmount"`
	AfterAmount    float64     `json:"afterAmount"`
	SerialId       string      `json:"serialId"`
	Type           int         `json:"type"`
	Symbol         string      `json:"symbol"`
	StartTime      interface{} `json:"startTime"`
	EndTime        interface{} `json:"endTime"`
	MinAmount      interface{} `json:"minAmount"`
	MaxAmount      interface{} `json:"maxAmount"`
	AdminParentIds string      `json:"adminParentIds"`
	OperateTime    interface{} `json:"operateTime"`
	Uamount        float64     `json:"uamount"`
}

// GetWalletRecordListReq 账变信息列表查询
type GetWalletRecordListReq struct {
	g.Meta   `path:"/bussiness/wallet/record/list" tags:"AdminReport" method:"get" summary:"获取系统账变流水信息"`
	PageNum  int    `json:"pageNum" d:"1"`
	PageSize int    `json:"pageSize" d:"20"`
	UserId   int64  `json:"userId" dc:"根据用户ID查询"`
	Symbol   string `json:"symbol" dc:"币种筛选"`
	Type     *int   `json:"type" dc:"账变类型筛选"`
}

type GetWalletRecordListRes struct {
	Total int                `json:"total"`
	Rows  []WalletRecordInfo `json:"data"`
}

type GetWalletRecordTypesReq struct {
	g.Meta `path:"/walletRecord/types" tags:"AdminReport" method:"get" summary:"获取账变类型字典"`
}

type GetWalletRecordTypesRes map[int]string

type GetWalletStatisticsReq struct {
	g.Meta `path:"/wallet/record/statisticsAmount" tags:"AdminReport" method:"get" summary:"获取全站非测试用户总账务流水统计"`
}

type GetWalletStatisticsRes struct {
	StatisticsAmount float64 `json:"statisticsAmount" dc:"统计总金额"`
}
