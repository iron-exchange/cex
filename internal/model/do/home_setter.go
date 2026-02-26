// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomeSetter is the golang structure of table t_home_setter for DAO operations like Where/Data.
type HomeSetter struct {
	g.Meta       `orm:"table:t_home_setter, do:true"`
	Id           any         //
	Title        any         // 标题
	Author       any         // 作者
	Content      any         // 内容
	CreateTime   *gtime.Time // 创建时间
	ImgUrl       any         // 图片地址
	Sort         any         // 排序
	IsShow       any         // 是否展示 0展示  2不展示
	LanguageName any         // 语言
	LikesNum     any         // 点赞数
	HomeType     any         // 类型（0 首页文本  1 问题列表）
	ModelType    any         // 功能（0=首页  1=defi挖矿 2=助力贷）
	SearchValue  any         //
}
