package router

import (
	"bootpkg/cmd/api/controller/activityControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, activityRouter)
}

func activityRouter() {
	r := routers.Group("/api/activity").Use(handler.UrlLocalCacheJson())
	r.GET("/list", activityControl.GetPromotionInfo)         //优惠活动列表
	r.GET("/detail", activityControl.GetPromotionInfoDetail) //优惠活动详情

	r2 := routers.Group("/api/user/activity").Use(handler.AuthApiMiddleware())
	r2.GET("/info", activityControl.ActivityInfo)      // 活动信息
	r2.POST("/reward", activityControl.ActivityReward) // 活动领取
}
