package router

import (
	"bootpkg/cmd/api/controller/customerControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, customerRouter)
}

func customerRouter() {
	r := routers.Group("/api/customer").Use(handler.UrlLocalCacheJson())

	r.GET("/get", customerControl.FindCustomer)
}
