package router

import (
	"bootpkg/cmd/web/controller/fcUserReport"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserReportRouter)
}

func fcUserReportRouter() {
	r := routers.Group("/api/fcUserReport").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserReport.FindPageFcUserReportControl)   //查询分页
	r.GET("/findByKey", fcUserReport.FindByKeyFcUserReportControl) //查询主键
	//r.POST("/save", fcUserReport.SaveFcUserReportControl)          //保存
	//r.POST("/update", fcUserReport.UpdateFcUserReportControl)      //修改
	//r.POST("/delete", fcUserReport.DeleteFcUserReportControl)      //删除
}
