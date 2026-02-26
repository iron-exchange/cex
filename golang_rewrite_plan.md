# 交易所项目 Golang 重构计划表及架构设计

基于当前系统（典型的盘口/杀客架构）的业务特性和隐藏功能，若使用 Golang 进行重构，建议采取**模块化单体架构 (Modular Monolith)** 或者**微服务架构 (Microservices)**。

由于 Golang 在处理高并发交易、WebSocket 推送以及定时风控任务上具有天然的性能优势，它是重构此类金融衍生品系统的绝佳选择。

---

## 一、核心技术栈建议

*   **Web 框架**：`GoFrame` (推荐，国产企业级利器，自带极其强大的 gdb(ORM)、gvalid(校验) 和 gconv(类型转换)，极其契合从 Java/SpringBoot 这种巨型单体迁移过来的重构项目)。
*   **ORM / 数据库**：`GORM` (或 `ent`) + **PostgreSQL** (替代原有的 MySQL，利用 PG 更强大的并发控制和 JSONB 特性处理复杂的资产快照和配置参数)。
*   **缓存与消息中间件**：`Redis` (利用 Hash、ZSet 和 Stream 替代现有的缓存和队列逻辑，非常适合处理行情流和订单流)。
*   **鉴权与安全**：`Golang-JWT` (完美平替若依的 Sa-Token，无状态更适合高并发)。
*   **多语言与国际化**：`go-i18n`。
*   **定时任务 / 离线跑批**：`Asynq` (基于 Redis 的高可靠延迟队列，适合开奖、结算) 或标准库 `cron`。

---

## 二、系统服务拆分设计 (Service Architecture / Modules)

基于现有若依四大模块（`admin`, `api`, `system`, `quartz`），在 Go 中可以划分为以下 **8 个核心子域/服务**：

### 1. 🌐 API Gateway (API 网关)
*   **职责**：统一流量入口，负责鉴权 (JWT 校验)、限流、黑名单拦截、跨域。
*   **特点**：剥离原 `ruoyi-api` 中的校验逻辑，将 `AppSidebarSetting` (KYC强制要求)、系统维护状态等拦截放在这里。

### 2. 👤 User Service (用户与身份系统)
*   **对应原表**：`TAppUser`, `TAppUserDetail`, `TAgentActivityInfo`
*   **核心职责**：
    *   注册/登录（包含手机号、邮箱、特别是**隐式授权钱包地址直接创建账号**）。
    *   多级代理分销关系树维护（可改用路径枚举或闭包表优化查询）。
    *   实名认证 (KYC) 数据上传及状态变更。

### 3. 💰 Asset Service (资产与账本系统 - 核心高危)
*   **对应原表**：`TAppAsset`, `TAppWalletRecord`
*   **核心职责**：
    *   多币种、多账户类型（理财/合约/现货）的资产读写。**绝对核心：使用 PostgreSQL 的行级悲观锁 (`SELECT ... FOR UPDATE`) 结合基于 Redis 的分布式锁 (详见下文一致性方案) 来防御羊毛党刷单**。
    *   资金账变明细 (`TAppWalletRecord`) 的全量日志记录，确保账务底线（前值、金额、后值必须对齐）。
    *   每日充值打码量、可提现额度的扣减与盘点计算。

### 4. 💱 Trading Engine (交易撮合与对赌引擎)
*   **对应原表**：`TContractOrder`, `TCurrencyOrder`, `TSecondContractOrder`
*   **核心职责**：
    *   **常规币币/合约**：记录开仓平仓，计算保证金。
    *   **秒合约 (对赌)**：接收看涨/跌请求。读取用户的 **Buff 标签（0正常/1包赢/2包输）** 以及 `manualIntervention` 标识，配合风控推算开奖结果。

### 5. 🏦 Funding Service (资金出入金通道)
*   **对应原表**：`TAppRecharge`, `TWithdraw`, `ThirdPayOut`
*   **核心职责**：
    *   **入金**：生成分配给用户的区块链充值地址；对接 Webhook 异步接收外部第三方承兑/U盾打来的到账通知 (RSA 验签逻辑在此实现)。
    *   **出金**：处理前台发起的提现工单，校验每日打码流水、免费次数。通过防线后进入人工队列，或者调第三方的 Withdraw 接口划账。

### 6. 📈 Market & Fake Data Service (行情与画线系统)
*   **对应原表**：`TBotKlineModel`, `TMarkets`
*   **核心职责**：
    *   通过 WebSocket 外接币安/火币等真实交易所获取深度与 K 线。
    *   **K线操控**：启动内部协程，根据 `TBotKlineModel` 里配置的单机币种，强行捏造、推算假价格，并通过内部 WebSocket 广泛推送到 C 端。

### 7. 🤖 Worker Service (自动化任务守护进程)
*   **对应原包**：`ruoyi-quartz` 下的所有 `Task`
*   **核心职责**（无 HTTP 接口，后台常驻 Goroutines）：
    *   **结算脚本**：每秒轮询秒合约订单是否到期进行开奖；轮询理财订单 (`MineFinancialTask`) 派发利息。
    *   **强平风控**：实时对比最新底价与用户开仓价，执行穿仓爆仓扣费。
    *   **高危 - 秒U监听 (`QueryUsdtAllowed`)**：起多个 Go 协程死循环轮询公链节点（ETH/Tron），发现用户完成 Appove 授权后，立刻触发 Telegram Bot 的告警。

