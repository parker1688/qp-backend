package router

import (
	"bootpkg/cmd/web/controller/fcAgentReport"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcAgentReportRouter)
}

func fcAgentReportRouter() {
	r := routers.Group("/api/fcAgentReport").Use(handler.AuthMiddleware())

	r.POST("/findAgentReport", fcAgentReport.FindAgentReport)             //统计推广
	r.POST("/findAgentReportDetail", fcAgentReport.FindDetailAgentReport) //统计推广详情
}
