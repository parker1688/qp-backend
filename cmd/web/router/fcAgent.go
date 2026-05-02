package router

import (
	"bootpkg/cmd/web/controller/fcAgent"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcAgentRouter)
}

func fcAgentRouter() {
	r := routers.Group("/api/fcAgent").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcAgent.FindPageFcAgentControl)   //查询分页
	r.GET("/findByKey", fcAgent.FindByKeyFcAgentControl) //查询主键
	r.GET("/inviteCode", fcAgent.InviteCode)             //查询主键
	r.POST("/save", fcAgent.SaveFcAgentControl)          //保存
	r.POST("/update", fcAgent.UpdateFcAgentControl)      //修改
	r.POST("/delete", fcAgent.DeleteFcAgentControl)      //删除
}
