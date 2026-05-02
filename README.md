# QP Backend

## 项目结构

```
├── bin                  --> 构建生成的可执行文件
├── cmd                  --> 各服务入口（api/web/cron 等）
├── common               --> 通用能力（配置、日志、中间件、工具）
├── gen                  --> 代码生成产物
├── pkg                  --> 业务核心
│   └── core/modules     --> 业务模块（dos/dto/enmus/vo）
├── rpc_files            --> grpc 接口定义
├── scripts              --> 测试与冒烟脚本
├── go.mod               --> 依赖管理
└── README.md
```

## 本地运行

```bash
cd /Users/parker/qp/qp-backend
go mod tidy

# Web 管理端
go run ./cmd/web/main.go -config=./conf.yaml

# API 服务
go run ./cmd/api/main.go -config=./conf.yaml
```

## 服务启动矩阵（生产必看）

当前仓库不是只跑 `api + web`，完整运行通常包含 5 个进程：

- `cmd/web/main.go`：管理后台 Web 服务。
- `cmd/api/main.go`：前台 API 与回调入口。
- `cmd/cron/main.go`：定时任务（报表汇总、代付调度、超时订单处理）。
- `cmd/consumidor/main.go`：用户报表类消费者（充值/提现/优惠/站内信）。
- `cmd/betconsumidor/main.go`：注单消费者（注单入库、去重、未结算处理）。

本地最小联调启动示例：

```bash
cd /Users/parker/qp/qp-backend

# 终端1: web
go run ./cmd/web/main.go -config=./conf.yaml

# 终端2: api
go run ./cmd/api/main.go -config=./conf.yaml

# 终端3: cron
go run ./cmd/cron/main.go -config=./conf.yaml

# 终端4: 用户报表消费者
go run ./cmd/consumidor/main.go -config=./conf.yaml

# 终端5: 注单消费者
go run ./cmd/betconsumidor/main.go -config=./conf.yaml
```

说明：如果只启动 `api/web`，会出现“回调成功但报表不更新”“注单写入延迟/缺失”“代付无人处理”等现象。

## 依赖开关与最小可运行配置

初始化逻辑在 `common/initialization.go`，核心依赖开关如下：

- `Redis.IsInit=true`：启用 Redis（登录态、锁、验证码、幂等缓存依赖）。
- `Mysql.IsInit=true`：启用主库。
- `MysqlSharding.IsInit=true/false`：按是否使用分库分表启用。
- `Mq.IsInit=true`：启用 Kafka 生产者（异步链路必需）。

`conf.example.yaml` 最小建议：

- `Mysql.IsInit: true`
- `Redis.IsInit: true`
- `MQ.IsInit: true`
- `MQ.Kafka.Addr` 指向可用集群（非占位地址）

如果只做静态页面联调，可临时关闭 `MQ.IsInit`；但资金、注单、报表相关链路测试必须开启 MQ。

## 初始化与迁移顺序（避免上线踩坑）

结论（必须统一口径）：可以一次性完成数据库表与字段初始化，但前提是严格执行“SQL 增量 + AutoMigrate 兜底 + 上线前校验”三步，不能只跑其中一步。

推荐顺序：

1. 备份数据库。
2. 执行 `sql/` 与 `sql/updated/` 的增量脚本。
3. 执行建表与全局键初始化工具（`migrate_db.go`）。
4. 再启动 5 个服务进程。

迁移工具示例：

```bash
cd /Users/parker/qp/qp-backend
go run ./migrate_db.go
```

`migrate_db.go` 会处理两类关键动作：

- `AutoMigrate` 核心业务表。
- 初始化 `fc_global` 中基础自增键（例如 `USER_ID_INCR`、`INVITE_CODE_INCR`）。

常见错误：只执行了 API 启动，没有先准备 `fc_global` 基础键，导致注册/邀请码相关逻辑报错。

### 数据库一次性初始化（强制流程）

上线前按以下顺序执行，避免缺表缺字段：

1. 执行 `sql/` 与 `sql/updated/` 的历史与增量脚本（保证业务特定索引/约束/修复项完整）。
2. 执行 `go run ./migrate_db.go`（保证模型对应表和新增字段自动补齐）。
3. 执行数据库校验 SQL（如下），不通过则禁止发布。

推荐在发布流水线中把第 2 步和第 3 步作为硬门禁。

### 上线前数据库校验 SQL（建议直接复制）

将下面 SQL 的 `<DB_NAME>` 替换为实际库名，例如 `hy_game`。

