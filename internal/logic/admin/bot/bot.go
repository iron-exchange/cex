package bot

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"
)

type sAdminBot struct{}

func New() *sAdminBot {
	return &sAdminBot{}
}

// GetBotModelList 获取K线控盘机器人配置列表
func (s *sAdminBot) GetBotModelList(ctx context.Context, req *v1.GetAdminBotKlineModelListReq) (*v1.GetAdminBotKlineModelListRes, error) {
	m := dao.BotKlineModel.Ctx(ctx)
	if req.Symbol != "" {
		m = m.WhereLike("symbol", "%"+req.Symbol+"%")
	}

	total, _ := m.Count()
	var list []entity.BotKlineModel
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminBotKlineModelInfo, 0, len(list))
	for _, b := range list {
		info := v1.AdminBotKlineModelInfo{
			Id:           b.Id,
			Symbol:       b.Symbol,
			Decline:      b.Decline,
			Increase:     b.Increase,
			Model:        b.Model,
			Granularity:  b.Granularity,
			PricePencent: b.PricePencent,
			ConPrice:     b.ConPrice,
			CreateTime:   b.CreateTime.Format("2006-01-02 15:04:05"),
		}
		if b.BeginTime != nil {
			info.BeginTime = b.BeginTime.Format("2006-01-02 15:04:05")
		}
		if b.EndTime != nil {
			info.EndTime = b.EndTime.Format("2006-01-02 15:04:05")
		}
		resList = append(resList, info)
	}

	return &v1.GetAdminBotKlineModelListRes{
		List:  resList,
		Total: total,
	}, nil
}

// GetBotModelDataList 获取机器人生成的K线打点记录
func (s *sAdminBot) GetBotModelDataList(ctx context.Context, req *v1.GetAdminBotKlineModelDataListReq) (*v1.GetAdminBotKlineModelDataListRes, error) {
	m := dao.BotKlineModelInfo.Ctx(ctx)
	if req.ModelId > 0 {
		m = m.Where("model_id", req.ModelId)
	}

	total, _ := m.Count()
	var list []entity.BotKlineModelInfo
	err := m.Page(req.Page, req.Size).OrderDesc("date_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AdminBotKlineModelDataInfo, 0, len(list))
	for _, d := range list {
		resList = append(resList, v1.AdminBotKlineModelDataInfo{
			Id:       d.Id,
			ModelId:  d.ModelId,
			DateTime: d.DateTime,
			Open:     d.Open,
			Close:    d.Close,
			High:     d.High,
			Low:      d.Low,
		})
	}

	return &v1.GetAdminBotKlineModelDataListRes{
		List:  resList,
		Total: total,
	}, nil
}
