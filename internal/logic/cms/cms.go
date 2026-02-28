package cms

import (
	"context"
	"time"

	v1 "GoCEX/app/api"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/os/gcache"
)

type sCms struct{}

func New() *sCms {
	return &sCms{}
}

// GetAllNoticeList 获取首页公告与轮播图 (带缓存)
func (s *sCms) GetAllNoticeList(ctx context.Context) (*v1.GetAllNoticeListRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:cms:notice_list", func(ctx context.Context) (interface{}, error) {
		var notices []entity.Notice
		// 只查询正常状态 (status=0) 的公告，并按 sort 排序
		err := dao.Notice.Ctx(ctx).Where(dao.Notice.Columns().Status, "0").OrderAsc(dao.Notice.Columns().Sort).Scan(&notices)
		if err != nil {
			return nil, err
		}

		res := &v1.GetAllNoticeListRes{
			List: make([]v1.NoticeInfo, 0, len(notices)),
		}

		for _, n := range notices {
			t := ""
			if n.CreateTime != nil {
				t = n.CreateTime.Format("Y-m-d H:i:s")
			}
			res.List = append(res.List, v1.NoticeInfo{
				NoticeId:      n.NoticeId,
				NoticeTitle:   n.NoticeTitle,
				NoticeType:    n.NoticeType,
				NoticeContent: n.NoticeContent,
				Cover:         n.Cover,
				CreateTime:    t,
			})
		}
		return res, nil
	}, 5*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(*v1.GetAllNoticeListRes), nil
}

// GetHelpCenterList 获取帮助中心与教学指南 (带缓存)
func (s *sCms) GetHelpCenterList(ctx context.Context) (*v1.GetHelpCenterListRes, error) {
	val, err := gcache.GetOrSetFunc(ctx, "app:cms:helpcenter_list", func(ctx context.Context) (interface{}, error) {
		var helps []entity.HelpCenter
		// 查询正常启用 (enable=1, del_flag=0)
		err := dao.HelpCenter.Ctx(ctx).Where(dao.HelpCenter.Columns().Enable, "1").
			Where(dao.HelpCenter.Columns().DelFlag, "0").Scan(&helps)
		if err != nil {
			return nil, err
		}

		res := &v1.GetHelpCenterListRes{
			List: make([]v1.HelpCenterInfo, 0, len(helps)),
		}

		for _, h := range helps {
			t := ""
			if h.CreateTime != nil {
				t = h.CreateTime.Format("Y-m-d H:i:s")
			}
			res.List = append(res.List, v1.HelpCenterInfo{
				Id:         h.Id,
				Title:      h.Title,
				Language:   h.Language,
				CreateTime: t,
			})
		}
		return res, nil
	}, 10*time.Minute)

	if err != nil {
		return nil, err
	}
	return val.Val().(*v1.GetHelpCenterListRes), nil
}

// GetUserMail 获取用户个人站内信列表 (不支持全局缓存，因人而异)
func (s *sCms) GetUserMail(ctx context.Context, userId uint64, in *v1.GetUserMailReq) (*v1.GetUserMailRes, error) {
	m := dao.AppMail.Ctx(ctx).Where(dao.AppMail.Columns().UserId, userId).
		Where(dao.AppMail.Columns().DelFlag, "0")

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var mails []entity.AppMail
	err = m.Page(in.Page, in.Size).OrderDesc(dao.AppMail.Columns().Id).Scan(&mails)
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserMailRes{
		List:  make([]v1.MailInfo, 0, len(mails)),
		Total: total,
	}

	for _, mail := range mails {
		t := ""
		if mail.CreateTime != nil {
			t = mail.CreateTime.Format("Y-m-d H:i:s")
		}
		res.List = append(res.List, v1.MailInfo{
			Id:         mail.Id,
			Title:      mail.Title,
			Content:    mail.Content,
			Type:       mail.Type,
			Status:     mail.Status,
			CreateTime: t,
		})
	}
	return res, nil
}
