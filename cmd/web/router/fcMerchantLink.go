package router

import (
	"bootpkg/cmd/web/controller/fcMerchantLink"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcMerchantLinkRouter)
}

func fcMerchantLinkRouter() {
	r := routers.Group("/api/fcMerchantLink").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcMerchantLink.FindPageFcMerchantLinkControl)   //查询分页
	r.GET("/findByKey", fcMerchantLink.FindByKeyFcMerchantLinkControl) //查询主键
	r.POST("/save", fcMerchantLink.SaveFcMerchantLinkControl)          //保存
	r.POST("/update", fcMerchantLink.UpdateFcMerchantLinkControl)      //修改
	r.POST("/delete", fcMerchantLink.DeleteFcMerchantLinkControl)      //删除
}
