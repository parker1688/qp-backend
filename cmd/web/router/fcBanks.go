package router

import (
	"bootpkg/cmd/web/controller/fcBanks"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcBanksRouter)
}

func fcBanksRouter() {
	r := routers.Group("/api/fcBanks").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcBanks.FindPageFcBanksControl)   //查询分页
	r.GET("/findByKey", fcBanks.FindByKeyFcBanksControl) //查询主键
	r.POST("/save", fcBanks.SaveFcBanksControl)          //保存
	r.POST("/update", fcBanks.UpdateFcBanksControl)      //修改
	r.POST("/delete", fcBanks.DeleteFcBanksControl)      //删除
}
