package router

import (
	"bootpkg/cmd/api/handler"
	"bootpkg/common/global"
	"bootpkg/common/middleware"
	"bootpkg/pkg/core/modules/enmus"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var routers *gin.Engine
var routersFun []func()

func NewRouter() *gin.Engine {
	if global.CONFIG.General.ENV == enmus.Release {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	routers = gin.Default()
	routers.NoRoute(func(c *gin.Context) {
		c.Status(200)
	})
	criticalPathLimit := map[string]int{
		"/api/user/login":              20,
		"/api/user/password":           20,
		"/api/money/walletWithdraw":    10,
		"/api/money/userChannel":       15,
		"/api/money/getPhoneVeryCode":  5,
	}
	routers.Use(handler.Recovery(), handler.Cors(), handler.RequestRateLimit(300, criticalPathLimit), handler.EncryptionDataMiddleware(), middleware.CSRFProtection())
	routers.Static("/api/upload", "./upload") //静态文件
	//自动执行方法Map
	if len(routersFun) > 0 {
		for _, f := range routersFun {
			f()
		}
	}
	routersFun = nil
	return routers
}
