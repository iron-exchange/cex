package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 币种列表 (IEO Coin List) ---------
type OwnCoinInfo struct {
	Id            int64   `json:"id"`
	Coin          string  `json:"coin"`
	Logo          string  `json:"logo"`
	ReferCoin     string  `json:"referCoin"`
	ShowSymbol    string  `json:"showSymbol"`
	Price         float64 `json:"price"`
	Proportion    float64 `json:"proportion"`
	RaisingAmount float64 `json:"raisingAmount"`
	RaisedAmount  float64 `json:"raisedAmount"`
	PurchaseLimit int     `json:"purchaseLimit"`
	TotalAmount   float64 `json:"totalAmount"`
	Status        int     `json:"status"` // 1.未发布  2.筹集中 3 筹集成功 4.筹集失败
	BeginTime     string  `json:"beginTime"`
	EndTime       string  `json:"endTime"`
}

type GetOwnCoinListReq struct {
	g.Meta `path:"/ieo/coin/list" tags:"AdminIEO" method:"get" summary:"获取IEO新币列表"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Coin   string `json:"coin" dc:"币种名称过滤"`
}

type GetOwnCoinListRes struct {
	List  []OwnCoinInfo `json:"list"`
	Total int           `json:"total"`
}

// --------- 申购订单 (IEO Subscription Orders) ---------
type OwnCoinSubscribeOrderInfo struct {
	Id          int64   `json:"id"`
	SubscribeId string  `json:"subscribeId"`
	UserId      int64   `json:"userId"`
	OrderId     string  `json:"orderId"`
	OwnCoin     string  `json:"ownCoin"`
	AmountLimit float64 `json:"amountLimit"`
	NumLimit    int     `json:"numLimit"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"` // 1订阅中、2订阅成功、3成功消息推送完成
	CreateTime  string  `json:"createTime"`
}

type GetOwnCoinSubscribeOrderListReq struct {
	g.Meta  `path:"/ieo/order/list" tags:"AdminIEO" method:"get" summary:"获取新币申购订单记录"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int64  `json:"userId" dc:"用户ID"`
	OwnCoin string `json:"ownCoin" dc:"申购币种"`
	Status  string `json:"status" dc:"状态"`
}

type GetOwnCoinSubscribeOrderListRes struct {
	List  []OwnCoinSubscribeOrderInfo `json:"list"`
	Total int                         `json:"total"`
}

// --------- 申购订单 (IEO OwnCoin Orders) ---------
type OwnCoinOrderInfo struct {
	Id             int64   `json:"id"`
	UserId         int64   `json:"userId"`
	OrderId        string  `json:"orderId"`
	OwnId          int64   `json:"ownId"`
	OwnCoin        string  `json:"ownCoin"`
	Number         int     `json:"number"`
	Price          float64 `json:"price"`
	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
	AdminUserIds   string  `json:"adminUserIds"`
	AdminParentIds string  `json:"adminParentIds"`
	CreateTime     string  `json:"createTime"`
	UpdateTime     string  `json:"updateTime"`
	CreateBy       string  `json:"createBy"`
	UpdateBy       string  `json:"updateBy"`
	Remark         string  `json:"remark"`
}

type GetOwnCoinOrderListReq struct {
	g.Meta   `path:"/bussiness/ownCoinOrder/list" tags:"AdminIEO" method:"get" summary:"获取IEO认购订单列表"`
	PageNum  int    `json:"pageNum" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	UserId   int64  `json:"userId" dc:"用户ID"`
	OrderId  string `json:"orderId" dc:"订单号"`
	OwnId    int64  `json:"ownId" dc:"新币ID"`
	Status   string `json:"status" dc:"状态"`
}

type GetOwnCoinOrderListRes struct {
	Rows  []OwnCoinOrderInfo `json:"data"`
	Total int                `json:"total"`
}

type EditOwnCoinOrderPlacingReq struct {
	g.Meta `path:"/bussiness/ownCoinOrder/editPlacing" tags:"AdminIEO" method:"post" summary:"审批/调整新币申购订单"`
	Id     int64 `json:"id" v:"required#订单ID不能为空"`
	Number int   `json:"number" v:"required#调整后的数量不能为空"`
}

type GetOwnCoinOrderReq struct {
	g.Meta `path:"/bussiness/ownCoinOrder/{id}" tags:"AdminIEO" method:"get" summary:"获取新币申购订单详情"`
	Id     int64 `path:"id"`
}

type CreateOwnCoinOrderReq struct {
	g.Meta `path:"/bussiness/ownCoinOrder" tags:"AdminIEO" method:"post" summary:"管理员手动补单"`
	UserId int64   `json:"userId" v:"required#用户ID不能为空"`
	OwnId  int64   `json:"ownId" v:"required#新币ID不能为空"`
	Number int     `json:"number" v:"required#申购数量不能为空"`
	Price  float64 `json:"price" v:"required#申购价格不能为空"`
	Remark string  `json:"remark"`
}

type DeleteOwnCoinOrderReq struct {
	g.Meta `path:"/bussiness/ownCoinOrder/{ids}" tags:"AdminIEO" method:"delete" summary:"批量删除申购记录"`
	Ids    []int64 `path:"ids" v:"required#ID数组不能为空"`
}

type ExportOwnCoinOrderReq struct {
	g.Meta  `path:"/bussiness/ownCoinOrder/export" tags:"AdminIEO" method:"post" summary:"导出申购明细"`
	UserId  int64  `json:"userId"`
	OrderId string `json:"orderId"`
	OwnId   int64  `json:"ownId"`
	Status  string `json:"status"`
}
