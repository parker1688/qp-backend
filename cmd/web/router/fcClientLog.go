package router

import (
	"bootpkg/cmd/web/controller/fcClientLog"
)

func init() {
	routersFun = append(routersFun, fcClientLogRouter)
}

func fcClientLogRouter() {
	r := routers.Group("/api/fcClientLog")

	r.GET("/findPage", fcClientLog.FindPageFcClinetLogControl) //查询分页
	r.POST("/save", fcClientLog.SaveFcClinetLogControl)        //写入日志
	r.POST("/update", fcClientLog.UpdateFcClientLogControl)    //更新
	r.GET("/statics", fcClientLog.Statics)                     //统计
}
