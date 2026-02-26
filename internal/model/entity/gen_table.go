// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTable is the golang structure for table gen_table.
type GenTable struct {
	TableId        int64       `json:"table_id"          orm:"table_id"          description:"编号"`
	TableName      string      `json:"table_name"        orm:"table_name"        description:"表名称"`
	TableComment   string      `json:"table_comment"     orm:"table_comment"     description:"表描述"`
	SubTableName   string      `json:"sub_table_name"    orm:"sub_table_name"    description:"关联子表的表名"`
	SubTableFkName string      `json:"sub_table_fk_name" orm:"sub_table_fk_name" description:"子表关联的外键名"`
	ClassName      string      `json:"class_name"        orm:"class_name"        description:"实体类名称"`
	TplCategory    string      `json:"tpl_category"      orm:"tpl_category"      description:"使用的模板（crud单表操作 tree树表操作）"`
	PackageName    string      `json:"package_name"      orm:"package_name"      description:"生成包路径"`
	ModuleName     string      `json:"module_name"       orm:"module_name"       description:"生成模块名"`
	BusinessName   string      `json:"business_name"     orm:"business_name"     description:"生成业务名"`
	FunctionName   string      `json:"function_name"     orm:"function_name"     description:"生成功能名"`
	FunctionAuthor string      `json:"function_author"   orm:"function_author"   description:"生成功能作者"`
	GenType        string      `json:"gen_type"          orm:"gen_type"          description:"生成代码方式（0zip压缩包 1自定义路径）"`
	GenPath        string      `json:"gen_path"          orm:"gen_path"          description:"生成路径（不填默认项目路径）"`
	Options        string      `json:"options"           orm:"options"           description:"其它生成选项"`
	CreateBy       string      `json:"create_by"         orm:"create_by"         description:"创建者"`
	CreateTime     *gtime.Time `json:"create_time"       orm:"create_time"       description:"创建时间"`
	UpdateBy       string      `json:"update_by"         orm:"update_by"         description:"更新者"`
	UpdateTime     *gtime.Time `json:"update_time"       orm:"update_time"       description:"更新时间"`
	Remark         string      `json:"remark"            orm:"remark"            description:"备注"`
}
