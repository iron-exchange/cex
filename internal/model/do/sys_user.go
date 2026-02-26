// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta      `orm:"table:sys_user, do:true"`
	UserId      any         // 用户ID
	DeptId      any         // 部门ID
	UserName    any         // 用户账号
	NickName    any         // 用户昵称
	UserType    any         // 用户类型（0普通用户 1 组长 2 代理）
	Email       any         // 用户邮箱
	Phonenumber any         // 手机号码
	Sex         any         // 用户性别（0男 1女 2未知）
	Avatar      any         // 头像地址
	Password    any         // 密码
	Status      any         // 帐号状态（0正常 1停用）
	GoogleKey   any         // 谷歌验证码key
	DelFlag     any         // 删除标志（0代表存在 2代表删除）
	LoginIp     any         // 最后登录IP
	LoginDate   *gtime.Time // 最后登录时间
	CreateBy    any         // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新者
	UpdateTime  *gtime.Time // 更新时间
	Remark      any         // 备注
	ParentId    any         // 组长ID
}
