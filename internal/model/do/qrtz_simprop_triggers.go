// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QrtzSimpropTriggers is the golang structure of table qrtz_simprop_triggers for DAO operations like Where/Data.
type QrtzSimpropTriggers struct {
	g.Meta       `orm:"table:qrtz_simprop_triggers, do:true"`
	SchedName    any // 调度名称
	TriggerName  any // qrtz_triggers表trigger_name的外键
	TriggerGroup any // qrtz_triggers表trigger_group的外键
	StrProp1     any // String类型的trigger的第一个参数
	StrProp2     any // String类型的trigger的第二个参数
	StrProp3     any // String类型的trigger的第三个参数
	IntProp1     any // int类型的trigger的第一个参数
	IntProp2     any // int类型的trigger的第二个参数
	LongProp1    any // long类型的trigger的第一个参数
	LongProp2    any // long类型的trigger的第二个参数
	DecProp1     any // decimal类型的trigger的第一个参数
	DecProp2     any // decimal类型的trigger的第二个参数
	BoolProp1    any // Boolean类型的trigger的第一个参数
	BoolProp2    any // Boolean类型的trigger的第二个参数
}
