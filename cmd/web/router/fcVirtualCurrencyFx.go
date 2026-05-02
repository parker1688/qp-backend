package router

import (
	"bootpkg/cmd/web/controller/fcVirtualCurrencyFx"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVirtualCurrencyFxRouter)
}

func fcVirtualCurrencyFxRouter() {
	r := routers.Group("/api/fcVirtualCurrencyFx").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVirtualCurrencyFx.FindPageFcVirtualCurrencyFxControl)   //查询分页
	r.GET("/findByKey", fcVirtualCurrencyFx.FindByKeyFcVirtualCurrencyFxControl) //查询主键
	r.POST("/save", fcVirtualCurrencyFx.SaveFcVirtualCurrencyFxControl)          //保存
	r.POST("/update", fcVirtualCurrencyFx.UpdateFcVirtualCurrencyFxControl)      //修改
	r.POST("/delete", fcVirtualCurrencyFx.DeleteFcVirtualCurrencyFxControl)      //删除
}
