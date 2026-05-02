package tool

import (
	"fmt"
	"os"
	"strings"
)

// SecretManager 密钥管理器
type SecretManager struct {
	SessionAuthToken string // Web session 认证令牌
	SHA256Salt       string // SHA256 盐值
	ApiSHA256Salt    string // API 盐值
	CryptoAuthToken  string // 加密认证令牌
}

// GetSecretManager 获取密钥管理器实例，从环境变量读取
func GetSecretManager() *SecretManager {
	sm := &SecretManager{
		SessionAuthToken: os.Getenv("SESSION_AUTH_TOKEN"),
		SHA256Salt:       os.Getenv("SHA256_SALT"),
		ApiSHA256Salt:    os.Getenv("API_SHA256_SALT"),
		CryptoAuthToken:  os.Getenv("CRYPTO_AUTH_TOKEN"),
	}

	// 如果环境变量未设置，使用配置文件中的默认值（仅用于向后兼容）
	// 生产环境应始终从环境变量设置
	if sm.SessionAuthToken == "" {
		sm.SessionAuthToken = "DEFAULT_SESSION_TOKEN_CHANGE_IN_PRODUCTION"
	}
	if sm.SHA256Salt == "" {
		sm.SHA256Salt = "DEFAULT_SHA256_SALT_CHANGE_IN_PRODUCTION"
	}
	if sm.ApiSHA256Salt == "" {
		sm.ApiSHA256Salt = "DEFAULT_API_SHA256_SALT_CHANGE_IN_PRODUCTION"
	}
	if sm.CryptoAuthToken == "" {
		sm.CryptoAuthToken = "DEFAULT_CRYPTO_TOKEN_CHANGE_IN_PRODUCTION"
	}

	return sm
}

// 全局单例
var globalSecretManager *SecretManager

// InitSecretManager 初始化全局密钥管理器
func InitSecretManager() {
	globalSecretManager = GetSecretManager()
}

// GetGlobalSecrets 获取全局密钥管理器
func GetGlobalSecrets() *SecretManager {
	if globalSecretManager == nil {
		InitSecretManager()
	}
	return globalSecretManager
}

func parseCSVEnv(key string) []string {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	items := make([]string, 0, len(parts))
	for _, part := range parts {
		v := strings.TrimSpace(part)
		if v != "" {
			items = append(items, v)
		}
	}
	return items
}

// GetAllowedCORSOrigins 获取允许的 CORS 域名列表
func GetAllowedCORSOrigins() []string {
	return parseCSVEnv("ALLOWED_CORS_ORIGINS")
}

// GetPaymentCallbackIPs 获取支付回调 IP 白名单
func GetPaymentCallbackIPs() []string {
	return parseCSVEnv("PAYMENT_CALLBACK_IPS")
}

// ValidateSecurityEnv 校验关键安全配置。
// strict=true 时用于生产环境，缺失关键配置会返回错误。
func ValidateSecurityEnv(strict bool) error {
	secrets := GetGlobalSecrets()

	if strings.TrimSpace(secrets.SessionAuthToken) == "" || strings.HasPrefix(secrets.SessionAuthToken, "DEFAULT_") {
		if strict {
			return fmt.Errorf("SESSION_AUTH_TOKEN is missing or using default value")
		}
	}
	if strings.TrimSpace(secrets.SHA256Salt) == "" || strings.HasPrefix(secrets.SHA256Salt, "DEFAULT_") {
		if strict {
			return fmt.Errorf("SHA256_SALT is missing or using default value")
		}
	}
	if strings.TrimSpace(secrets.ApiSHA256Salt) == "" || strings.HasPrefix(secrets.ApiSHA256Salt, "DEFAULT_") {
		if strict {
			return fmt.Errorf("API_SHA256_SALT is missing or using default value")
		}
	}

	if strict && len(GetAllowedCORSOrigins()) == 0 {
		return fmt.Errorf("ALLOWED_CORS_ORIGINS is required in release mode")
	}
	if strict && len(GetPaymentCallbackIPs()) == 0 {
		return fmt.Errorf("PAYMENT_CALLBACK_IPS is required in release mode")
	}

	return nil
}
