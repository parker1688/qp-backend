package router

import (
	"bootpkg/cmd/web/controller/fcVip"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVipRouter)
}

func fcVipRouter() {
	r := routers.Group("/api/fcVip").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVip.FindPageFcVipControl)   //查询分页
	r.GET("/findByKey", fcVip.FindByKeyFcVipControl) //查询主键
	r.POST("/save", fcVip.SaveFcVipControl)          //保存
	r.POST("/update", fcVip.UpdateFcVipControl)      //修改
	r.POST("/delete", fcVip.DeleteFcVipControl)      //删除
}
