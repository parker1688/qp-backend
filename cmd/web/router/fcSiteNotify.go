package router

import (
	"bootpkg/cmd/web/controller/fcSiteNotify"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcSiteNotifyRouter)
}

func fcSiteNotifyRouter() {
	r := routers.Group("/api/fcSiteNotify").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcSiteNotify.FindPageFcSiteNotifyControl)   //查询分页
	r.GET("/findByKey", fcSiteNotify.FindByKeyFcSiteNotifyControl) //查询主键
	r.POST("/save", fcSiteNotify.SaveFcSiteNotifyControl)          //保存
	r.POST("/update", fcSiteNotify.UpdateFcSiteNotifyControl)      //修改
	r.POST("/delete", fcSiteNotify.DeleteFcSiteNotifyControl)      //删除

	r.POST("/save/content", fcSiteNotify.UpdateFcSiteNotifyContentControl) //保存内容
}
