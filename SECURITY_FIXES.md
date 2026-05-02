# QP 项目安全修复实施指南

## 概述
本文档说明如何集成已实施的安全修复到应用中。这些修复解决了 P0/P1 级别的安全漏洞。

## 已完成的修复

### P0 级别 - 关键漏洞（必须修复）

#### 1. Cookie 安全加固 ✅
**文件**: `cmd/web/handler/middleware.go`, `cmd/api/handler/middleware.go`

**修复**:
- 添加 `HttpOnly=true`: 防止 JavaScript XSS 攻击窃取
- 添加 `Secure=true`: 仅限 HTTPS 传输，防止中间人攻击

**验证方法**:
```bash
# 检查 cookie 标志
curl -i http://localhost:8080/api/endpoint
# 查看 Set-Cookie 响应头中是否包含 HttpOnly; Secure;
```

---

#### 2. CORS 过度宽松修复 ✅
**文件**: `cmd/api/handler/middleware.go`, `cmd/web/handler/middleware.go`

**修复前**:
```go
c.Header("Access-Control-Allow-Origin", "*")         // 危险！允许任何域名
c.Header("Access-Control-Allow-Headers", "*")        // 危险！允许任何请求头
```

**修复后**:
```go
allowedOrigins := map[string]bool{
    "http://localhost:3000": true,
    "https://app.example.com": true,
}
if allowedOrigins[origin] {
    c.Header("Access-Control-Allow-Origin", origin)
    c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-CSRF-Token")
}
```

**集成步骤**:
1. 在 `.env` 文件中配置允许的前端域名
2. 在 middleware 中从环境变量读取

---

#### 3. CSRF 防护机制 ✅
**文件**: `common/middleware/csrf.go`

**功能**:
- 为每个用户会话生成唯一的 CSRF token
- GET/HEAD/OPTIONS 请求时自动生成 token
- POST/PUT/DELETE 请求时验证 token

**集成到项目**:
1. 在路由中添加 CSRF 中间件
2. 前端在请求时同时发送 token

**后端集成示例**:
```go
// 主路由
api := router.Group("/api")
api.Use(middleware.CSRFProtection())  // 在认证之后添加

// 或仅在特定路由添加
router.POST("/api/transfer", middleware.CSRFProtection(), controller.Transfer)
```

**前端集成示例**:
```javascript
// 获取 CSRF token
async function getCsrfToken() {
    const response = await fetch('/api/csrf-token', { method: 'GET' });
    const data = await response.json();
    return data.csrf_token;
}

// 发送 POST 请求时提交 token
const token = await getCsrfToken();
fetch('/api/transfer', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'X-CSRF-Token': token,  // 在请求头中提交
    },
    body: JSON.stringify(payload),
    credentials: 'include'  // 包含 cookie
});
```

---

#### 4. 硬编码密钥迁移至环境变量 ✅
**文件**: `common/tool/secrets.go`, `cmd/web/handler/middleware.go`

**修复**:
- 将硬编码的 `sessionIdAuthToken` 移至环境变量
- 支持从 `.env` 文件或系统环境变量读取

**集成步骤**:

1. **创建 `.env` 文件** (复制 `.env.example`):
```bash
cp .env.example .env
```

2. **编辑 `.env`，填入强随机密钥**:
```bash
# 使用 openssl 生成强密钥
openssl rand -hex 32 > session_token.txt
openssl rand -hex 32 > sha256_salt.txt
```

3. **在 main.go 中初始化**:
```go
import "bootpkg/common/tool"

func init() {
    tool.InitSecretManager()
}
```

4. **验证**:
```go
// 获取密钥（会从环境变量读取）
secrets := tool.GetGlobalSecrets()
fmt.Println(secrets.SessionAuthToken) // 应该显示从 .env 读取的值
```

**部署步骤**:
```bash
# 生产环境：设置环境变量（不使用.env文件）
export SESSION_AUTH_TOKEN="your-strong-random-token"
export SHA256_SALT="your-sha256-salt"
export API_SHA256_SALT="your-api-salt"
export CRYPTO_AUTH_TOKEN="your-crypto-token"

# 启动应用
./app
```

---

#### 5. 支付回调验证加强 ✅
**文件**: `pkg/service/payment/validator.go`

**功能**:
- IP 白名单验证（支持 CIDR 段）
- 金额匹配检查
- 商户代码验证
- 重复回调检测

**集成示例**:
```go
import "bootpkg/pkg/service/payment"

// 在回调处理函数中调用
validator := payment.NewPaymentCallbackValidator()

// 配置支付渠道 IP 白名单
validator.AllowedIPs = map[string][]string{
    "alipay": {"220.248.137.0/24", "110.84.101.0/24"},
    "wechat": {"203.119.29.0/24"},
}

// 验证回调
err := validator.ComprehensiveCallbackValidation(
    orderInfo,           // 订单信息
    "alipay",            // 支付渠道
    callbackAmount,      // 回调金额
    callbackMerchantCode,// 商户代码
    c.ClientIP(),        // 客户端 IP
    0.01,                // 允许的金额差异
)

if err != nil {
    log.Printf("Callback validation failed: %v", err)
    c.String(200, "fail")
    return
}

// 继续处理回调
```

---

### P1 级别 - 高危问题

#### 6. 密码算法升级至 bcrypt ✅
**文件**: `common/tool/password.go`

**特性**:
- 支持 bcrypt 密码加密（推荐）
- 向后兼容旧的 SHA256 密码
- 自动提示需要升级的密码

**集成步骤**:

1. **在应用启动时初始化**:
```go
import "bootpkg/common/tool"

func main() {
    // 从配置读取 SHA256 盐值用于向后兼容
    tool.InitPasswordHasher(global.CONFIG.General.ApiSHA256Salt)
    
    // ... 启动应用
}
```

