package router

import (
	"bootpkg/cmd/web/controller/fcVipReport"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVipReportRouter)
}

func fcVipReportRouter() {
	r := routers.Group("/api/fcVipReport").Use(handler.AuthMiddleware())

	r.POST("/list", fcVipReport.VipReport) //vip统计报告
}
