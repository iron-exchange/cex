// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppuserLoginLog is the golang structure of table t_appuser_login_log for DAO operations like Where/Data.
type AppuserLoginLog struct {
	g.Meta        `orm:"table:t_appuser_login_log, do:true"`
	Id            any         // 主键ID
	UserId        any         // 登录用户ID
	Username      any         // 登录用户名
	Ipaddr        any         // 访问IP
	LoginLocation any         // 访问位置
	Browser       any         // 浏览器
	Os            any         // 系统OS
	Status        any         // 登录状态（0成功 1失败）
	Msg           any         //
	LoginTime     *gtime.Time // 访问时间
}
