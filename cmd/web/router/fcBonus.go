package router

import (
	"bootpkg/cmd/web/controller/fcBonus"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcBonusRouter)
}

func fcBonusRouter() {
	r := routers.Group("/api/fcBonus").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcBonus.FindPageFcBonusControl)   //查询分页
	r.GET("/findByKey", fcBonus.FindByKeyFcBonusControl) //查询主键
	r.POST("/save", fcBonus.SaveFcBonusControl)          //保存
	r.POST("/update", fcBonus.UpdateFcBonusControl)      //修改
	r.POST("/delete", fcBonus.DeleteFcBonusControl)      //删除
}
