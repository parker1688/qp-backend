package router

import (
	"bootpkg/cmd/web/controller/fcUserRebateRecords"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserRebateRecordsRouter)
}

func fcUserRebateRecordsRouter() {
	r := routers.Group("/api/fcUserRebateRecords").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserRebateRecords.FindPageFcUserRebateRecordsControl)   //查询分页
	r.GET("/findByKey", fcUserRebateRecords.FindByKeyFcUserRebateRecordsControl) //查询主键
	r.POST("/save", fcUserRebateRecords.SaveFcUserRebateRecordsControl)          //保存
	r.POST("/update", fcUserRebateRecords.UpdateFcUserRebateRecordsControl)      //修改
	r.POST("/delete", fcUserRebateRecords.DeleteFcUserRebateRecordsControl)      //删除
}
