package v1

import "github.com/gogf/gf/v2/frame/g"

type CollectionOrderInfo struct {
	Id         int     `json:"id"`
	UserId     int64   `json:"userId"`
	Amount     float64 `json:"amount"`
	Coin       string  `json:"coin"`
	Address    string  `json:"address"`
	TxHash     string  `json:"txHash"`
	Status     string  `json:"status"`
	CreateTime string  `json:"createTime"`
}

// GetCollectionOrderListReq 获取归集订单列表
type GetCollectionOrderListReq struct {
	g.Meta `path:"/collectionOrder/list" tags:"AdminCollection" method:"get" summary:"获取资金归集订单列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	UserId int64  `json:"userId" dc:"按用户ID搜索"`
	Status *int   `json:"status" dc:"按状态搜索"`
	TxHash string `json:"txHash" dc:"按交易哈希搜索"`
}

type GetCollectionOrderListRes struct {
	List  []CollectionOrderInfo `json:"list"`
	Total int                   `json:"total"`
}
