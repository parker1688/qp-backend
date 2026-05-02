package router

import (
	"bootpkg/cmd/web/controller/fcBetRecordUnsettled"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcBetRecordUnsettledRouter)
}

func fcBetRecordUnsettledRouter() {
	r := routers.Group("/api/fcBetRecordUnsettled").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcBetRecordUnsettled.FindPageFcBetRecordControl)   //查询分页
	r.GET("/findByKey", fcBetRecordUnsettled.FindByKeyFcBetRecordControl) //查询主键
	r.POST("/save", fcBetRecordUnsettled.SaveFcBetRecordControl)          //保存
	r.POST("/update", fcBetRecordUnsettled.UpdateFcBetRecordControl)      //修改
	r.POST("/delete", fcBetRecordUnsettled.DeleteFcBetRecordControl)      //删除

	r.POST("/playback", fcBetRecordUnsettled.VenuePlaybackControl) // 场馆录像
}