```sql
-- 1) 核心表存在性检查（可按业务继续追加）
SELECT t.required_table,
			 CASE WHEN s.table_name IS NULL THEN 'MISSING' ELSE 'OK' END AS table_status
FROM (
	SELECT 'fc_global' AS required_table UNION ALL
	SELECT 'fc_order_deposit' UNION ALL
	SELECT 'fc_order_withdraw' UNION ALL
	SELECT 'fc_order_withdraw_payment_out' UNION ALL
	SELECT 'fc_pay_channel' UNION ALL
	SELECT 'fc_payment' UNION ALL
	SELECT 'fc_pay_channel_out' UNION ALL
	SELECT 'fc_payment_out' UNION ALL
	SELECT 'fc_bet_record' UNION ALL
	SELECT 'fc_bet_record_unsettled' UNION ALL
	SELECT 'fc_customer_link' UNION ALL
	SELECT 'fc_site_link' UNION ALL
	SELECT 'fc_sms_channel'
) t
LEFT JOIN information_schema.tables s
	ON s.table_schema = '<DB_NAME>'
 AND s.table_name = t.required_table;

-- 2) 关键字段存在性检查（防止“表有但字段缺”）
SELECT c.required_table,
			 c.required_column,
			 CASE WHEN s.column_name IS NULL THEN 'MISSING' ELSE 'OK' END AS column_status
FROM (
	SELECT 'fc_global' AS required_table, 'key' AS required_column UNION ALL
	SELECT 'fc_global', 'value' UNION ALL
	SELECT 'fc_order_withdraw_payment_out', 'status' UNION ALL
	SELECT 'fc_order_withdraw_payment_out', 'withdraw_status' UNION ALL
	SELECT 'fc_order_withdraw_payment_out', 'channel_code' UNION ALL
	SELECT 'fc_order_withdraw', 'status' UNION ALL
	SELECT 'fc_order_withdraw', 'order_sn' UNION ALL
	SELECT 'fc_order_deposit', 'status' UNION ALL
	SELECT 'fc_order_deposit', 'order_sn' UNION ALL
	SELECT 'fc_agent_domain', 'customer_link' UNION ALL
	SELECT 'fc_bet_record', 'venue_code' UNION ALL
	SELECT 'fc_bet_record', 'order_sn' UNION ALL
	SELECT 'fc_bet_record_unsettled', 'venue_code' UNION ALL
	SELECT 'fc_bet_record_unsettled', 'order_sn'
) c
LEFT JOIN information_schema.columns s
	ON s.table_schema = '<DB_NAME>'
 AND s.table_name = c.required_table
 AND s.column_name = c.required_column;
```

判定标准：上述结果中出现任意 `MISSING`，本次发布直接阻断，先补迁移再发布。

## Kafka Topic 与消费者映射

Topic 常量定义在 `pkg/service/channelData/enums.go`，建议在接入文档中固定维护映射关系：

- `userrecharge` -> `cmd/consumidor/controller/userAddCredit.go`
- `userwithdrawal` -> `cmd/consumidor/controller/userAddWithdrawal.go`
- `userpromotion` -> `cmd/consumidor/controller/userAddPromotion.go`
- `usersitemsg` -> `cmd/consumidor/controller/message.go`
- `betrecorddata` -> `cmd/betconsumidor/controller/betRecordData.go`

发送侧保护：`pkg/service/channelData/userReport.go` 中 `safeSend` 会在发送失败后自动重试一次；但这不是“最终一致性保障”，仍需依赖消费者幂等和补偿策略。

## 定时任务职责（与资金链路强相关）

`cmd/cron/crontab` 中高频关键任务：

- `rechargeDeposit.go`：超时未支付充值订单转失败。
- `withdrawPaymentOut.go`：轮询待代付订单并下发三方。
- `complexReport.go`：今日/昨日综合报表汇总。

因此，生产环境禁用 `cron` 进程会直接影响订单状态推进和报表数据完整性。

## 常见故障排查（建议加入值班手册）

- 现象：回调成功但用户报表不更新。
	- 检查：`cmd/consumidor` 是否在运行，`MQ.IsInit` 是否开启，Kafka 地址是否可达。
- 现象：注单未入库或延迟明显。
	- 检查：`cmd/betconsumidor` 是否在运行，topic `betrecorddata` 是否有堆积。
- 现象：提现审核后长时间不打款。
	- 检查：`cmd/cron` 是否在运行，`withdrawPaymentOut` 任务是否正常执行。
- 现象：注册/邀请码相关接口报基础键缺失。
	- 检查：是否执行过 `go run ./migrate_db.go` 并成功写入 `fc_global`。
- 现象：测试脚本 `finance/merchant` 无法执行。
	- 检查：Redis 管理员会话是否存在（先登录后台），并确认 API/Web 可访问。

## 编译

```bash
go build ./cmd/web/main.go
go build ./cmd/api/main.go
```

## 部署上线（生产建议）

### 1. 服务器与运行环境

