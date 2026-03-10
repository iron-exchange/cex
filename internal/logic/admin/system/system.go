package system

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type sAdminSystem struct{}

func New() *sAdminSystem {
	return &sAdminSystem{}
}

// GetJobList 获取定时任务列表
func (s *sAdminSystem) GetJobList(ctx context.Context, req *v1.GetAdminSysJobListReq) (*v1.GetAdminSysJobListRes, error) {
	m := dao.SysJob.Ctx(ctx)
	if req.JobName != "" {
		m = m.WhereLike("job_name", "%"+req.JobName+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.SysJob
	err := m.Page(req.Page, req.Size).Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysJobInfo, 0, len(list))
	for _, j := range list {
		resList = append(resList, v1.AdminSysJobInfo{
			JobId:          j.JobId,
			JobName:        j.JobName,
			JobGroup:       j.JobGroup,
			InvokeTarget:   j.InvokeTarget,
			CronExpression: j.CronExpression,
			Status:         j.Status,
			CreateTime:     j.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysJobListRes{
		List:  resList,
		Total: total,
	}, nil
}

// UpdateJobStatus 修改定时任务状态
func (s *sAdminSystem) UpdateJobStatus(ctx context.Context, req *v1.UpdateAdminSysJobStatusReq) (*v1.UpdateAdminSysJobStatusRes, error) {
	_, err := dao.SysJob.Ctx(ctx).Data(g.Map{
		"status": req.Status,
	}).Where("job_id", req.JobId).Update()
	if err != nil {
		return nil, err
	}
	// TODO: Runtime Quartz integration to pause/resume actual jobs can be placed here
	return &v1.UpdateAdminSysJobStatusRes{}, nil
}

// GetMenuList 获取菜单列表 (不分页树形/列表)
func (s *sAdminSystem) GetMenuList(ctx context.Context, req *v1.GetAdminSysMenuListReq) (*v1.GetAdminSysMenuListRes, error) {
	m := dao.SysMenu.Ctx(ctx)
	if req.MenuName != "" {
		m = m.WhereLike("menu_name", "%"+req.MenuName+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	var list []entity.SysMenu
	err := m.OrderAsc("parent_id").OrderAsc("order_num").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysMenuInfo, 0, len(list))
	for _, m := range list {
		resList = append(resList, v1.AdminSysMenuInfo{
			MenuId:    m.MenuId,
			MenuName:  m.MenuName,
			ParentId:  m.ParentId,
			OrderNum:  m.OrderNum,
			Path:      m.Path,
			Component: m.Component,
			MenuType:  m.MenuType,
			Visible:   m.Visible,
			Status:    m.Status,
			Perms:     m.Perms,
			Icon:      m.Icon,
		})
	}
	return &v1.GetAdminSysMenuListRes{List: resList}, nil
}

// GetDictTypeList 获取字典类型
func (s *sAdminSystem) GetDictTypeList(ctx context.Context, req *v1.GetAdminSysDictTypeListReq) (*v1.GetAdminSysDictTypeListRes, error) {
	m := dao.SysDictType.Ctx(ctx)
	if req.DictName != "" {
		m = m.WhereLike("dict_name", "%"+req.DictName+"%")
	}
	if req.DictType != "" {
		m = m.WhereLike("dict_type", "%"+req.DictType+"%")
	}

	total, _ := m.Count()
	var list []entity.SysDictType
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysDictTypeInfo, 0, len(list))
	for _, d := range list {
		resList = append(resList, v1.AdminSysDictTypeInfo{
			DictId:     d.DictId,
			DictName:   d.DictName,
			DictType:   d.DictType,
			Status:     d.Status,
			Remark:     d.Remark,
			CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysDictTypeListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetDictDataList 获取字典数据
func (s *sAdminSystem) GetDictDataList(ctx context.Context, req *v1.GetAdminSysDictDataListReq) (*v1.GetAdminSysDictDataListRes, error) {
	m := dao.SysDictData.Ctx(ctx)
	if req.DictType != "" {
		m = m.Where("dict_type", req.DictType)
	}
	if req.DictLabel != "" {
		m = m.WhereLike("dict_label", "%"+req.DictLabel+"%")
	}

	total, _ := m.Count()
	var list []entity.SysDictData
	err := m.Page(req.Page, req.Size).OrderAsc("dict_sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysDictDataInfo, 0, len(list))
	for _, d := range list {
		resList = append(resList, v1.AdminSysDictDataInfo{
			DictCode:   d.DictCode,
			DictSort:   d.DictSort,
			DictLabel:  d.DictLabel,
			DictValue:  d.DictValue,
			DictType:   d.DictType,
			CssClass:   d.CssClass,
			ListClass:  d.ListClass,
			IsDefault:  d.IsDefault,
			Status:     d.Status,
			CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysDictDataListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetDictDataByType 根据 Type 获取字典选项数组
func (s *sAdminSystem) GetDictDataByType(ctx context.Context, req *v1.GetAdminDictDataByTypeReq) (*v1.GetAdminDictDataByTypeRes, error) {
	var list []entity.SysDictData
	err := dao.SysDictData.Ctx(ctx).
		Where("dict_type", req.DictType).
		Where("status", "0").
		OrderAsc("dict_sort").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysDictDataInfo, 0, len(list))
	for _, d := range list {
		resList = append(resList, v1.AdminSysDictDataInfo{
			DictCode:  d.DictCode,
			DictLabel: d.DictLabel,
			DictValue: d.DictValue,
			ListClass: d.ListClass,
			CssClass:  d.CssClass,
			IsDefault: d.IsDefault,
		})
	}
	return &v1.GetAdminDictDataByTypeRes{List: resList}, nil
}

// GetConfigList 获取系统参数列表
func (s *sAdminSystem) GetConfigList(ctx context.Context, req *v1.GetAdminSysConfigListReq) (*v1.GetAdminSysConfigListRes, error) {
	m := dao.SysConfig.Ctx(ctx)
	if req.ConfigName != "" {
		m = m.WhereLike("config_name", "%"+req.ConfigName+"%")
	}
	if req.ConfigKey != "" {
		m = m.WhereLike("config_key", "%"+req.ConfigKey+"%")
	}

	total, _ := m.Count()
	var list []entity.SysConfig
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysConfigInfo, 0, len(list))
	for _, c := range list {
		resList = append(resList, v1.AdminSysConfigInfo{
			ConfigId:    c.ConfigId,
			ConfigName:  c.ConfigName,
			ConfigKey:   c.ConfigKey,
			ConfigValue: c.ConfigValue,
			ConfigType:  c.ConfigType,
			Remark:      c.Remark,
			CreateTime:  c.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysConfigListRes{
		List:  resList,
		Total: total,
	}, nil
}

// UpdateConfig 修改参数配置
func (s *sAdminSystem) UpdateConfig(ctx context.Context, req *v1.UpdateAdminSysConfigReq) (*v1.UpdateAdminSysConfigRes, error) {
	_, err := dao.SysConfig.Ctx(ctx).Data(g.Map{
		"config_value": req.ConfigValue,
	}).Where("config_id", req.ConfigId).Update()
	if err != nil {
		return nil, err
	}
	return &v1.UpdateAdminSysConfigRes{}, nil
}
