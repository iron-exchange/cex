package user

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sAdminUser struct{}

func New() *sAdminUser {
	return &sAdminUser{}
}

func (s *sAdminUser) GetAppUserList(ctx context.Context, req *v1.GetAppUserListReq) (*v1.GetAppUserListRes, error) {
	m := dao.AppUser.Ctx(ctx)
	if req.LoginName != "" {
		m = m.WhereLike("login_name", "%"+req.LoginName+"%")
	}
	if req.Phone != "" {
		m = m.WhereLike("phone", "%"+req.Phone+"%")
	}
	if req.Address != "" {
		m = m.WhereLike("address", "%"+req.Address+"%")
	}
	if req.AdminParentIds != "" {
		m = m.WhereLike("admin_parent_ids", "%"+req.AdminParentIds+"%")
	}

	total, _ := m.Count()
	var list []entity.AppUser
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AppUserInfo, 0, len(list))
	for _, u := range list {
		resList = append(resList, v1.AppUserInfo{
			UserId:         u.UserId,
			LoginName:      u.LoginName,
			Phone:          u.Phone,
			Email:          u.Email,
			Address:        u.Address,
			Status:         u.Status,
			Buff:           u.Buff,
			TotalAmount:    u.TotleAmont,
			AdminParentIds: u.AdminParentIds,
			CreateTime:     u.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAppUserListRes{
		List:  resList,
		Total: total,
	}, nil
}

func (s *sAdminUser) FreezeUser(ctx context.Context, req *v1.FreezeUserReq) error {
	_, err := dao.AppUser.Ctx(ctx).Where("user_id", req.UserId).Update(map[string]interface{}{
		"status": req.Status,
	})
	if err != nil {
		return gerror.Wrap(err, "更新用户状态失败")
	}
	return nil
}

func (s *sAdminUser) UpdateParent(ctx context.Context, req *v1.UpdateUserParentReq) error {
	_, err := dao.AppUser.Ctx(ctx).Where("user_id", req.UserId).Update(map[string]interface{}{
		"admin_parent_ids": req.AdminParentIds,
	})
	if err != nil {
		return gerror.Wrap(err, "修改用户代理上级失败")
	}
	return nil
}
