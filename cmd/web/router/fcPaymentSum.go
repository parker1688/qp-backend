package router

import (
	"bootpkg/cmd/web/controller/fcPaymentSum"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPaymentSumRouter)
}

func fcPaymentSumRouter() {
	r := routers.Group("/api/fcPaymentSum").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPaymentSum.FindPageFcPaymentSumControl)   //查询分页
	r.GET("/findByKey", fcPaymentSum.FindByKeyFcPaymentSumControl) //查询主键
	r.POST("/save", fcPaymentSum.SaveFcPaymentSumControl)          //保存
	r.POST("/update", fcPaymentSum.UpdateFcPaymentSumControl)      //修改
	r.POST("/delete", fcPaymentSum.DeleteFcPaymentSumControl)      //删除
}
