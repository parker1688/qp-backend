package router

import (
	"bootpkg/cmd/web/controller/fcCustomerOrderType"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcCustomerOrderTypeRouter)
}

func fcCustomerOrderTypeRouter() {
	r := routers.Group("/api/fcCustomerOrderType").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcCustomerOrderType.FindPageFcCustomerOrderTypeControl)   //查询分页
	r.GET("/findByKey", fcCustomerOrderType.FindByKeyFcCustomerOrderTypeControl) //查询主键
	r.POST("/save", fcCustomerOrderType.SaveFcCustomerOrderTypeControl)          //保存
	r.POST("/update", fcCustomerOrderType.UpdateFcCustomerOrderTypeControl)      //修改
	r.POST("/delete", fcCustomerOrderType.DeleteFcCustomerOrderTypeControl)      //删除
}
