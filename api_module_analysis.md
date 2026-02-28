# `ruoyi-api` 模块接口深度分析及 Golang 重写蓝图

本文档基于 `ruoyi-api` 原始 Java 代码，提取了核心控制器的每一个主要功能接口，并**详细列出了请求和响应的核心字段**，作为 Golang 版本接口参数绑定结构体（Req/Res 结构体）的直接参考。

---

## 1. 用户认证与基础体系 (`/api/user`)
主要控制器：`TAppUserController`
用于所有 C 端用户的注册、登录、安全设置以及身份认证。

### 1.1 用户登录
*   **路径**：`POST /api/user/login`
*   **功能**：支持账号/邮箱/手机登录，成功后返回 token (`StpUtil.getTokenValue()`)。
*   **请求字段 (TAppUser)**：
    *   `signType` (String, 必填): 决定登录方式 `LOGIN` (账号), `PHONE` (手机), `EMAIL` (邮箱)
    *   `loginName`/`phone`/`email` (String): 根据 `signType` 对应传入
    *   `loginPassword` (String)
    *   `code` (String): 如果是手机或邮箱登录，需传入验证码
*   **响应体**：
    *   `token` / 或者原框架自定义的 Token_Name 字段对应的值。

### 1.2 用户注册
*   **路径**：`POST /api/user/register`
*   **功能**：处理 App 端注册逻辑，根据 `signType` 判断不同的注册类型。
*   **请求字段 (TAppUser)**：
    *   `signType` (String): `LOGIN` / `PHONE` / `EMAIL` / `ADDRESS` (钱包地址直签授权注册)
    *   `loginName` (String): 用户名
    *   `loginPassword` (String): 登录密码（未传默认 123456）
    *   `phone` (String): 手机号
    *   `email` (String): 邮箱
    *   `address` (String): ERC20/TRC20 钱包地址
    *   `code` (String): 验证码
*   **响应体**：
    *   整个 `TAppUser` 用户对象脱敏字段后返回。

### 1.3 密码及安全设置大类
*   **修改/设置登录密码**：`POST /api/user/pwdSett` 
    *   `pwd` (String)
*   **修改交易密码（资金密码）**：`POST /api/user/tardPwdSet`
    *   `pwd` (String)
*   **绑定手机**：`POST /api/user/bindPhone`
    *   `phone` (String), `code` (String)
*   **绑定邮箱**：`POST /api/user/bindEmail`
    *   `email` (String), `emailCode` (String)
*   **地址静默绑定**：`POST /api/user/updateUserAddress`
    *   `address` (String): 钱包地址
    *   `userId` (Long)
    *   `type` (String): 钱包类型 (`ETH`, `TRON`)

### 1.4 KYC 实名认证
*   **路径**：`POST /api/user/uploadKYC`
*   **请求字段**：
    *   `realName` (String): 真实姓名
    *   `idCard` (String): 证件或者护照号码
    *   `frontUrl` (String): 正面照 URL
    *   `backUrl` (String): 反面照 URL
    *   `handelUrl` (String): 手持证件照 URL
    *   `country` (String): 国家代码
    *   `cardType` (String): 证件类型
    *   `flag` (String): "2" 会触发基于 `APP_SIDEBAR_SETTING` 的高级实名开关校验。

---

## 2. 资金出入金枢纽
主要控制器：`TAppRechargeController` 和 `TAppWithdrawController`

### 2.1 会员充值提交
*   **路径**：`POST /api/recharge/submit`
*   **功能**：会员发起一笔充值订单。
*   **请求类型**：`@RequestBody Map<String, Object>`
    *   `amount` (BigDecimal或String): 充值数量
    *   `type` (String): 充值数字货币名称（对应表配置的 `ASSET_COIN`），例如 USDT, BTC
    *   `address` / 其他扩展字段会透传给 `TAppRecharge` 对象。
*   **行为**：后台检查是否在 `rechargeMin` 到 `rechargeMax` 之间，然后入库，并且打入 Redis Stream (`CacheConstants.RECHARGE_KEY`)，由后台程序或客服接手处理审核。

