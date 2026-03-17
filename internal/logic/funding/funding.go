package funding

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"

	adminV1 "GoCEX/api/admin/v1"
	v1 "GoCEX/app/api"
	"GoCEX/internal/codes"
	"GoCEX/internal/consts"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"
	"GoCEX/internal/service/websocket"
)

type sFunding struct{}

func New() *sFunding {
	return &sFunding{}
}

// formatTime 强制格式化时间，避免返回 Go 的 2006 模板字符串
func formatTime(t *gtime.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format("Y-m-d H:i:s")
}

// RechargeSubmit 申请充提
func (s *sFunding) RechargeSubmit(ctx context.Context, userId uint64, in *v1.RechargeSubmitReq) error {
	// 获取用户信息填充冗余项
	var u entity.AppUser
	if err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).Scan(&u); err != nil {
		return gerror.NewCode(codes.UserNotFound, "找不到此用户")
	}

	// （扩展预留）根据后台系统配置 ASSET_COIN (rechargeMin/rechargeMax) 校验是否达标

	// 插入充值记录 (待定状态 0)
	record := g.Map{
		dao.AppRecharge.Columns().UserId:     userId,
		dao.AppRecharge.Columns().Username:   u.LoginName,
		dao.AppRecharge.Columns().Amount:     in.Amount,
		dao.AppRecharge.Columns().Type:       in.Type,
		dao.AppRecharge.Columns().Coin:       in.Type,
		dao.AppRecharge.Columns().Address:    in.Address,
		dao.AppRecharge.Columns().TxId:       in.TxId,
		dao.AppRecharge.Columns().FileName:   in.Pic,
		dao.AppRecharge.Columns().Status:     0,                                                          // 0 = 待审核
		dao.AppRecharge.Columns().SerialId:   "RC" + gtime.Now().Format("YmdHis") + gconv.String(userId), // 充值流水订单号
		dao.AppRecharge.Columns().CreateBy:   u.LoginName,
		dao.AppRecharge.Columns().CreateTime: gtime.Now(),
		dao.AppRecharge.Columns().UpdateTime: gtime.Now(),
	}

	result, err := dao.AppRecharge.Ctx(ctx).Data(record).Insert()
	if err != nil {
		return err
	}

	recordId, _ := result.LastInsertId()

	// 充值信号打入 Redis 阻塞队列，交由后台其他监听器异步分配地址或等客服审批
	_, err = g.Redis().Do(ctx, "XADD", "CEX:STREAM:RECHARGE", "*", "id", recordId)
	if err != nil {
		g.Log().Error(ctx, "推送充值消息至Redis失败", err)
	}

	// 立刻通过 WebSocket 推送至后台 Admin 前端页面，eventType=1 (举例:1充值)
	websocket.PublishToAdmin(ctx, 1)

	return nil
}

// WithdrawChannelSetting 提现通道配置结构
type WithdrawChannelSetting struct {
	RechargeName  string          `json:"rechargeName"`
	RechargeType  string          `json:"rechargeType"`
	Ratio         decimal.Decimal `json:"ratio"`         // 百分比手续费率
	Fee           decimal.Decimal `json:"fee"`           // 固定单笔手续费
	WithdrawalMix decimal.Decimal `json:"withdrawalMix"` // 最低提现金额
	WithdrawalMax decimal.Decimal `json:"withdrawalMax"` // 最高提现金额
	FreeNum       int             `json:"freeNum"`       // 每日免手续费次数
}

func (s *sFunding) getWithdrawConfig(ctx context.Context, coinType string) (*WithdrawChannelSetting, error) {
	record, err := dao.Setting.Ctx(ctx).Where("id", "WITHDRAWAL_CHANNEL_SETTING").One()
	if err != nil || record.IsEmpty() {
		return nil, gerror.New("未找到提现通道配置 (WITHDRAWAL_CHANNEL_SETTING)")
	}

	var settings []WithdrawChannelSetting
	if err := gjson.DecodeTo(record["setting_value"].String(), &settings); err != nil {
		return nil, gerror.New("解析提现配置失败: " + err.Error())
	}

	for _, item := range settings {
		// 优先匹配具体的通道名 (如 USDT-TRC)，忽略大小写
		if strings.EqualFold(item.RechargeName, coinType) {
			return &item, nil
		}
	}
	return nil, gerror.New("未找到币种 [" + coinType + "] 对应的提现费率配置")
}

