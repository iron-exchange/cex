package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CollectionOrderInfo struct {
	Id          int         `json:"id"`
	OrderId     string      `json:"orderId"`
	UserId      int64       `json:"userId"`
	Address     string      `json:"address"`
	Chain       string      `json:"chain"`
	Coin        string      `json:"coin"`
	Hash        string      `json:"hash"`
	Amount      float64     `json:"amount"`
	Status      string      `json:"status"`
	ClientName  string      `json:"clientName"`
	CreateTime  *gtime.Time `json:"createTime"`
	CreateBy    string      `json:"createBy"`
	UpdateTime  *gtime.Time `json:"updateTime"`
	UpdateBy    interface{} `json:"updateBy"`
	Remark      string      `json:"remark"`
	SearchValue string      `json:"searchValue"`
}

// GetCollectionOrderListReq 获取归集订单列表
type GetCollectionOrderListReq struct {
	g.Meta      `path:"/collectionOrder/list" tags:"AdminCollection" method:"get" summary:"获取归集订单列表"`
	PageNum     int     `json:"pageNum" d:"1"`
	PageSize    int     `json:"pageSize" d:"10"`
	OrderId     string  `json:"orderId"`
	UserId      int64   `json:"userId"`
	Address     string  `json:"address"`
	Chain       string  `json:"chain"`
	Coin        string  `json:"coin"`
	Hash        string  `json:"hash"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	ClientName  string  `json:"clientName"`
	SearchValue string  `json:"searchValue"`
}

type GetCollectionOrderListRes struct {
	Total int                   `json:"total"`
	Rows  []CollectionOrderInfo `json:"data"`
}
