package router

import (
	"bootpkg/cmd/web/controller/fcCustomerOrder"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcCustomerOrderRouter)
}

func fcCustomerOrderRouter() {
	r := routers.Group("/api/fcCustomerOrder").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcCustomerOrder.FindPageFcCustomerOrderControl)   //查询分页
	r.GET("/findByKey", fcCustomerOrder.FindByKeyFcCustomerOrderControl) //查询主键
	r.POST("/save", fcCustomerOrder.SaveFcCustomerOrderControl)          //保存
	r.POST("/update", fcCustomerOrder.UpdateFcCustomerOrderControl)      //修改
	r.POST("/delete", fcCustomerOrder.DeleteFcCustomerOrderControl)      //删除

	r.POST("/update/status", fcCustomerOrder.UpdateFcCustomerOrderStatusControl) // 更改状态
}
