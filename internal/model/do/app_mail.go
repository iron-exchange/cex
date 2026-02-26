// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppMail is the golang structure of table t_app_mail for DAO operations like Where/Data.
type AppMail struct {
	g.Meta      `orm:"table:t_app_mail, do:true"`
	Id          any         //
	UserId      any         //
	Title       any         // 标题
	Content     any         // 内容
	Type        any         // 消息类型 1=普通消息 2=全站消息
	Status      any         // 状态（0 未读 1已读）
	OpertorId   any         // 操作人
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	SearchValue any         //
	DelFlag     any         // 0正常 2删除
}
