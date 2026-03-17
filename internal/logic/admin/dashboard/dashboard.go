package dashboard

import (
	"context"
	"fmt"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminDashboard struct{}

func New() *sAdminDashboard {
	return &sAdminDashboard{}
}

func (s *sAdminDashboard) GetDashboard(ctx context.Context, req *v1.GetDashboardReq) (*v1.GetDashboardRes, error) {
	userId := gconv.Int64(ctx.Value("adminId"))

	// 生成过去 7 天的日期序列
	dates := make([]string, 7)
	now := gtime.Now().StartOfDay()
	for i := 0; i < 7; i++ {
		dates[6-i] = now.AddDate(0, 0, -i).Format("Y-m-d")
	}

	res := make(v1.GetDashboardRes, 0)

	// 1. 平台收支概况
	profitSeries, err := s.getProfitStatistics(ctx, userId, dates)
	if err != nil {
		return nil, err
	}
	res = append(res, profitSeries)

	// 2. 用户增长统计
	userSeries, err := s.getUserStatistics(ctx, userId, dates)
	if err != nil {
		return nil, err
	}
	res = append(res, userSeries)

	// 3. 充值详情统计
	rechargeSeries, err := s.getRechargeStatistics(ctx, userId, dates)
	if err != nil {
		return nil, err
	}
	res = append(res, rechargeSeries)

	// 4. 提现详情统计
	withdrawSeries, err := s.getWithdrawStatistics(ctx, userId, dates)
	if err != nil {
		return nil, err
	}
	res = append(res, withdrawSeries)

	return &res, nil
}

// applyAgentFilter 统一处理代理商数据隔离逻辑
func (s *sAdminDashboard) applyAgentFilter(m *gdb.Model, userId int64, tableAlias string) *gdb.Model {
	if userId == 1 {
		return m
	}
	// PostgreSQL 兼容性查询：判断 admin_parent_ids 是否包含该 ID (逗号分隔字符串)
	column := "admin_parent_ids"
	if tableAlias != "" {
		column = fmt.Sprintf("%s.admin_parent_ids", tableAlias)
	}
	// 利用正则表达式或字符串包含逻辑过滤。此处假设采用 Ruoyi 默认的 admin_parent_ids 存储方式
	return m.Where(fmt.Sprintf("%s ~ ('(^|,)' || %d || '(,|$)')", column, userId))
}

func (s *sAdminDashboard) getProfitStatistics(ctx context.Context, userId int64, dates []string) (v1.DashboardSeries, error) {
	series := v1.DashboardSeries{
		Title:        1,
		RedLineName:  "充值",
		BlueLineName: "提现",
		RedLine:      make(map[string]float64),
		BlueLine:     make(map[string]float64),
	}

	// 累计总充值成功
	totalRecharge, _ := s.applyAgentFilter(dao.AppRecharge.Ctx(ctx), userId, "").Where("status", 2).Sum("amount")
	// 累计总提现成功
	totalWithdraw, _ := s.applyAgentFilter(dao.Withdraw.Ctx(ctx), userId, "").Where("status", 1).Sum("amount")
	series.TotalNum = totalRecharge - totalWithdraw

	// 每日充值 (红线)
	rRes, _ := s.applyAgentFilter(dao.AppRecharge.Ctx(ctx), userId, "").
		Where("status", 2).
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, SUM(amount) as val").
		Group("date").All()
	rMap := rRes.MapKeyStr("date")

	// 每日提现 (蓝线)
	wRes, _ := s.applyAgentFilter(dao.Withdraw.Ctx(ctx), userId, "").
		Where("status", 1).
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, SUM(amount) as val").
		Group("date").All()
	wMap := wRes.MapKeyStr("date")

	for _, d := range dates {
		series.RedLine[d] = gconv.Float64(rMap[d]["val"])
		series.BlueLine[d] = gconv.Float64(wMap[d]["val"])
	}

	return series, nil
}

func (s *sAdminDashboard) getUserStatistics(ctx context.Context, userId int64, dates []string) (v1.DashboardSeries, error) {
	series := v1.DashboardSeries{
		Title:        2,
		RedLineName:  "注册",
		BlueLineName: "冻结",
		RedLine:      make(map[string]float64),
		BlueLine:     make(map[string]float64),
	}

	// 总玩家数 (剔除测试)
	totalUser, _ := s.applyAgentFilter(dao.AppUser.Ctx(ctx), userId, "").Where("is_test", 0).Count()
	series.TotalNum = float64(totalUser)

	// 每日注册 (红线)
	rRes, _ := s.applyAgentFilter(dao.AppUser.Ctx(ctx), userId, "").
		Where("is_test", 0).
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, COUNT(1) as val").
		Group("date").All()
	rMap := rRes.MapKeyStr("date")

	// 每日冻结用户 (蓝线 - 状态为 2=冻结)
	fRes, _ := s.applyAgentFilter(dao.AppUser.Ctx(ctx), userId, "").
		Where("is_test", 0).Where("is_freeze", "2").
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, COUNT(1) as val").
		Group("date").All()
	fMap := fRes.MapKeyStr("date")

	for _, d := range dates {
		series.RedLine[d] = gconv.Float64(rMap[d]["val"])
		series.BlueLine[d] = gconv.Float64(fMap[d]["val"])
	}

	return series, nil
}

func (s *sAdminDashboard) getRechargeStatistics(ctx context.Context, userId int64, dates []string) (v1.DashboardSeries, error) {
	series := v1.DashboardSeries{
		Title:        3,
		RedLineName:  "充值成功",
		BlueLineName: "充值失败",
		RedLine:      make(map[string]float64),
		BlueLine:     make(map[string]float64),
	}

	total, _ := s.applyAgentFilter(dao.AppRecharge.Ctx(ctx), userId, "").Where("status", 2).Sum("amount")
	series.TotalNum = total

	// 充值成功 (红线)
	sRes, _ := s.applyAgentFilter(dao.AppRecharge.Ctx(ctx), userId, "").
		Where("status", 2).
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, SUM(amount) as val").
		Group("date").All()
	sMap := sRes.MapKeyStr("date")

	// 充值失败/驳回 (蓝线 - 假设 status=3)
	eRes, _ := s.applyAgentFilter(dao.AppRecharge.Ctx(ctx), userId, "").
		Where("status", 3).
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, SUM(amount) as val").
		Group("date").All()
	eMap := eRes.MapKeyStr("date")

	for _, d := range dates {
		series.RedLine[d] = gconv.Float64(sMap[d]["val"])
		series.BlueLine[d] = gconv.Float64(eMap[d]["val"])
	}

	return series, nil
}

func (s *sAdminDashboard) getWithdrawStatistics(ctx context.Context, userId int64, dates []string) (v1.DashboardSeries, error) {
	series := v1.DashboardSeries{
		Title:        4,
		RedLineName:  "提现成功",
		BlueLineName: "提现失败",
		RedLine:      make(map[string]float64),
		BlueLine:     make(map[string]float64),
	}

	total, _ := s.applyAgentFilter(dao.Withdraw.Ctx(ctx), userId, "").Where("status", 1).Sum("amount")
	series.TotalNum = total

	// 提现成功 (红线 - status=1)
	sRes, _ := s.applyAgentFilter(dao.Withdraw.Ctx(ctx), userId, "").
		Where("status", 1).
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, SUM(amount) as val").
		Group("date").All()
	sMap := sRes.MapKeyStr("date")

	// 提现失败 (蓝线 - status=2)
	eRes, _ := s.applyAgentFilter(dao.Withdraw.Ctx(ctx), userId, "").
		Where("status", 2).
		Where("create_time >= ?", dates[0]).
		Fields("DATE(create_time) as date, SUM(amount) as val").
		Group("date").All()
	eMap := eRes.MapKeyStr("date")

	for _, d := range dates {
		series.RedLine[d] = gconv.Float64(sMap[d]["val"])
		series.BlueLine[d] = gconv.Float64(eMap[d]["val"])
	}

	return series, nil
}
