package router

import (
	"bootpkg/cmd/web/controller/fcOrderDeposit"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcOrderDepositRouter)
}

func fcOrderDepositRouter() {
	r := routers.Group("/api/fcOrderDeposit").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcOrderDeposit.FindPageFcOrderDepositControl)   //查询分页
	r.GET("/findByKey", fcOrderDeposit.FindByKeyFcOrderDepositControl) //查询主键
	r.POST("/save", fcOrderDeposit.SaveFcOrderDepositControl)          //保存
	r.POST("/update", fcOrderDeposit.UpdateFcOrderDepositControl)      //修改
	r.POST("/delete", fcOrderDeposit.DeleteFcOrderDepositControl)      //删除

	r.POST("/no", fcOrderDeposit.OrderDepositNo)  //拒绝
	r.POST("/yes", fcOrderDeposit.OrderDepositOk) //通过

	r.POST("/push", fcOrderDeposit.OrderDepositPush) //通过

}
