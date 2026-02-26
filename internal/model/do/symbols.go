// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Symbols is the golang structure of table t_symbols for DAO operations like Where/Data.
type Symbols struct {
	g.Meta   `orm:"table:t_symbols, do:true"`
	Id       any //
	Slug     any // 币种名称（ID）
	Symbol   any // 币种符号
	Fullname any // 币种全称
	LogoUrl  any // 图标链接
	Fiat     any // 是否法定货币
}
