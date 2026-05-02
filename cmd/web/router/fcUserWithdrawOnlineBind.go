package router

import (
	"bootpkg/cmd/web/controller/fcUserWithdrawOnlineBind"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserWithdrawOnlineBindRouter)
}

func fcUserWithdrawOnlineBindRouter() {
	r := routers.Group("/api/fcUserWithdrawOnlineBind").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserWithdrawOnlineBind.FindPageFcUserWithdrawOnlineBindControl)   //查询分页
	r.GET("/findByKey", fcUserWithdrawOnlineBind.FindByKeyFcUserWithdrawOnlineBindControl) //查询主键
	r.POST("/save", fcUserWithdrawOnlineBind.SaveFcUserWithdrawOnlineBindControl)          //保存
	r.POST("/update", fcUserWithdrawOnlineBind.UpdateFcUserWithdrawOnlineBindControl)      //修改
	r.POST("/delete", fcUserWithdrawOnlineBind.DeleteFcUserWithdrawOnlineBindControl)      //删除
}