func (s *sFunding) getReceiptPrice(ctx context.Context, coin string) decimal.Decimal {
	if coin == "" {
		return decimal.NewFromInt(1)
	}
	// 尝试从 Redis 获取汇率价格 (针对法定货币或银行币)
	// 逻辑参考 Java: CURRENCY_PRICE:符号，统一大写
	priceV, _ := g.Redis().Get(ctx, "CURRENCY_PRICE:"+strings.ToUpper(coin))
	if !priceV.IsEmpty() {
		p, err := decimal.NewFromString(priceV.String())
		if err == nil && !p.IsZero() {
			return p
		}
	}
	return decimal.NewFromInt(1)
}

// WithdrawSubmit 提现动作
func (s *sFunding) WithdrawSubmit(ctx context.Context, userId uint64, in *v1.WithdrawSubmitReq) (err error) {
	// 1. 校验密码
	var detail entity.AppUserDetail
	if err := dao.AppUserDetail.Ctx(ctx).Where(dao.AppUserDetail.Columns().UserId, userId).Scan(&detail); err != nil || detail.UserTardPwd == "" {
		return gerror.NewCode(codes.Failed, "请先设置资金交易密码")
	}
	salt := "cex_salt"
	hashedPassword, _ := gmd5.EncryptString(in.Pwd + salt)
	if hashedPassword != detail.UserTardPwd {
		return gerror.NewCode(codes.Failed, "资金交易密码错误")
	}

	// 2. 强校验用户状态及打码量流水 (这里我们加入预留判断)
	var u entity.AppUser
	err = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).Scan(&u)
	if err != nil || u.UserId == 0 {
		return gerror.NewCode(codes.UserNotFound, "找不到此用户")
	}
	if u.Status == 1 || u.IsFreeze == "2" {
		return gerror.NewCode(codes.UserDisabled, "您的账户状态异常，禁止提现")
	}
	// TODO: 判断 u.TotleAmont 是否 >= 打码量阈值

	// 3. 获取费率配置并执行校验
	config, err := s.getWithdrawConfig(ctx, in.CoinType)
	if err != nil {
		return err
	}

	amountDec := decimal.NewFromFloat(in.Amount)

	if amountDec.LessThan(config.WithdrawalMix) {
		return gerror.NewCodef(codes.Failed, "提现金额低于最低限制: %s", config.WithdrawalMix.String())
	}
	if !config.WithdrawalMax.IsZero() && amountDec.GreaterThan(config.WithdrawalMax) {
		return gerror.NewCodef(codes.Failed, "提现金额超过最高限制: %s", config.WithdrawalMax.String())
	}

	// 4. 计算手续费
	// 判断手续费减免：查询用户今日已发起的提现笔数（包含审核中和已成功的）
	today := gtime.Now().Format("Y-m-d")
	count, _ := dao.Withdraw.Ctx(ctx).
		Where(dao.Withdraw.Columns().UserId, userId).
		WhereGTE(dao.Withdraw.Columns().CreateTime, today+" 00:00:00").
		WhereLTE(dao.Withdraw.Columns().CreateTime, today+" 23:59:59").
		Count()
	// 注意：只要发起过申请（无论审核中、成功、锁定、甚至驳回失败），均视为消耗一次免费名额

	var feeDec decimal.Decimal
	var appliedRatio, appliedFixedFee float64

	if config.FreeNum > 0 && count < config.FreeNum {
		feeDec = decimal.Zero
		appliedRatio = 0
		appliedFixedFee = 0
	} else {
		// fee = amount * (ratio / 100) + fixedFee
		ratioDec := config.Ratio.Div(decimal.NewFromInt(100))
		feeDec = amountDec.Mul(ratioDec).Add(config.Fee)
		appliedRatio, _ = config.Ratio.Float64()
		appliedFixedFee, _ = config.Fee.Float64()
	}

	realAmountDec := decimal.NewFromFloat(in.Amount).Sub(feeDec)
	if realAmountDec.LessThanOrEqual(decimal.Zero) {
		return gerror.NewCode(codes.Failed, "提现金额不足以支付手续费")
	}

	fee, _ := feeDec.Float64()
	realAmount, _ := realAmountDec.Float64()

	// 5. 处理收款折算 (receiptAmount / receiptRealAmount / receiptCoin)
	// 逻辑：如果提现的是非主流币或银行币，可能存在汇率折算
	receiptCoin := in.Coin
	if config.RechargeType != "" {
		receiptCoin = config.RechargeType // 假设配置中的 RechargeType 是收款币种 (如 VND)
	}

	// 获取收款币种对原币种的汇率 (例如 1 USDT = 25000 VND)
	// 注意：这里的 getReceiptPrice 逻辑根据业务定义，如果是 VND 则取 CURRENCY_PRICE:VND
	receiptPrice := s.getReceiptPrice(ctx, receiptCoin)

	receiptAmountDec := amountDec.Mul(receiptPrice)
	receiptRealAmountDec := receiptAmountDec.Sub(feeDec.Mul(receiptPrice))

	receiptAmount, _ := receiptAmountDec.Float64()
	receiptRealAmount, _ := receiptRealAmountDec.Float64()
	exchangeRate, _ := receiptPrice.Float64()

	// 6. 构建提现记录
	record := g.Map{
		dao.Withdraw.Columns().UserId:            userId,
		dao.Withdraw.Columns().Username:          u.LoginName,
		dao.Withdraw.Columns().Amount:            in.Amount,
		dao.Withdraw.Columns().Fee:               fee,
		dao.Withdraw.Columns().RealAmount:        realAmount,
		dao.Withdraw.Columns().Type:              in.CoinType,
		dao.Withdraw.Columns().Coin:              in.Coin,
		dao.Withdraw.Columns().ToAdress:          in.Address,
		dao.Withdraw.Columns().Status:            0, // 发起为 0: 审核中
		dao.Withdraw.Columns().CreateBy:          u.LoginName,
		dao.Withdraw.Columns().CreateTime:        gtime.Now(),
		dao.Withdraw.Columns().UpdateTime:        gtime.Now(),
		dao.Withdraw.Columns().Ratio:             appliedRatio,
		dao.Withdraw.Columns().FixedFee:          appliedFixedFee,
		dao.Withdraw.Columns().ReceiptCoin:       receiptCoin,
		dao.Withdraw.Columns().ReceiptAmount:     receiptAmount,
		dao.Withdraw.Columns().ReceiptRealAmount: receiptRealAmount,
		dao.Withdraw.Columns().ExchangeRate:      exchangeRate,
	}

	// 在双锁防穿透的环境下执行资金冻结扣减
	rate := asset.New().GetExchangeRate(ctx, in.Coin)
	uAmount := decimal.NewFromFloat(in.Amount).Mul(rate).InexactFloat64()

	freezeReq := &adminV1.SubAmountReq{
		UserId:     int64(userId),
		Symbol:     in.Coin,
		Amount:     in.Amount,
		RecordType: consts.RecordTypeWithdraw,
		Remark:     "提现申请冻结",
		UAmount:    uAmount,
	}

	res, err := asset.New().FreezeAmount(ctx, freezeReq)
	if err != nil {
		// 因为并发资源不足导致的 Freeze 抛错会自动捕获，直接将失败扔回客户端
		return err
	}

	record[dao.Withdraw.Columns().SerialId] = res.RecordId

	_, err = dao.Withdraw.Ctx(ctx).Data(record).Insert()
	if err != nil {
		return err
	}

	// 提现信号打入 Redis 阻塞队列，通知人工或后台风控审核中心
	_, err = g.Redis().Do(ctx, "XADD", "CEX:STREAM:WITHDRAW", "*", "userId", userId, "amount", in.Amount)
	if err != nil {
		g.Log().Error(ctx, "推送提现消息至Redis失败", err)
	}

	// 立刻通过 WebSocket 推送至后台 Admin 前端页面，eventType=2 (举例:2提现)
	websocket.PublishToAdmin(ctx, 2)

	return nil
}

