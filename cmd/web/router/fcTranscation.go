package router

import (
	"bootpkg/cmd/web/controller/fcTranscation"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcTranscationRouter)
}

func fcTranscationRouter() {
	r := routers.Group("/api/fcTranscation").Use(handler.AuthMiddleware())

	r.GET("/cashflow", fcTranscation.CashFlowFcTranscationControl)   //查询分页
	r.GET("/findPage", fcTranscation.FindPageFcTranscationControl)   //查询分页
	r.GET("/findByKey", fcTranscation.FindByKeyFcTranscationControl) //查询主键
	r.POST("/save", fcTranscation.SaveFcTranscationControl)          //保存
	r.POST("/update", fcTranscation.UpdateFcTranscationControl)      //修改
	r.POST("/delete", fcTranscation.DeleteFcTranscationControl)      //删除
}
