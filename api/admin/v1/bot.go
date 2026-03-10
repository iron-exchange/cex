package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 控盘机器人配置 (Bot Kline Model) ---------
type AdminBotKlineModelInfo struct {
	Id           int64   `json:"id"`
	Symbol       string  `json:"symbol"`
	Decline      int     `json:"decline"`     // 最大跌幅
	Increase     int     `json:"increase"`    // 最大涨幅
	Model        int     `json:"model"`       // 控盘策略
	Granularity  int     `json:"granularity"` // 粒度
	PricePencent int     `json:"pricePencent"`
	ConPrice     float64 `json:"conPrice"`
	BeginTime    string  `json:"beginTime"`
	EndTime      string  `json:"endTime"`
	CreateTime   string  `json:"createTime"`
}

type GetAdminBotKlineModelListReq struct {
	g.Meta `path:"/bot/model/list" tags:"AdminBot" method:"get" summary:"获取K线控盘机器人配置"`
	Page   int    `json:"page" d:"1"`
	Size   int    `json:"size" d:"20"`
	Symbol string `json:"symbol" dc:"交易对"`
}

type GetAdminBotKlineModelListRes struct {
	List  []AdminBotKlineModelInfo `json:"list"`
	Total int                      `json:"total"`
}

// --------- 控盘记录 (Bot Kline Model Info) ---------
type AdminBotKlineModelDataInfo struct {
	Id       int64   `json:"id"`
	ModelId  int64   `json:"modelId"`
	DateTime int64   `json:"dateTime"` // 时间戳
	Open     float64 `json:"open"`
	Close    float64 `json:"close"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
}

type GetAdminBotKlineModelDataListReq struct {
	g.Meta  `path:"/bot/model/data/list" tags:"AdminBot" method:"get" summary:"获取机器人的打点K线数据记录"`
	Page    int   `json:"page" d:"1"`
	Size    int   `json:"size" d:"20"`
	ModelId int64 `json:"modelId" dc:"模型ID"`
}

type GetAdminBotKlineModelDataListRes struct {
	List  []AdminBotKlineModelDataInfo `json:"list"`
	Total int                          `json:"total"`
}
