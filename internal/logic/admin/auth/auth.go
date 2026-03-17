package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type sAdminAuth struct{}

func New() *sAdminAuth {
	return &sAdminAuth{}
}

// checkPassword 支持 BCrypt（RuoYi 默认）和 MD5 双模式匹配
func checkPassword(plain, hashed string) bool {
	// BCrypt 格式以 $2a$ 或 $2b$ 开头，长度通常 60 位
	if len(hashed) > 4 && (hashed[:4] == "$2a$" || hashed[:4] == "$2b$") {
		err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
		return err == nil
	}
	// 兼容老库 MD5（32位 hex）
	h := md5.New()
	h.Write([]byte(plain))
	return hex.EncodeToString(h.Sum(nil)) == hashed
}

// Login 后台管理员登录
func (s *sAdminAuth) Login(ctx context.Context, req *v1.AdminLoginReq) (*v1.AdminLoginRes, error) {
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).Where("user_name", req.Username).Where("del_flag", "0").Scan(&user)
	if err != nil {
		// 直接暴露真实 DB 错误信息，方便调试
		return nil, gerror.Newf("数据库查询失败: %v", err)
	}

	if user.UserId == 0 {
		return nil, gerror.New("用户不存在或密码错误")
	}
	// 兼容 RuoYi 两种激活状态：有些库用 "0" 正常，有些是空字符串
	// 只要非 "1"(停用) 就放行
	if user.Status == "1" {
		return nil, gerror.New("账户已停用，请联系管理员")
	}
	if !checkPassword(req.Password, user.Password) {
		// 临时调试：输出数据库里存的密码前20位字符（调试完山删除）
		hashPreview := user.Password
		if len(hashPreview) > 20 {
			hashPreview = hashPreview[:20]
		}
		return nil, gerror.Newf("密码错误, DB哈希字头: [%s...]长度:%d", hashPreview, len(user.Password))
	}

	// Sign JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  user.UserId,
		"account": user.UserName, // reuse existing gf-jwt logic if any, else define standalone
	})
	tokenString, err := token.SignedString([]byte("your-256-bit-secret")) // Update with real loaded config
	if err != nil {
		return nil, err
	}

	return &v1.AdminLoginRes{Token: tokenString}, nil
}

