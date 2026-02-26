// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MingProductUser is the golang structure of table t_ming_product_user for DAO operations like Where/Data.
type MingProductUser struct {
	g.Meta      `orm:"table:t_ming_product_user, do:true"`
	Id          any         //
	ProductId   any         // 产品id
	AppUserId   any         // 玩家用户id
	PledgeNum   any         // 限购次数
	CreateBy    any         // 创建人
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    any         // 更新人员
	UpdateTime  *gtime.Time // 更新时间
	SearchValue any         // 币种
	Remark      any         // 标签
}
