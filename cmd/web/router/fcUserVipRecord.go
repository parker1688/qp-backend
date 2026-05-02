package router

import (
	"bootpkg/cmd/web/controller/fcUserVipRecord"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserVipRecordRouter)
}

func fcUserVipRecordRouter() {
	r := routers.Group("/api/fcUserVipRecord").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserVipRecord.FindPageFcUserVipRecordControl)   //查询分页
	r.GET("/findByKey", fcUserVipRecord.FindByKeyFcUserVipRecordControl) //查询主键
	r.POST("/save", fcUserVipRecord.SaveFcUserVipRecordControl)          //保存
	r.POST("/update", fcUserVipRecord.UpdateFcUserVipRecordControl)      //修改
	r.POST("/delete", fcUserVipRecord.DeleteFcUserVipRecordControl)      //删除
}
