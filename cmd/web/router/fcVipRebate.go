package router

import (
	"bootpkg/cmd/web/controller/fcVipRebate"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVipRebateRouter)
}

func fcVipRebateRouter() {
	r := routers.Group("/api/fcVipRebate").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVipRebate.FindPageFcVipRebateControl)   //查询分页
	r.GET("/findByKey", fcVipRebate.FindByKeyFcVipRebateControl) //查询主键
	r.POST("/save", fcVipRebate.SaveFcVipRebateControl)          //保存
	r.POST("/update", fcVipRebate.UpdateFcVipRebateControl)      //修改
	r.POST("/delete", fcVipRebate.DeleteFcVipRebateControl)      //删除
}
