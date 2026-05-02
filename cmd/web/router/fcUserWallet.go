package router

import (
	"bootpkg/cmd/web/controller/fcUserWallet"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserWalletRouter)
}

func fcUserWalletRouter() {
	r := routers.Group("/api/fcUserWallet").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserWallet.FindPageFcUserWalletControl)   //查询分页
	r.GET("/findByKey", fcUserWallet.FindByKeyFcUserWalletControl) //查询主键
	r.POST("/save", fcUserWallet.SaveFcUserWalletControl)          //保存
	r.POST("/update", fcUserWallet.UpdateFcUserWalletControl)      //修改
	r.POST("/delete", fcUserWallet.DeleteFcUserWalletControl)      //删除

	r.GET("/fcOrderManageOpt/findPage", fcUserWallet.FindPageFcOrderManageOptControl) //查询分页
	r.POST("/managerOpt", fcUserWallet.WalletAmountOpt)                               // 人工上下分
	r.POST("/clearAmountOpt", fcUserWallet.ClearWalletAmountOpt)                      //管理员清空金额
}
