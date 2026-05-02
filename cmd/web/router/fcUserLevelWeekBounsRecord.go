package router

import (
	"bootpkg/cmd/web/controller/fcUserLevelWeekBounsRecord"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserLevelWeekBounsRecordRouter)
}

func fcUserLevelWeekBounsRecordRouter() {
	r := routers.Group("/api/fcUserLevelWeekBounsRecord").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserLevelWeekBounsRecord.FindPageFcUserLevelWeekBounsRecordControl)   //查询分页
	r.GET("/findByKey", fcUserLevelWeekBounsRecord.FindByKeyFcUserLevelWeekBounsRecordControl) //查询主键
	r.POST("/save", fcUserLevelWeekBounsRecord.SaveFcUserLevelWeekBounsRecordControl)          //保存
	r.POST("/update", fcUserLevelWeekBounsRecord.UpdateFcUserLevelWeekBounsRecordControl)      //修改
	r.POST("/delete", fcUserLevelWeekBounsRecord.DeleteFcUserLevelWeekBounsRecordControl)      //删除
}
