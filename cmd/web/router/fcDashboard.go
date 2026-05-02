package router

import (
	"bootpkg/cmd/web/controller/fcDashboard"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcDashboardRouter)
}

func fcDashboardRouter() {
	r := routers.Group("/api/fcDashboard").Use(handler.AuthMiddleware())

	r.POST("/pics", fcDashboard.DashboardPics) //柱状图-6个值，3个时间周期，每个前后周期对比
	r.POST("/hour", fcDashboard.DashboardHour) //24小时波动曲线。
	r.POST("/user", fcDashboard.DashboardUser) //用户分层
}
