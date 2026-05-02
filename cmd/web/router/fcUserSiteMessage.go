package router

import (
	"bootpkg/cmd/web/controller/fcUserSiteMessage"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserSiteMessageRouter)
}

func fcUserSiteMessageRouter() {
	r := routers.Group("/api/fcUserSiteMessage").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserSiteMessage.FindPageFcUserSiteMessageControl)   //查询分页
	r.GET("/findByKey", fcUserSiteMessage.FindByKeyFcUserSiteMessageControl) //查询主键
	r.POST("/save", fcUserSiteMessage.SaveFcUserSiteMessageControl)          //保存
	r.POST("/update", fcUserSiteMessage.UpdateFcUserSiteMessageControl)      //修改
	r.POST("/delete", fcUserSiteMessage.DeleteFcUserSiteMessageControl)      //删除
}