- OS: Linux（建议 Ubuntu 22.04 LTS / CentOS 7+）
- CPU/内存（最小）:
	- 测试环境: 2C4G
	- 生产环境: 4C8G 起步（按并发和模块数量扩展）
- 磁盘: SSD，至少 100GB（含日志与备份空间）
- 时区: 统一 `Asia/Shanghai`
- Go 版本: 与 `go.mod` 保持一致
- 进程管理: `systemd`（推荐），禁止直接裸跑

建议开放端口（按实际环境调整）：

- Web 管理端: `11080`
- API 服务: `11072`
- Crypto 端口: `11073`（如启用）
- MySQL: `3306`（仅内网）
- Redis: `6379`（仅内网）

### 2. 数据库配置（MySQL）

推荐要求：

- 版本: MySQL 8.0+
- 字符集: `utf8mb4`
- 排序规则: `utf8mb4_general_ci`（或按现网统一规则）
- 账号权限: 按库最小权限，不使用 root 直连业务服务

初始化建议：

1. 创建业务库（例如 `hy_game`）。
2. 执行历史 SQL 与增量迁移（目录见 `sql/`）。
3. 校验关键表、索引、唯一约束是否齐全。
4. 为上线前准备全量备份与回滚点。

### 3. Redis 配置

推荐要求：

- 版本: Redis 6+
- 开启密码与内网访问控制
- 持久化策略按环境选择（AOF/RDB）
- 预留足够内存，避免频繁淘汰登录态与业务缓存

说明：

- 管理后台登录态和会话依赖 Redis。
- 冒烟中的 `finance`、`merchant` 依赖管理员会话可用。

### 4. 应用配置（conf.yaml / .env）

上线前必须确认：

- 所有密钥类配置均为非默认值：
	- `SESSION_AUTH_TOKEN`
	- `SHA256_SALT`
	- `API_SHA256_SALT`
	- `CRYPTO_AUTH_TOKEN`
- 跨域与回调白名单符合生产要求：
	- `ALLOWED_CORS_ORIGINS`
	- `PAYMENT_CALLBACK_IPS`
- 数据库、Redis、外部依赖地址指向生产内网。

建议：

- 生产环境使用独立配置文件，不要复用 `.env.example`。
- 密钥通过安全配置中心或 CI/CD Secret 注入，不要硬编码进仓库。

### 5. 上线步骤（标准流程）

1. 代码检查：确认目标分支与发布 tag。
2. 编译产物：

```bash
cd /path/to/qp-backend
go mod tidy
go build -o ./bin/web ./cmd/web/main.go
go build -o ./bin/api ./cmd/api/main.go
```

3. 执行数据库迁移（先备份后执行）。
4. 部署配置文件（生产 conf + env + 权限）。
5. 重启服务（建议灰度/分批）。
6. 执行冒烟：

```bash
API_BASE_URL=http://127.0.0.1:11072 \
WEB_BASE_URL=http://127.0.0.1:11080 \
ENV_FILE=./.env.example \
./scripts/test_suite.sh all
```

7. 验证核心链路（登录、充值、提现、回调、报表）。

### 6. systemd 示例（建议）

`api.service` 示例：

```ini
[Unit]
Description=QP Backend API
After=network.target

[Service]
Type=simple
WorkingDirectory=/data/qp/qp-backend
ExecStart=/data/qp/qp-backend/bin/api -config=/data/qp/qp-backend/conf.yaml -port=11072 -cryptoPort=11073
Restart=always
RestartSec=5
User=www-data
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
```

`web.service` 示例：

```ini
[Unit]
Description=QP Backend Web
After=network.target

[Service]
Type=simple
WorkingDirectory=/data/qp/qp-backend
ExecStart=/data/qp/qp-backend/bin/web -config=/data/qp/qp-backend/conf.yaml -port=11080
Restart=always
RestartSec=5
User=www-data
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
```

启停命令：

```bash
sudo systemctl daemon-reload
sudo systemctl enable api web
sudo systemctl restart api web
sudo systemctl status api web
```

### 7. 回滚策略

- 保留最近 N 个版本二进制与配置快照。
- 每次上线前先做 DB 备份。
- 回滚顺序建议：
	1. 停止新版本服务
	2. 切回上一版二进制与配置
	3. 必要时恢复数据库（按变更类型评估）
	4. 执行核心冒烟验证

### 8. 监控与告警建议

- 进程状态、端口存活
- API 5xx 比例、响应时延
- MySQL 连接数与慢查询
- Redis 内存与命中率
- 关键业务成功率（登录、支付回调、提现流转）

## 发布执行清单（T-1h / T+10m / T+30m）

以下清单可直接用于发布值班，建议按阶段打勾执行，不要跳步。

### T-1h（发布前 1 小时）

