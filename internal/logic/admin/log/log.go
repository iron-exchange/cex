package log

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminLog struct{}

func New() *sAdminLog {
	return &sAdminLog{}
}

func (s *sAdminLog) GetLoginLogList(ctx context.Context, req *v1.GetLoginLogListReq) (*v1.GetLoginLogListRes, error) {
	m := dao.AppuserLoginLog.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Ipaddr != "" {
		m = m.WhereLike("ipaddr", "%"+req.Ipaddr+"%")
	}

	total, _ := m.Count()
	var list []entity.AppuserLoginLog
	err := m.Page(req.Page, req.Size).OrderDesc("login_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.LoginLogInfo, 0, len(list))
	for _, l := range list {
		resList = append(resList, v1.LoginLogInfo{
			Id:            l.Id,
			UserId:        l.UserId,
			Ipaddr:        l.Ipaddr,
			LoginLocation: l.LoginLocation,
			Browser:       l.Browser,
			Os:            l.Os,
			Status:        l.Status,
			Msg:           l.Msg,
			LoginTime:     l.LoginTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetLoginLogListRes{
		List:  resList,
		Total: total,
	}, nil
}
