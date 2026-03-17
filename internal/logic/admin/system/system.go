package system

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
	userId := gconv.Int64(ctx.Value("adminId"))
	m := dao.SysDictData.Ctx(ctx).Where("dict_type", req.DictType).Where("status", "0")

	// 特殊逻辑：sys_user_type / user_type 权限过滤
	if req.DictType == "sys_user_type" || req.DictType == "user_type" {
		if userId != 1 {
			// 查询当前后台用户的类型
			var user entity.SysUser
			_ = dao.SysUser.Ctx(ctx).Where("user_id", userId).Scan(&user)

			if user.UserType == "1" {
				// 代理商级管理员被过滤，只能看到特定选项 (如：只能看到 dictValue='2' 的选项)
				m = m.Where("dict_value", "2")
			} else if userId != 1 {
				// 其他非超级管理员如果没有特殊定义，这里可能返回空，遵循 Ruoyi 逻辑
				return &v1.GetAdminDictDataByTypeRes{List: []v1.AdminSysDictDataInfo{}}, nil
			}
		}
	}

	var list []entity.SysDictData
	err := m.OrderAsc("dict_sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysDictDataInfo, 0, len(list))
	for _, d := range list {
		resList = append(resList, v1.AdminSysDictDataInfo{
			DictCode:  d.DictCode,
			DictSort:  d.DictSort,
			DictLabel: d.DictLabel,
			DictValue: d.DictValue,
			DictType:  d.DictType,
			ListClass: d.ListClass,
			CssClass:  d.CssClass,
			IsDefault: d.IsDefault,
			Status:    d.Status,
			Remark:    d.Remark,
			Default:   d.IsDefault == "Y",
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

// UpdateSetting 统一修改大盘单项配置 (类似原来 PUT /setting/put/{key})
func (s *sAdminSystem) UpdateSetting(ctx context.Context, req *v1.AdminUpdateSettingReq) (*v1.AdminUpdateSettingRes, error) {
	key := req.Key
	value := req.Value

	// 敏感配置鉴权拦截: 只有超级管理员能修改
	if key == "WITHDRAWAL_CHANNEL_SETTING" || key == "ASSET_COIN" {
		r := g.RequestFromCtx(ctx)
		adminId := r.GetCtxVar("adminId").Int()
		adminAccount := r.GetCtxVar("adminAccount").String()
		if adminId != 1 && adminAccount != "admin" {
			return nil, gerror.New("您没有操作权限，请联系管理员！")
		}
	}

	count, err := dao.Setting.Ctx(ctx).Where("id", key).Count()
	if err != nil {
		return nil, err
	}

	// APP_SIDEBAR_SETTING 特殊合并逻辑: 合并新增/覆盖，而不是直接全量替换
	if key == "APP_SIDEBAR_SETTING" && count > 0 {
		var config entity.Setting
		err = dao.Setting.Ctx(ctx).Where("id", key).Scan(&config)
		if err == nil && config.SettingValue != "" {
			// 解析旧的 JSON Array
			oldJson, err1 := gjson.DecodeToJson(config.SettingValue)
			newJson, err2 := gjson.DecodeToJson(value)

			if err1 == nil && err2 == nil {
				// 转换为 interface{} 切片
				oldSlice := oldJson.Interfaces()
				newSlice := newJson.Interfaces()

				mergedMap := make(map[string]interface{})

				for _, v := range oldSlice {
					vm, ok := v.(map[string]interface{})
					if !ok {
						continue
					}
					if name, hasName := vm["name"]; hasName {
						mergedMap[gconv.String(name)] = vm
					} else if path, hasPath := vm["path"]; hasPath {
						mergedMap[gconv.String(path)] = vm
					}
				}

				for _, v := range newSlice {
					vm, ok := v.(map[string]interface{})
					if !ok {
						continue
					}
					if name, hasName := vm["name"]; hasName {
						mergedMap[gconv.String(name)] = vm
					} else if path, hasPath := vm["path"]; hasPath {
						mergedMap[gconv.String(path)] = vm
					} else {
						// 找不到特征键兜底
						mergedMap[gconv.String(vm)] = vm
					}
				}

				mergedArray := make([]interface{}, 0, len(mergedMap))
				for _, v := range mergedMap {
					mergedArray = append(mergedArray, v)
				}

				finalJsonBytes, _ := gjson.Encode(mergedArray)
				value = string(finalJsonBytes)
			}
		}
	}

	if count > 0 {
		// Update
		_, err = dao.Setting.Ctx(ctx).Data(g.Map{"setting_value": value}).Where("id", key).Update()
	} else {
		// Insert
		_, err = dao.Setting.Ctx(ctx).Data(g.Map{
			"id":            key,
			"setting_value": value,
		}).Insert()
	}

	if err != nil {
		return nil, err
	}

	return &v1.AdminUpdateSettingRes{}, nil
}
func (s *sAdminSystem) GetSetting(ctx context.Context, req *v1.AdminGetSettingReq) (res v1.AdminGetSettingRes, err error) {
	var config entity.Setting
	err = dao.Setting.Ctx(ctx).Where("id", req.Key).Scan(&config)
	if err != nil {
		return nil, err
	}

	if config.SettingValue == "" {
		return nil, nil
	}

	// 尝试作为 JSON 解析
	if jsonWrap, err := gjson.Decode(config.SettingValue); err == nil {
		if m, ok := jsonWrap.(map[string]interface{}); ok {
			return m, nil
		}
		// 如果是数组或其他 JSON 类型，包装成 map 返回以适配 AdminGetSettingRes 类型并防止 Swagger 崩溃
		return v1.AdminGetSettingRes{"value": jsonWrap}, nil
	}

	// 普通字符串也包装成 map
	return v1.AdminGetSettingRes{"value": config.SettingValue}, nil
}
