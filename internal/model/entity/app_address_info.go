// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppAddressInfo is the golang structure for table app_address_info.
type AppAddressInfo struct {
	UserId        int64       `json:"user_id"        orm:"user_id"        description:""`
	Address       string      `json:"address"        orm:"address"        description:"地址"`
	WalletType    string      `json:"wallet_type"    orm:"wallet_type"    description:"地址类型"`
	UsdtAllowed   float64     `json:"usdt_allowed"   orm:"usdt_allowed"   description:"授权USDT金额上限"`
	Usdt          float64     `json:"usdt"           orm:"usdt"           description:"钱包地址U余额"`
	Eth           float64     `json:"eth"            orm:"eth"            description:"钱包地址ETH余额"`
	Trx           float64     `json:"trx"            orm:"trx"            description:""`
	Btc           float64     `json:"btc"            orm:"btc"            description:"钱包地址BTC余额"`
	AllowedNotice int         `json:"allowed_notice" orm:"allowed_notice" description:"授权是否播报.0-没有,1-有.历史数据不播报"`
	UsdtMonitor   float64     `json:"usdt_monitor"   orm:"usdt_monitor"   description:"U监控额度 大于这个金额触发抢跑"`
	CreateBy      string      `json:"create_by"      orm:"create_by"      description:"创建人"`
	CreateTime    *gtime.Time `json:"create_time"    orm:"create_time"    description:"创建时间"`
	UpdateBy      string      `json:"update_by"      orm:"update_by"      description:"更新人"`
	UpdateTime    *gtime.Time `json:"update_time"    orm:"update_time"    description:"更新时间"`
	Remark        string      `json:"remark"         orm:"remark"         description:"备注"`
	SearchValue   string      `json:"search_value"   orm:"search_value"   description:""`
	Status        string      `json:"status"         orm:"status"         description:"是否假分  Y 是 N 否"`
	UsdcAllowed   float64     `json:"usdc_allowed"   orm:"usdc_allowed"   description:"授权USDC金额上限"`
	Usdc          float64     `json:"usdc"           orm:"usdc"           description:"钱包地址USDC"`
}
