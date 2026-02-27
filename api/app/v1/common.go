package v1

import "github.com/gogf/gf/v2/frame/g"

type CommonConfigReq struct {
	g.Meta `path:"/common/config" tags:"Common" method:"get" summary:"获取系统基础配置"`
}

type CommonConfigRes struct {
	CustomerServiceUrl string   `json:"customerServiceUrl" dc:"在线客服地址"`
	RechargeMinAmount  string   `json:"rechargeMinAmount" dc:"起充门槛"`
	Banners            []string `json:"banners" dc:"轮播图"`
}

type MarketTickerWsReq struct {
	g.Meta `path:"/market/ticker/ws" tags:"Market" method:"get" summary:"行情 WebSocket 订阅"`
}

type MarketTickerWsRes struct{}
