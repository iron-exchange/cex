// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppAddressInfo is the golang structure of table t_app_address_info for DAO operations like Where/Data.
type AppAddressInfo struct {
	g.Meta        `orm:"table:t_app_address_info, do:true"`
	UserId        any         //
	Address       any         // 地址
	WalletType    any         // 地址类型
	UsdtAllowed   any         // 授权USDT金额上限
	Usdt          any         // 钱包地址U余额
	Eth           any         // 钱包地址ETH余额
	Trx           any         //
	Btc           any         // 钱包地址BTC余额
	AllowedNotice any         // 授权是否播报.0-没有,1-有.历史数据不播报
	UsdtMonitor   any         // U监控额度 大于这个金额触发抢跑
	CreateBy      any         // 创建人
	CreateTime    *gtime.Time // 创建时间
	UpdateBy      any         // 更新人
	UpdateTime    *gtime.Time // 更新时间
	Remark        any         // 备注
	SearchValue   any         //
	Status        any         // 是否假分  Y 是 N 否
	UsdcAllowed   any         // 授权USDC金额上限
	Usdc          any         // 钱包地址USDC
}
