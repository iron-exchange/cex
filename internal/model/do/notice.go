// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure of table t_notice for DAO operations like Where/Data.
type Notice struct {
	g.Meta        `orm:"table:t_notice, do:true"`
	NoticeId      any         // 公告ID
	NoticeTitle   any         // 标题
	NoticeType    any         // 公告类型 1=公告信息 2=活动公告 3=首页滚动公告
	ModelType     any         // 模块类型 1=公告信息 2=活动公告 3=首页滚动公告   1={1=链接弹窗 2=图文弹窗},   2={1=首页轮播活动 2=Defi挖矿活动图},  3={1=首页滚动公告}  注:没有二级的默认给1  二级联动
	NoticeContent any         // 内容
	CommentsNum   any         // 评论数
	Cover         any         // 图片
	ViewNum       any         // 浏览数
	ExpireTime    *gtime.Time // 公告截止时间
	ImgUrl        any         // 图片地址
	ChainedUrl    any         // 链接地址
	DetailUrl     any         // 详情页
	LanguageId    any         // zh:1,cht:2,en:3,pt:4,sa:5,ko:6,ja:7,es:8,th:9,ms:10,id:11,fr:12,ru:13
	Status        any         // 公告状态（0正常 1关闭）
	Source        any         // 展示端1=pc 2=h5
	Sort          any         // 排序
	Remark        any         // 备注
	CreateBy      any         // 创建人
	CreateTime    *gtime.Time // 创建时间
	UpdateBy      any         // 修改人
	UpdateTime    *gtime.Time // 更新时间
}
