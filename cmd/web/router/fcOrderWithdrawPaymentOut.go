package router

import (
	"bootpkg/cmd/web/controller/fcOrderWithdrawPaymentOut"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcOrderWithdrawPaymentOutRouter)
}

func fcOrderWithdrawPaymentOutRouter() {
	r := routers.Group("/api/fcOrderWithdrawPaymentOut").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcOrderWithdrawPaymentOut.FindPageFcOrderWithdrawPaymentOutControl)                //查询分页
	r.GET("/findByKey", fcOrderWithdrawPaymentOut.FindByKeyFcOrderWithdrawPaymentOutControl)              //查询主键
	r.POST("/save", fcOrderWithdrawPaymentOut.SaveFcOrderWithdrawPaymentOutControl)                       //保存
	r.POST("/update", fcOrderWithdrawPaymentOut.UpdateFcOrderWithdrawPaymentOutControl)                   //修改
	r.POST("/delete", fcOrderWithdrawPaymentOut.DeleteFcOrderWithdrawPaymentOutControl)                   //删除
	r.POST("/upWithdrawStats", fcOrderWithdrawPaymentOut.UpWithdrawStatsFcOrderWithdrawPaymentOutControl) //修改打款状态

	r.POST("/batch/anotherPay", fcOrderWithdrawPaymentOut.PayOutOk) //批量申请代付
	r.POST("/batch/no", fcOrderWithdrawPaymentOut.PayOutFail)       //批量拒绝
}
