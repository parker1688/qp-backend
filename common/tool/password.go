package tool

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher 密码加密器 - 支持 bcrypt 和向后兼容 SHA256
type PasswordHasher struct {
	bcryptCost int    // bcrypt 计算成本因子 (10-12 recommended)
	sha256Salt string // SHA256 盐值（用于兼容旧密码）
}

// NewPasswordHasher 创建密码加密器实例
func NewPasswordHasher(sha256Salt string) *PasswordHasher {
	return &PasswordHasher{
		bcryptCost: 12,
		sha256Salt: sha256Salt,
	}
}

// HashPassword 使用 bcrypt 加密密码 - 推荐用于新密码
func (ph *PasswordHasher) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), ph.bcryptCost)
	if err != nil {
		log.Printf("[PasswordHasher] Failed to hash password: %v", err)
		return "", err
	}
	return string(hash), nil
}

// VerifyPassword 验证密码 - 首先尝试 bcrypt，降级到 SHA256（向后兼容）
func (ph *PasswordHasher) VerifyPassword(password string, hash string) (bool, bool) {
	// 第一次返回值：是否匹配
	// 第二次返回值：是否需要升级（true = 旧的 SHA256，需要升级到 bcrypt）

	// 首先尝试 bcrypt
	if strings.HasPrefix(hash, "$2a$") || strings.HasPrefix(hash, "$2b$") || strings.HasPrefix(hash, "$2y$") {
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		if err != nil {
			return false, false
		}
		return true, false
	}

	// 向后兼容：尝试 SHA256 方式
	sha256Hash := HashSHA256WithSalt(password, ph.sha256Salt)
	if sha256Hash == hash {
		return true, true
	}

	return false, false
}

// ShouldUpgradePassword 判断密码是否应该升级
func (ph *PasswordHasher) ShouldUpgradePassword(hash string) bool {
	return !(strings.HasPrefix(hash, "$2a$") || strings.HasPrefix(hash, "$2b$") || strings.HasPrefix(hash, "$2y$"))
}

// HashSHA256WithSalt 使用 SHA256 + 盐值加密（兼容旧系统）
func HashSHA256WithSalt(password string, salt string) string {
	return Sha256Encryption(password + salt)
}

// UpgradePasswordToBcrypt 将 SHA256 密码升级到 bcrypt
func (ph *PasswordHasher) UpgradePasswordToBcrypt(oldHash string, password string) (string, error) {
	verifyOk, isLegacy := ph.VerifyPassword(password, oldHash)
	if !verifyOk || !isLegacy {
		log.Printf("[PasswordHasher] Cannot upgrade password - invalid or already bcrypt")
		return "", nil
	}
	newHash, err := ph.HashPassword(password)
	if err != nil {
		return "", err
	}
	return newHash, nil
}

var globalPasswordHasher *PasswordHasher

// InitPasswordHasher 初始化全局密码加密器
func InitPasswordHasher(sha256Salt string) {
	globalPasswordHasher = NewPasswordHasher(sha256Salt)
}

// GetGlobalPasswordHasher 获取全局密码加密器
func GetGlobalPasswordHasher() *PasswordHasher {
	if globalPasswordHasher == nil {
		log.Println("[PasswordHasher] Warning: not initialized, using default")
		globalPasswordHasher = NewPasswordHasher("")
	}
	return globalPasswordHasher
}