1. 冻结发布范围
- 确认发布 commit/tag。
- 确认本次涉及模块（API/Web/Cron/消费者/SQL）。

2. 环境与配置核对
- 对比 `conf.yaml` 与目标环境配置差异。
- 核对关键安全项：`SESSION_AUTH_TOKEN`、`SHA256_SALT`、`API_SHA256_SALT`、`CRYPTO_AUTH_TOKEN`。
- 核对回调白名单与跨域白名单。

3. 依赖健康检查
- MySQL 可连接、主从延迟可接受。
- Redis 可连接、内存余量充足。
- Kafka broker 可达，topic 正常。

4. 数据库准备
- 做全量备份并记录备份文件名。
- 预跑 SQL（先测试库），确认无语法/索引冲突。
- 明确回滚 SQL 或回滚策略。

5. 发布前冒烟（旧版本基线）

```bash
cd /path/to/qp-backend
./scripts/test_suite.sh release-env
./scripts/test_suite.sh security
```

放行标准：`release-env` 与 `security` 均通过。

### T0 ~ T+10m（发布窗口）

1. 部署顺序（建议）
- 先执行 SQL。
- 再发布二进制与配置。
- 按顺序重启：`web -> api -> cron -> consumidor -> betconsumidor`。

2. 启动后基础检查
- 检查 5 个进程均存活。
- 检查 API/Web 健康接口可访问。
- 检查日志无持续 panic/连接失败。

3. 快速功能验收

```bash
cd /path/to/qp-backend
API_BASE_URL=http://127.0.0.1:11072 WEB_BASE_URL=http://127.0.0.1:11080 ./scripts/test_suite.sh all
```

放行标准：`test_suite all` 通过，且无持续错误日志。

阻断条件（任一命中即暂停放量）：
- 资金相关接口出现连续 5xx。
- 回调验签异常持续出现。
- Kafka 消费明显堆积且持续增长。

### T+10m ~ T+30m（观察期）

1. 指标观察
- API 错误率、P95 延迟。
- 充值成功率、提现成功率、回调失败率。
- Kafka 消费 lag 是否回落。
- MySQL 慢查询与连接数是否异常。

2. 账务一致性抽检
- 抽检 3-5 笔充值单：订单状态、账变、报表一致。
- 抽检 3-5 笔提现单：审核、代付状态、账变一致。
- 抽检注单入库：`venue_code + order_sn` 无重复。

3. 值班结论
- 记录发布结果、异常与处理动作。
- 明确“继续观察”或“发布完成”。

### 回滚触发阈值（建议）

满足以下任一条件，建议立即回滚：

- 核心资金链路连续异常超过 5 分钟。
- 回调失败率明显高于基线且无法快速修复。
- 出现资金不一致且无法在短时间内确认影响范围。

回滚后必须执行：

1. 核心冒烟（`security`、`finance`、`merchant`）。
2. 资金链路抽检（充值/提现/回调）。
3. 记录事故时间线与根因初判。

## 三方游戏对接



### 1. 新增三方渠道时，通常要改哪些文件

API 侧（启动、转账、回调、查单）：

- `pkg/service/venues/venueDetail/base.go`：统一接口与请求/响应结构定义。
- `pkg/service/venues/deposit.go`：转入场馆逻辑（含状态流转、失败补偿）。
- `pkg/service/venues/withdraw.go`：从场馆转回逻辑。
- `cmd/api/controller/venueControl` 下相关控制器：登录场馆、余额、转账、恢复、记录查询。
- `cmd/api/controller/callbackControl` 下回调处理：三方回调入站与验签流程。

拉单与注单消费（跨项目，强相关）：

- `../qp-betrecord/service/venues/pull.go`：按场馆拉单入口与分片发送。
- `../qp-betrecord/controller/betrecord.go`：手动补单规则、场馆时间窗策略。
- `cmd/betconsumidor/controller/betRecordData.go`：注单消费、去重、并发锁、入库。

数据库与模型：

- `pkg/core/modules/dos/fcBetRecord.go`
- `pkg/core/modules/dos/fcBetRecordUnsettled.go`
- `sql/` 下增量迁移（例如唯一索引、新字段、兼容修复）

### 2. 对接改动顺序（推荐）

1. 在 `venueDetail` 增加/完善该渠道能力定义。
2. 完成场馆 SDK 封装（鉴权、签名、请求重试、错误码映射）。
3. 接入 API 控制器层（登录、余额、转账、回调确认）。
4. 打通拉单与消费链路（拉单 -> Kafka -> 入库）。
5. 增加 SQL 迁移与唯一约束，防止重复注单。
6. 执行全链路冒烟与回归（含并发、幂等、重复回调场景）。

### 3. 上线前必查清单（建议）

