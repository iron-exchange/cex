// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectionOrder is the golang structure of table t_collection_order for DAO operations like Where/Data.
type CollectionOrder struct {
	g.Meta      `orm:"table:t_collection_order, do:true"`
	Id          any         // 主键ID
	OrderId     any         // 订单号
	UserId      any         // 用户ID
	Address     any         // 归集地址
	Chain       any         // 地址类型
	Hash        any         // hash
	Coin        any         // 币种
	Amount      any         // 归集金额
	Status      any         // 1  进行中   2 归集成功  3 归集失败
	ClientName  any         // 客户端名称
	CreateTime  *gtime.Time // 创建时间
	CreateBy    any         // 创建人
	UpdateTime  *gtime.Time // 修改时间
	UpdateBy    *gtime.Time // 修改人
	Remark      any         // 备注
	SearchValue any         //
}
