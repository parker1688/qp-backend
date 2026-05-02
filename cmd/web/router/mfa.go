package router

import (
	"bootpkg/cmd/web/controller/gooleMFA"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, mfaRouter)
}

func mfaRouter() {
	r := routers.Group("/api/mfa").Use(handler.AuthTokenNotSafeMiddleware())

	r.GET("/getMFA", gooleMFA.GenerateTOTP)         //生成OTP身份验证器
	r.POST("/validate_bind", gooleMFA.ValidateTOTP) //绑定验证
	r.POST("/validate", gooleMFA.ValidateUserTOTP)  //绑定验证
}