- 回调地址、白名单、签名密钥已按环境配置。
- 订单号幂等策略有效（重复回调不重复入账）。
- 注单去重策略有效（`venue_code + order_sn` 唯一）。
- 转账状态机完整（成功、失败、处理中、冲突重试）。
- 对账链路可用（三方账单、平台流水、用户资金一致）。

### 4. 文档维护建议

每新增一个三方渠道，建议在 README 追加：

- 场馆 code
- 接口文档链接/版本
- 改动文件列表
- 数据库变更脚本
- 验证结果与回滚要点

### 5. 单个场馆接入模板（建议复制使用）

以下模板建议每新增一个场馆都填写一份，直接追加在本节后。

```md
#### 场馆对接记录: <VENUE_CODE>

- 场馆名称: <名称>
- 对接负责人: <姓名>
- 联调时间: <YYYY-MM-DD>
- API 文档: <链接/版本>

1. 配置项
- 新增/修改配置: <conf.yaml / 环境变量键>
- 密钥来源: <配置中心/运维下发>
- 回调白名单: <IP 或域名>

2. 代码改动
- 统一接口定义: pkg/service/venues/venueDetail/base.go
- 场馆 SDK/实现: <具体文件>
- API 控制器: cmd/api/controller/venueControl/<具体文件>
- 回调处理: cmd/api/controller/callbackControl/<具体文件>
- 拉单与消费: ../qp-betrecord/service/venues/<具体文件>、../qp-betrecord/controller/<具体文件>
- 注单消费入库: cmd/betconsumidor/controller/betRecordData.go

3. 数据库变更
- SQL 文件: sql/<YYYY-MM-DD>.sql
- 变更内容: <表/索引/唯一约束>
- 回滚方案: <DROP INDEX / 回滚脚本>

4. 验证结果
- 登录场馆: 通过/失败
- 余额查询: 通过/失败
- 转入转出: 通过/失败
- 回调验签: 通过/失败
- 注单拉取与入库: 通过/失败
- 幂等与并发: 通过/失败

5. 上线与回滚
- 上线版本: <tag/commit>
- 上线窗口: <时间>
- 回滚触发条件: <错误率/资金异常阈值>
```

### 6. 三方对接常见漏项（强烈建议上线前复查）

- 只改 API 未改拉单消费，导致前台可玩但报表无注单。
- 只加代码未加唯一索引，导致重复回调/重复注单入账。
- 回调白名单或签名密钥未区分环境，导致生产回调失败。
- 只做单线程联调，未做并发和重放（replay）幂等验证。
- 未补文档，后续同场馆二次变更难以追溯。

### 7. 真实样板（PG 场馆示例）

#### 场馆对接记录: PGDZ

- 场馆名称: PG 电子
- 对接负责人: 待补充
- 联调时间: 2026-04-25
- API 文档: 供应商 PG API 文档（版本待补充）

1. 配置项
- 核心配置: PG 商户号、密钥、API 域名、回调地址
- 配置位置: 渠道配置文件 + 运行时环境变量
- 白名单: 支付/注单回调 IP 与域名白名单

2. 代码改动（示例）
- 统一能力定义: `pkg/service/venues/venueDetail/base.go`
- 注单拉取与调度: `../qp-betrecord/controller/betrecord.go`（`model.PGDZ` 手动拉单计划）
- 拉单实现入口: `../qp-betrecord/service/venues/pull.go`
- 注单消费与入库: `cmd/betconsumidor/controller/betRecordData.go`
- API 侧场馆能力编排: `pkg/service/venues/deposit.go`、`pkg/service/venues/withdraw.go`

3. 数据库变更（示例）
- SQL 文件: `sql/2026-04-23.sql`
- 变更内容: 为注单表增加幂等唯一索引
	- `fc_bet_record` 增加 `(venue_code, order_sn)` 唯一约束
	- `fc_bet_record_unsettled` 增加 `(venue_code, order_sn)` 唯一约束

4. 验证结果（示例）
- 登录场馆: 通过
- 余额查询: 通过
- 转入转出: 通过
- 注单拉取与入库: 通过
- 幂等与并发: 通过（重复订单不重复入库）

5. 回滚要点（示例）
- 应用回滚: 切回上一版本二进制与配置
- 数据回滚: 评估后再处理索引变更（避免直接破坏历史幂等保障）
- 验证: 回滚后执行核心冒烟与注单链路核验

## 充值/提现通道接入（建议写进 README）


### 1. 新增充值/提现通道时，通常要改哪些文件

对外路由与入口：

- `cmd/api/router/wallet.go`：钱包、充值、提现、绑卡等前台路由。
- `cmd/api/router/callback.go`：`/api/callback/pay/:paymentType` 与 `/api/callback/payOut/:paymentType` 回调入口。

前台业务控制器：

