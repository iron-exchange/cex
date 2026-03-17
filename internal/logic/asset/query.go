package asset

import (
	"context"
	"strings"

	v1admin "GoCEX/api/admin/v1"
	v1 "GoCEX/app/api"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// GetAssetList 获取用户资产大盘 (带 USDT 折合计算)
func (s *sAsset) GetAssetList(ctx context.Context, userId uint64) (*v1.AssetListRes, error) {
	var assets []entity.AppAsset
	err := dao.AppAsset.Ctx(ctx).Where(dao.AppAsset.Columns().UserId, userId).Scan(&assets)
	if err != nil {
		return nil, err
	}

	res := &v1.AssetListRes{
		List:               make([]v1.AssetInfo, 0, len(assets)),
		TotalUsdtValuation: "0",
	}

	totalUsdt := decimal.NewFromInt(0)

	for _, a := range assets {
		// AppAsset 中的 Amout 代表总数量，AvailableAmount 为可用，差值为冻结
		total := decimal.NewFromFloat(a.Amout)
		available := decimal.NewFromFloat(a.AvailableAmount)
		frozen := total.Sub(available)

		price := decimal.NewFromInt(1)
		symbol := strings.ToUpper(a.Symbol)
		if symbol != "USDT" {
			// 从 Redis 取出实时行情。暂定存放在 CEX:PRICE:符号 中
			priceStr, _ := g.Redis().Do(ctx, "GET", "CEX:PRICE:"+symbol)
			if !priceStr.IsEmpty() {
				p, errParse := decimal.NewFromString(priceStr.String())
				if errParse == nil && !p.IsZero() {
					price = p
				} else {
					price = decimal.NewFromInt(0)
				}
			} else {
				price = decimal.NewFromInt(0)
			}
		}

		// 使用 Decimal 防御乘法溢出与粉尘异常
		valuation := total.Mul(price)
		totalUsdt = totalUsdt.Add(valuation)

		res.List = append(res.List, v1.AssetInfo{
			Symbol:          a.Symbol,
			AvailableAmount: available.String(),
			FrozenAmount:    frozen.String(),
			TotalAmount:     total.String(),
			UsdtValuation:   valuation.StringFixed(4), // 保留4位常用估值小数
		})
	}

	res.TotalUsdtValuation = totalUsdt.StringFixed(4)
	return res, nil
}

// GetWalletRecords 获取个人的财务账变流水
func (s *sAsset) GetWalletRecords(ctx context.Context, in *v1.WalletRecordReq, userId uint64) (*v1.WalletRecordRes, error) {
	m := dao.AppWalletRecord.Ctx(ctx).Where(dao.AppWalletRecord.Columns().UserId, userId)

	// 枚举类型过滤，实现高性能下钻条件查询
	if in.Type != 0 {
		m = m.Where(dao.AppWalletRecord.Columns().Type, in.Type)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var records []entity.AppWalletRecord
	// 利用 id DESC 进行排序（配合将来建立复合索引加快深度翻页）
	err = m.Page(in.Page, in.Size).OrderDesc(dao.AppWalletRecord.Columns().Id).Scan(&records)
	if err != nil {
		return nil, err
	}

	res := &v1.WalletRecordRes{
		List:  make([]v1.WalletRecordInfo, 0, len(records)),
		Total: total,
	}

	for _, r := range records {
		res.List = append(res.List, v1.WalletRecordInfo{
			Id:           r.Id,
			Symbol:       r.Symbol,
			Amount:       decimal.NewFromFloat(r.Amount).String(),
			BeforeAmount: decimal.NewFromFloat(r.BeforeAmount).String(),
			AfterAmount:  decimal.NewFromFloat(r.AfterAmount).String(),
			Type:         r.Type,
			Remark:       r.Remark,
			CreateTime:   r.CreateTime.Format("Y-m-d H:i:s"),
		})
	}

	return res, nil
}
func (s *sAsset) GetAppAssetList(ctx context.Context, req *v1admin.GetAppAssetListReq) (*v1admin.GetAppAssetListRes, error) {
	m := dao.AppAsset.Ctx(ctx).As("asset").LeftJoin("t_app_user u", "asset.user_id = u.user_id")

	// 1. 权限过滤 (仅超级管理员可看全部)
	adminId := gconv.Int64(ctx.Value("adminId"))
	if adminId != 1 && adminId > 0 {
		m = m.WhereLike("u.admin_parent_ids", "%"+gconv.String(adminId)+"%")
	}

	// 2. 基础过滤
	if req.UserId > 0 {
		m = m.Where("asset.user_id", req.UserId)
	}
	if req.Adress != "" {
		m = m.WhereLike("asset.adress", "%"+req.Adress+"%")
	}
	if req.Symbol != "" {
		m = m.Where("asset.symbol", strings.ToLower(req.Symbol))
	}
	if req.Type > 0 {
		m = m.Where("asset.type", req.Type)
	}
	if req.SearchValue != "" {
		m = m.WhereLike("asset.adress", "%"+req.SearchValue+"%")
	}

	// 3. 金额区间过滤 (Params)
	if req.Params.AmountMin != "" {
		m = m.WhereGTE("asset.amout", req.Params.AmountMin)
	}
	if req.Params.AmountMax != "" {
		m = m.WhereLTE("asset.amout", req.Params.AmountMax)
	}
	if req.Params.AvailableAmountMin != "" {
		m = m.WhereGTE("asset.available_amount", req.Params.AvailableAmountMin)
	}
	if req.Params.AvailableAmountMax != "" {
		m = m.WhereLTE("asset.available_amount", req.Params.AvailableAmountMax)
	}
	if req.Params.OccupiedAmountMin != "" {
		m = m.WhereGTE("asset.occupied_amount", req.Params.OccupiedAmountMin)
	}
	if req.Params.OccupiedAmountMax != "" {
		m = m.WhereLTE("asset.occupied_amount", req.Params.OccupiedAmountMax)
	}

	// 4. 时间过滤
	if req.Params.BeginTime != "" {
		m = m.WhereGTE("asset.create_time", req.Params.BeginTime)
	}
	if req.Params.EndTime != "" {
		m = m.WhereLTE("asset.create_time", req.Params.EndTime)
	}

	total, err := m.Count()
	if err != nil || total == 0 {
		return &v1admin.GetAppAssetListRes{Total: 0, Rows: []v1admin.AppAssetInfo{}}, nil
	}

	// 5. 聚合查询
	var list []struct {
		entity.AppAsset
		AdminParentIds string `orm:"admin_parent_ids"`
	}

	err = m.Page(req.PageNum, req.PageSize).Fields("asset.*, u.admin_parent_ids").OrderDesc("asset.create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resRows := make([]v1admin.AppAssetInfo, 0, len(list))
	for _, a := range list {
		row := v1admin.AppAssetInfo{
			CreateBy:             a.CreateBy,
			CreateTime:           a.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:             a.UpdateBy,
			UpdateTime:           a.UpdateTime.Format("2006-01-02 15:04:05"),
			Remark:               a.Remark,
			UserId:               a.UserId,
			Adress:               &a.Adress,
			Symbol:               a.Symbol,
			Amout:                a.Amout,
			OccupiedAmount:       a.OccupiedAmount,
			AvailableAmount:      a.AvailableAmount,
			AvailableAmountDaily: a.AvailableAmountDaily,
			CodingVolumeDaily:    a.CodingVolumeDaily,
			Type:                 gconv.Int(a.Type),
			ExchageAmount:        nil,
			AdminParentIds:       a.AdminParentIds,
			Loge:                 nil,
		}
		if a.Adress == "" {
			row.Adress = nil
		}
		resRows = append(resRows, row)
	}

	return &v1admin.GetAppAssetListRes{
		Total: int(total),
		Rows:  resRows,
	}, nil
}
