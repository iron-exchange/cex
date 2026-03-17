package funding

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	adminV1 "GoCEX/api/admin/v1"
	v1 "GoCEX/app/api"
	"GoCEX/internal/codes"
	"GoCEX/internal/consts"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"

	"github.com/shopspring/decimal"
)

type sRecall struct{}

func NewRecall() *sRecall {
	return &sRecall{}
}

// UncCallback 优盾代收网关主验签逻辑
func (s *sRecall) UncCallback(ctx context.Context, in *v1.UncCallbackReq) error {
	// 1. 第一关: IP 白名单防御
	clientIp := g.RequestFromCtx(ctx).GetClientIp()

	// 从配置中读取允许回调的 IP 列表 (如果为空默认全放行)
	allowedIps := g.Cfg().MustGet(ctx, "udun.whitelistIps").Strings()
	if len(allowedIps) > 0 {
		isAllowed := false
		for _, ip := range allowedIps {
			if clientIp == ip {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			g.Log().Warningf(ctx, "[优盾防御] 未授权的 IP 访问回调: %s", clientIp)
			return gerror.NewCode(codes.Unauthorized, "IP forbidden")
		}
	}

	// 2. 第二关: 签名防御 (MD5: body + apiKey + nonce + timestamp)
	apiKey := "my_udun_cloud_secret_key_123" // 兜底默认值
	record, err := dao.Setting.Ctx(ctx).Where("id", "UDUN_SETTING").One()
	if err == nil && !record.IsEmpty() {
		if j, err := gjson.DecodeToJson(record["setting_value"].String()); err == nil {
			apiKey = j.Get("merchantKey").String()
		}
	}

	rawString := in.Body + apiKey + in.Nonce + in.Timestamp
	expectedSign, _ := gmd5.EncryptString(rawString)

	if expectedSign != in.Sign {
		g.Log().Warningf(ctx, "[优盾防御] 签名篡改预警 - IP: %s", clientIp)
		return gerror.NewCode(codes.Unauthorized, "Sign verify failed")
	}

	// 3. 第三关: 业务解析
	var udunBody v1.UdunBody
	if err := json.Unmarshal([]byte(in.Body), &udunBody); err != nil {
		g.Log().Errorf(ctx, "[优盾解析] 反序列化错误: %v", err)
		return gerror.NewCode(codes.ClientError, "Parse body failed")
	}

	// 统一计算金额 (amount / 10^decimals)
	amountDec, _ := decimal.NewFromString(udunBody.Amount)
	precisionDec := decimal.New(1, int32(udunBody.Decimals))
	finalAmount := amountDec.Div(precisionDec).String()

	// 4. 第四关: 逻辑分支处理
	if udunBody.TradeType == 2 {
		return s.handleWithdrawCallback(ctx, &udunBody, finalAmount)
	} else if udunBody.TradeType == 1 {
		return s.handleRechargeCallback(ctx, &udunBody, finalAmount)
	}

	return nil
}

// handleWithdrawCallback 处理提现/代付回调
func (s *sRecall) handleWithdrawCallback(ctx context.Context, udunBody *v1.UdunBody, amount string) error {
	return dao.Withdraw.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var order entity.Withdraw
		err := dao.Withdraw.Ctx(ctx).TX(tx).LockUpdate().Where("serial_id", udunBody.BusinessId).Scan(&order)
		if err != nil {
			g.Log().Warningf(ctx, "[优盾代付] 找不到订单: %s", udunBody.BusinessId)
			return nil
		}

		// 幂等校验：如果是终态则跳过
		if order.Status == 1 || order.Status == 2 {
			return nil
		}

		if udunBody.Status == 3 {
			// 提现成功
			_, err = dao.Withdraw.Ctx(ctx).TX(tx).Data(g.Map{
				dao.Withdraw.Columns().Status: 1,
				dao.Withdraw.Columns().Fee:    udunBody.Fee,
			}).Where("id", order.Id).Update()
			return err
		} else if udunBody.Status == 2 || udunBody.Status == 4 {
			// 提现失败/驳回 -> 资金回滚
			rate := asset.New().GetExchangeRate(ctx, order.Coin)
			unAmountDec, _ := decimal.NewFromString(gconv.String(order.Amount))
			uAmount := unAmountDec.Mul(rate).InexactFloat64()

			unfreezeReq := &adminV1.SubAmountReq{
				UserId:     int64(order.UserId),
				Symbol:     order.Coin,
				AmountStr:  gconv.String(order.Amount),
				RecordType: consts.RecordTypeWithdrawFail, // 提现失败退回
				Remark:     "优盾提现失败，资金退回",
				UAmount:    uAmount,
			}
			_, err = asset.New().UnfreezeAmount(ctx, unfreezeReq)
			if err != nil {
				return err
			}

			_, err = dao.Withdraw.Ctx(ctx).TX(tx).Data(g.Map{
				dao.Withdraw.Columns().Status:         2,
				dao.Withdraw.Columns().WithDrawRemark: "优盾提现失败",
			}).Where("id", order.Id).Update()
			return err
		}
		return nil
	})
}