- `cmd/api/controller/walletControl/recharge.go`：充值通道查询、下单、订单状态查询。
- `cmd/api/controller/walletControl/withdraw.go`：提现申请、提现通道与代付通道查询。
- `cmd/api/controller/walletControl/bind.go`：银行卡/虚拟币/在线账号绑定、通道图片资源。

回调与资金状态流转：

- `cmd/api/controller/callbackControl/pay.go`：充值回调与代付回调处理。
- `pkg/service/userTransfer/deposit.go`：充值成功后的账变入账、任务/活动联动。
- `pkg/service/userTransfer/withdraw.go`：提现失败/退回等账变处理。
- `pkg/service/paymentOut/withdraw.go`：代付成功/失败状态推进与补偿。

通道适配层（充值/代付）：

- `pkg/service/payment/payment.go`：充值通道工厂与统一接口。
- `pkg/service/payment/*.go`：具体充值通道实现（如本地银行、本地虚拟币、三方通道）。
- `pkg/service/paymentOut/payment.go`：代付通道工厂与统一接口。
- `pkg/service/paymentOut/*.go`：具体代付实现（银行卡/虚拟币/支付宝等）。

后台配置与审核：

- `cmd/web/router/fcPayChannel.go`、`cmd/web/router/fcPayment.go`、`cmd/web/router/fcPaymentSetting.go`：充值渠道与充值通道管理。
- `cmd/web/router/fcPayChannelOut.go`、`cmd/web/router/fcPaymentOut.go`：提现渠道与代付通道管理。
- `cmd/web/router/fcOrderWithdraw.go`、`cmd/web/router/fcOrderWithdrawPaymentOut.go`：提现审核、批量代付、打款状态处理。
- `cmd/web/router/fcChannelBankImg.go`：代付相关图标/素材配置。

定时任务与异步处理：

- `cmd/cron/crontab/withdrawPaymentOut.go`：代付队列轮询与下发三方。

数据库与模型：

- `pkg/core/modules/dos/fcPayChannel.go`、`pkg/core/modules/dos/fcPayment.go`
- `pkg/core/modules/dos/fcPayChannelOut.go`、`pkg/core/modules/dos/fcPaymentOut.go`
- `pkg/core/modules/dos/fcOrderDeposit.go`、`pkg/core/modules/dos/fcOrderWithdraw.go`、`pkg/core/modules/dos/fcOrderWithdrawPaymentOut.go`
- `sql/*.sql` 与 `sql/updated/*.sql`：新增字段、索引、状态位、兼容迁移脚本。

### 2. 接入顺序（推荐）

1. 在后台先建渠道与通道（充值 `fcPayChannel/fcPayment`，提现 `fcPayChannelOut/fcPaymentOut`），确认币种、等级、限额、手续费、排序。
2. 实现通道适配器（充值放 `pkg/service/payment`，代付放 `pkg/service/paymentOut`），统一签名、请求、错误码映射。
3. 打通前台下单与通道查询（`walletControl/recharge.go`、`walletControl/withdraw.go`）。
4. 打通回调入口与验签逻辑（`callbackControl/pay.go`），保证回调可重放且幂等。
5. 打通提现审核与代付链路（`fcOrderWithdraw` + `fcOrderWithdrawPaymentOut` + `withdrawPaymentOut` 定时任务）。
6. 增加数据库迁移与必要索引，然后执行全链路联调。

### 3. 上线前必查清单（充值/提现）

- 回调地址已区分环境：测试/生产不共用域名与密钥。
- 充值回调幂等：同一 `order_sn` 重复通知不会重复入账。
- 提现代付幂等：同一代付单重复回调不会重复推进状态。
- 状态机完整：待支付/待处理/成功/失败路径都可闭环。
- 代付任务可恢复：`Prepare -> Progress -> Success/Failed` 可追踪并可补偿。
- 手续费计算与入账口径一致（订单金额、到账金额、手续费、报表口径一致）。
- 后台权限可控：通道管理、审核、批量操作有权限隔离。
- 对账可落地：三方账单、平台订单、账变流水三方一致。

### 4. 新通道接入模板（建议复制）

