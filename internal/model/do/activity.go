// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Activity is the golang structure of table t_activity for DAO operations like Where/Data.
type Activity struct {
	g.Meta       `orm:"table:t_activity, do:true"`
	Id           any         // 卡ID
	CreateBy     any         //
	CreateTime   *gtime.Time // 创建时间
	UpdateBy     any         //
	UpdateTime   *gtime.Time // 更新时间
	SearchValue  any         //
	Status       any         // 状态0编辑1上线
	Remark       any         //
	DisplayUrl   any         //
	DetailUrl    any         // 详情页
	DetailStatus any         //
	LoadingUrl   any         //
	Tags         any         //
	Type         any         // 1广告2活动
	Source       any         // 展示端,1:PC,2:H5
	LanguageId   any         // zh:1,cht:2,en:3,pt:4,sa:5,ko:6,ja:7,es:8,th:9,ms:10,id:11,fr:12,ru:13
	Agent        any         // 0代理无关，1相关
	JumpType     any         // 跳转类型 0 内链 1外链
	SortNum      any         // 排序
}