### 2.2 用户提现提交
*   **路径**：`POST /api/withdraw/submit`
*   **功能**：会员发起提现申请，扣减余额，冻结资金。
*   **请求字段 (Form 表单)**：
    *   `amount` (BigDecimal): 提现额度
    *   `coinType` (String): U本位或某种数字代币通道 (TRC20/ERC20)
    *   `pwd` (String): 资金密码
    *   `adress` (String): 提现目标钱包地址
    *   `coin` (String): 具体的提现币种（比如 USDT）
*   **行为**：强校验打码量(`t_app_user`表的`totleAmont / rechargeAmont`)，强校验是否实名，之后进入提现审核待定记录，并推 Socket 消息。

---

## 3. 交易打单接口群
平台利润和核心玩法的入口。包含了传统的币币交易、U本位合约挂单、以及平台自创的“秒合约”(期权/极速合约)。

### 3.1 现货币币交易
*   **路径**：`POST /api/currency/order/submit` 等（推测，由业务包名 `TCurrencyOrderController` 驱动）
*   **请求字段 (TCurrencyOrder)**：
    *   `symbol` (String): 交易对，如 `btc_usdt`
    *   `tradeDirection` (Integer/String): 方向 1买 2卖
    *   `price` (BigDecimal): 挂单价格（如果是限价单）
    *   `quantity` (BigDecimal): 买入/卖出数量
    *   `orderType` (Integer): 订单类型 (市价/限价)
*   **风控**：会对用户是否是被风控标记的人进行检测拦截。

### 3.2 杠杆合约挂单
*   **路径**： `/api/contract/order/buyContractOrder` 等
*   **请求字段 (TContractOrder)**：
    *   `symbol` (String): 期货交易对，例如 `btc_usdt`
    *   `leverage` (BigDecimal/Integer): 杠杆倍数
    *   `direction` (String/Integer): 多(`BUY`) / 空(`SELL`)
    *   `price` (BigDecimal): 下单价
    *   `margin` (BigDecimal): 指定的保证金
    *   `orderType` (Integer): 限价/市价仓位
*   **结算机制**：由后端的 Quartz 定时任务结合实时的 K 线涨跌对该表进行轮询爆仓或止盈止损。

### 3.3 极速期权（秒合约）下单
*   **路径**：`POST /api/secondContractOrder/createSecondContractOrder`
*   **功能**：俗称“猜大小单 / 二元期权单”。
*   **请求字段 (TSecondContractOrder)**：
    *   `symbol` (String): 购买币种
    *   `direction` (String): 涨或跌
    *   `buyAmount` (BigDecimal): 金额
    *   `period` / `secondConfigId` (Long): 买具体的哪一个时间段（如 60秒、3分钟）
    *   `profitRate` (BigDecimal): 固定赔率收益率（前端传或后端读配置）
*   **结算机制**：**黑产最核心表**。在 `t_app_user` 对象中有一个隐藏极深的 `buff` 字段（`0正常 1包赢 2包输`），秒合约结算会在倒计时结束那根 K 线时，强行检查如果玩家 `buff=1/2` 或者全盘正在执行“杀大赔小”的风控逻辑时，进行客损交割。

---

## 4. API 响应体共识 (`AjaxResult`)
Golang 重写时，所有的对外 JSON HTTP 响应都应当维持原 Java 版 `AjaxResult` (或标准 Result) 的字段风格，避免前端 App 大规模重构：

```json
{
  "code": 200,          // 业务响应码 (原样保留 200 成功，其他值为失败，前端可全局拦截)
  "msg": "操作成功",     // 给前台用户的提示文案
  "data": { ... }       // 具体的对象数据 (List / Map / POJO)
}
```

