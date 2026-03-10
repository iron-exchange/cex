package rbac

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminRBAC struct{}

func New() *sAdminRBAC {
	return &sAdminRBAC{}
}

// GetDeptList 获取部门
func (s *sAdminRBAC) GetDeptList(ctx context.Context, req *v1.GetAdminSysDeptListReq) (*v1.GetAdminSysDeptListRes, error) {
	m := dao.SysDept.Ctx(ctx).Where("del_flag", "0")
	if req.DeptName != "" {
		m = m.WhereLike("dept_name", "%"+req.DeptName+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	var list []entity.SysDept
	err := m.OrderAsc("parent_id").OrderAsc("order_num").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysDeptInfo, 0, len(list))
	for _, d := range list {
		resList = append(resList, v1.AdminSysDeptInfo{
			DeptId:     d.DeptId,
			ParentId:   d.ParentId,
			DeptName:   d.DeptName,
			OrderNum:   d.OrderNum,
			Leader:     d.Leader,
			Phone:      d.Phone,
			Email:      d.Email,
			Status:     d.Status,
			CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysDeptListRes{List: resList}, nil
}

// GetRoleList 获取角色
func (s *sAdminRBAC) GetRoleList(ctx context.Context, req *v1.GetAdminSysRoleListReq) (*v1.GetAdminSysRoleListRes, error) {
	m := dao.SysRole.Ctx(ctx).Where("del_flag", "0")
	if req.RoleName != "" {
		m = m.WhereLike("role_name", "%"+req.RoleName+"%")
	}
	if req.RoleKey != "" {
		m = m.WhereLike("role_key", "%"+req.RoleKey+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.SysRole
	err := m.Page(req.Page, req.Size).OrderAsc("role_sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysRoleInfo, 0, len(list))
	for _, r := range list {
		resList = append(resList, v1.AdminSysRoleInfo{
			RoleId:     r.RoleId,
			RoleName:   r.RoleName,
			RoleKey:    r.RoleKey,
			RoleSort:   r.RoleSort,
			Status:     r.Status,
			Remark:     r.Remark,
			CreateTime: r.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysRoleListRes{List: resList, Total: total}, nil
}

// GetPostList 获取岗位
func (s *sAdminRBAC) GetPostList(ctx context.Context, req *v1.GetAdminSysPostListReq) (*v1.GetAdminSysPostListRes, error) {
	m := dao.SysPost.Ctx(ctx)
	if req.PostCode != "" {
		m = m.WhereLike("post_code", "%"+req.PostCode+"%")
	}
	if req.PostName != "" {
		m = m.WhereLike("post_name", "%"+req.PostName+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.SysPost
	err := m.Page(req.Page, req.Size).OrderAsc("post_sort").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysPostInfo, 0, len(list))
	for _, p := range list {
		resList = append(resList, v1.AdminSysPostInfo{
			PostId:     p.PostId,
			PostCode:   p.PostCode,
			PostName:   p.PostName,
			PostSort:   p.PostSort,
			Status:     p.Status,
			Remark:     p.Remark,
			CreateTime: p.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysPostListRes{List: resList, Total: total}, nil
}

// GetUserList 获取后台用户
func (s *sAdminRBAC) GetUserList(ctx context.Context, req *v1.GetAdminSysUserListReq) (*v1.GetAdminSysUserListRes, error) {
	m := dao.SysUser.Ctx(ctx).Where("del_flag", "0")
	if req.UserName != "" {
		m = m.WhereLike("user_name", "%"+req.UserName+"%")
	}
	if req.Phonenumber != "" {
		m = m.WhereLike("phonenumber", "%"+req.Phonenumber+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}
	if req.DeptId > 0 {
		m = m.Where("dept_id", req.DeptId)
	}

	total, _ := m.Count()
	var list []entity.SysUser
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysUserInfo, 0, len(list))
	for _, u := range list {
		info := v1.AdminSysUserInfo{
			UserId:      u.UserId,
			DeptId:      u.DeptId,
			UserName:    u.UserName,
			NickName:    u.NickName,
			UserType:    u.UserType,
			Email:       u.Email,
			Phonenumber: u.Phonenumber,
			Sex:         u.Sex,
			Status:      u.Status,
			LoginIp:     u.LoginIp,
			CreateTime:  u.CreateTime.Format("2006-01-02 15:04:05"),
		}
		if u.LoginDate != nil {
			info.LoginDate = u.LoginDate.Format("2006-01-02 15:04:05")
		}
		resList = append(resList, info)
	}
	return &v1.GetAdminSysUserListRes{List: resList, Total: total}, nil
}

// GetLogininforList 获取登录日志
func (s *sAdminRBAC) GetLogininforList(ctx context.Context, req *v1.GetAdminSysLogininforListReq) (*v1.GetAdminSysLogininforListRes, error) {
	m := dao.SysLogininfor.Ctx(ctx)
	if req.Ipaddr != "" {
		m = m.WhereLike("ipaddr", "%"+req.Ipaddr+"%")
	}
	if req.UserName != "" {
		m = m.WhereLike("user_name", "%"+req.UserName+"%")
	}
	if req.Status != "" {
		m = m.Where("status", req.Status)
	}

	total, _ := m.Count()
	var list []entity.SysLogininfor
	err := m.Page(req.Page, req.Size).OrderDesc("login_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysLogininforInfo, 0, len(list))
	for _, l := range list {
		resList = append(resList, v1.AdminSysLogininforInfo{
			InfoId:        l.InfoId,
			UserName:      l.UserName,
			Ipaddr:        l.Ipaddr,
			LoginLocation: l.LoginLocation,
			Browser:       l.Browser,
			Os:            l.Os,
			Status:        l.Status,
			Msg:           l.Msg,
			LoginTime:     l.LoginTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminSysLogininforListRes{List: resList, Total: total}, nil
}

// GetOperLogList 获取操作日志
func (s *sAdminRBAC) GetOperLogList(ctx context.Context, req *v1.GetAdminSysOperLogListReq) (*v1.GetAdminSysOperLogListRes, error) {
	m := dao.SysOperLog.Ctx(ctx)
	if req.Title != "" {
		m = m.WhereLike("title", "%"+req.Title+"%")
	}
	if req.OperName != "" {
		m = m.WhereLike("oper_name", "%"+req.OperName+"%")
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}
	if req.BusinessType != nil {
		m = m.Where("business_type", *req.BusinessType)
	}

	total, _ := m.Count()
	var list []entity.SysOperLog
	err := m.Page(req.Page, req.Size).OrderDesc("oper_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSysOperLogInfo, 0, len(list))
	for _, l := range list {
		info := v1.AdminSysOperLogInfo{
			OperId:        l.OperId,
			Title:         l.Title,
			BusinessType:  l.BusinessType,
			Method:        l.Method,
			RequestMethod: l.RequestMethod,
			OperName:      l.OperName,
			DeptName:      l.DeptName,
			OperUrl:       l.OperUrl,
			OperIp:        l.OperIp,
			OperLocation:  l.OperLocation,
			Status:        l.Status,
			CostTime:      l.CostTime,
		}
		if l.OperTime != nil {
			info.OperTime = l.OperTime.Format("2006-01-02 15:04:05")
		}
		resList = append(resList, info)
	}
	return &v1.GetAdminSysOperLogListRes{List: resList, Total: total}, nil
}
