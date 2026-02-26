// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LoadOrder is the golang structure of table t_load_order for DAO operations like Where/Data.
type LoadOrder struct {
	g.Meta         `orm:"table:t_load_order, do:true"`
	Id             any         // 主键
	ProId          any         // 贷款商品表id
	UserId         any         // 用户id
	CreateTime     *gtime.Time // 购买时间
	Amount         any         // 贷款金额
	Rate           any         // 贷款利率
	Interest       any         // 利息
	Status         any         // 0=待审核 1=审核通过  2=审核拒绝  3=已结清  4=已逾期
	FinalRepayTime *gtime.Time // 最后还款日
	DisburseTime   *gtime.Time // 放款日期
	ReturnTime     *gtime.Time // 还款日期
	DisburseAmount any         // 审批金额
	AdminParentIds any         // 后台代理ids
	CardUrl        any         // 手持身份证
	CardBackUrl    any         // 身份证正面
	CapitalUrl     any         // 身份证反面
	LicenseUrl     any         //
	OrderNo        any         //
	CycleType      any         // 还款周期  多少天
	Remark         any         // 用户备注
	CreateBy       any         //
	UpdateBy       any         //
	UpdateTime     *gtime.Time // 更新时间
	SearchValue    any         //
}
