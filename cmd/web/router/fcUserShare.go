package router

import (
	"bootpkg/cmd/web/controller/fcUserShare"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserShareRouter)
}

func fcUserShareRouter() {
	r := routers.Group("/api/fcUserShare").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserShare.FindPageFcUserShareControl)   //查询分页
	r.GET("/findByKey", fcUserShare.FindByKeyFcUserShareControl) //查询主键
	r.POST("/save", fcUserShare.SaveFcUserShareControl)          //保存
	r.POST("/update", fcUserShare.UpdateFcUserShareControl)      //修改
	r.POST("/delete", fcUserShare.DeleteFcUserShareControl)      //删除
}
