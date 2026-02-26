// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QrtzSimpropTriggers is the golang structure for table qrtz_simprop_triggers.
type QrtzSimpropTriggers struct {
	SchedName    string  `json:"sched_name"    orm:"sched_name"    description:"调度名称"`
	TriggerName  string  `json:"trigger_name"  orm:"trigger_name"  description:"qrtz_triggers表trigger_name的外键"`
	TriggerGroup string  `json:"trigger_group" orm:"trigger_group" description:"qrtz_triggers表trigger_group的外键"`
	StrProp1     string  `json:"str_prop_1"    orm:"str_prop_1"    description:"String类型的trigger的第一个参数"`
	StrProp2     string  `json:"str_prop_2"    orm:"str_prop_2"    description:"String类型的trigger的第二个参数"`
	StrProp3     string  `json:"str_prop_3"    orm:"str_prop_3"    description:"String类型的trigger的第三个参数"`
	IntProp1     int     `json:"int_prop_1"    orm:"int_prop_1"    description:"int类型的trigger的第一个参数"`
	IntProp2     int     `json:"int_prop_2"    orm:"int_prop_2"    description:"int类型的trigger的第二个参数"`
	LongProp1    int64   `json:"long_prop_1"   orm:"long_prop_1"   description:"long类型的trigger的第一个参数"`
	LongProp2    int64   `json:"long_prop_2"   orm:"long_prop_2"   description:"long类型的trigger的第二个参数"`
	DecProp1     float64 `json:"dec_prop_1"    orm:"dec_prop_1"    description:"decimal类型的trigger的第一个参数"`
	DecProp2     float64 `json:"dec_prop_2"    orm:"dec_prop_2"    description:"decimal类型的trigger的第二个参数"`
	BoolProp1    string  `json:"bool_prop_1"   orm:"bool_prop_1"   description:"Boolean类型的trigger的第一个参数"`
	BoolProp2    string  `json:"bool_prop_2"   orm:"bool_prop_2"   description:"Boolean类型的trigger的第二个参数"`
}
