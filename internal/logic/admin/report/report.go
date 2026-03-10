package report

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminReport struct{}

func New() *sAdminReport {
	return &sAdminReport{}
}

// GetDailyData 获取每日统计数据 (聚合查询示例)
func (s *sAdminReport) GetDailyData(ctx context.Context, req *v1.GetDailyDataReq) (*v1.GetDailyDataRes, error) {
	// TODO: 完整的每日数据由复杂的定时任务生成报表，此处仅展示 SQL 聚合统计框架
	// Select DATE(create_time) as date, count(1) as new_users ... Group by DATE(create_time)
	// 由于这属于高级功能，在初期没有实际报表库时，这里先构建占位返回结构。

	resList := []v1.DailyDataInfo{
		{
			Date:          "2026-03-01",
			NewUsers:      12,
			TotalRecharge: 50000.0,
			TotalWithdraw: 12000.0,
			CompanyProfit: 3800.0,
		},
	}

	return &v1.GetDailyDataRes{
		List:  resList,
		Total: 1,
	}, nil
}

// GetAgentData 获取代理数据统计
func (s *sAdminReport) GetAgentData(ctx context.Context, req *v1.GetAgentDataReq) (*v1.GetAgentDataRes, error) {
	// TODO: 基于 `app_user` 的 admin_parent_ids 往下推算汇总，或者读取建立的代理日结报表
	resList := []v1.AgentDataInfo{
		{
			AgentId:       1001,
			AgentName:     "sysadmin",
			SubUsersCount: 156,
			TotalRecharge: 231500.0,
			TotalWithdraw: 89000.0,
			AgentProfit:   45000.0,
		},
	}

	return &v1.GetAgentDataRes{
		List:  resList,
		Total: 1,
	}, nil
}

// GetPlayerData 单玩家历史充提统计
func (s *sAdminReport) GetPlayerData(ctx context.Context, req *v1.GetPlayerDataReq) (*v1.GetPlayerDataRes, error) {
	m := dao.AppUser.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}

	total, _ := m.Count()
	records, err := m.Page(req.Page, req.Size).All()
	if err != nil {
		return nil, err
	}

	resList := make([]v1.PlayerDataInfo, 0, len(records))
	for _, record := range records {
		resList = append(resList, v1.PlayerDataInfo{
			UserId:         gconv.Int64(record["user_id"]),
			LoginName:      gconv.String(record["login_name"]),
			TotalRecharge:  gconv.Float64(record["recharge_amont"]), // 使用冗余字段或自行 sum
			TotalWithdraw:  0.0,                                     // TODO: SELECT SUM(amount) FROM withdraw WHERE user_id = ?
			NetProfit:      gconv.Float64(record["recharge_amont"]), // TODO: recharge - withdraw - balance
			CurrentBalance: 0.0,                                     // TODO: 从 app_asset 表拉取折合 U
		})
	}

	return &v1.GetPlayerDataRes{
		List:  resList,
		Total: total,
	}, nil
}
