package router

import (
	"bootpkg/cmd/web/controller/mailTemplate"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, mailTemplateRouter)
}

func mailTemplateRouter() {
	r := routers.Group("/api/mailTemplate").Use(handler.AuthMiddleware())

	r.GET("/findPage", mailTemplate.FindPageMailTemplateControl)   //查询分页
	r.GET("/findByKey", mailTemplate.FindByKeyMailTemplateControl) //查询主键
	r.POST("/save", mailTemplate.SaveMailTemplateControl)          //保存
	r.POST("/update", mailTemplate.UpdateMailTemplateControl)      //修改
	r.POST("/delete", mailTemplate.DeleteMailTemplateControl)      //删除
}
