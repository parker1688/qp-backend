package router

import (
	"bootpkg/cmd/web/controller/fcUserWithdrawBankBind"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserWithdrawBankBindRouter)
}

func fcUserWithdrawBankBindRouter() {
	r := routers.Group("/api/fcUserWithdrawBankBind").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserWithdrawBankBind.FindPageFcUserWithdrawBankBindControl)   //查询分页
	r.GET("/findByKey", fcUserWithdrawBankBind.FindByKeyFcUserWithdrawBankBindControl) //查询主键
	r.POST("/save", fcUserWithdrawBankBind.SaveFcUserWithdrawBankBindControl)          //保存
	r.POST("/update", fcUserWithdrawBankBind.UpdateFcUserWithdrawBankBindControl)      //修改
	r.POST("/delete", fcUserWithdrawBankBind.DeleteFcUserWithdrawBankBindControl)      //删除
}
