// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzCronTriggers is the golang structure of table qrtz_cron_triggers for DAO operations like Where/Data.
type QrtzCronTriggers struct {
	g.Meta         `orm:"table:qrtz_cron_triggers, do:true"`
	SchedName      any // 调度名称
	TriggerName    any // qrtz_triggers表trigger_name的外键
	TriggerGroup   any // qrtz_triggers表trigger_group的外键
	CronExpression any // cron表达式
	TimeZoneId     any // 时区
}
