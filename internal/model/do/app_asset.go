// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppAsset is the golang structure of table t_app_asset for DAO operations like Where/Data.
type AppAsset struct {
	g.Meta               `orm:"table:t_app_asset, do:true"`
	Id                   any         //
	UserId               any         //
	Adress               any         // 地址
	Symbol               any         // 币种
	Amout                any         // 资产总额
	OccupiedAmount       any         // 占用资产
	AvailableAmount      any         // 可用资产
	AvailableAmountDaily any         // 每日余额（0点时分的余额，提现会减少）
	CodingVolumeDaily    any         // 每日打码量（24点之前，提现会减少）
	Type                 any         // 资产类型 1=平台资产 2=理财资产 3=合约账户
	CreateBy             any         // 创建人
	CreateTime           *gtime.Time // 创建时间
	UpdateBy             any         // 更新人
	UpdateTime           *gtime.Time // 更新时间
	Remark               any         // 备注
	SearchValue          any         //
}