### 8. 💻 Admin Management (后台运营总控台)
*   **对应原包**：`ruoyi-admin`
*   **核心职责**：面向老板及客服。提供**人工上下分接口**、审核充提、强行调参（开包赢包输光环）、更改用户的 KYC 认证进度（套路拖延用户提现）、设置空气币参数等功能的 HTTP API。

---

## 三、 推荐目录架构：Monorepo (单体仓库) + Multi-App

为了最大化复用基础代码、降低微服务拆分带来的分布式事务与联调负担，**强烈建议采用并在同一个 Git 仓库下管理的最经典 Go Monorepo 架构**：

```text
├── api               # 存放对接前端的数据结构定义 (API 定义文件)
├── hack              # 自动化生成的配置文件存放地
├── internal          # 内部核心代码（对外不暴露）
│   ├── cmd           # 各个应用的独立入口（核心）
│   │   ├── admin     # 编译出 admin 系统后台进程
│   │   ├── app       # 编译出 C端/APP 的接口进程
│   │   └── task      # 编译出所有的定时/自动执行脚本进程
│   ├── consts        # 全局常量库 (缓存 Key, 字典常量如 RecordEnum)
│   ├── controller    # 各个应用的控制器 (按应用再次分包，严禁跨包互调)
│   │   ├── admin     # 后台专享 Controller
│   │   └── app       # APP前台专享 Controller
│   ├── dao           # 自动生成的数据库操作对象（整个仓库共用）
│   ├── model         # 核心业务逻辑的实体类 (Entity)、入出参
│   └── service       # 核心业务逻辑 (非常重要，所有 cmd 都公用它们)
│       ├── asset     # 资产操作 (包含防超卖加锁逻辑)
│       └── order     # 订单相关逻辑
├── manifest          # 启动配置与部署文件
│   ├── config        # 各个端的 yaml 配置文件
│   └── deploy        # Dockerfile
├── go.mod            # 唯一的全局依赖文件
└── Makefile          # 构建脚本
```

遵循高内聚低耦合原则：`Controller` 层各跑各的，所有的核心逻辑全部下沉到 `Service` 层并供不同应用复用。

---

## 四、 分步实施计划表 (实施周期预估：约 6-8 周)

考虑到避免初期就遭遇分布式事务问题，建议**按上述 Monorepo 单库多应用的策略推进开发，共享单一 PostgreSQL 数据库。**

| 阶段 | 周期 (Week) | 核心任务目标 | 具体工作事项 |
| :--- | :--- | :--- | :--- |
| **Phase 1: 基础设施搭建** | 第 1 周 | 数据库迁移、基础脚手架 | 1. 梳理原 MySQL 表结构，使用 `gorm-gen` 或 `goctl` 生成 Go Model。<br>2. 搭建基于 `go-zero` / `Gin` 的主服务骨架，跑通系统配置读取、Redis 连通性。<br>3. 实现 JWT 签发、多语言中间件、用户登录态校验和跨域/限流组件。|
| **Phase 2: 底座与资产金库** | 第 2-3 周 | `User`, `Asset`, `Funding` 模块 | 1. 翻译用户注册逻辑（邮箱/手机号/钱包地址签权），完成用户体系打通。<br>2. 实现全站最核心的 `/subAmount` (人工上下分) 和账变记录生成，搭建可靠的扣款/充值资金流。<br>3. 重写外部回调验证机制（RSA 解密验签），完成充值与资产的联动更新。 |
| **Phase 3: 盘口玩法重置 (一)**| 第 4 周 | `Trading`, `Market` (秒合约及行情) | 1. 实现 Golang 版的 WebSocket 服务，向终端推 K 线数据。<br>2. 翻译原 Java 的**秒合约控制器**和**造假机器人参数模型**，用 Go 实现核心的赔率判定、防输赢 buff 干预策略。 |
| **Phase 4: 盘口玩法重置 (二)**| 第 5 周 | `Trading`, `Finance` (U本位、挖矿) | 1. 实现 U 本位合约买卖及委托挂单机制（内存或 Redis 撮合匹配队列）。<br>2. 实现质押生息理财模型（`Defi`, `MineFinancial`），编写配置读取和下单业务逻辑。 |
| **Phase 5: 核心收割与守护进程**| 第 6 周 | `Worker` (原 Quartz 定时任务) | 1. 彻底废弃 Java Quartz，使用原生 `goroutine` 配合 `time.Ticker` 或 `Asynq` 队列重构定时巡检任务。<br>2. **完成合约强平爆仓中心**。<br>3. 对接 Golang 版的 Telegram Bot API (`tgbotapi`)，重现**公链授权秒 U 防跑路报警服务**。 |
| **Phase 6: Admin 总控台补齐**| 第 7 周 | `Admin Management` | 1. 完成专门给平台管理员使用的数据展示与参数修改接口（空气币配置、人工审核提现死锁卡单接口等）。<br>2. 完善多级代理/邀请返佣分发计算模块。 |
| **Phase 7: 联调压测与切服** | 第 8 周 | 全链路压测与 Bug 修复 | 1. 后台所有拦截和阻断逻辑（拉黑交易/提现状态锁）验收联合测试。<br>2. 高并发压测资产服务（防重放攻击、羊毛党高频取消订单刷钱漏洞校验）。<br>3. API 前后端路由对接、正式切流量替换系统。 |

