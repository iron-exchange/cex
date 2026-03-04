package defi

import (
	"context"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDefi struct{}

func New() *sDefi {
	return &sDefi{}
}

// GetDefiRate 获取分红收益列表
func (s *sDefi) GetDefiRate(ctx context.Context) (res *v1.GetDefiRateRes, err error) {
	var list []entity.DefiRate
	err = dao.DefiRate.Ctx(ctx).Order("min_amount asc").Scan(&list)
	if err != nil {
		return nil, err
	}

	res = &v1.GetDefiRateRes{
		List: make([]v1.DefiRateInfo, 0, len(list)),
	}
	for _, item := range list {
		res.List = append(res.List, v1.DefiRateInfo{
			Id:         item.Id,
			MinAmount:  item.MinAmount,
			MaxAmount:  item.MaxAmount,
			DailyRate:  item.Rate,
			Symbol:     "USDT", // 默认固定锁 USDT
			RewardCoin: "ETH",  // 默认固定投 ETH
		})
	}
	return
}

// SendApproveHash 处理用户提交的授权 Hash
func (s *sDefi) SendApproveHash(ctx context.Context, req *v1.SendApproveHashReq) (err error) {
	// 在原系统中，这里通常是把授权 Hash 更新到地址表或插入一个专门的记录表
	// 我们更新 t_app_address_info 的 remark 或其他字段标记已提交
	_, err = dao.AppAddressInfo.Ctx(ctx).
		Data(g.Map{"remark": "AuthorizedHash: " + req.Hash}).
		Where("address", req.Address).
		Update()
	return
}

// ShowIncome 展示每日收益明细
func (s *sDefi) ShowIncome(ctx context.Context, userId uint64) (res *v1.ShowIncomeRes, err error) {
	var list []entity.DefiOrder
	m := dao.DefiOrder.Ctx(ctx).Where("user_id", userId).Order("create_time desc")

	total, _ := m.Count()
	err = m.Page(1, 20).Scan(&list) // 暂不处理前端传来的 Page/Size，优先跑通逻辑
	if err != nil {
		return nil, err
	}

	res = &v1.ShowIncomeRes{
		Total: total,
		List:  make([]v1.DefiIncomeInfo, 0, len(list)),
	}
	for _, item := range list {
		res.List = append(res.List, v1.DefiIncomeInfo{
			Date:       item.CreateTime.Format("2006-01-02"),
			Amount:     item.TotleAmount,
			Reward:     item.Amount,
			RewardCoin: "ETH",
		})
	}
	return
}

// ShowOrder 获取用户正进行的 DeFi 授权订单
func (s *sDefi) ShowOrder(ctx context.Context, userId uint64) (res *v1.ShowOrderRes, err error) {
	// 按照 ruoyi-api, 这里的订单通常指授权过的钱包地址状态
	var list []entity.AppAddressInfo
	err = dao.AppAddressInfo.Ctx(ctx).
		Where("user_id", userId).
		Where("usdt_allowed > 0").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	res = &v1.ShowOrderRes{
		List: make([]v1.DefiOrderInfo, 0, len(list)),
	}
	for _, item := range list {
		res.List = append(res.List, v1.DefiOrderInfo{
			Id:         gconv.Int64(item.UserId),
			Address:    item.Address,
			UsdtAmount: item.Usdt,
			Status:     0,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return
}
