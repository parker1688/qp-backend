package router

import (
	"bootpkg/cmd/web/controller/fcComplexReport"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcComplexReportRouter)
}

func fcComplexReportRouter() {
	r := routers.Group("/api/fcComplexReport").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcComplexReport.FindPageFcComplexReportControl)   //查询分页
	r.GET("/findByKey", fcComplexReport.FindByKeyFcComplexReportControl) //查询主键
	r.POST("/save", fcComplexReport.SaveFcComplexReportControl)          //保存
	r.POST("/update", fcComplexReport.UpdateFcComplexReportControl)      //修改
	r.POST("/delete", fcComplexReport.DeleteFcComplexReportControl)      //删除
}
