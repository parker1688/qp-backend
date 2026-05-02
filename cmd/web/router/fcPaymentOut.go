package router

import (
	"bootpkg/cmd/web/controller/fcPaymentOut"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPaymentOutRouter)
}

func fcPaymentOutRouter() {
	r := routers.Group("/api/fcPaymentOut").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPaymentOut.FindPageFcPaymentOutControl)   //查询分页
	r.GET("/findByKey", fcPaymentOut.FindByKeyFcPaymentOutControl) //查询主键
	r.POST("/save", fcPaymentOut.SaveFcPaymentOutControl)          //保存
	r.POST("/update", fcPaymentOut.UpdateFcPaymentOutControl)      //修改
	r.POST("/delete", fcPaymentOut.DeleteFcPaymentOutControl)      //删除
}
