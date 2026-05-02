package router

import (
	"bootpkg/cmd/web/controller/fcUserWithdrawBlockchainBind"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserWithdrawBlockchainBindRouter)
}

func fcUserWithdrawBlockchainBindRouter() {
	r := routers.Group("/api/fcUserWithdrawBlockchainBind").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserWithdrawBlockchainBind.FindPageFcUserWithdrawBlockchainBindControl)   //查询分页
	r.GET("/findByKey", fcUserWithdrawBlockchainBind.FindByKeyFcUserWithdrawBlockchainBindControl) //查询主键
	r.POST("/save", fcUserWithdrawBlockchainBind.SaveFcUserWithdrawBlockchainBindControl)          //保存
	r.POST("/update", fcUserWithdrawBlockchainBind.UpdateFcUserWithdrawBlockchainBindControl)      //修改
	r.POST("/delete", fcUserWithdrawBlockchainBind.DeleteFcUserWithdrawBlockchainBindControl)      //删除
}
