package router

import (
	"bootpkg/cmd/web/controller/fcVenueConfig"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVenueConfigRouter)
}

func fcVenueConfigRouter() {
	r := routers.Group("/api/fcVenueConfig").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVenueConfig.FindPageFcVenueConfigControl)   //查询分页
	r.GET("/findByKey", fcVenueConfig.FindByKeyFcVenueConfigControl) //查询主键
	r.POST("/save", fcVenueConfig.SaveFcVenueConfigControl)          //保存
	r.POST("/update", fcVenueConfig.UpdateFcVenueConfigControl)      //修改
	r.POST("/delete", fcVenueConfig.DeleteFcVenueConfigControl)      //删除
}
