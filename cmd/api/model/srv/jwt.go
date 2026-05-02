package srv

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	memberTokenTTL      = 3 * 24 * time.Hour
	memberInfoRenewalTT = 3 * time.Hour
)

// 单点登录生成Token
func GenTokenJWTToken(userNameM string) string {
	var token = strings.Replace(uuid.NewString(), "-", "", -1)
	signUserName, _ := tool.EncryptAES([]byte(userNameM), []byte(token[:16]))
	iss := base64.StdEncoding.EncodeToString(signUserName)
	now := time.Now()
	jwtSub := &jwt.RegisteredClaims{
		ID:     token,
		Issuer: iss,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(memberTokenTTL)),
	}
	jwtToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtSub).SignedString([]byte(global.CONFIG.SHA256Salt))
	tokenRedisKey := fmt.Sprintf(enmus.REDIS_MEMBER_LOGIN_TOKEN, userNameM)
	global.G_REDIS.Set(context.Background(), tokenRedisKey, token, memberTokenTTL)

	return jwtToken
}

// VerifyJWTToken
//
//	@Description: 验证Token
//	@param tokenStr Token字符串
//	@return bool 是否正确
//	@return string userName
func VerifyJWTToken(tokenStr string) (bool, string) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.SHA256Salt), nil
	})
	if err != nil || !token.Valid {
		return false, ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, ""
	}
	id, ok := claims["jti"].(string)
	if !ok || len(id) < 16 {
		return false, ""
	}
	iss, ok := claims["iss"].(string)
	if !ok || iss == "" {
		return false, ""
	}
	issBytes, err := base64.StdEncoding.DecodeString(iss)
	if err != nil {
		return false, ""
	}
	signUserName, err := tool.DecryptAES(issBytes, []byte(id[:16]))
	if err != nil {
		return false, ""
	}
	userNameM := string(signUserName)
	tokenRedisKey := fmt.Sprintf(enmus.REDIS_MEMBER_LOGIN_TOKEN, userNameM)
	if global.G_REDIS.Get(context.Background(), tokenRedisKey).Val() != id {
		return false, ""
	}
	if global.G_REDIS.TTL(context.Background(), tokenRedisKey).Val().Seconds() < 3600 {
		memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
		global.G_REDIS.Expire(context.Background(), memberRedisKey, memberInfoRenewalTT)
	}
	return true, userNameM
}
