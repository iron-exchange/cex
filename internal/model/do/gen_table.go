// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTable is the golang structure of table gen_table for DAO operations like Where/Data.
type GenTable struct {
	g.Meta         `orm:"table:gen_table, do:true"`
	TableId        any         // 编号
	TableName      any         // 表名称
	TableComment   any         // 表描述
	SubTableName   any         // 关联子表的表名
	SubTableFkName any         // 子表关联的外键名
	ClassName      any         // 实体类名称
	TplCategory    any         // 使用的模板（crud单表操作 tree树表操作）
	PackageName    any         // 生成包路径
	ModuleName     any         // 生成模块名
	BusinessName   any         // 生成业务名
	FunctionName   any         // 生成功能名
	FunctionAuthor any         // 生成功能作者
	GenType        any         // 生成代码方式（0zip压缩包 1自定义路径）
	GenPath        any         // 生成路径（不填默认项目路径）
	Options        any         // 其它生成选项
	CreateBy       any         // 创建者
	CreateTime     *gtime.Time // 创建时间
	UpdateBy       any         // 更新者
	UpdateTime     *gtime.Time // 更新时间
	Remark         any         // 备注
}
