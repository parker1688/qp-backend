package router

import (
	"bootpkg/cmd/web/controller/fcMerchant"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcMerchantRouter)
}

func fcMerchantRouter() {
	r := routers.Group("/api/fcMerchant").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcMerchant.FindPageFcMerchantControl)   //查询分页
	r.GET("/findByKey", fcMerchant.FindByKeyFcMerchantControl) //查询主键
	r.POST("/save", fcMerchant.SaveFcMerchantControl)          //保存
	r.POST("/update", fcMerchant.UpdateFcMerchantControl)      //修改
	r.POST("/delete", fcMerchant.DeleteFcMerchantControl)      //删除
}
