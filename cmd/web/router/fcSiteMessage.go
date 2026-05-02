package router

import (
	"bootpkg/cmd/web/controller/fcSiteMessage"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcSiteMessageRouter)
}

func fcSiteMessageRouter() {
	r := routers.Group("/api/fcSiteMessage").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcSiteMessage.FindPageFcSiteMessageControl)   //查询分页
	r.GET("/findByKey", fcSiteMessage.FindByKeyFcSiteMessageControl) //查询主键
	r.POST("/save", fcSiteMessage.SaveFcSiteMessageControl)          //保存
	r.POST("/update", fcSiteMessage.UpdateFcSiteMessageControl)      //修改
	r.POST("/delete", fcSiteMessage.DeleteFcSiteMessageControl)      //删除
	// 系统邮件
	r.GET("/system/findPage", fcSiteMessage.FindPageSystemMailControl) // 系统邮件获取
	r.POST("/system/save", fcSiteMessage.SaveSystemMailControl)        // 系统邮件新增
	r.POST("/system/update", fcSiteMessage.UpdateSystemMailControl)    // 系统邮件更新
	r.POST("/system/delete", fcSiteMessage.DeleteSystemMailControl)    // 系统邮件删除
	// 邮件列表
	r.GET("/mail/findPage", fcSiteMessage.FindPageMailControl) // 邮件列表
	r.POST("/mail/delete", fcSiteMessage.DeleteMailControl)    // 邮件列表
}
