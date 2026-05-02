package router

import (
	"bootpkg/cmd/web/controller/adminUser"
	"bootpkg/cmd/web/controller/commonControl"
	"bootpkg/cmd/web/handler"
	"bootpkg/common/global"
	"bootpkg/common/middleware"
	"bootpkg/pkg/core/modules/enmus"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var routers *gin.Engine
var routersFun []func()

const (
	auth_token = "mu0wrXAxh1KivV5BAuXDUyJ7n"
)

func NewRouter() *gin.Engine {
	if global.CONFIG.General.ENV == enmus.Release {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	routers = gin.Default()
	routers.NoRoute(func(c *gin.Context) {
		c.Status(404)
	})
	routers.Static("/api/upload", "./upload") //静态文件
	//handler.EncryptionDataMiddleware()
	routers.Use(handler.Recovery(), cors.Default(), handler.Cors(), handler.WhiteIPMiddleware(), handler.MerchantCodeMiddleware(), middleware.CSRFProtection())

	BaseRouter()
	//自动执行方法Map
	if len(routersFun) > 0 {
		for _, f := range routersFun {
			f()
		}
	}

	return routers
	//return NewCryptoRouter()
}

func BaseRouter() {
	r := routers.Group("/api/base/nologin")
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("login", adminUser.AdminUserLogin)           //用户登录
	r.GET("/getCaptcha", adminUser.GetCaptchaData)      //获取验证码
	r.GET("/csrf-token", commonControl.GetCSRFToken)     //获取 CSRF Token
	r.POST("/checkCaptcha", adminUser.CheckCaptchaData) //验证验证码
	r.POST("/getIP", adminUser.GetIp)                   //获取客户端IP
	//r.GET("/getCaptchaSlide", adminUser.GetCaptchaDataSlide)         //获取滑动验证码
	//r.POST("/checkCaptchaSlide", adminUser.GetCaptchaDataSlideCheck) //获取滑动验证码

}
