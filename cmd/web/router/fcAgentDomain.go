package router

import (
	"bootpkg/cmd/web/controller/fcAgentDomain"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcAgentDomainRouter)
}

func fcAgentDomainRouter() {
	r := routers.Group("/api/fcAgentDomain").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcAgentDomain.FindPageFcAgentDomainControl)        //查询分页
	r.GET("/findByKey", fcAgentDomain.FindByKeyFcAgentDomainControl)      //查询主键
	r.POST("/save", fcAgentDomain.SaveFcAgentDomainControl)               //保存
	r.POST("/update", fcAgentDomain.UpdateFcAgentDomainControl)           //修改
	r.POST("/delete", fcAgentDomain.DeleteFcAgentDomainControl)           //删除
	r.GET("/type/options", fcAgentDomain.TypeOptionsFcAgentDomainControl) // 类型列表
}
