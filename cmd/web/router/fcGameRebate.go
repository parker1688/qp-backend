package router

import (
	"bootpkg/cmd/web/controller/fcGameRebate"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcGameRebateRouter)
}

func fcGameRebateRouter() {
	r := routers.Group("/api/fcGameRebate").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcGameRebate.FindPageFcGameRebateControl)   //查询分页
	r.GET("/findByKey", fcGameRebate.FindByKeyFcGameRebateControl) //查询主键
	r.POST("/save", fcGameRebate.SaveFcGameRebateControl)          //保存
	r.POST("/update", fcGameRebate.UpdateFcGameRebateControl)      //修改
	r.POST("/delete", fcGameRebate.DeleteFcGameRebateControl)      //删除
}
