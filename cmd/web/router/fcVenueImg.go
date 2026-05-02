package router

import (
	"bootpkg/cmd/web/controller/fcVenueImg"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVenueImgRouter)
}

func fcVenueImgRouter() {
	r := routers.Group("/api/fcVenueImg").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVenueImg.FindPageFcVenueImgControl)   //查询分页
	r.GET("/findByKey", fcVenueImg.FindByKeyFcVenueImgControl) //查询主键
	r.POST("/save", fcVenueImg.SaveFcVenueImgControl)          //保存
	r.POST("/update", fcVenueImg.UpdateFcVenueImgControl)      //修改
	r.POST("/delete", fcVenueImg.DeleteFcVenueImgControl)      //删除
}
