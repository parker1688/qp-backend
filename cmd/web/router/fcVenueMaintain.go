package router

import (
	"bootpkg/cmd/web/controller/fcVenueMaintain"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVenueMaintainRouter)
}

func fcVenueMaintainRouter() {
	r := routers.Group("/api/fcVenueMaintain").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVenueMaintain.FindPageFcVenueMaintainControl)   //查询分页
	r.GET("/findByKey", fcVenueMaintain.FindByKeyFcVenueMaintainControl) //查询主键
	r.POST("/save", fcVenueMaintain.SaveFcVenueMaintainControl)          //保存
	r.POST("/update", fcVenueMaintain.UpdateFcVenueMaintainControl)      //修改
	r.POST("/delete", fcVenueMaintain.DeleteFcVenueMaintainControl)      //删除
}
