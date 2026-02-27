package asset

import (
	"context"
	"strings"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
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
