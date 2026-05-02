package router

import (
	"bootpkg/cmd/web/controller/fcMerchantVenue"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcMerchantVenueRouter)
}

func fcMerchantVenueRouter() {
	r := routers.Group("/api/fcMerchantVenue").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcMerchantVenue.FindPageFcMerchantVenueControl)   //查询分页
	r.GET("/findByKey", fcMerchantVenue.FindByKeyFcMerchantVenueControl) //查询主键
	r.POST("/save", fcMerchantVenue.SaveFcMerchantVenueControl)          //保存
	r.POST("/update", fcMerchantVenue.UpdateFcMerchantVenueControl)      //修改
	r.POST("/delete", fcMerchantVenue.DeleteFcMerchantVenueControl)      //删除

	r.POST("/copy", fcMerchantVenue.CopyVenue) //复制商户场馆信息
}
