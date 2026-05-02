package router

import (
	"bootpkg/cmd/api/controller/versionControl"
	"bootpkg/cmd/web/controller/fcClientLog"
)

func init() {
	routersFun = append(routersFun, ServerRouter)
}
func ServerRouter() {
	r := routers.Group("/api/server")
	r.GET("/version", versionControl.VersionInfo)

	r1 := routers.Group("/api/fcClientLog")
	r1.POST("/save", fcClientLog.SaveFcClinetLogControl)     //写入日志
	r1.POST("/update", fcClientLog.UpdateFcClientLogControl) //更新
}
