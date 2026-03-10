package system

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/system"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// --------- 定时任务 (Sys Job) ---------

func (c *Controller) GetJobList(ctx context.Context, req *v1.GetAdminSysJobListReq) (res *v1.GetAdminSysJobListRes, err error) {
	return system.New().GetJobList(ctx, req)
}

func (c *Controller) UpdateJobStatus(ctx context.Context, req *v1.UpdateAdminSysJobStatusReq) (res *v1.UpdateAdminSysJobStatusRes, err error) {
	return system.New().UpdateJobStatus(ctx, req)
}

// --------- 菜单管理 (Sys Menu) ---------

func (c *Controller) GetMenuList(ctx context.Context, req *v1.GetAdminSysMenuListReq) (res *v1.GetAdminSysMenuListRes, err error) {
	return system.New().GetMenuList(ctx, req)
}

// --------- 字典管理 (Sys Dict Type & Data) ---------

func (c *Controller) GetDictTypeList(ctx context.Context, req *v1.GetAdminSysDictTypeListReq) (res *v1.GetAdminSysDictTypeListRes, err error) {
	return system.New().GetDictTypeList(ctx, req)
}

func (c *Controller) GetDictDataList(ctx context.Context, req *v1.GetAdminSysDictDataListReq) (res *v1.GetAdminSysDictDataListRes, err error) {
	return system.New().GetDictDataList(ctx, req)
}

func (c *Controller) GetDictDataByType(ctx context.Context, req *v1.GetAdminDictDataByTypeReq) (res *v1.GetAdminDictDataByTypeRes, err error) {
	return system.New().GetDictDataByType(ctx, req)
}

// --------- 参数设置 (Sys Config) ---------

func (c *Controller) GetConfigList(ctx context.Context, req *v1.GetAdminSysConfigListReq) (res *v1.GetAdminSysConfigListRes, err error) {
	return system.New().GetConfigList(ctx, req)
}

func (c *Controller) UpdateConfig(ctx context.Context, req *v1.UpdateAdminSysConfigReq) (res *v1.UpdateAdminSysConfigRes, err error) {
	return system.New().UpdateConfig(ctx, req)
}
