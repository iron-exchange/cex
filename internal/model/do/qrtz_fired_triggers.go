// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzFiredTriggers is the golang structure of table qrtz_fired_triggers for DAO operations like Where/Data.
type QrtzFiredTriggers struct {
	g.Meta           `orm:"table:qrtz_fired_triggers, do:true"`
	SchedName        any // 调度名称
	EntryId          any // 调度器实例id
	TriggerName      any // qrtz_triggers表trigger_name的外键
	TriggerGroup     any // qrtz_triggers表trigger_group的外键
	InstanceName     any // 调度器实例名
	FiredTime        any // 触发的时间
	SchedTime        any // 定时器制定的时间
	Priority         any // 优先级
	State            any // 状态
	JobName          any // 任务名称
	JobGroup         any // 任务组名
	IsNonconcurrent  any // 是否并发
	RequestsRecovery any // 是否接受恢复执行
}