// GetInfo 获取管理员权限信息
func (s *sAdminAuth) GetInfo(ctx context.Context, req *v1.AdminGetInfoReq) (*v1.AdminGetInfoRes, error) {
	userId := gconv.Int64(ctx.Value("adminId"))
	if userId == 0 {
		userId = g.RequestFromCtx(ctx).GetCtxVar("adminId").Int64()
	}

	// 1. 获取用户信息
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).Where("user_id", userId).Scan(&user)
	if err != nil {
		return nil, err
	}

	// 2. 获取部门信息
	var dept entity.SysDept
	_ = dao.SysDept.Ctx(ctx).Where("dept_id", user.DeptId).Scan(&dept)

	// 3. 获取角色信息与权限标识 (roles 字符串数组)
	var roles []entity.SysRole
	err = dao.SysRole.Ctx(ctx).As("r").
		InnerJoin(dao.SysUserRole.Table(), "ur", "ur.role_id = r.role_id").
		Where("ur.user_id", userId).
		Where("r.status", "0").
		Where("r.del_flag", "0").
		Scan(&roles)

	// [Fix] 超级管理员兜底逻辑：如果数据库没关联角色，手动补全 admin 角色
	if userId == 1 && len(roles) == 0 {
		_ = dao.SysRole.Ctx(ctx).Where("role_key", "admin").Scan(&roles)
		// 如果数据库连 admin 角色记录都没有，硬编码补一个
		if len(roles) == 0 {
			roles = append(roles, entity.SysRole{
				RoleId:   1,
				RoleName: "超级管理员",
				RoleKey:  "admin",
				Status:   "0",
			})
		}
	}

	roleStrings := make([]string, 0, len(roles))
	apiRoles := make([]v1.AdminInfoRole, 0, len(roles))
	for _, r := range roles {
		roleStrings = append(roleStrings, r.RoleKey)
		apiRole := v1.AdminInfoRole{
			RoleId:            r.RoleId,
			RoleName:          r.RoleName,
			RoleKey:           r.RoleKey,
			RoleSort:          r.RoleSort,
			DataScope:         r.DataScope,
			MenuCheckStrictly: r.MenuCheckStrictly == 1,
			DeptCheckStrictly: r.DeptCheckStrictly == 1,
			Status:            r.Status,
			DelFlag:           nil, // 线上返回里大部分是 null 或不传，这里先对齐 Java 样例
			Admin:             r.RoleKey == "admin",
			Flag:              false,
			CreateBy:          r.CreateBy,
			Remark:            r.Remark,
			MenuIds:           nil,
			DeptIds:           nil,
			Permissions:       nil,
		}
		if r.CreateBy == "" {
			apiRole.CreateBy = nil
		}
		if r.Remark == "" {
			apiRole.Remark = nil
		}
		if r.CreateTime != nil {
			apiRole.CreateTime = r.CreateTime.String()
		} else {
			apiRole.CreateTime = nil
		}
		if r.UpdateTime != nil {
			apiRole.UpdateTime = r.UpdateTime.String()
		} else {
			apiRole.UpdateTime = nil
		}
		apiRoles = append(apiRoles, apiRole)
	}

	// 4. 获取权限标识 (permissions 字符串数组)
	permissions := make([]string, 0)
	if userId == 1 {
		permissions = append(permissions, "*:*:*")
	} else {
		// 查询菜单对应的 perms 字段
		err = dao.SysMenu.Ctx(ctx).As("m").
			InnerJoin(dao.SysRoleMenu.Table(), "rm", "rm.menu_id = m.menu_id").
			InnerJoin(dao.SysUserRole.Table(), "ur", "ur.role_id = rm.role_id").
			Where("ur.user_id", userId).
			Where("m.status", "0").
			Fields("m.perms").
			Group("m.perms").
			Scan(&permissions)
	}

	// 组装返回结果
	res := &v1.AdminGetInfoRes{
		Roles:       roleStrings,
		Permissions: permissions,
		User: v1.AdminInfoUser{
			UserId:      user.UserId,
			DeptId:      user.DeptId,
			UserName:    user.UserName,
			NickName:    user.NickName,
			UserType:    user.UserType,
			Email:       user.Email,
			Phonenumber: user.Phonenumber,
			Sex:         user.Sex,
			Avatar:      user.Avatar,
			Password:    user.Password,
			Status:      user.Status,
			DelFlag:     user.DelFlag,
			LoginIp:     user.LoginIp,
			Remark:      user.Remark,
			Admin:       userId == 1,
			GoogleKey:   user.GoogleKey,
			CreateBy:    user.CreateBy,
			Dept: v1.AdminInfoDept{
				DeptId:     dept.DeptId,
				ParentId:   nil,
				DeptName:   dept.DeptName,
				OrderNum:   dept.OrderNum,
				Leader:     dept.Leader,
				Status:     dept.Status,
				Ancestors:  dept.Ancestors,
				DelFlag:    nil,
				Children:   make([]v1.AdminInfoDept, 0),
				CreateBy:   nil,
				CreateTime: nil,
				UpdateBy:   nil,
				UpdateTime: nil,
				Remark:     nil,
				Phone:      nil,
				Email:      nil,
				ParentName: nil,
			},
			Roles:   apiRoles,
			RoleIds: nil,
			PostIds: nil,
			RoleId:  nil,
		},
	}
	if user.Remark == "" {
		res.User.Remark = nil
	}
	if user.ParentId == 0 {
		res.User.ParentId = nil
	} else {
		res.User.ParentId = user.ParentId
	}
	if dept.ParentId != 0 {
		res.User.Dept.ParentId = dept.ParentId
	}
	if user.LoginDate != nil {
		res.User.LoginDate = user.LoginDate.String()
	}
	if user.CreateTime != nil {
		res.User.CreateTime = user.CreateTime.String()
	}
	if user.UpdateTime != nil {
		res.User.UpdateTime = user.UpdateTime.String()
	} else {
		res.User.UpdateTime = nil
	}

	return res, nil
}

