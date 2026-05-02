package router

import (
	"bootpkg/cmd/web/controller/fcBindBankType"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcBindBankTypeRouter)
}

func fcBindBankTypeRouter() {
	r := routers.Group("/api/fcBindBankType").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcBindBankType.FindPageFcBindBankTypeControl)   //查询分页
	r.GET("/findByKey", fcBindBankType.FindByKeyFcBindBankTypeControl) //查询主键
	r.POST("/save", fcBindBankType.SaveFcBindBankTypeControl)          //保存
	r.POST("/update", fcBindBankType.UpdateFcBindBankTypeControl)      //修改
	r.POST("/delete", fcBindBankType.DeleteFcBindBankTypeControl)      //删除
}
