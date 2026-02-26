// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzCalendars is the golang structure for table qrtz_calendars.
type QrtzCalendars struct {
	SchedName    string `json:"sched_name"    orm:"sched_name"    description:"调度名称"`
	CalendarName string `json:"calendar_name" orm:"calendar_name" description:"日历名称"`
	Calendar     string `json:"calendar"      orm:"calendar"      description:"存放持久化calendar对象"`
}
