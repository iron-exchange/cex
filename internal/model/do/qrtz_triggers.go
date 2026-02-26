// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzTriggers is the golang structure of table qrtz_triggers for DAO operations like Where/Data.
type QrtzTriggers struct {
	g.Meta       `orm:"table:qrtz_triggers, do:true"`
	SchedName    any // 调度名称
	TriggerName  any // 触发器的名字
	TriggerGroup any // 触发器所属组的名字
	JobName      any // qrtz_job_details表job_name的外键
	JobGroup     any // qrtz_job_details表job_group的外键
	Description  any // 相关介绍
	NextFireTime any // 上一次触发时间（毫秒）
	PrevFireTime any // 下一次触发时间（默认为-1表示不触发）
	Priority     any // 优先级
	TriggerState any // 触发器状态
	TriggerType  any // 触发器的类型
	StartTime    any // 开始时间
	EndTime      any // 结束时间
	CalendarName any // 日程表名称
	MisfireInstr any // 补偿执行的策略
	JobData      any // 存放持久化job对象
}
