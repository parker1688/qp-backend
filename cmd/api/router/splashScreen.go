package router

import "bootpkg/cmd/api/controller/splashScreenControl"

func init() {
	routersFun = append(routersFun, splashScreenRouter)
}

func splashScreenRouter() {
	r := routers.Group("/api/splashScreen")
	r.GET("/info", splashScreenControl.SplashScreenInfo) // 开屏信息
}
