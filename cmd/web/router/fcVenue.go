package router

import (
	"bootpkg/cmd/web/controller/fcVenue"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVenueRouter)
}

func fcVenueRouter() {
	r := routers.Group("/api/fcVenue").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVenue.FindPageFcVenueControl)   //查询分页
	r.GET("/findByKey", fcVenue.FindByKeyFcVenueControl) //查询主键
	r.POST("/save", fcVenue.SaveFcVenueControl)          //保存
	r.POST("/update", fcVenue.UpdateFcVenueControl)      //修改
	r.POST("/delete", fcVenue.DeleteFcVenueControl)      //删除
}