// GetRechargeList 获取充值记录列表
func (s *sFunding) GetRechargeList(ctx context.Context, userId uint64, req *v1.RechargeListReq) (*v1.RechargeListRes, error) {
	m := dao.AppRecharge.Ctx(ctx).Where(dao.AppRecharge.Columns().UserId, userId)

	if req.Coin != "" {
		m = m.Where(dao.AppRecharge.Columns().Coin, req.Coin)
	}

	total, err := m.Count()
	if err != nil || total == 0 {
		return &v1.RechargeListRes{Rows: []v1.RechargeItem{}, Total: 0}, nil
	}

	var list []entity.AppRecharge
	err = m.Page(req.PageNum, req.PageSize).OrderDesc(dao.AppRecharge.Columns().CreateTime).Scan(&list)
	if err != nil {
		return nil, err
	}

	rows := make([]v1.RechargeItem, 0, len(list))
	for _, item := range list {
		// 关联查询流水表捞取 realAmount (uamount)
		realAmount, _ := dao.AppWalletRecord.Ctx(ctx).
			Where(dao.AppWalletRecord.Columns().SerialId, item.SerialId).
			Where(dao.AppWalletRecord.Columns().Type, 1).
			Fields(dao.AppWalletRecord.Columns().UAmount).
			Value()

		rows = append(rows, v1.RechargeItem{
			Id:         item.Id,
			SerialId:   item.SerialId,
			Amount:     item.Amount,
			RealAmount: realAmount.Float64(),
			Coin:       item.Coin,
			Type:       item.OrderType,
			Status:     item.Status,
			Address:    item.Address,
			TxId:       item.TxId,
			Remark:     item.Remark,
			CreateTime: formatTime(item.CreateTime),
		})
	}

	return &v1.RechargeListRes{
		Total: total,
		Rows:  rows,
	}, nil
}

