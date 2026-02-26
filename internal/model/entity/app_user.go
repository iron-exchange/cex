// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUser is the golang structure for table app_user.
type AppUser struct {
	UserId         int64       `json:"user_id"          orm:"user_id"          description:""`
	LoginName      string      `json:"login_name"       orm:"login_name"       description:"姓名"`
	LoginPassword  string      `json:"login_password"   orm:"login_password"   description:"登陆密码"`
	Phone          string      `json:"phone"            orm:"phone"            description:"手机号"`
	IsTest         int         `json:"is_test"          orm:"is_test"          description:"0-正常 1-测试"`
	Address        string      `json:"address"          orm:"address"          description:"地址"`
	WalletType     string      `json:"wallet_type"      orm:"wallet_type"      description:"地址类型 ETH TRC"`
	Status         int         `json:"status"           orm:"status"           description:"0正常1冻结"`
	TotleAmont     float64     `json:"totle_amont"      orm:"totle_amont"      description:"总打码量"`
	RechargeAmont  float64     `json:"recharge_amont"   orm:"recharge_amont"   description:"充值打码量"`
	Buff           int         `json:"buff"             orm:"buff"             description:"0正常 1包赢 2包输"`
	AppParentIds   string      `json:"app_parent_ids"   orm:"app_parent_ids"   description:"app代理ids"`
	AdminParentIds string      `json:"admin_parent_ids" orm:"admin_parent_ids" description:"后台代理ids"`
	ActiveCode     string      `json:"active_code"      orm:"active_code"      description:"邀请码"`
	RegisterIp     string      `json:"register_ip"      orm:"register_ip"      description:"注册ip"`
	Host           string      `json:"host"             orm:"host"             description:"注册域名"`
	Email          string      `json:"email"            orm:"email"            description:"邮箱"`
	Level          int         `json:"level"            orm:"level"            description:"vip等级"`
	IsFreeze       string      `json:"is_freeze"        orm:"is_freeze"        description:"是否冻结  1=正常 2=冻结"`
	CreateBy       string      `json:"create_by"        orm:"create_by"        description:"创建人"`
	CreateTime     *gtime.Time `json:"create_time"      orm:"create_time"      description:"创建时间"`
	UpdateBy       string      `json:"update_by"        orm:"update_by"        description:"更新人"`
	UpdateTime     *gtime.Time `json:"update_time"      orm:"update_time"      description:"更新时间"`
	Remark         string      `json:"remark"           orm:"remark"           description:"备注"`
	SearchValue    string      `json:"search_value"     orm:"search_value"     description:""`
	IsBlack        int         `json:"is_black"         orm:"is_black"         description:"黑名单 1=正常 2拉黑"`
	BinanceEmail   string      `json:"binance_email"    orm:"binance_email"    description:"币安子账号邮箱"`
}
