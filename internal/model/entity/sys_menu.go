// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	MenuId     int64       `json:"menu_id"     orm:"menu_id"     description:"菜单ID"`
	MenuName   string      `json:"menu_name"   orm:"menu_name"   description:"菜单名称"`
	ParentId   int64       `json:"parent_id"   orm:"parent_id"   description:"父菜单ID"`
	OrderNum   int         `json:"order_num"   orm:"order_num"   description:"显示顺序"`
	Path       string      `json:"path"        orm:"path"        description:"路由地址"`
	Component  string      `json:"component"   orm:"component"   description:"组件路径"`
	Query      string      `json:"query"       orm:"query"       description:"路由参数"`
	IsFrame    int         `json:"is_frame"    orm:"is_frame"    description:"是否为外链（0是 1否）"`
	IsCache    int         `json:"is_cache"    orm:"is_cache"    description:"是否缓存（0缓存 1不缓存）"`
	MenuType   string      `json:"menu_type"   orm:"menu_type"   description:"菜单类型（M目录 C菜单 F按钮）"`
	Visible    string      `json:"visible"     orm:"visible"     description:"菜单状态（0显示 1隐藏）"`
	Status     string      `json:"status"      orm:"status"      description:"菜单状态（0正常 1停用）"`
	Perms      string      `json:"perms"       orm:"perms"       description:"权限标识"`
	Icon       string      `json:"icon"        orm:"icon"        description:"菜单图标"`
	CreateBy   string      `json:"create_by"   orm:"create_by"   description:"创建者"`
	CreateTime *gtime.Time `json:"create_time" orm:"create_time" description:"创建时间"`
	UpdateBy   string      `json:"update_by"   orm:"update_by"   description:"更新者"`
	UpdateTime *gtime.Time `json:"update_time" orm:"update_time" description:"更新时间"`
	Remark     string      `json:"remark"      orm:"remark"      description:"备注"`
}