## Golang 重写的重要建议（Tips）：
1. **参数绑定最佳实践**：GoFrame 有一个非常强大的 `g.Map()` 和 `gvalid`。类似 `TAppUser` 这个实体在注册时复用了很多字段，建议在 Golang 的 `controller` 层定义如 `UserRegisterReq` 的结构体，通过 `v:"required"` 来做强校验，不要把整个大的 `entity` 或 `DO` 直接暴露给前台。
2. **多登录渠道的解耦**：像 `/register` 方法里面大量用了 `if (LoginOrRegisterEnum.EMAIL.getCode().equals(signType))`。在 Go 中可以用 `switch-case` 结合设计模式或者接口实现进行策略解耦，代码看起来会清爽很多。
### 6.4 门户与内容管理 (CMS)
*   **基础路径**: `/api/notice`, `/api/helpcenter`, `/api/mail`, `/api/app/noticeTop`
*   **核心接口**:
    *   `/api/notice/getAllNoticeList`: 获取首页轮播图、走马灯系统公告。
    *   `/api/helpcenter/list`: 帮助中心文档（客服指引、新手教程等）。
    *   `/api/mail/listByUserId`: 个人站内信推送列表（后台客服针对单个用户发送的话术消息）。

---

## 7. 全局公共配置与基础组件 (Common & App & Blockcc)

在 `ruoyi-api` 中，除了 `bussiness` 文件夹下的核心业务，还有几个非常关键的高频被叫网关接口。

### 7.1 系统全局动态配置 (`CommonController`)
*   **基础路径**: `/api/common`
*   **核心特性**: 这些接口**请求频率极高**，决定了 App 端 UI 的长相以及充提拉新的各种门槛。在 Golang 重构时**必须全部接入 Redis/LocalCache 多级缓存**，强烈建议采用 Pub/Sub 机制热更，坚决不能直接扫查 DB。
*   **核心接口**:
    *   `POST /getAllSetting`: 将全站参数一揽子下发给前端（减少 HTTP 握手）。
    *   `POST /getAppSidebarSetting` / `getHomeCoinSetting`: 侧边栏及首页显示哪些币种。
    *   `POST /getAppCurrencyList` / `getWithDrawCoinList`: 充提通道开关与额度费率下发。
    *   `POST /upload/OSS`: 统一的阿里云 OSS 文件（头像、KYC 照片）上传凭证与中转。

### 7.2 行情与 K线控制 (`BlockccController`, `TMarketController`)
*   **基础路径**: `/api`, `/api/app/market`
*   **核心接口**:
    *   `GET /api/app/market/list`: 获取盘口大厅的币种行情列表。
    *   `POST /api/kline` & `/api/newKline`: K 线深度数据拉取口。在原 Java 架构中，前端依赖这个接口进行高频 HTTP 轮询。**Golang 重构时，这部分逻辑应全盘作废，强制替换为 WebSocket 定时下发 (Push) 架构**，极大降低服务器解析压力。

3. **安全加密**：Java 版用了 `SecurityUtils.encryptPassword()`，你在 Go 里面替换成标准的 `golang.org/x/crypto/bcrypt` 会非常完美且更安全。


## 5. API 迁移重写进度表 (Golang版)

