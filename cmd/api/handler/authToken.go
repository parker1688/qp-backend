package handler

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/srv"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/ecode"
	"bootpkg/langs"
	"bootpkg/pkg/core/modules/enmus"
	"github.com/gin-gonic/gin"
)

// 只验证登录不验证权限
func AuthApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isSuc, userName := authToken(c)
		if !isSuc {
			response.FailErrJSON(c, ecode.Unauthorized, langs.GetWithLocaleGin(c, "login_fail"))
			c.Abort()
			return
		}
		user := srv.GetUserMaterial(userName)
		if len(user.UserId) == 0 {
			response.FailErrJSON(c, ecode.Unauthorized, langs.GetWithLocaleGin(c, "login_fail"))
			c.Abort()
			return
		}
		c.Set(vo.USER_NAME_INFO_G, user)
		c.Next()
	}
}

func authToken(c *gin.Context) (bool, string) {
	tokenStr := c.Request.Header.Get(enmus.LOGIN_TOKEN)
	if len(tokenStr) == 0 {
		return false, ""
	}
	return srv.VerifyJWTToken(tokenStr)
}
