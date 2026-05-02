package router

import (
	"bootpkg/cmd/web/controller/fcVirtualCurrency"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVirtualCurrencyRouter)
}

func fcVirtualCurrencyRouter() {
	r := routers.Group("/api/fcVirtualCurrency").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVirtualCurrency.FindPageFcVirtualCurrencyControl)   //查询分页
	r.GET("/findByKey", fcVirtualCurrency.FindByKeyFcVirtualCurrencyControl) //查询主键
	r.POST("/save", fcVirtualCurrency.SaveFcVirtualCurrencyControl)          //保存
	r.POST("/update", fcVirtualCurrency.UpdateFcVirtualCurrencyControl)      //修改
	r.POST("/delete", fcVirtualCurrency.DeleteFcVirtualCurrencyControl)      //删除
}
