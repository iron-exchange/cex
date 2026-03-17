package system

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/logic/admin/system"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// --------- 定时任务 (Sys Job) ---------

func (c *Controller) GetJobList(ctx context.Context, req *v1.GetAdminSysJobListReq) (res *v1.GetAdminSysJobListRes, err error) {
	data, err := system.New().GetJobList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "操作成功",
		"total": data.Total,
		"data":  data.List,
	})
	return nil, nil
}

func (c *Controller) UpdateJobStatus(ctx context.Context, req *v1.UpdateAdminSysJobStatusReq) (res *v1.UpdateAdminSysJobStatusRes, err error) {
	_, err = system.New().UpdateJobStatus(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

// --------- 菜单管理 (Sys Menu) ---------

func (c *Controller) GetMenuList(ctx context.Context, req *v1.GetAdminSysMenuListReq) (res *v1.GetAdminSysMenuListRes, err error) {
	data, err := system.New().GetMenuList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": data.List,
	})
	return nil, nil
}

// --------- 字典管理 (Sys Dict Type & Data) ---------

func (c *Controller) GetDictTypeList(ctx context.Context, req *v1.GetAdminSysDictTypeListReq) (res *v1.GetAdminSysDictTypeListRes, err error) {
	data, err := system.New().GetDictTypeList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "操作成功",
		"total": data.Total,
		"data":  data.List,
	})
	return nil, nil
}

func (c *Controller) GetDictDataList(ctx context.Context, req *v1.GetAdminSysDictDataListReq) (res *v1.GetAdminSysDictDataListRes, err error) {
	data, err := system.New().GetDictDataList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "操作成功",
		"total": data.Total,
		"data":  data.List,
	})
	return nil, nil
}

func (c *Controller) GetDictDataByType(ctx context.Context, req *v1.GetAdminDictDataByTypeReq) (res *v1.GetAdminDictDataByTypeRes, err error) {
	data, err := system.New().GetDictDataByType(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}

	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": data.List,
	})
	return nil, nil
}

// --------- 参数设置 (Sys Config) ---------

func (c *Controller) GetConfigList(ctx context.Context, req *v1.GetAdminSysConfigListReq) (res *v1.GetAdminSysConfigListRes, err error) {
	data, err := system.New().GetConfigList(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code":  200,
		"msg":   "操作成功",
		"total": data.Total,
		"data":  data.List,
	})
	return nil, nil
}

func (c *Controller) UpdateConfig(ctx context.Context, req *v1.UpdateAdminSysConfigReq) (res *v1.UpdateAdminSysConfigRes, err error) {
	_, err = system.New().UpdateConfig(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

func (c *Controller) UpdateSetting(ctx context.Context, req *v1.AdminUpdateSettingReq) (res *v1.AdminUpdateSettingRes, err error) {
	r := g.RequestFromCtx(ctx)
	// (Existing logic for req.Value...)
	if req.Value == "" {
		m := r.GetMap()
		delete(m, "key")
		if len(m) > 0 {
			if val, ok := m["value"]; ok {
				req.Value = gconv.String(val)
			} else {
				req.Value = gjson.New(m).String()
			}
		} else {
			req.Value = r.GetBodyString()
		}
	}

	_, err = system.New().UpdateSetting(ctx, req)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{"code": 200, "msg": "操作成功"})
	return nil, nil
}

func (c *Controller) GetSetting(ctx context.Context, req *v1.AdminGetSettingReq) (res v1.AdminGetSettingRes, err error) {
	data, err := system.New().GetSetting(ctx, req)
	r := g.RequestFromCtx(ctx)
	if err != nil {
		r.Response.WriteJson(g.Map{"code": 500, "msg": err.Error()})
		return nil, nil
	}
	r.Response.WriteJson(g.Map{
		"code": 200,
		"msg":  "操作成功",
		"data": data,
	})
	return nil, nil
}
