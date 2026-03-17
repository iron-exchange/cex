package wallet_record

import (
	"context"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/internal/consts"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminWalletRecord struct{}

func New() *sAdminWalletRecord {
	return &sAdminWalletRecord{}
}

// applyAgentFilter 统一处理代理商数据隔离逻辑
func (s *sAdminWalletRecord) applyAgentFilter(m *gdb.Model, userId int64, adminParentIdsField string) *gdb.Model {
	if userId == 1 {
		return m
	}
	// 代理商只能看自己的下级 (包含自己)
	// PostgreSQL 语法: admin_parent_ids ~ '(^|,)(userId)(,|$)'
	pattern := "(^|,)" + gconv.String(userId) + "(,|$)"
	if adminParentIdsField == "" {
		adminParentIdsField = "admin_parent_ids"
	}
	return m.Where(adminParentIdsField+" ~ ?", pattern)
}

func (s *sAdminWalletRecord) GetWalletRecordList(ctx context.Context, req *v1.GetWalletRecordListReq) (*v1.GetWalletRecordListRes, error) {
	userId := gconv.Int64(ctx.Value("adminId"))
	if userId == 0 {
		userId = g.RequestFromCtx(ctx).GetCtxVar("adminId").Int64()
	}

	m := s.applyAgentFilter(dao.AppWalletRecord.Ctx(ctx).As("awr"), userId, "awr.admin_parent_ids")
	if req.UserId > 0 {
		m = m.Where("awr.user_id", req.UserId)
	}
	if req.Symbol != "" {
		m = m.Where("awr.symbol", req.Symbol)
	}
	if req.Type != nil {
		m = m.Where("awr.type", *req.Type)
	}

	total, _ := m.Count()
	var list []struct {
		entity.AppWalletRecord
		IsTest int `json:"is_test"`
	}
	err := m.LeftJoin("t_app_user u", "u.user_id = awr.user_id").
		Fields("awr.*, u.is_test").
		Page(req.PageNum, req.PageSize).OrderDesc("awr.create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.WalletRecordInfo, 0, len(list))
	for _, r := range list {
		info := v1.WalletRecordInfo{
			SearchValue:    nil,
			Id:             r.Id,
			UserId:         r.UserId,
			Symbol:         r.Symbol,
			Type:           r.Type,
			Amount:         r.Amount,
			Uamount:        r.UAmount,
			BeforeAmount:   r.BeforeAmount,
			AfterAmount:    r.AfterAmount,
			SerialId:       r.SerialId,
			CreateBy:       r.CreateBy,
			IsTest:         gconv.String(r.IsTest),
			Remark:         r.Remark,
			AdminParentIds: r.AdminParentIds,
			UpdateBy:       nil,
			UpdateTime:     nil,
			StartTime:      nil,
			EndTime:        nil,
			MinAmount:      nil,
			MaxAmount:      nil,
			OperateTime:    nil,
		}
		if r.OperateTime != nil {
			info.OperateTime = r.OperateTime.Format("2006-01-02 15:04:05")
		}
		if r.CreateTime != nil {
			info.CreateTime = r.CreateTime.Format("2006-01-02 15:04:05")
		}
		if r.UpdateTime != nil {
			info.UpdateTime = r.UpdateTime.Format("2006-01-02 15:04:05")
		}
		if r.UpdateBy != "" {
			info.UpdateBy = r.UpdateBy
		}
		resList = append(resList, info)
	}

	return &v1.GetWalletRecordListRes{
		Rows:  resList,
		Total: int(total),
	}, nil
}

func (s *sAdminWalletRecord) GetWalletRecordTypes(ctx context.Context, req *v1.GetWalletRecordTypesReq) (v1.GetWalletRecordTypesRes, error) {
	// 复用全局定义的账变类型字典
	res := make(v1.GetWalletRecordTypesRes)
	for k, v := range consts.RecordTypeMap {
		res[k] = v
	}
	return res, nil
}
func (s *sAdminWalletRecord) GetWalletStatistics(ctx context.Context, req *v1.GetWalletStatisticsReq) (*v1.GetWalletStatisticsRes, error) {
	// 根据用户要求：统计所有非测试用户的累计账变总金额
	// SELECT sum( r.amount ) AS statisticsAmount FROM t_app_user a LEFT JOIN t_app_wallet_record r ON r.user_id = a.user_id WHERE a.is_test = 0;
	val, err := dao.AppUser.Ctx(ctx).As("a").
		LeftJoin(dao.AppWalletRecord.Table(), "r", "r.user_id = a.user_id").
		Where("a.is_test", 0).
		Value("SUM(r.amount)")
	if err != nil {
		return nil, err
	}

	return &v1.GetWalletStatisticsRes{
		StatisticsAmount: val.Float64(),
	}, nil
}