```md
#### 充值/提现通道接入记录: <CHANNEL_CODE>

- 通道名称: <名称>
- 商户/项目: <merchant_code>
- 对接负责人: <姓名>
- 联调时间: <YYYY-MM-DD>
- 三方文档: <链接/版本>

1. 配置项
- 回调地址(充值): /api/callback/pay/<paymentType>
- 回调地址(代付): /api/callback/payOut/<paymentType>
- 三方商户号/密钥: <配置键>
- 白名单/IP: <要求>

2. 代码改动
- 路由: cmd/api/router/wallet.go, cmd/api/router/callback.go
- 充值: cmd/api/controller/walletControl/recharge.go
- 提现: cmd/api/controller/walletControl/withdraw.go
- 回调: cmd/api/controller/callbackControl/pay.go
- 充值适配器: pkg/service/payment/<channel>.go
- 代付适配器: pkg/service/paymentOut/<channel>.go
- 账变: pkg/service/userTransfer/deposit.go, pkg/service/userTransfer/withdraw.go
- 代付状态推进: pkg/service/paymentOut/withdraw.go
- 定时任务: cmd/cron/crontab/withdrawPaymentOut.go

3. 数据库变更
- SQL: sql/<YYYY-MM-DD>.sql
- 变更项: <字段/索引/状态>
- 回滚方案: <回滚脚本>

4. 验证结果
- 充值下单: 通过/失败
- 充值回调入账: 通过/失败
- 提现申请: 通过/失败
- 代付下发: 通过/失败
- 代付回调: 通过/失败
- 幂等重放: 通过/失败
- 对账核验: 通过/失败

5. 上线与回滚
- 上线版本: <tag/commit>
- 观察指标: 充值成功率/提现成功率/回调失败率
- 回滚触发条件: <阈值>
```

### 5. 常见漏项（上线前再看一遍）

- 只接了充值回调，遗漏提现回调与代付状态推进。
- 只改了 API，未配置后台通道（前台看得到但不可用）。
- 回调签名通过但未做订单金额校验/商户校验。
- 回调成功后未做事务性账变，导致订单状态与余额不一致。
- 忽略并发重放，导致重复入账或重复出款。
- 未补 SQL 与索引，线上才发现唯一约束缺失。

## 短信通道接入（建议写进 README）


### 1. 关键文件清单

配置与SDK：

- `common/conf/conf.go`：`Sms` 配置结构（`AccessKeyId/AccessKeySecret/Endpoint/SignName/TemplateCode`）。
- `common/sms/smsapi.go`：阿里云短信发送实现（`sms.Handle`）。
- `conf.yaml` 及 `config/hy-*/conf.yaml`：各服务环境的 `Sms` 配置块。

前台API链路：

- `cmd/api/router/userInfo.go`：
	- `POST /api/userinfo/phone/veryCode`（发送验证码）
	- `POST /api/userinfo/phone/veryCodeSub`（提交验证码）
- `cmd/api/controller/userControl/userinfo.go`：验证码发送、Redis 限流与校验逻辑。

后台管理链路：

- `cmd/web/router/fcSmsChannel.go`：短信通道后台 CRUD 路由。
- `pkg/core/modules/dos/fcSmsChannel.go`：短信通道模型（`fc_sms_channel`）。
- `pkg/core/modules/fcSmsChannel.go`：短信通道查询与维护。

### 2. 接入顺序（推荐）

1. 配置短信厂商参数（AK/SK、签名、模板、Endpoint），按环境隔离。
2. 后台创建并启用短信通道（`fcSmsChannel`），设置等级范围、排序、状态。
3. 打通发送接口（`phone/veryCode`）与验证码提交接口（`phone/veryCodeSub`）。
4. 验证 Redis 过期策略：验证码 5 分钟、发送锁 2 分钟。
5. 联调注册/绑卡/找回密码等依赖验证码的业务入口。

### 3. 上线前必查清单（短信）

- 生产环境 `Sms` 配置不为空，且非测试模板。
- 短信发送失败会返回明确错误，不吞异常。
- 发送频控生效（同手机号短时间不可重复发送）。
- 验证码校验只在有效期内通过，过期后拒绝。
- 可通过字典配置 `SmsVerification` 开关控制是否启用短信校验（联调环境可关闭）。

### 4. 短信接入记录模板

```md
#### 短信通道接入记录: <SMS_CODE>

- 供应商: <阿里云/其他>
- 对接负责人: <姓名>
- 联调时间: <YYYY-MM-DD>

1. 配置
- conf键: Sms.AccessKeyId / Sms.AccessKeySecret / Sms.SignName / Sms.TemplateCode
- 环境: <dev/test/prod>

2. 代码改动
- 发送实现: common/sms/smsapi.go
- 发送入口: cmd/api/controller/userControl/userinfo.go (PhoneVeryCode)
- 校验入口: cmd/api/controller/userControl/userinfo.go (Verification)
- 后台通道: cmd/web/router/fcSmsChannel.go

3. 验证
- 发送成功: 通过/失败
- 频控限制: 通过/失败
- 错误码/文案: 通过/失败
```

## 客服系统对接（建议写进 README）

### 1. 关键文件清单

前台客服查询接口：

- `cmd/api/router/customer.go`：`GET /api/customer/get`。
- `cmd/api/controller/customerControl/customer.go`：按 `merchant_code` 返回站点客服链接（facebook/whatsapp/twitter/telegram/skype）。
- `cmd/api/router/user.go`：`GET /api/customer/link`。
- `cmd/api/controller/userControl/user.go`：客服链接与推广域名信息返回。

