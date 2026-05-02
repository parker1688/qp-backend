package srv

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"github.com/gin-gonic/gin"
)

func ExitLogin(c *gin.Context) bool {
	sessionId, err := c.Cookie(enmus.LOGIN_COOKIE)
	if err != nil {
		return false
	}
	//获取用户信息
	tokenKey := enmus.REDIS_LOGIN_TOKEN + sessionId
	global.G_REDIS.Del(context.Background(), tokenKey)
	return true
}
