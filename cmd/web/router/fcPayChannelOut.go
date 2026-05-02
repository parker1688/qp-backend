package router

import (
	"bootpkg/cmd/web/controller/fcPayChannelOut"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPayChannelOutRouter)
}

func fcPayChannelOutRouter() {
	r := routers.Group("/api/fcPayChannelOut").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPayChannelOut.FindPageFcPayChannelOutControl)   //查询分页
	r.GET("/findByKey", fcPayChannelOut.FindByKeyFcPayChannelOutControl) //查询主键
	r.POST("/save", fcPayChannelOut.SaveFcPayChannelOutControl)          //保存
	r.POST("/update", fcPayChannelOut.UpdateFcPayChannelOutControl)      //修改
	r.POST("/delete", fcPayChannelOut.DeleteFcPayChannelOutControl)      //删除
}