// handleRechargeCallback 处理充值回调
func (s *sRecall) handleRechargeCallback(ctx context.Context, udunBody *v1.UdunBody, amount string) error {
	if udunBody.Status != 3 {
		return nil
	}

	// 1. 定位玩家
	var userAddr entity.AppUserAddress
	err := dao.AppUserAddress.Ctx(ctx).
		Where(dao.AppUserAddress.Columns().Address, udunBody.Address).
		Where(dao.AppUserAddress.Columns().Symbol, udunBody.Coin).
		Scan(&userAddr)

	if err != nil || userAddr.UserId == 0 {
		g.Log().Errorf(ctx, "[优盾风控] 找不到持有地址 %s 且币种为 %s 的用户", udunBody.Address, udunBody.Coin)
		return gerror.NewCode(codes.Failed, "Address Unknown Or Coin Mismatch")
	}

	var user entity.AppUser
	_ = dao.AppUser.Ctx(ctx).Where(dao.AppUser.Columns().UserId, userAddr.UserId).Scan(&user)

	// 2. 幂等与资产注资
	rate := asset.New().GetExchangeRate(ctx, udunBody.Coin)
	uAmountDec, _ := decimal.NewFromString(amount)
	uAmount := uAmountDec.Mul(rate).InexactFloat64()

	subReq := &adminV1.SubAmountReq{
		UserId:     userAddr.UserId,
		Symbol:     udunBody.Coin,
		AmountStr:  amount,
		RecordType: consts.RecordTypeRecharge,
		Remark:     "优盾代收自动到账",
		UAmount:    uAmount,
	}

	_, err = asset.New().AddAmount(ctx, subReq, func(ctx context.Context, tx gdb.TX) error {
		count, txErr := dao.AppRecharge.Ctx(ctx).TX(tx).Where(dao.AppRecharge.Columns().TxId, udunBody.TxId).Count()
		if txErr != nil {
			return txErr
		}
		if count > 0 {
			return gerror.NewCode(codes.Failed, "Idempotency Rejected")
		}

		// 创建充值记录
		record := g.Map{
			dao.AppRecharge.Columns().UserId:   userAddr.UserId,
			dao.AppRecharge.Columns().Username: user.LoginName,
			dao.AppRecharge.Columns().Amount:   amount,
			dao.AppRecharge.Columns().Type:     udunBody.Coin,
			dao.AppRecharge.Columns().Coin:     udunBody.Coin,
			dao.AppRecharge.Columns().Address:  udunBody.Address,
			dao.AppRecharge.Columns().TxId:     udunBody.TxId,
			dao.AppRecharge.Columns().Status:   1, // 成功
		}
		_, txErr = dao.AppRecharge.Ctx(ctx).TX(tx).Data(record).Insert()
		return txErr
	})

	if err != nil {
		if gerror.Code(err) == codes.Failed && err.Error() == "Idempotency Rejected" {
			return nil
		}
		return err
	}

	return nil
}
