package router

import (
	"bootpkg/cmd/web/controller/fcSmsChannel"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcSmsChannelRouter)
}

func fcSmsChannelRouter() {
	r := routers.Group("/api/fcSmsChannel").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcSmsChannel.FindPageFcSmsChannelControl)   //查询分页
	r.GET("/findByKey", fcSmsChannel.FindByKeyFcSmsChannelControl) //查询主键
	r.POST("/save", fcSmsChannel.SaveFcSmsChannelControl)          //保存
	r.POST("/update", fcSmsChannel.UpdateFcSmsChannelControl)      //修改
	r.POST("/delete", fcSmsChannel.DeleteFcSmsChannelControl)      //删除
}
