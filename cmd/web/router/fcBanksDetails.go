package router

import (
	"bootpkg/cmd/web/controller/fcBanksDetails"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcBanksDetailsRouter)
}

func fcBanksDetailsRouter() {
	r := routers.Group("/api/fcBanksDetails").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcBanksDetails.FindPageFcBanksDetailsControl)   //查询分页
	r.GET("/findByKey", fcBanksDetails.FindByKeyFcBanksDetailsControl) //查询主键
	r.POST("/save", fcBanksDetails.SaveFcBanksDetailsControl)          //保存
	r.POST("/update", fcBanksDetails.UpdateFcBanksDetailsControl)      //修改
	r.POST("/delete", fcBanksDetails.DeleteFcBanksDetailsControl)      //删除
}
