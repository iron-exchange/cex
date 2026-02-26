// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJobLog is the golang structure for table sys_job_log.
type SysJobLog struct {
	JobLogId      int64       `json:"job_log_id"     orm:"job_log_id"     description:"任务日志ID"`
	JobName       string      `json:"job_name"       orm:"job_name"       description:"任务名称"`
	JobGroup      string      `json:"job_group"      orm:"job_group"      description:"任务组名"`
	InvokeTarget  string      `json:"invoke_target"  orm:"invoke_target"  description:"调用目标字符串"`
	JobMessage    string      `json:"job_message"    orm:"job_message"    description:"日志信息"`
	Status        string      `json:"status"         orm:"status"         description:"执行状态（0正常 1失败）"`
	ExceptionInfo string      `json:"exception_info" orm:"exception_info" description:"异常信息"`
	CreateTime    *gtime.Time `json:"create_time"    orm:"create_time"    description:"创建时间"`
}
