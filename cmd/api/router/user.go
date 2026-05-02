package router

import (
	"bootpkg/cmd/api/controller/userControl"
)

func init() {
	routersFun = append(routersFun, UserRouter)
}

func UserRouter() {
	r := routers.Group("/api/")
	r.GET("/merchant/code", userControl.FindMerchantCode)
	r.POST("/register", userControl.Register)
	r.POST("/login", userControl.Login)
	r.POST("/login/abnormal/recover", userControl.LoginAbnormalRecover)
	r.POST("/genCode", userControl.GenVerificationCodeImage) //生成验证码图片
	r.POST("/invite/save", userControl.InviteSave)           // 推广页统计
	r.POST("/invite/link", userControl.InviteLink)           // 推广页获取跳转链接
	r.GET("/customer/link", userControl.CustomerLink)        // 客服链接
	r.GET("/guide/info", userControl.GuideInfo)              // 文本导航信息
	r.GET("/adsCarousel/info", userControl.AdsCarouselInfo)  // 广告栏信息
	r.GET("/currency/fx", userControl.GetCurrencyFx)         // 获取货币汇率
	r.POST("/clientlogs", userControl.TempClientLogs)        // 前端临时日志
	r.GET("/ip/info", userControl.GetClientIpInfo)           // 获取客户的ip信息
}
