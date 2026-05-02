package router

import (
	"bootpkg/cmd/web/controller/opRecord"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, opRecordRouter)
}

func opRecordRouter() {
	r := routers.Group("/api/opRecord").Use(handler.AuthMiddleware())

	r.GET("/findPage", opRecord.FindPageOpRecord) //查询分页
}
