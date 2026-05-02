package router

import (
	"bootpkg/cmd/web/controller/log"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, LogRouter)
}

func LogRouter() {
	r := routers.Group("/api/log").Use(handler.AuthMiddleware())
	r.GET("/userLogin/findPage", log.AdminUserLoginLog)
	r.GET("/userAction/findPage", log.FindPageActionControl)
}
