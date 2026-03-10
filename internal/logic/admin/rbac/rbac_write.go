package rbac

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strings"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// md5Password is copied/shared conceptually from auth logic
func md5Password(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// --------- 菜单管理 (Sys Menu Write) ---------
func (s *sAdminRBAC) AddMenu(ctx context.Context, req *v1.AddAdminSysMenuReq) (*v1.AddAdminSysMenuRes, error) {
	_, err := dao.SysMenu.Ctx(ctx).Data(entity.SysMenu{
		ParentId:  req.ParentId,
		MenuName:  req.MenuName,
		OrderNum:  req.OrderNum,
		Path:      req.Path,
		Component: req.Component,
		IsFrame:   req.IsFrame,
		IsCache:   req.IsCache,
		MenuType:  req.MenuType,
		Visible:   req.Visible,
		Status:    req.Status,
		Perms:     req.Perms,
		Icon:      req.Icon,
	}).Insert()
	return &v1.AddAdminSysMenuRes{}, err
}

func (s *sAdminRBAC) EditMenu(ctx context.Context, req *v1.EditAdminSysMenuReq) (*v1.EditAdminSysMenuRes, error) {
	_, err := dao.SysMenu.Ctx(ctx).Where("menu_id", req.MenuId).Data(entity.SysMenu{
		ParentId:  req.ParentId,
		MenuName:  req.MenuName,
		OrderNum:  req.OrderNum,
		Path:      req.Path,
		Component: req.Component,
		IsFrame:   req.IsFrame,
		IsCache:   req.IsCache,
		MenuType:  req.MenuType,
		Visible:   req.Visible,
		Status:    req.Status,
		Perms:     req.Perms,
		Icon:      req.Icon,
	}).Update()
	return &v1.EditAdminSysMenuRes{}, err
}

func (s *sAdminRBAC) DeleteMenu(ctx context.Context, req *v1.DeleteAdminSysMenuReq) (*v1.DeleteAdminSysMenuRes, error) {
	// 简易判断：是否有子级
	count, _ := dao.SysMenu.Ctx(ctx).Where("parent_id", req.MenuId).Count()
	if count > 0 {
		return nil, gerror.New("存在子菜单,不允许删除")
	}
	_, err := dao.SysMenu.Ctx(ctx).Where("menu_id", req.MenuId).Delete()
	// 也应该一并清理 sys_role_menu，简单起见：
	dao.SysRoleMenu.Ctx(ctx).Where("menu_id", req.MenuId).Delete()
	return &v1.DeleteAdminSysMenuRes{}, err
}

// --------- 角色管理 (Sys Role Write) ---------
func (s *sAdminRBAC) AddRole(ctx context.Context, req *v1.AddAdminSysRoleReq) (*v1.AddAdminSysRoleRes, error) {
	return nil, dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := dao.SysRole.Ctx(ctx).Data(entity.SysRole{
			RoleName: req.RoleName,
			RoleKey:  req.RoleKey,
			RoleSort: req.RoleSort,
			Status:   req.Status,
			Remark:   req.Remark,
			DelFlag:  "0",
		}).InsertAndGetId()
		if err != nil {
			return err
		}
		// Insert Role-Menu mappings
		return s.insertRoleMenus(ctx, id, req.MenuIds)
	})
}

func (s *sAdminRBAC) EditRole(ctx context.Context, req *v1.EditAdminSysRoleReq) (*v1.EditAdminSysRoleRes, error) {
	return nil, dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.SysRole.Ctx(ctx).Where("role_id", req.RoleId).Data(entity.SysRole{
			RoleName: req.RoleName,
			RoleKey:  req.RoleKey,
			RoleSort: req.RoleSort,
			Status:   req.Status,
			Remark:   req.Remark,
		}).Update()
		if err != nil {
			return err
		}
		// Clear old mappings
		dao.SysRoleMenu.Ctx(ctx).Where("role_id", req.RoleId).Delete()
		// Insert new Role-Menu mappings
		return s.insertRoleMenus(ctx, req.RoleId, req.MenuIds)
	})
}

