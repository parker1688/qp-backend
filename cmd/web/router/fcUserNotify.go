package router

import (
	"bootpkg/cmd/web/controller/fcUserNotify"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserNotifyRouter)
}

func fcUserNotifyRouter() {
	r := routers.Group("/api/fcUserNotify").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserNotify.FindPageFcUserNotifyControl)   //查询分页
	r.GET("/findByKey", fcUserNotify.FindByKeyFcUserNotifyControl) //查询主键
	r.POST("/save", fcUserNotify.SaveFcUserNotifyControl)          //保存
	r.POST("/update", fcUserNotify.UpdateFcUserNotifyControl)      //修改
	r.POST("/delete", fcUserNotify.DeleteFcUserNotifyControl)      //删除
}