// GetRouters 获取动态菜单树
func (s *sAdminAuth) GetRouters(ctx context.Context, req *v1.AdminGetRoutersReq) (*v1.AdminGetRoutersRes, error) {
	rawId := ctx.Value("adminId")
	userId := gconv.Int64(rawId)

	g.Log().Debugf(ctx, "[GetRouters] Context 读取结果: rawId=%v(%T), userId=%d", rawId, rawId, userId)

	// 如果从 ctx 拿不到，尝试从 GF Request 拿 (兼容两种注入方式)
	if userId == 0 {
		userId = g.RequestFromCtx(ctx).GetCtxVar("adminId").Int64()
		g.Log().Debugf(ctx, "[GetRouters] 后备方案 GF Request 读取结果: userId=%d", userId)
	}

	var list []entity.SysMenu
	m := dao.SysMenu.Ctx(ctx).As("m").Fields("m.*").Where("m.status", "0").WhereIn("m.menu_type", []string{"M", "C"})

	if userId == 1 {
		// 超级管理员：获取全量菜单
		err := m.OrderAsc("m.parent_id").OrderAsc("m.order_num").Scan(&list)
		if err != nil {
			return nil, err
		}
	} else {
		// 普通用户/代理：按角色过滤
		err := m.InnerJoin(dao.SysRoleMenu.Table(), "rm", "rm.menu_id = m.menu_id").
			InnerJoin(dao.SysUserRole.Table(), "ur", "ur.role_id = rm.role_id").
			InnerJoin(dao.SysRole.Table(), "ro", "ro.role_id = ur.role_id").
			Where("ur.user_id", userId).
			Where("ro.status", "0").
			Where("ro.del_flag", "0").
			Group("m.menu_id").
			OrderAsc("m.parent_id").OrderAsc("m.order_num").
			Scan(&list)
		if err != nil {
			return nil, err
		}
	}

	g.Log().Debugf(ctx, "[GetRouters] 从数据库查询到菜单数量: %d", len(list))

	res := v1.AdminGetRoutersRes(s.buildMenuTree(0, list))
	return &res, nil
}

func (s *sAdminAuth) buildMenuTree(parentId int64, menus []entity.SysMenu) []v1.AdminRouterInfo {
	res := make([]v1.AdminRouterInfo, 0)
	for _, m := range menus {
		if m.ParentId == parentId {
			router := v1.AdminRouterInfo{
				Name:      s.getRouterName(m),
				Path:      s.getRouterPath(m),
				Hidden:    m.Visible == "1",
				Component: s.getComponent(m),
				Meta: v1.RouterMeta{
					Title:   m.MenuName,
					Icon:    m.Icon,
					NoCache: m.IsCache == 1,
				},
			}

			// Ruoyi 特色路由处理
			if m.MenuType == "M" {
				router.AlwaysShow = true
				router.Redirect = "noRedirect"
			}

			children := s.buildMenuTree(m.MenuId, menus)
			if len(children) > 0 {
				router.Children = children
			}
			res = append(res, router)
		}
	}
	return res
}

func (s *sAdminAuth) getRouterName(m entity.SysMenu) string {
	// 如果是目录，Ruoyi 通常首字母大写 Path
	if m.Path == "" || m.Path == "/" {
		return gstr.CaseCamel(m.MenuName)
	}
	name := gstr.CaseCamel(m.Path)

	// 如果 Path 是纯数字或其他奇怪字符导致 Camel 为空，使用 MenuName
	if name == "" {
		return gstr.CaseCamel(m.MenuName)
	}
	return name
}

func (s *sAdminAuth) getRouterPath(m entity.SysMenu) string {
	path := m.Path
	// 处理一级目录
	if m.ParentId == 0 {
		if !gstr.HasPrefix(path, "/") && !gstr.HasPrefix(path, "http") {
			return "/" + path
		}
	} else if gstr.HasPrefix(path, "/") {
		// 非一级菜单不应该以 / 开头，否则前端拼接会出问题 (除非是外链)
		if !gstr.HasPrefix(path, "http") {
			return gstr.TrimLeft(path, "/")
		}
	}

	// 防御：防止返回双斜杠 //
	if path == "/" {
		return "/"
	}

	return path
}

func (s *sAdminAuth) getComponent(m entity.SysMenu) string {
	if m.Component != "" {
		return m.Component
	}
	if m.MenuType == "M" {
		return "Layout"
	}
	// 默认 Layout 或 ParentView
	if m.ParentId == 0 {
		return "Layout"
	}
	return "ParentView"
}
