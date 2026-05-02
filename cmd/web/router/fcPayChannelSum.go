package router

import (
	"bootpkg/cmd/web/controller/fcPayChannelSum"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPayChannelSumRouter)
}

func fcPayChannelSumRouter() {
	r := routers.Group("/api/fcPayChannelSum").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPayChannelSum.FindPageFcPayChannelSumControl)   //查询分页
	r.GET("/findByKey", fcPayChannelSum.FindByKeyFcPayChannelSumControl) //查询主键
	r.POST("/save", fcPayChannelSum.SaveFcPayChannelSumControl)          //保存
	r.POST("/update", fcPayChannelSum.UpdateFcPayChannelSumControl)      //修改
	r.POST("/delete", fcPayChannelSum.DeleteFcPayChannelSumControl)      //删除
}
