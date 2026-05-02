package router

import (
	"bootpkg/cmd/web/controller/fcUserGameRebate"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserGameRebateRouter)
}

func fcUserGameRebateRouter() {
	r := routers.Group("/api/fcUserGameRebate").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserGameRebate.FindPageFcUserGameRebateControl)   //查询分页
	r.GET("/findByKey", fcUserGameRebate.FindByKeyFcUserGameRebateControl) //查询主键
	r.POST("/save", fcUserGameRebate.SaveFcUserGameRebateControl)          //保存
	r.POST("/update", fcUserGameRebate.UpdateFcUserGameRebateControl)      //修改
	r.POST("/delete", fcUserGameRebate.DeleteFcUserGameRebateControl)      //删除
}