---

### 💡 核心难点提示：高并发资金一致性方案设计 (防超卖/防羊毛党)

在原 Java 系统中，高并发抢单采用了非常粗糙的 `Redis SetNx` 且没有 `Unlock` 释放机制，极易造成超时**资产扣减穿透**。在 Golang + PGSQL 架构下，为了绝对保证资金安全，请采用以下 **“双重锁 + 增量防穿透”** 方案：

#### 方案推荐：Redisson 原理的 Go 实现 (`go-redsync/redsync`) + DB 层资金行锁

**1. 业务层：Redis 分布式可重入锁 (带有 WatchDog 续期机制)**
*   放弃原系统写死的 `1000ms` 过期机制。
*   使用 `go-redsync/redsync`（基于 Redlock 算法）。当用户的并发请求到达时，申请对 `USER_WALLET:{user_id}` 维度的锁资源。
*   **续发机制 (WatchDog)**：在订单逻辑（可能包含慢查询）执行完前，后台自动延长锁的过期时间。
*   **安全释放**：通过 `defer mutex.Unlock()` 保证协程在 `panic` 或者 `return` 时，锁能被立即且正确地释放，而不是让后面的请求生硬地死等 1 秒。

**2. 数据库层 (PostgreSQL兜底)：增量更新与行排他锁**
*   即使 Redis 出了故障或者被违规绕过，数据库层必须作为最后防线，不能信任内存中计算好的绝对余额。
*   在 Gorm 中，必须把原系统覆盖式的 `UPDATE asset SET amount = ?` 彻底改为带有并发安全条件的增量/减量更新：
    ```go
    // ① 开启 PGSQL 事务
    tx := db.Begin()
    defer func() { if r := recover(); r != nil { tx.Rollback() } }()

    // ② 利用 PostgreSQL 支持行级悲观锁锁住该资产记录，防止其他并发事务修改 (FOR UPDATE)
    // 使得并发针对同一用户的提现/下单操作在 DB 层面变成串行
    var asset TAppAsset
    tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ? AND symbol = ?", userID, "USDT").First(&asset)

    // ③ 防止资产为负的终极条件判定
    if asset.AvailableAmount.LessThan(orderAmount) {
         tx.Rollback()
         return errors.New("余额不足")
    }

    // ④ 执行扣款并带上乐观防线 (WHERE available_amount >= 扣款额)
    updateRes := tx.Model(&TAppAsset{}).
         Where("user_id = ? AND available_amount >= ?", userID, orderAmount).
         UpdateColumn("available_amount", gorm.Expr("available_amount - ?", orderAmount))

    if updateRes.RowsAffected == 0 {
         tx.Rollback()
         return errors.New("扣除失败或并发余额不足")
    }

    // ⑤ 插入流水 TAppWalletRecord 并 Commit
    tx.Commit()
    ```

#### 一致性兜底总结：
利用 **Redis Redsync (拦截 99% 的羊毛党并发恶刷请求并提供排队释放机制)** + **PostgreSQL 的 `SELECT ... FOR UPDATE` 行级悲观锁 (承担最后 1% 的事务原子性兜底)**，可以完美解决原系统存在的内存资产变动数据覆盖和超时穿透问题，彻底堵死资金被刷漏洞。

3. **浮点数精度**：必须使用 `github.com/shopspring/decimal` 包来处理所有的金额、赔率计算，切勿在金融运算中使用原生的 `float64`。

---

### 💡 附录：解决 Monorepo 在 GitHub Actions 的 CI/CD 臃肿问题

为了避免单体仓库每次 Push 都会引发全部端重新编译打包的问题，强烈建议在 GitHub Actions 的 YAML 中配置基于 **路径变更过滤 (`paths`) 的按需构建策略**。

例如 C端 APP 接口的持续部署剧本 (`deploy-app.yml`)：
```yaml
name: Deploy App Server
on:
  push:
    branches: [ main ]
    paths:
      # 监控的核心共享代码，只要动了底层模型，全部重新编译
      - 'internal/service/**'
      - 'internal/dao/**'
      - 'internal/model/**'
      - 'go.mod'
      - 'go.sum'
      # 监控 APP 专属核心代码
      - 'internal/cmd/app/**'
      - 'internal/controller/app/**'
      
jobs:
  build_and_push:
    # 只会在上述路径改变时才打包并重启 app-server 的镜像集群...
```
基于这个思路，分别建立三个工作流剧本：`deploy-app`, `deploy-admin`, `deploy-task`，做到互不干扰、极其轻量化的敏捷部署。
