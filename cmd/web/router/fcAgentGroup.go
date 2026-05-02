package router

import (
	"bootpkg/cmd/web/controller/fcAgentGroup"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcAgentGroupRouter)
}

func fcAgentGroupRouter() {
	r := routers.Group("/api/fcAgentGroup").Use(handler.AuthMiddleware())

	r.GET("/list", fcAgentGroup.ListFcAgentGroup)           //查询主键
	r.GET("/findByKey", fcAgentGroup.FindByKeyFcAgentGroup) //查询主键
	r.POST("/save", fcAgentGroup.SaveFcAgentGroup)          //保存
	r.POST("/update", fcAgentGroup.UpdateFcAgentGroup)      //修改
	r.POST("/delete", fcAgentGroup.DeleteFcAgentGroup)      //删除
}
