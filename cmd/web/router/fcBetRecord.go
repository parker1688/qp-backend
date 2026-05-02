package router

import (
	"bootpkg/cmd/web/controller/fcBetRecord"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcBetRecordRouter)
}

func fcBetRecordRouter() {
	r := routers.Group("/api/fcBetRecord").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcBetRecord.FindPageFcBetRecordControl)   //查询分页
	r.GET("/findByKey", fcBetRecord.FindByKeyFcBetRecordControl) //查询主键
	r.POST("/save", fcBetRecord.SaveFcBetRecordControl)          //保存
	r.POST("/update", fcBetRecord.UpdateFcBetRecordControl)      //修改
	r.POST("/delete", fcBetRecord.DeleteFcBetRecordControl)      //删除

	r.POST("/manualPullRecord", fcBetRecord.ManualPullRecord) //手工拉单
	r.POST("/playback", fcBetRecord.VenuePlaybackControl)     // 场馆录像
}
