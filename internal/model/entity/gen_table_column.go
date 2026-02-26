// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTableColumn is the golang structure for table gen_table_column.
type GenTableColumn struct {
	ColumnId      int64       `json:"column_id"      orm:"column_id"      description:"编号"`
	TableId       string      `json:"table_id"       orm:"table_id"       description:"归属表编号"`
	ColumnName    string      `json:"column_name"    orm:"column_name"    description:"列名称"`
	ColumnComment string      `json:"column_comment" orm:"column_comment" description:"列描述"`
	ColumnType    string      `json:"column_type"    orm:"column_type"    description:"列类型"`
	JavaType      string      `json:"java_type"      orm:"java_type"      description:"JAVA类型"`
	JavaField     string      `json:"java_field"     orm:"java_field"     description:"JAVA字段名"`
	IsPk          string      `json:"is_pk"          orm:"is_pk"          description:"是否主键（1是）"`
	IsIncrement   string      `json:"is_increment"   orm:"is_increment"   description:"是否自增（1是）"`
	IsRequired    string      `json:"is_required"    orm:"is_required"    description:"是否必填（1是）"`
	IsInsert      string      `json:"is_insert"      orm:"is_insert"      description:"是否为插入字段（1是）"`
	IsEdit        string      `json:"is_edit"        orm:"is_edit"        description:"是否编辑字段（1是）"`
	IsList        string      `json:"is_list"        orm:"is_list"        description:"是否列表字段（1是）"`
	IsQuery       string      `json:"is_query"       orm:"is_query"       description:"是否查询字段（1是）"`
	QueryType     string      `json:"query_type"     orm:"query_type"     description:"查询方式（等于、不等于、大于、小于、范围）"`
	HtmlType      string      `json:"html_type"      orm:"html_type"      description:"显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）"`
	DictType      string      `json:"dict_type"      orm:"dict_type"      description:"字典类型"`
	Sort          int         `json:"sort"           orm:"sort"           description:"排序"`
	CreateBy      string      `json:"create_by"      orm:"create_by"      description:"创建者"`
	CreateTime    *gtime.Time `json:"create_time"    orm:"create_time"    description:"创建时间"`
	UpdateBy      string      `json:"update_by"      orm:"update_by"      description:"更新者"`
	UpdateTime    *gtime.Time `json:"update_time"    orm:"update_time"    description:"更新时间"`
}
