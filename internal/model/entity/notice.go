// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure for table notice.
type Notice struct {
	NoticeId      int         `json:"notice_id"      orm:"notice_id"      description:"公告ID"`
	NoticeTitle   string      `json:"notice_title"   orm:"notice_title"   description:"标题"`
	NoticeType    string      `json:"notice_type"    orm:"notice_type"    description:"公告类型 1=公告信息 2=活动公告 3=首页滚动公告"`
	ModelType     string      `json:"model_type"     orm:"model_type"     description:"模块类型 1=公告信息 2=活动公告 3=首页滚动公告   1={1=链接弹窗 2=图文弹窗},   2={1=首页轮播活动 2=Defi挖矿活动图},  3={1=首页滚动公告}  注:没有二级的默认给1  二级联动"`
	NoticeContent string      `json:"notice_content" orm:"notice_content" description:"内容"`
	CommentsNum   int         `json:"comments_num"   orm:"comments_num"   description:"评论数"`
	Cover         string      `json:"cover"          orm:"cover"          description:"图片"`
	ViewNum       int         `json:"view_num"       orm:"view_num"       description:"浏览数"`
	ExpireTime    *gtime.Time `json:"expire_time"    orm:"expire_time"    description:"公告截止时间"`
	ImgUrl        string      `json:"img_url"        orm:"img_url"        description:"图片地址"`
	ChainedUrl    string      `json:"chained_url"    orm:"chained_url"    description:"链接地址"`
	DetailUrl     string      `json:"detail_url"     orm:"detail_url"     description:"详情页"`
	LanguageId    string      `json:"language_id"    orm:"language_id"    description:"zh:1,cht:2,en:3,pt:4,sa:5,ko:6,ja:7,es:8,th:9,ms:10,id:11,fr:12,ru:13"`
	Status        string      `json:"status"         orm:"status"         description:"公告状态（0正常 1关闭）"`
	Source        string      `json:"source"         orm:"source"         description:"展示端1=pc 2=h5"`
	Sort          int         `json:"sort"           orm:"sort"           description:"排序"`
	Remark        string      `json:"remark"         orm:"remark"         description:"备注"`
	CreateBy      string      `json:"create_by"      orm:"create_by"      description:"创建人"`
	CreateTime    *gtime.Time `json:"create_time"    orm:"create_time"    description:"创建时间"`
	UpdateBy      string      `json:"update_by"      orm:"update_by"      description:"修改人"`
	UpdateTime    *gtime.Time `json:"update_time"    orm:"update_time"    description:"更新时间"`
}