// GetWithdrawList 获取提现记录列表
func (s *sFunding) GetWithdrawList(ctx context.Context, userId uint64, req *v1.WithdrawListReq) (*v1.WithdrawListRes, error) {
	m := dao.Withdraw.Ctx(ctx).Where(dao.Withdraw.Columns().UserId, userId)

	if req.Coin != "" {
		m = m.Where(dao.Withdraw.Columns().Coin, req.Coin)
	}

	total, err := m.Count()
	if err != nil || total == 0 {
		return &v1.WithdrawListRes{Rows: []v1.WithdrawItem{}, Total: 0}, nil
	}

	var list []entity.Withdraw
	err = m.Page(req.PageNum, req.PageSize).OrderDesc(dao.Withdraw.Columns().CreateTime).Scan(&list)
	if err != nil {
		return nil, err
	}

	rows := make([]v1.WithdrawItem, 0, len(list))
	for _, item := range list {
		rows = append(rows, v1.WithdrawItem{
			Id:         item.Id,
			SerialId:   item.SerialId,
			Amount:     item.Amount,
			Type:       item.Type,
			Coin:       item.Coin,
			Fee:        item.Fee,
			RealAmount: item.RealAmount,
			Address:    item.ToAdress,
			Status:     item.Status,
			Remark:     item.WithDrawRemark,
			CreateTime: formatTime(item.CreateTime),
		})
	}

	return &v1.WithdrawListRes{
		Total: total,
		Rows:  rows,
	}, nil
}