2. **注册新用户时使用 bcrypt**:
```go
func RegisterUser(password string) error {
    hasher := tool.GetGlobalPasswordHasher()
    
    // 使用 bcrypt 加密新密码
    hash, err := hasher.HashPassword(password)
    if err != nil {
        return err
    }
    
    // 存储到数据库
    user.Password = hash
    return db.Save(&user).Error
}
```

3. **登录时支持密码升级**:
```go
func Login(username, password string) (*User, error) {
    user := &User{}
    db.Where("username = ?", username).First(user)
    
    hasher := tool.GetGlobalPasswordHasher()
    matches, needsUpgrade := hasher.VerifyPassword(password, user.Password)
    
    if !matches {
        return nil, errors.New("invalid password")
    }
    
    // 如果需要升级，在后台升级密码
    if needsUpgrade {
        newHash, err := hasher.UpgradePasswordToBcrypt(user.Password, password)
        if err == nil && newHash != "" {
            db.Model(user).Update("password", newHash)
            log.Printf("Password upgraded to bcrypt for user: %s", username)
        }
    }
    
    return user, nil
}
```

4. **数据迁移** (已有用户):
```bash
# 登录时自动升级密码，无需额外工作
# 旧密码继续有效，但不推荐
# 建议在修改密码时强制使用 bcrypt
```

---

#### 7. 安全响应头补充 ✅
**文件**: `cmd/api/handler/middleware.go`, `cmd/web/handler/middleware.go`

**已添加的头**:
- `X-Content-Type-Options: nosniff` - 防止 MIME 错判
- `X-Frame-Options: DENY` - 防止点击劫持
- `X-XSS-Protection: 1; mode=block` - 浏览器 XSS 保护
- `Strict-Transport-Security` - HTTPS 强制

**验证方法**:
```bash
curl -i http://localhost:8080/api/endpoint | grep -E "X-|Strict"
```

---

## 环境变量配置

### 开发环境 (.env)
```
SESSION_AUTH_TOKEN=development-token-change-in-production
SHA256_SALT=dev-salt-value
API_SHA256_SALT=dev-api-salt
CRYPTO_AUTH_TOKEN=dev-crypto-token
ALLOWED_CORS_ORIGINS=http://localhost:3000,http://localhost:8080
PAYMENT_CALLBACK_IPS=127.0.0.1
```

### 生产环境

绝对不要在代码中硬编码密钥！使用以下方式之一：

**方式1: 系统环境变量**
```bash
export SESSION_AUTH_TOKEN="$(openssl rand -hex 32)"
export SHA256_SALT="$(openssl rand -hex 32)"
export API_SHA256_SALT="$(openssl rand -hex 32)"
export CRYPTO_AUTH_TOKEN="$(openssl rand -hex 32)"
```

**方式2: Docker Secrets (Swarm)**
```yaml
services:
  api:
    environment:
      - SESSION_AUTH_TOKEN_FILE=/run/secrets/session_token
    secrets:
      - session_token
secrets:
  session_token:
    external: true
```

**方式3: Kubernetes Secrets**
```yaml
env:
  - name: SESSION_AUTH_TOKEN
    valueFrom:
      secretKeyRef:
        name: qp-secrets
        key: session-token
```

---

## 测试验证

### 安全测试清单

- [ ] Cookie 设置了 HttpOnly 和 Secure 标志
- [ ] CORS 只允许配置的前端域名
- [ ] CSRF 令牌正确生成和验证
- [ ] 支付回调验证了 IP、金额、商户代码
- [ ] 环境变量中的密钥在生产环境生效
- [ ] 新注册用户使用 bcrypt 密码
- [ ] 旧用户登录时密码自动升级到 bcrypt

### 测试命令

```bash
# 测试 Cookie 属性
curl -i http://localhost:8080/login | grep Set-Cookie

# 测试 CORS 限制
curl -H "Origin: http://malicious.com" http://localhost:8080/api/data

# 测试 CSRF 保护
curl -X POST http://localhost:8080/api/transfer -H "X-CSRF-Token: wrong" 

# 测试密码哈希
go test ./common/tool -v
```

---

## 后续建议

### 立即实施
1. 配置环境变量（.env 或系统变量）
2. 在应用启动时初始化安全模块
3. 在路由中启用 CSRF 中间件
4. 从生产配置文件中移除硬编码密钥

### 中期计划
1. 对现有密码进行审计和升级计划
2. 实施密码强度验证
3. 添加账户锁定机制（多次错误登录）
4. 实施 API 速率限制

### 长期计划
1. 考虑 OAuth2 / OpenID Connect 集成
2. 实施多因素认证 (MFA)
3. 定期安全审计和渗透测试
4. 完整的日志审计系统

---

## 故障排除

### 问题: CSRF token 不匹配
**原因**: Cookie 未正确设置或前端未提交 token
**解决**: 
1. 检查浏览器 DevTools 中的 Cookie
2. 确保前端在请求头中包含 token
3. 检查跨域请求是否正确设置 credentials

### 问题: 环境变量未读取
**原因**: 应用启动时未调用初始化函数
**解决**: 在 main.go 中添加 `tool.InitSecretManager()`

### 问题: 密码升级失败
**原因**: 旧密码格式不匹配或盐值配置错误
**解决**: 检查配置中的 SHA256 盐值是否正确

---

## 参考资源

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [bcrypt hashing](https://en.wikipedia.org/wiki/Bcrypt)
- [CSRF Prevention Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html)
- [Secure Coding Guidelines](https://www.securecoding.cert.org/)
