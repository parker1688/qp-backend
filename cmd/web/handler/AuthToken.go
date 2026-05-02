package handler

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"context"

	"github.com/gin-gonic/gin"
)

// 只验证登录不验证权限
func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isSuc, user := authToken(c)
		if !isSuc {
			response.FailErrJSON(c, ecode.Unauthorized, "")
			c.Abort()
			return
		}
		if user.EnforcePwd == 1 { //需要强制修改密码
			response.FailErrJSON(c, ecode.RestPassword, "")
			c.Abort()
			return
		}

		if user.EnforcePwd == 0 {
			/*if modules.UseGoogleMfa() && len(user.Mfa) == 0 {
				global.G_LOG.Errorf("user.EnforcePwd=%v, user.Mfa=%v\n", user.EnforcePwd, user.Mfa)
				response.FailErrJSON(c, ecode.GoogleMFABind, "") //需要绑定google验证码 MFA
				c.Abort()
				return
			}*/
			//global.G_LOG.Errorf("user.EnforcePwd=%v, user.Mfa=%v\n", user.EnforcePwd, user.Mfa)
			//response.FailErrJSON(c, ecode.GoogleMFABind, "") //需要绑定google验证码 MFA
			//c.Abort()
			//return
		}

		tokenV := global.G_REDIS.Get(context.Background(), user.Token).Val()
		if tokenV != "1" { //谷歌验证码
			response.FailErrJSON(c, ecode.GoogleMFAVerify, "")
			c.Abort()
			return
		}
		c.Set("UserInfo", user)
		c.Next()
	}
}

// 只验证登录不验证、权限、安全密码
func AuthTokenNotSafeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isSuc, user := authToken(c)
		if !isSuc {
			response.FailErrJSON(c, ecode.Unauthorized, "")
			c.Abort()
			return
		}
		c.Set("UserInfo", user)
		c.Next()
	}
}
