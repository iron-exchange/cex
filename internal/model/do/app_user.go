// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUser is the golang structure of table t_app_user for DAO operations like Where/Data.
type AppUser struct {
	g.Meta         `orm:"table:t_app_user, do:true"`
	UserId         any         //
	LoginName      any         // 姓名
	LoginPassword  any         // 登陆密码
	Phone          any         // 手机号
	IsTest         any         // 0-正常 1-测试
	Address        any         // 地址
	WalletType     any         // 地址类型 ETH TRC
	Status         any         // 0正常1冻结
	TotleAmont     any         // 总打码量
	RechargeAmont  any         // 充值打码量
	Buff           any         // 0正常 1包赢 2包输
	AppParentIds   any         // app代理ids
	AdminParentIds any         // 后台代理ids
	ActiveCode     any         // 邀请码
	RegisterIp     any         // 注册ip
	Host           any         // 注册域名
	Email          any         // 邮箱
	Level          any         // vip等级
	IsFreeze       any         // 是否冻结  1=正常 2=冻结
	CreateBy       any         // 创建人
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       any         // 更新人
	UpdateTime     *gtime.Time // 更新时间
	Remark         any         // 备注
	SearchValue    any         //
	IsBlack        any         // 黑名单 1=正常 2拉黑
	BinanceEmail   any         // 币安子账号邮箱
}
