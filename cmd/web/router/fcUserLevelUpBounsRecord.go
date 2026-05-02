package router

import (
	"bootpkg/cmd/web/controller/fcUserLevelUpBounsRecord"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserLevelUpBounsRecordRouter)
}

func fcUserLevelUpBounsRecordRouter() {
	r := routers.Group("/api/fcUserLevelUpBounsRecord").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserLevelUpBounsRecord.FindPageFcUserLevelUpBounsRecordControl)   //查询分页
	r.GET("/findByKey", fcUserLevelUpBounsRecord.FindByKeyFcUserLevelUpBounsRecordControl) //查询主键
	r.POST("/save", fcUserLevelUpBounsRecord.SaveFcUserLevelUpBounsRecordControl)          //保存
	r.POST("/update", fcUserLevelUpBounsRecord.UpdateFcUserLevelUpBounsRecordControl)      //修改
	r.POST("/delete", fcUserLevelUpBounsRecord.DeleteFcUserLevelUpBounsRecordControl)      //删除
}
