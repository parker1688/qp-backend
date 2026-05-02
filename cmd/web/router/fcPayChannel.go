package router

import (
	"bootpkg/cmd/web/controller/fcPayChannel"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPayChannelRouter)
}

func fcPayChannelRouter() {
	r := routers.Group("/api/fcPayChannel").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPayChannel.FindPageFcPayChannelControl)   //查询分页
	r.GET("/findByKey", fcPayChannel.FindByKeyFcPayChannelControl) //查询主键
	r.POST("/save", fcPayChannel.SaveFcPayChannelControl)          //保存
	r.POST("/update", fcPayChannel.UpdateFcPayChannelControl)      //修改
	r.POST("/delete", fcPayChannel.DeleteFcPayChannelControl)      //删除
}
