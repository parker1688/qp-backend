package router

import (
	"bootpkg/cmd/api/controller/dailyBonusControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, dailyBonusRouter)
}

func dailyBonusRouter() {
	r := routers.Group("/api/dailyBonus").Use(handler.AuthApiMiddleware())
	r.GET("/info", dailyBonusControl.DailyBonusInfo)      //签到信息
	r.POST("/reward", dailyBonusControl.DailyBonusReward) // 签到领取
}
