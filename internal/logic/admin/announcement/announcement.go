package announcement

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type sAdminAnnouncement struct{}

func New() *sAdminAnnouncement {
	return &sAdminAnnouncement{}
}

// GetMailList 获取站内信
func (s *sAdminAnnouncement) GetMailList(ctx context.Context, req *v1.GetAdminAppMailListReq) (*v1.GetAdminAppMailListRes, error) {
	m := dao.AppMail.Ctx(ctx).Where("del_flag", "0")
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Title != "" {
		m = m.WhereLike("title", "%"+req.Title+"%")
	}

	total, _ := m.Count()
	var list []entity.AppMail
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminAppMailInfo, 0, len(list))
	for _, l := range list {
		resList = append(resList, v1.AdminAppMailInfo{
			Id:         l.Id,
			UserId:     l.UserId,
			Title:      l.Title,
			Content:    l.Content,
			Type:       l.Type,
			Status:     l.Status,
			CreateTime: l.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminAppMailListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetNoticeList 获取通知公告
func (s *sAdminAnnouncement) GetNoticeList(ctx context.Context, req *v1.GetAdminNoticeListReq) (*v1.GetAdminNoticeListRes, error) {
	m := dao.Notice.Ctx(ctx)
	if req.NoticeTitle != "" {
		m = m.WhereLike("notice_title", "%"+req.NoticeTitle+"%")
	}
	if req.NoticeType != "" {
		m = m.Where("notice_type", req.NoticeType)
	}

	total, _ := m.Count()
	var list []entity.Notice
	err := m.Page(req.Page, req.Size).OrderDesc("sort").OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminNoticeInfo, 0, len(list))
	for _, n := range list {
		resList = append(resList, v1.AdminNoticeInfo{
			NoticeId:      n.NoticeId,
			NoticeTitle:   n.NoticeTitle,
			NoticeType:    n.NoticeType,
			NoticeContent: n.NoticeContent,
			Cover:         n.Cover,
			ViewNum:       n.ViewNum,
			LanguageId:    n.LanguageId,
			Status:        n.Status,
			Sort:          n.Sort,
			CreateTime:    n.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminNoticeListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetSettingList 获取前台参数配置
func (s *sAdminAnnouncement) GetSettingList(ctx context.Context, req *v1.GetAdminSettingListReq) (*v1.GetAdminSettingListRes, error) {
	m := dao.Setting.Ctx(ctx).Where("delete_flag", "false")
	if req.Id != "" {
		m = m.WhereLike("id", "%"+req.Id+"%")
	}

	total, _ := m.Count()
	var list []entity.Setting
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminSettingInfo, 0, len(list))
	for _, st := range list {
		resList = append(resList, v1.AdminSettingInfo{
			Id:           st.Id,
			SettingValue: st.SettingValue,
			CreateTime:   st.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminSettingListRes{
		List:  resList,
		Total: total,
	}, nil
}

// UpdateSetting 修改前台参数配置
func (s *sAdminAnnouncement) UpdateSetting(ctx context.Context, req *v1.UpdateAdminSettingReq) (*v1.UpdateAdminSettingRes, error) {
	_, err := dao.Setting.Ctx(ctx).Data(g.Map{
		"setting_value": req.SettingValue,
	}).Where("id", req.Id).Update()
	if err != nil {
		return nil, err
	}
	return &v1.UpdateAdminSettingRes{}, nil
}

// GetHomeSetterList 获取规则说明列表
func (s *sAdminAnnouncement) GetHomeSetterList(ctx context.Context, req *v1.GetAdminHomeSetterListReq) (*v1.GetAdminHomeSetterListRes, error) {
	m := dao.HomeSetter.Ctx(ctx)
	if req.Title != "" {
		m = m.WhereLike("title", "%"+req.Title+"%")
	}
	if req.LanguageName != "" {
		m = m.Where("language_name", req.LanguageName)
	}
	if req.ModelType != nil {
		m = m.Where("model_type", *req.ModelType)
	}

	total, _ := m.Count()
	var list []entity.HomeSetter
	err := m.Page(req.Page, req.Size).OrderDesc("sort").OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminHomeSetterInfo, 0, len(list))
	for _, l := range list {
		resList = append(resList, v1.AdminHomeSetterInfo{
			Id:           l.Id,
			Title:        l.Title,
			Author:       l.Author,
			Content:      l.Content,
			ImgUrl:       l.ImgUrl,
			Sort:         l.Sort,
			IsShow:       l.IsShow,
			LanguageName: l.LanguageName,
			HomeType:     l.HomeType,
			ModelType:    l.ModelType,
			CreateTime:   l.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.GetAdminHomeSetterListRes{
		List:  resList,
		Total: total,
	}, nil
}

// AddHomeSetter 添加规则说明
func (s *sAdminAnnouncement) AddHomeSetter(ctx context.Context, req *v1.AddAdminHomeSetterReq) (*v1.AddAdminHomeSetterRes, error) {
	_, err := dao.HomeSetter.Ctx(ctx).Data(g.Map{
		"title":         req.Title,
		"content":       req.Content,
		"language_name": req.LanguageName,
		"home_type":     req.HomeType,
		"model_type":    req.ModelType,
		"sort":          req.Sort,
		"is_show":       req.IsShow,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return &v1.AddAdminHomeSetterRes{}, nil
}

// UpdateHomeSetter 修改规则说明
func (s *sAdminAnnouncement) UpdateHomeSetter(ctx context.Context, req *v1.UpdateAdminHomeSetterReq) (*v1.UpdateAdminHomeSetterRes, error) {
	data := g.Map{}
	if req.Title != "" {
		data["title"] = req.Title
	}
	if req.Content != "" {
		data["content"] = req.Content
	}
	if req.LanguageName != "" {
		data["language_name"] = req.LanguageName
	}
	data["sort"] = req.Sort
	data["is_show"] = req.IsShow

	_, err := dao.HomeSetter.Ctx(ctx).Data(data).Where("id", req.Id).Update()
	if err != nil {
		return nil, err
	}
	return &v1.UpdateAdminHomeSetterRes{}, nil
}

// GetHelpCenterList 获取帮助中心大类
func (s *sAdminAnnouncement) GetHelpCenterList(ctx context.Context, req *v1.GetAdminHelpCenterListReq) (*v1.GetAdminHelpCenterListRes, error) {
	m := dao.HelpCenter.Ctx(ctx).Where("del_flag", "0")
	if req.Title != "" {
		m = m.WhereLike("title", "%"+req.Title+"%")
	}

	total, _ := m.Count()
	var list []entity.HelpCenter
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminHelpCenterInfo, 0, len(list))
	for _, h := range list {
		resList = append(resList, v1.AdminHelpCenterInfo{
			Id:         h.Id,
			Title:      h.Title,
			Language:   h.Language,
			Enable:     h.Enable,
			ShowSymbol: h.ShowSymbol,
			CreateTime: h.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminHelpCenterListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetHelpCenterArticleList 获取帮助中心问题列表
func (s *sAdminAnnouncement) GetHelpCenterArticleList(ctx context.Context, req *v1.GetAdminHelpCenterArticleListReq) (*v1.GetAdminHelpCenterArticleListRes, error) {
	m := dao.HelpCenterInfo.Ctx(ctx).Where("del_flag", "0")
	if req.HelpCenterId > 0 {
		m = m.Where("help_center_id", req.HelpCenterId)
	}
	if req.Question != "" {
		m = m.WhereLike("question", "%"+req.Question+"%")
	}

	total, _ := m.Count()
	var list []entity.HelpCenterInfo
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminHelpCenterArticleInfo, 0, len(list))
	for _, h := range list {
		resList = append(resList, v1.AdminHelpCenterArticleInfo{
			Id:           h.Id,
			HelpCenterId: h.HelpCenterId,
			Question:     h.Question,
			Content:      h.Content,
			Language:     h.Language,
			Enable:       h.Enable,
			CreateTime:   h.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAdminHelpCenterArticleListRes{
		List:  resList,
		Total: total,
	}, nil
}
