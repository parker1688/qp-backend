package router

import (
	"bootpkg/cmd/web/controller/fcChannel"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcChannelRouter)
}

func fcChannelRouter() {
	r := routers.Group("/api/fcChannel").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcChannel.FindPageFcChannelControl)   //查询分页
	r.GET("/findByKey", fcChannel.FindByKeyFcChannelControl) //查询主键
	r.POST("/save", fcChannel.SaveFcChannelControl)          //保存
	r.POST("/update", fcChannel.UpdateFcChannelControl)      //修改
	r.POST("/delete", fcChannel.DeleteFcChannelControl)      //删除
}
