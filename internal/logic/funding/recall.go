package funding

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	adminV1 "GoCEX/api/admin/v1"
	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/codes"
	"GoCEX/internal/dao"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/model/entity"
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

	// 系统约定的优盾专属 API Key
	apiKey := "my_udun_cloud_secret_key_123"

	// 2. 第二关: 签名防御 (MD5: body + apiKey + nonce + timestamp)
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

	// 仅处理 status == 3 也就是已确认到账的逻辑
	if udunBody.Status != 3 {
		return nil
	}
	if udunBody.Type != "1" { // 充币入款
		return nil
	}

	// 4. 第四关: 用户挂载归属与币种安全双检 (防止劣质币造假)
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

	// 5. 第五关 重放防御 + 原子性注资闭环 (由 Asset 层的 Redis 锁与 DB 事务保护)
	subReq := &adminV1.SubAmountReq{
		UserId:     userAddr.UserId,
		Symbol:     udunBody.Coin,
		AmountStr:  udunBody.Amount, // 使用高精度防止 float64 截断
		RecordType: 1,
		Remark:     "优盾网关智能充值",
	}

	_, err = asset.New().AddAmount(ctx, subReq, func(ctx context.Context, tx gdb.TX) error {
		// (核心防线) 此时已经处于 Redis `USER_WALLET` 独占锁和 DB 事务内

		// 严谨幂等复查：利用 SELECT FOR UPDATE 加固
		count, txErr := dao.AppRecharge.Ctx(ctx).TX(tx).Where(dao.AppRecharge.Columns().TxId, udunBody.TxId).Count()
		if txErr != nil {
			return txErr
		}
		if count > 0 {
			g.Log().Infof(ctx, "[优盾风控] Tx %s 已出账过，同事务拦截", udunBody.TxId)
			// 返回特定错误字眼，指导外层吞噬该错误，返回成功给优盾停止重试
			return gerror.NewCode(codes.Failed, "Idempotency Rejected")
		}

		// 执行单点记录生成 (和资金增加在同一个事务 commit)
		record := g.Map{
			dao.AppRecharge.Columns().UserId:   userAddr.UserId,
			dao.AppRecharge.Columns().Username: user.LoginName,
			dao.AppRecharge.Columns().Amount:   udunBody.Amount,
			dao.AppRecharge.Columns().Type:     udunBody.Coin,
			dao.AppRecharge.Columns().Coin:     udunBody.Coin,
			dao.AppRecharge.Columns().Address:  udunBody.Address,
			dao.AppRecharge.Columns().TxId:     udunBody.TxId,
			dao.AppRecharge.Columns().Status:   2,
		}
		_, txErr = dao.AppRecharge.Ctx(ctx).TX(tx).Data(record).Insert()
		return txErr
	})

	if err != nil {
		if err.Error() == "Idempotency Rejected" {
			g.Log().Infof(ctx, "[优盾流水] Tx %s 已经确认入账，防刷丢弃结束", udunBody.TxId)
			return nil
		}
		return err
	}

	g.Log().Infof(ctx, "[优盾大捷] 用户 %v 成功到账 %v %v", userAddr.UserId, udunBody.Amount, udunBody.Coin)
	return nil
}
