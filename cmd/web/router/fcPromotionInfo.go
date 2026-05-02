package router

import (
	"bootpkg/cmd/web/controller/fcPromotionInfo"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPromotionInfoRouter)
}

func fcPromotionInfoRouter() {
	r := routers.Group("/api/fcPromotionInfo").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPromotionInfo.FindPageFcPromotionInfoControl)   //查询分页
	r.GET("/findByKey", fcPromotionInfo.FindByKeyFcPromotionInfoControl) //查询主键
	r.POST("/save", fcPromotionInfo.SaveFcPromotionInfoControl)          //保存
	r.POST("/update", fcPromotionInfo.UpdateFcPromotionInfoControl)      //修改
	r.POST("/delete", fcPromotionInfo.DeleteFcPromotionInfoControl)      //删除

	r.POST("/save/content", fcPromotionInfo.UpdateFcPromotionInfoContentControl) //保存内容
	r.POST("/update/status", fcPromotionInfo.UpdateFcPromotionInfoStatusControl) //活动开关
}
