// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTableColumn is the golang structure of table gen_table_column for DAO operations like Where/Data.
type GenTableColumn struct {
	g.Meta        `orm:"table:gen_table_column, do:true"`
	ColumnId      any         // 编号
	TableId       any         // 归属表编号
	ColumnName    any         // 列名称
	ColumnComment any         // 列描述
	ColumnType    any         // 列类型
	JavaType      any         // JAVA类型
	JavaField     any         // JAVA字段名
	IsPk          any         // 是否主键（1是）
	IsIncrement   any         // 是否自增（1是）
	IsRequired    any         // 是否必填（1是）
	IsInsert      any         // 是否为插入字段（1是）
	IsEdit        any         // 是否编辑字段（1是）
	IsList        any         // 是否列表字段（1是）
	IsQuery       any         // 是否查询字段（1是）
	QueryType     any         // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      any         // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      any         // 字典类型
	Sort          any         // 排序
	CreateBy      any         // 创建者
	CreateTime    *gtime.Time // 创建时间
	UpdateBy      any         // 更新者
	UpdateTime    *gtime.Time // 更新时间
}
