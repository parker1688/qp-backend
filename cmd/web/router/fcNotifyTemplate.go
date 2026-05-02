package router

import (
	"bootpkg/cmd/web/controller/fcNotifyTemplate"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcNotifyTemplateRouter)
}

func fcNotifyTemplateRouter() {
	r := routers.Group("/api/fcNotifyTemplate").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcNotifyTemplate.FindPageFcNotifyTemplateControl)   //查询分页
	r.GET("/findByKey", fcNotifyTemplate.FindByKeyFcNotifyTemplateControl) //查询主键
	r.POST("/save", fcNotifyTemplate.SaveFcNotifyTemplateControl)          //保存
	r.POST("/update", fcNotifyTemplate.UpdateFcNotifyTemplateControl)      //修改
	r.POST("/delete", fcNotifyTemplate.DeleteFcNotifyTemplateControl)      //删除
}
