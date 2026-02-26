// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MineFinancial is the golang structure of table t_mine_financial for DAO operations like Where/Data.
type MineFinancial struct {
	g.Meta              `orm:"table:t_mine_financial, do:true"`
	Id                  any         //
	Title               any         // 标题
	Icon                any         // 图标
	Status              any         // 启用禁用(展示在前端)1开0关
	Days                any         // 天数(如 7,10,30)
	DefaultOdds         any         // 违约利率
	MinOdds             any         // 最小日利率百分比
	MaxOdds             any         // 最大日利率百分比
	TimeLimit           any         // 每人限购次数，0表示不限
	LimitMin            any         // 最小金额
	LimitMax            any         // 最大金额
	IsHot               any         // 是否热销1是0否
	Sort                any         // 排序
	CreateBy            any         // 创建人
	CreateTime          *gtime.Time // 创建时间
	UpdateBy            any         // 更新人员
	UpdateTime          *gtime.Time // 更新时间
	BuyPurchase         any         // 购买次数
	AvgRate             any         // 日平均利率
	Coin                any         // 币种
	Classify            any         // 分类（0 普通  1 vip  2 增值）
	BasicInvestAmount   any         // 平台基础投资金额
	TotalInvestAmount   any         // 平台总投资额
	Level               any         // VIP等级
	Process             any         // 项目进度
	RemainAmount        any         // 剩余金额
	Remark              any         // 标签
	PurchasedAmount     any         // 易购金额
	Problem             any         // 常见问题
	ProdectIntroduction any         // 产品介绍
}
