package router

import (
	"bootpkg/cmd/web/controller/fcFirstOrderDeposit"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcFirstOrderDepositRouter)
}

func fcFirstOrderDepositRouter() {
	r := routers.Group("/api/fcFirstOrderDeposit").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcFirstOrderDeposit.FindPageFcFirstOrderDepositControl)   //查询分页
	r.GET("/findByKey", fcFirstOrderDeposit.FindByKeyFcFirstOrderDepositControl) //查询主键
	r.POST("/save", fcFirstOrderDeposit.SaveFcFirstOrderDepositControl)          //保存
	r.POST("/update", fcFirstOrderDeposit.UpdateFcFirstOrderDepositControl)      //修改
	r.POST("/delete", fcFirstOrderDeposit.DeleteFcFirstOrderDepositControl)      //删除
}
