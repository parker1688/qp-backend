package router

import (
	"bootpkg/cmd/web/controller/fcVenueTransfer"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVenueTransferRouter)
}

func fcVenueTransferRouter() {
	r := routers.Group("/api/fcVenueTransfer").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVenueTransfer.FindPageFcVenueTransferControl)   //查询分页
	r.GET("/findByKey", fcVenueTransfer.FindByKeyFcVenueTransferControl) //查询主键
	r.POST("/save", fcVenueTransfer.SaveFcVenueTransferControl)          //保存
	r.POST("/update", fcVenueTransfer.UpdateFcVenueTransferControl)      //修改
	r.POST("/delete", fcVenueTransfer.DeleteFcVenueTransferControl)      //删除

	//先去掉这几个人工影响订单的接口。
	r.POST("/deposit/no", fcVenueTransfer.DepositNo)    //订单存款失败
	r.POST("/deposit/yes", fcVenueTransfer.DepositOk)   //订单存款成功
	r.POST("/withdraw/no", fcVenueTransfer.WithdrawNo)  //订单提款失败
	r.POST("/withdraw/yes", fcVenueTransfer.WithdrawOk) //订单提款成功
}
