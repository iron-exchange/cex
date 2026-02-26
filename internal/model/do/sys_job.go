// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJob is the golang structure of table sys_job for DAO operations like Where/Data.
type SysJob struct {
	g.Meta         `orm:"table:sys_job, do:true"`
	JobId          any         // 任务ID
	JobName        any         // 任务名称
	JobGroup       any         // 任务组名
	InvokeTarget   any         // 调用目标字符串
	CronExpression any         // cron执行表达式
	MisfirePolicy  any         // 计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     any         // 是否并发执行（0允许 1禁止）
	Status         any         // 状态（0正常 1暂停）
	CreateBy       any         // 创建者
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       any         // 更新者
	UpdateTime     *gtime.Time // 更新时间
	Remark         any         // 备注信息
}
