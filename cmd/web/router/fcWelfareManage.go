package router

import (
	"bootpkg/cmd/web/controller/fcWelfareManage"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcWelfareManageRouter)
}

func fcWelfareManageRouter() {
	r := routers.Group("/api/fcWelfareManage").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcWelfareManage.FindPageFcWelfareManageControl)   //查询分页
	r.GET("/findByKey", fcWelfareManage.FindByKeyFcWelfareManageControl) //查询主键
	r.POST("/save", fcWelfareManage.SaveFcWelfareManageControl)          //保存
	r.POST("/update", fcWelfareManage.UpdateFcWelfareManageControl)      //修改
	r.POST("/delete", fcWelfareManage.DeleteFcWelfareManageControl)      //删除

	r.POST("/update/flowMultiple", fcWelfareManage.UpdateFcWelfareManageFlowMultipleControl) // 修改流水倍数
}
