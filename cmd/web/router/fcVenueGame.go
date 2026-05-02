package router

import (
	"bootpkg/cmd/web/controller/fcVenueGame"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVenueGameRouter)
}

func fcVenueGameRouter() {
	r := routers.Group("/api/fcVenueGame").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVenueGame.FindPageFcVenueGameControl)   //查询分页
	r.GET("/findByKey", fcVenueGame.FindByKeyFcVenueGameControl) //查询主键
	r.POST("/save", fcVenueGame.SaveFcVenueGameControl)          //保存
	r.POST("/update", fcVenueGame.UpdateFcVenueGameControl)      //修改
	r.POST("/delete", fcVenueGame.DeleteFcVenueGameControl)      //删除

	r.GET("/find/venue", fcVenueGame.FindFcVenueControl) //查询场馆所有配置
}
