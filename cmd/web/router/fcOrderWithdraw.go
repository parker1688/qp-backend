package router

import (
	"bootpkg/cmd/web/controller/fcOrderWithdraw"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcOrderWithdrawRouter)
}

func fcOrderWithdrawRouter() {
	r := routers.Group("/api/fcOrderWithdraw").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcOrderWithdraw.FindPageFcOrderWithdrawControl)   //查询分页
	r.GET("/findByKey", fcOrderWithdraw.FindByKeyFcOrderWithdrawControl) //查询主键
	r.POST("/save", fcOrderWithdraw.SaveFcOrderWithdrawControl)          //保存
	r.POST("/update", fcOrderWithdraw.UpdateFcOrderWithdrawControl)      //修改
	r.POST("/delete", fcOrderWithdraw.DeleteFcOrderWithdrawControl)      //删除

	r.POST("/yes", fcOrderWithdraw.OrderWithdrawOk)       //通过
	r.POST("/no", fcOrderWithdraw.OrderWithdrawNo)        //拒绝
	r.POST("/audit", fcOrderWithdraw.OrderWithdrawAudit2) //审核通过

	r.GET("/getAnotherPay", fcOrderWithdraw.OrderWithdrawGetAnotherPay) //获取符合代付通道
	//r.POST("/anotherPay", fcOrderWithdraw.OrderWithdrawAnotherPay)      //申请代付

	r.POST("/batch/anotherPay", fcOrderWithdraw.OrderWithdrawAnotherBatchPay) //批量申请代付
	r.POST("/batch/no", fcOrderWithdraw.OrderWithdrawBatchNo)                 //批量拒绝
	r.POST("/batch/audit", fcOrderWithdraw.OrderWithdrawBatchAudit)           //批量审核通过
}
