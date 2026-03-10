package v1

import "github.com/gogf/gf/v2/frame/g"

// --------- 定时任务 (Sys Job) ---------
type AdminSysJobInfo struct {
	JobId          int64  `json:"jobId"`
	JobName        string `json:"jobName"`
	JobGroup       string `json:"jobGroup"`
	InvokeTarget   string `json:"invokeTarget"`
	CronExpression string `json:"cronExpression"`
	Status         string `json:"status"` // 0正常 1暂停
	CreateTime     string `json:"createTime"`
}

type GetAdminSysJobListReq struct {
	g.Meta  `path:"/system/job/list" tags:"AdminSystem" method:"get" summary:"获取定时任务列表"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	JobName string `json:"jobName" dc:"任务名称"`
	Status  string `json:"status" dc:"状态"`
}

type GetAdminSysJobListRes struct {
	List  []AdminSysJobInfo `json:"list"`
	Total int               `json:"total"`
}

type UpdateAdminSysJobStatusReq struct {
	g.Meta `path:"/system/job/changeStatus" tags:"AdminSystem" method:"post" summary:"修改定时任务状态"`
	JobId  int64  `json:"jobId" v:"required#任务ID不能为空"`
	Status string `json:"status" v:"required#状态不能为空"`
}
type UpdateAdminSysJobStatusRes struct{}

// --------- 菜单管理 (Sys Menu) ---------
type AdminSysMenuInfo struct {
	MenuId    int64  `json:"menuId"`
	MenuName  string `json:"menuName"`
	ParentId  int64  `json:"parentId"`
	OrderNum  int    `json:"orderNum"`
	Path      string `json:"path"`
	Component string `json:"component"`
	MenuType  string `json:"menuType"` // M目录 C菜单 F按钮
	Visible   string `json:"visible"`  // 0显示 1隐藏
	Status    string `json:"status"`   // 0正常 1停用
	Perms     string `json:"perms"`
	Icon      string `json:"icon"`
}

type GetAdminSysMenuListReq struct {
	g.Meta   `path:"/system/menu/list" tags:"AdminSystem" method:"get" summary:"获取系统菜单列表"`
	MenuName string `json:"menuName" dc:"菜单名称"`
	Status   string `json:"status" dc:"状态"`
}

type GetAdminSysMenuListRes struct {
	List []AdminSysMenuInfo `json:"list"`
}

// --------- 字典类型 (Sys Dict Type) ---------
type AdminSysDictTypeInfo struct {
	DictId     int64  `json:"dictId"`
	DictName   string `json:"dictName"`
	DictType   string `json:"dictType"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
}

type GetAdminSysDictTypeListReq struct {
	g.Meta   `path:"/system/dict/type/list" tags:"AdminSystem" method:"get" summary:"获取字典类型列表"`
	Page     int    `json:"page" d:"1"`
	Size     int    `json:"size" d:"20"`
	DictName string `json:"dictName" dc:"字典名称"`
	DictType string `json:"dictType" dc:"字典类型"`
}

type GetAdminSysDictTypeListRes struct {
	List  []AdminSysDictTypeInfo `json:"list"`
	Total int                    `json:"total"`
}

// --------- 字典数据 (Sys Dict Data) ---------
type AdminSysDictDataInfo struct {
	DictCode   int64  `json:"dictCode"`
	DictSort   int    `json:"dictSort"`
	DictLabel  string `json:"dictLabel"`
	DictValue  string `json:"dictValue"`
	DictType   string `json:"dictType"`
	CssClass   string `json:"cssClass"`
	ListClass  string `json:"listClass"`
	IsDefault  string `json:"isDefault"` // Y/N
	Status     string `json:"status"`    // 0正常 1停用
	CreateTime string `json:"createTime"`
}

type GetAdminSysDictDataListReq struct {
	g.Meta    `path:"/system/dict/data/list" tags:"AdminSystem" method:"get" summary:"获取字典数据列表"`
	Page      int    `json:"page" d:"1"`
	Size      int    `json:"size" d:"20"`
	DictType  string `json:"dictType" dc:"字典类型"`
	DictLabel string `json:"dictLabel" dc:"字典标签"`
}

type GetAdminSysDictDataListRes struct {
	List  []AdminSysDictDataInfo `json:"list"`
	Total int                    `json:"total"`
}

// 供前端无感查询字典数据的接口 (例如 /dict/data/type/{dictType})
type GetAdminDictDataByTypeReq struct {
	g.Meta   `path:"/system/dict/data/type/{dictType}" tags:"AdminSystem" method:"get" summary:"根据字典类型查询字典数据"`
	DictType string `json:"dictType" in:"path" v:"required#字典类型不能为空"`
}

type GetAdminDictDataByTypeRes struct {
	List []AdminSysDictDataInfo `json:"data"` // 若配合框架，此字段常映射为 root 数组或 data
}

// --------- 参数设置 (Sys Config) ---------
type AdminSysConfigInfo struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
	ConfigType  string `json:"configType"` // Y是 N否
	Remark      string `json:"remark"`
	CreateTime  string `json:"createTime"`
}

type GetAdminSysConfigListReq struct {
	g.Meta     `path:"/system/config/list" tags:"AdminSystem" method:"get" summary:"获取参数配置列表"`
	Page       int    `json:"page" d:"1"`
	Size       int    `json:"size" d:"20"`
	ConfigName string `json:"configName" dc:"参数名称"`
	ConfigKey  string `json:"configKey" dc:"参数键名"`
}

type GetAdminSysConfigListRes struct {
	List  []AdminSysConfigInfo `json:"list"`
	Total int                  `json:"total"`
}

type UpdateAdminSysConfigReq struct {
	g.Meta      `path:"/system/config/update" tags:"AdminSystem" method:"post" summary:"修改参数配置"`
	ConfigId    int    `json:"configId" v:"required#配置ID不能为空"`
	ConfigValue string `json:"configValue" v:"required#配置值不能为空"`
}
type UpdateAdminSysConfigRes struct{}
