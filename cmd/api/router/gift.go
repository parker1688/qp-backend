package router

import (
	"bootpkg/cmd/api/controller/vipgiftControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, VipRouter)
}
func VipRouter() {
	r := routers.Group("/api/gift/").Use(handler.AuthApiMiddleware())

	r.GET("/week/apply", vipgiftControl.WeekGiftApply)
	r.GET("/month/apply", vipgiftControl.MonthGiftApply)
	r.GET("/vipUp/apply", vipgiftControl.VipUpGiftApply)
}
