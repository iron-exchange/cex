package funding

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	adminV1 "GoCEX/api/admin/v1"
	v1 "GoCEX/app/api"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"
)

type sFunding struct{}

func New() *sFunding {
	return &sFunding{}
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
		dao.AppRecharge.Columns().UserId:   userId,
		dao.AppRecharge.Columns().Username: u.LoginName,
		dao.AppRecharge.Columns().Amount:   in.Amount,
		dao.AppRecharge.Columns().Type:     in.Type,
		dao.AppRecharge.Columns().Coin:     in.Type,
		dao.AppRecharge.Columns().Address:  in.Address,
		dao.AppRecharge.Columns().TxId:     in.TxId,
		dao.AppRecharge.Columns().FileName: in.Pic,
		dao.AppRecharge.Columns().Status:   0, // 0 = 待审核
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

	return nil
}

// WithdrawSubmit 提现动作
func (s *sFunding) WithdrawSubmit(ctx context.Context, userId uint64, in *v1.WithdrawSubmitReq) error {
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
	if err := dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userId).Scan(&u); err != nil {
		return gerror.NewCode(codes.UserNotFound, "找不到此用户")
	}
	if u.Status == 1 || u.IsFreeze == "2" {
		return gerror.NewCode(codes.UserDisabled, "您的账户状态异常，禁止提现")
	}
	// TODO: 判断 u.TotleAmont 是否 >= 打码量阈值

	// 3. 构建提现记录
	// 简单的手续费与实际到账换算假定
	fee := 0.0
	realAmount := in.Amount - fee

	record := g.Map{
		dao.Withdraw.Columns().UserId:     userId,
		dao.Withdraw.Columns().Username:   u.LoginName,
		dao.Withdraw.Columns().Amount:     in.Amount,
		dao.Withdraw.Columns().Fee:        fee,
		dao.Withdraw.Columns().RealAmount: realAmount,
		dao.Withdraw.Columns().Type:       in.CoinType,
		dao.Withdraw.Columns().Coin:       in.Coin,
		dao.Withdraw.Columns().ToAdress:   in.Address,
		dao.Withdraw.Columns().Status:     0, // 发起为 0: 审核中
	}

	// 在双锁防穿透的环境下执行资金冻结扣减
	freezeReq := &adminV1.SubAmountReq{
		UserId:     int64(userId),
		Symbol:     in.Coin,
		Amount:     in.Amount, // 正数由于我们内部逻辑，会走可用减，冻结加
		RecordType: 21,        // RecordType 21 假设为 "发起提现冻结"
		Remark:     "发起提现",
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

	return nil
}
