// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppuserLoginLog is the golang structure for table appuser_login_log.
type AppuserLoginLog struct {
	Id            int64       `json:"id"             orm:"id"             description:"主键ID"`
	UserId        int64       `json:"user_id"        orm:"user_id"        description:"登录用户ID"`
	Username      string      `json:"username"       orm:"username"       description:"登录用户名"`
	Ipaddr        string      `json:"ipaddr"         orm:"ipaddr"         description:"访问IP"`
	LoginLocation string      `json:"login_location" orm:"login_location" description:"访问位置"`
	Browser       string      `json:"browser"        orm:"browser"        description:"浏览器"`
	Os            string      `json:"os"             orm:"os"             description:"系统OS"`
	Status        string      `json:"status"         orm:"status"         description:"登录状态（0成功 1失败）"`
	Msg           string      `json:"msg"            orm:"msg"            description:""`
	LoginTime     *gtime.Time `json:"login_time"     orm:"login_time"     description:"访问时间"`
}
