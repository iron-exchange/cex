// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBank is the golang structure of table t_user_bank for DAO operations like Where/Data.
type UserBank struct {
	g.Meta         `orm:"table:t_user_bank, do:true"`
	Id             any         //
	UserName       any         // 姓名
	CardNumber     any         // 银行卡号
	BankName       any         // 开户银行名称
	BankAddress    any         // 开户省市
	BankBranch     any         // 开户网点
	UserId         any         // 用户名称
	AdminParentIds any         //
	CreateBy       any         //
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       any         //
	UpdateTime     *gtime.Time // 更新时间
	Remark         any         //
	SearchValue    any         //
	BankCode       any         // 银行编码
	UserAddress    any         // 用户地址
}
