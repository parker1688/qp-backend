package router

import (
	"bootpkg/cmd/web/controller/dailyBonus"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, dailyBonusRouter)
}

func dailyBonusRouter() {
	r := routers.Group("/api/dailyBonus").Use(handler.AuthMiddleware())

	r.GET("/findPage", dailyBonus.FindPageDailyBonusControl)   //查询分页
	r.GET("/findByKey", dailyBonus.FindByKeyDailyBonusControl) //查询主键
	r.POST("/save", dailyBonus.SaveDailyBonusControl)          //保存
	r.POST("/update", dailyBonus.UpdateDailyBonusControl)      //修改
	r.POST("/delete", dailyBonus.DeleteDailyBonusControl)      //删除
}
