// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzSchedulerState is the golang structure of table qrtz_scheduler_state for DAO operations like Where/Data.
type QrtzSchedulerState struct {
	g.Meta          `orm:"table:qrtz_scheduler_state, do:true"`
	SchedName       any // 调度名称
	InstanceName    any // 实例名称
	LastCheckinTime any // 上次检查时间
	CheckinInterval any // 检查间隔时间
}
