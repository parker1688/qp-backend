package router

import (
	"bootpkg/cmd/api/controller/commonControl"
)

func init() {
	routersFun = append(routersFun, CommonRouter)
}

func CommonRouter() {
	r := routers.Group("/api")
	r.GET("/csrf-token", commonControl.GetCSRFToken)
}
