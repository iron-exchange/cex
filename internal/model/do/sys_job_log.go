// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJobLog is the golang structure of table sys_job_log for DAO operations like Where/Data.
type SysJobLog struct {
	g.Meta        `orm:"table:sys_job_log, do:true"`
	JobLogId      any         // 任务日志ID
	JobName       any         // 任务名称
	JobGroup      any         // 任务组名
	InvokeTarget  any         // 调用目标字符串
	JobMessage    any         // 日志信息
	Status        any         // 执行状态（0正常 1失败）
	ExceptionInfo any         // 异常信息
	CreateTime    *gtime.Time // 创建时间
}
