package router

import (
	"bootpkg/cmd/web/controller/fcPayment"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPaymentRouter)
}

func fcPaymentRouter() {
	r := routers.Group("/api/fcPayment").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPayment.FindPageFcPaymentControl)   //查询分页
	r.GET("/findByKey", fcPayment.FindByKeyFcPaymentControl) //查询主键
	r.POST("/save", fcPayment.SaveFcPaymentControl)          //保存
	r.POST("/update", fcPayment.UpdateFcPaymentControl)      //修改
	r.POST("/delete", fcPayment.DeleteFcPaymentControl)      //删除

	r.POST("/sync", fcPayment.SyncFcPayment) //同步渠道配置信息

	r.POST("/batch/status", fcPayment.FcPaymentBatchStatus) //批量修改状态

}