后台配置与维护：

- `cmd/web/router/fcSiteLink.go`：客服社媒站点链接维护（`fc_site_link`）。
- `cmd/web/router/fcCustomerLink.go`：统一客服链接维护（`fc_customer_link`）。
- `cmd/web/router/fcCustomerOrder.go`：客服工单管理（`fc_customer_order`）。
- `cmd/web/router/fcCustomerOrderType.go`：工单类型管理。
- `cmd/web/router/fcAgentDomain.go`：域名配置，包含客服链接同步相关。

模型与同步逻辑：

- `pkg/core/modules/dos/fcSiteLink.go`、`pkg/core/modules/dos/fcCustomerLink.go`。
- `pkg/core/modules/fcAgentDomain.go`：`SyncFcAgentDomainCustomerLink` 同步客服链接到推广域名。

数据库迁移：

- `sql/2025-05-30.sql` 与 `sql/updated/prod20250628.sql`：`fc_agent_domain` 增加 `customer_link`。
- 需确保存在 `fc_customer_link`、`fc_site_link`、`fc_customer_order`、`fc_customer_order_type` 相关表。

### 2. 接入顺序（推荐）

1. 先在后台配置 `fcCustomerLink`（统一客服链接）与 `fcSiteLink`（社媒客服入口）。
2. 配置并校验推广域名，触发/执行客服链接同步逻辑。
3. 验证前台接口：`/api/customer/get` 和 `/api/customer/link` 返回正确商户数据。
4. 配置客服工单类型后，联调工单提交与状态流转。
5. 在 H5/LP/APP 各端确认客服入口展示一致。

### 3. 上线前必查清单（客服）

- 商户维度客服链接已配置，不为空。
- 社媒客服 key 完整：`facebookCustomer/whatsappCustomer/twitterCustomer/telegramCustomer/skypeCustomer`。
- `fc_customer_link` 表已创建，否则域名同步会报错。
- 前台客服接口按商户隔离返回，不串商户数据。
- 工单状态流转与后台权限可用（查询、更新状态、类型管理）。

### 4. 客服对接记录模板

```md
#### 客服系统对接记录: <MERCHANT_CODE>

- 商户: <merchant_code>
- 负责人: <姓名>
- 联调时间: <YYYY-MM-DD>

1. 配置
- 统一客服链接: fc_customer_link
- 社媒客服链接: fc_site_link(app_key)
- 推广域名客服字段: fc_agent_domain.customer_link

2. 代码链路
- 前台查询: cmd/api/controller/customerControl/customer.go
- 客服link查询: cmd/api/controller/userControl/user.go
- 同步逻辑: pkg/core/modules/fcAgentDomain.go
- 后台维护: cmd/web/router/fcCustomerLink.go, cmd/web/router/fcSiteLink.go

3. 验证
- /api/customer/get: 通过/失败
- /api/customer/link: 通过/失败
- 域名同步: 通过/失败
- 工单流程: 通过/失败
```

## 统一测试入口

项目仅保留一个测试脚本入口：

```bash
./scripts/test_suite.sh <command>
```

支持命令：

- `security`: 安全冒烟（CORS / 响应头 / CSRF / 回调 CSRF 豁免）
- `release-env`: 发布前安全环境变量检查
- `finance`: 财务冒烟批次
- `merchant`: 商户模块冒烟
- `all`: 顺序执行 `release-env + security + finance + merchant`

示例：

```bash
# 安全冒烟
./scripts/test_suite.sh security

# 发布前环境校验
ENV_FILE=./.env.example ./scripts/test_suite.sh release-env

# 全量冒烟
./scripts/test_suite.sh all
```

前置条件：

- `release-env` 默认读取 `qp-backend/.env`，建议先准备有效环境变量（可用 `ENV_FILE=./.env.example` 做快速校验）。
- `security` 依赖 `API_BASE_URL` 和 `WEB_BASE_URL` 可访问，且服务安全策略应与检查项一致。
- `finance` 和 `merchant` 依赖 Redis 中存在管理员登录态（先从后台完成一次管理员登录）。

常用环境变量：

- `API_BASE_URL`（默认 `http://127.0.0.1:10072`）
- `WEB_BASE_URL`（默认 `http://127.0.0.1:10080`）
- `ALLOWED_ORIGIN`（默认 `http://localhost:5173`）
- `BLOCKED_ORIGIN`（默认 `http://malicious.invalid`）
- `ENV_FILE`（默认 `qp-backend/.env`）

注意：请始终在 `qp-backend` 根目录执行命令，避免 Go module 相对路径错误。