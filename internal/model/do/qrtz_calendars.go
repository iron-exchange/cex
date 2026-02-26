// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzCalendars is the golang structure of table qrtz_calendars for DAO operations like Where/Data.
type QrtzCalendars struct {
	g.Meta       `orm:"table:qrtz_calendars, do:true"`
	SchedName    any // 调度名称
	CalendarName any // 日历名称
	Calendar     any // 存放持久化calendar对象
}
