package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CommonConfigReq struct {
	g.Meta `path:"/common/config" tags:"Common" method:"get" summary:"获取系统基础配置"`
}

type CommonConfigRes struct {
	CustomerServiceUrl string   `json:"customerServiceUrl" dc:"在线客服地址"`
	RechargeMinAmount  string   `json:"rechargeMinAmount" dc:"起充门槛"`
	Banners            []string `json:"banners" dc:"轮播图"`
}

// GetAllSettingReq 获取全站配置参数
type GetAllSettingReq struct {
	g.Meta `path:"/common/getAllSetting" tags:"Common" method:"post,get" summary:"获取全站配置参数"`
}
type SettingInfo struct {
	Key   string `json:"key" dc:"配置键名"`
	Value string `json:"value" dc:"配置键值"`
	Desc  string `json:"desc" dc:"说明"`
}
type GetAllSettingRes struct {
	Settings map[string]string `json:"settings" dc:"KV配置集合"`
	List     []SettingInfo     `json:"list" dc:"详细配置列表"`
}

// GetAppSidebarSettingReq 获取侧边栏显示的币种配置
type GetAppSidebarSettingReq struct {
	g.Meta `path:"/common/getAppSidebarSetting" tags:"Common" method:"post,get" summary:"获取侧边栏显示的币种"`
}
type GetAppSidebarSettingRes struct {
	List []string `json:"list" dc:"侧边栏币种集合"`
}

// GetHomeCoinSettingReq 获取首页主推的币种配置
type GetHomeCoinSettingReq struct {
	g.Meta `path:"/common/getHomeCoinSetting" tags:"Common" method:"post,get" summary:"获取首页主推币种"`
}
type GetHomeCoinSettingRes struct {
	List []string `json:"list" dc:"首页主推币种集合"`
}

// GetAppCurrencyListReq 获取充值的通道与开关列表
type GetAppCurrencyListReq struct {
	g.Meta `path:"/common/getAppCurrencyList" tags:"Common" method:"post,get" summary:"获取充值的通道与开关列表"`
}
type CurrencyChannelInfo struct {
	CoinName       string `json:"coinName" dc:"通道名称, 例: USDT"`
	Type           string `json:"type" dc:"通道类型, 例: TRC20"`
	MinLimit       string `json:"minLimit" dc:"最小限额"`
	MaxLimit       string `json:"maxLimit" dc:"最大限额"`
	DepositAddress string `json:"depositAddress" dc:"平台收款地址"`
	IsOpen         int    `json:"isOpen" dc:"通道状态: 1开启 0关闭"`
}
type GetAppCurrencyListRes struct {
	List []CurrencyChannelInfo `json:"list" dc:"充值通道列表"`
}

// GetWithDrawCoinListReq 获取提现的通道与手续费列表
type GetWithDrawCoinListReq struct {
	g.Meta `path:"/common/getWithDrawCoinList" tags:"Common" method:"post,get" summary:"获取提现的通道与手续费列表"`
}
type WithdrawCoinInfo struct {
	CoinName string `json:"coinName" dc:"提现代币, 例: USDT"`
	Type     string `json:"type" dc:"网络类型, 例: TRC20"`
	MinLimit string `json:"minLimit" dc:"最小限额"`
	MaxLimit string `json:"maxLimit" dc:"最大限额"`
	FeeRate  string `json:"feeRate" dc:"提现费率"`
	IsOpen   int    `json:"isOpen" dc:"通道状态: 1开启 0关闭"`
}
type GetWithDrawCoinListRes struct {
	List []WithdrawCoinInfo `json:"list" dc:"提现通道列表"`
}

type MarketTickerWsReq struct {
	g.Meta `path:"/market/ticker/ws" tags:"Market" method:"get" summary:"行情 WebSocket 订阅"`
}
type MarketTickerWsRes struct{}