func (s *sAdminRBAC) DeleteRole(ctx context.Context, req *v1.DeleteAdminSysRoleReq) (*v1.DeleteAdminSysRoleRes, error) {
	roleIds := strings.Split(req.RoleIds, ",")
	return nil, dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.SysRole.Ctx(ctx).WhereIn("role_id", roleIds).Data(g.Map{"del_flag": "2"}).Update() // 2 means deleted in Ruoyi usually
		if err != nil {
			return err
		}
		// 清理权限映射 (可选物理删除)
		_, err = dao.SysRoleMenu.Ctx(ctx).WhereIn("role_id", roleIds).Delete()
		return err
	})
}

func (s *sAdminRBAC) insertRoleMenus(ctx context.Context, roleId int64, menuIds []int64) error {
	if len(menuIds) == 0 {
		return nil
	}
	var maps []entity.SysRoleMenu
	for _, menuId := range menuIds {
		maps = append(maps, entity.SysRoleMenu{
			RoleId: roleId,
			MenuId: menuId,
		})
	}
	_, err := dao.SysRoleMenu.Ctx(ctx).Data(maps).Insert()
	return err
}

// --------- 用户管理 (Sys User Write) ---------
func (s *sAdminRBAC) AddUser(ctx context.Context, req *v1.AddAdminSysUserReq) (*v1.AddAdminSysUserRes, error) {
	return nil, dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := dao.SysUser.Ctx(ctx).Data(entity.SysUser{
			DeptId:      req.DeptId,
			UserName:    req.UserName,
			NickName:    req.NickName,
			UserType:    "00",
			Email:       req.Email,
			Phonenumber: req.Phonenumber,
			Sex:         req.Sex,
			Password:    md5Password(req.Password), // Needs bcrypt or whatever mechanism RuoYi expects
			Status:      req.Status,
			DelFlag:     "0",
		}).InsertAndGetId()
		if err != nil {
			return err
		}
		return s.insertUserRoles(ctx, id, req.RoleIds)
	})
}

func (s *sAdminRBAC) EditUser(ctx context.Context, req *v1.EditAdminSysUserReq) (*v1.EditAdminSysUserRes, error) {
	if req.UserId == 1 { // 保护超级管理员
		return nil, gerror.New("不允许修改超级管理员用户")
	}
	return nil, dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.SysUser.Ctx(ctx).Where("user_id", req.UserId).Data(entity.SysUser{
			DeptId:      req.DeptId,
			UserName:    req.UserName,
			NickName:    req.NickName,
			Email:       req.Email,
			Phonenumber: req.Phonenumber,
			Sex:         req.Sex,
			Status:      req.Status,
		}).Update()
		if err != nil {
			return err
		}
		dao.SysUserRole.Ctx(ctx).Where("user_id", req.UserId).Delete()
		return s.insertUserRoles(ctx, req.UserId, req.RoleIds)
	})
}

func (s *sAdminRBAC) DeleteUser(ctx context.Context, req *v1.DeleteAdminSysUserReq) (*v1.DeleteAdminSysUserRes, error) {
	userIdsArr := strings.Split(req.UserIds, ",")
	for _, uid := range userIdsArr {
		if gconv.Int64(uid) == 1 {
			return nil, gerror.New("不允许删除超级管理员用户")
		}
	}
	return nil, dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.SysUser.Ctx(ctx).WhereIn("user_id", userIdsArr).Data(g.Map{"del_flag": "2"}).Update()
		if err != nil {
			return err
		}
		// 物理删除用户角色关连
		_, err = dao.SysUserRole.Ctx(ctx).WhereIn("user_id", userIdsArr).Delete()
		return err
	})
}

func (s *sAdminRBAC) ResetUserPwd(ctx context.Context, req *v1.ResetAdminSysUserPwdReq) (*v1.ResetAdminSysUserPwdRes, error) {
	if req.UserId == 1 {
		return nil, gerror.New("不允许重置超级管理员用户")
	}
	_, err := dao.SysUser.Ctx(ctx).Where("user_id", req.UserId).Data(g.Map{
		"password": md5Password(req.Password),
	}).Update()
	return &v1.ResetAdminSysUserPwdRes{}, err
}

func (s *sAdminRBAC) insertUserRoles(ctx context.Context, userId int64, roleIds []int64) error {
	if len(roleIds) == 0 {
		return nil
	}
	var maps []entity.SysUserRole
	for _, roleId := range roleIds {
		maps = append(maps, entity.SysUserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}
	_, err := dao.SysUserRole.Ctx(ctx).Data(maps).Insert()
	return err
}
