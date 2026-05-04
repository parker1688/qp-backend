package middleware

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	CSRFTokenHeaderKey  = "X-CSRF-Token"
	CSRFTokenCookieName = "CSRF-Token"
	CSRFTokenLength     = 32
)

var csrfSkipPathPrefixes = []string{
	"/api/callback/",
	"/api/updateProductStatus",
	"/api/queryProductList",
	"/api/login",
	"/api/user/login",
	"/api/base/nologin/login",
}

var csrfBypassOrigins = map[string]struct{}{
	"http://localhost:5173": {},
	"http://127.0.0.1:5173": {},
	"http://localhost:5174": {},
	"http://127.0.0.1:5174": {},
}

var csrfBypassRefererPrefixes = []string{
	"http://localhost:5173/",
	"http://127.0.0.1:5173/",
	"http://localhost:5174/",
	"http://127.0.0.1:5174/",
}

// GenerateCSRFToken 生成 CSRF token
func GenerateCSRFToken() (string, error) {
	bytes := make([]byte, CSRFTokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// SetCSRFTokenCookie 在响应中设置 CSRF token cookie
func SetCSRFTokenCookie(c *gin.Context) error {
	if token, err := c.Cookie(CSRFTokenCookieName); err == nil && token != "" {
		c.Header(CSRFTokenHeaderKey, token)
		return nil
	}

	token, err := GenerateCSRFToken()
	if err != nil {
		global.G_LOG.Errorf("[CSRF] Failed to generate token: %v", err)
		return err
	}
	// Local dev usually runs on http; Secure cookies would be dropped by browser and break CSRF flow.
	secureCookie := c.Request.TLS != nil || strings.EqualFold(c.GetHeader("X-Forwarded-Proto"), "https")
	c.SetCookie(CSRFTokenCookieName, token, 0, "/", "", secureCookie, false)
	c.Header(CSRFTokenHeaderKey, token)
	return nil
}

// VerifyCSRFToken 验证请求中的 CSRF token
func VerifyCSRFToken(c *gin.Context) bool {
	tokenFromHeader := c.GetHeader(CSRFTokenHeaderKey)
	if tokenFromHeader == "" {
		tokenFromHeader = c.PostForm("_csrf_token")
	}

	tokenFromCookie, err := c.Cookie(CSRFTokenCookieName)
	if err != nil {
		global.G_LOG.Warnf("[CSRF] Failed to read token from cookie: %v", err)
		return false
	}

	if tokenFromHeader == "" || tokenFromHeader != tokenFromCookie {
		global.G_LOG.Warnf("[CSRF] Token mismatch")
		return false
	}
	return true
}

func shouldSkipCSRF(path string) bool {
	for _, prefix := range csrfSkipPathPrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}

func shouldBypassCSRFByOrigin(c *gin.Context) bool {
	origin := strings.TrimSpace(c.GetHeader("Origin"))
	if origin != "" {
		_, ok := csrfBypassOrigins[origin]
		if ok {
			return true
		}
	}

	referer := strings.TrimSpace(c.GetHeader("Referer"))
	if referer == "" {
		return false
	}
	for _, prefix := range csrfBypassRefererPrefixes {
		if strings.HasPrefix(referer, prefix) {
			return true
		}
	}
	return false
}

// CSRFProtection CSRF 保护中间件
func CSRFProtection() gin.HandlerFunc {
	return func(c *gin.Context) {
		if shouldSkipCSRF(c.Request.URL.Path) {
			c.Next()
			return
		}

		if shouldBypassCSRFByOrigin(c) {
			c.Next()
			return
		}

		method := strings.ToUpper(c.Request.Method)
		if method == "GET" || method == "HEAD" || method == "OPTIONS" {
			if err := SetCSRFTokenCookie(c); err != nil {
				global.G_LOG.Errorf("[CSRF] Failed to set token: %v", err)
			}
			c.Next()
			return
		}

		if !VerifyCSRFToken(c) {
			global.G_LOG.Warnf("[CSRF] Verification failed for %s %s", c.Request.Method, c.Request.RequestURI)
			response.FailJSON(c, ecode.RequestErr)
			c.Abort()
			return
		}
		c.Next()
	}
}
