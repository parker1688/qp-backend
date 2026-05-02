package router

import (
	"bootpkg/cmd/web/controller/fcCurrency"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcCurrencyRouter)
}

func fcCurrencyRouter() {
	r := routers.Group("/api/fcCurrency").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcCurrency.FindPageFcCurrencyControl)   //查询分页
	r.GET("/findByKey", fcCurrency.FindByKeyFcCurrencyControl) //查询主键
	r.POST("/save", fcCurrency.SaveFcCurrencyControl)          //保存
	r.POST("/update", fcCurrency.UpdateFcCurrencyControl)      //修改
	r.POST("/delete", fcCurrency.DeleteFcCurrencyControl)      //删除
}