| 模块 | 原 API 路径 | 接口功能 | 当前进度 | 备注说明 |
| :--- | :--- | :--- | :--- | :--- |
| **基础认证** | `/api/user/login` | 多渠道登录 (账号/手机/邮箱/Web3) | ✅ 已完成 | 包含 JWT 签发、黑名单校验、地址类型静默登入 |
| **基础认证** | `/api/user/register` | 用户注册 (含 Web3 静默绑包) | ✅ 已完成 | 已实现邀请码生成和代理关系继承，支持各渠道注册 |
| **基础设置** | `/api/user/pwdSett` | 设置登录密码 | ✅ 已完成 | 已实现加盐哈希更新密码逻辑 |
| **基础设置** | `/api/user/tardPwdSet` | 设置资金(交易)密码 | ✅ 已完成 | 已实现自动绑定并插入到 `app_user_detail` 明细表中 |
| **基础设置** | `/api/user/bindPhone` | 绑定手机 | ✅ 已完成 | 已添加严格查重逻辑，支持安全绑定 |
| **基础设置** | `/api/user/bindEmail` | 绑定邮箱 | ✅ 已完成 | 已添加严格查重逻辑，支持安全绑定 |
| **基础设置** | `/api/user/updateUserAddress` | 静默绑定钱包地址 | ✅ 已完成 | 已实现双向绑定：同时更新主表快速出账冗余与 `app_user_address` 地址库子表 |
| **基础设置** | `/api/user/uploadKYC` | 上传 KYC 实名认证 | ✅ 已完成 | 已实现包含图床 URL 与基本身份信息的映射存证到 `app_user_detail` |
| **资金出入** | `/api/recharge/submit` | 充值订单申请 | ✅ 已完成 | 已实现充值待审记录入库，并通过 XADD 发送至 `CEX:STREAM:RECHARGE` 阻塞队列 |
| **资金出入** | `/api/withdraw/submit` | 提现订单申请 | ✅ 已完成 | 强校验资金安全密码，并调用 Asset 底层 `FreezeAmount` 进行并发双锁安全冻结扣款 |
| **币币交易** | `/api/currency/order/submit` | 现货挂单 (市价/限价) | ✅ 已完成 | 已实现委托拦截、`Asset`资金安全双重锁扣减，及撮合引擎的 `XADD` 消息分发桥接 |
| **期权/秒合约** | `/api/secondContractOrder/...` | 猜大小秒跌单 (支持 Buff) | ✅ 已完成 | 已实现下注扣款、下注周期锁定，并完成核心黑场逻辑的自动包输/包赢 `Buff/Sign` 特权拦截打标 |
| **永续合约** | `/api/contract/order/buyContractOrder` | U本位做多/做空 (插针系统) | ✅ 已完成 | 已实现利用配置杠杆倍数乘区自动计算强平保证金，完成 `Asset` U本位保证金冻结抽取及 `Redis` 打包 |
| **回调引擎** | `/api/recall/unc` | 第三方代收付 Webhook | ✅ 已完成 | 已开放 API 路由并绕过 JWT，通过构建 `APIKey` MD5 五重签名认证防线并结合 `TxId` 幂等检测直达底层发钱 |

### 待发掘的 C端 (App) 遗漏核心接口群

在完成了**“资金流”**和**“指令流”**（注册、充提、下单）的构建后，要让整个前台 App 顺利运转闭环，我们还缺以下**“查询流”与“状态管理流”**接口（建议列入下阶段开发计划）：

| 模块 | 推荐 API 路径 | 接口功能 | 当前进度 | 备注说明 |
| :--- | :--- | :--- | :--- | :--- |
| **资产大盘** | `/api/asset/list` | 钱包余额查询 | ❌ 未开始 | 获取用户所有币种的可用、冻结、折合 USDT 总资产 |
| **账单流水** | `/api/wallet/records` | 个人财务流水/账变明细 | ❌ 未开始 | 查询 `t_app_wallet_record`，展示上下分、充提、手续费明细 |
| **持仓记录** | `/api/contract/order/list` | 永续/秒合约/现货持仓列表 | ❌ 未开始 | 查询历史单、活跃仓位，供用户查看盈亏 |
| **平仓撤单** | `/api/currency/order/cancel` | 现货撤单 / 合约手动平仓 | ❌ 未开始 | 需要解冻退回保证金或结算盈亏（重要交互节点） |
| **充提记录** | `/api/recharge/list` 等 | 充值与提现历史查询 | ❌ 未开始 | 供用户查验自己的提现进度是否通过 |
| **行情大厅** | `/api/market/kline` 等 | K 线、Ticker、各币种实时价 | ❌ 未开始 | 供前端绘制闪兑行情与列表页 (通常直接读 Redis) |
| **公共配置** | `/api/common/config` 等 | 获取客服链接、最低充提阈值 | ❌ 未开始 | 供 App 启动时拉取基础配置、轮播图 Banner、系统公告等 |
