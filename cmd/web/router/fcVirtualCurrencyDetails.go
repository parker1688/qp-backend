package router

import (
	"bootpkg/cmd/web/controller/fcVirtualCurrencyDetails"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVirtualCurrencyDetailsRouter)
}

func fcVirtualCurrencyDetailsRouter() {
	r := routers.Group("/api/fcVirtualCurrencyDetails").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVirtualCurrencyDetails.FindPageFcVirtualCurrencyDetailsControl)   //查询分页
	r.GET("/findByKey", fcVirtualCurrencyDetails.FindByKeyFcVirtualCurrencyDetailsControl) //查询主键
	r.POST("/save", fcVirtualCurrencyDetails.SaveFcVirtualCurrencyDetailsControl)          //保存
	r.POST("/update", fcVirtualCurrencyDetails.UpdateFcVirtualCurrencyDetailsControl)      //修改
	r.POST("/delete", fcVirtualCurrencyDetails.DeleteFcVirtualCurrencyDetailsControl)      //删除
}
