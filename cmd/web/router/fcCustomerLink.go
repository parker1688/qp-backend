package router

import (
	"bootpkg/cmd/web/controller/fcCustomerLink"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcCustomerLinkRouter)
}

func fcCustomerLinkRouter() {
	r := routers.Group("/api/fcCustomerLink").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcCustomerLink.FindPageFcCustomerLinkControl)   //查询分页
	r.GET("/findByKey", fcCustomerLink.FindByKeyFcCustomerLinkControl) //查询主键
	r.POST("/save", fcCustomerLink.SaveFcCustomerLinkControl)          //保存
	r.POST("/update", fcCustomerLink.UpdateFcCustomerLinkControl)      //修改
	r.POST("/delete", fcCustomerLink.DeleteFcCustomerLinkControl)      //删除
}
