package router

import (
	"bootpkg/cmd/web/controller/analysis"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, analysisRouter)
}

func analysisRouter() {
	r := routers.Group("/api/analysisRetention").Use(handler.AuthMiddleware())
	r.GET("/findByKey", analysis.FindByKeyAnalysisRetentionControl) //查询分页
}
