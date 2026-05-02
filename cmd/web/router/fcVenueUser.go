package router

import (
	"bootpkg/cmd/web/controller/fcVenueUser"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVenueUserRouter)
}

func fcVenueUserRouter() {
	r := routers.Group("/api/fcVenueUser").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVenueUser.FindPageFcVenueUserControl)   //查询分页
	r.GET("/findByKey", fcVenueUser.FindByKeyFcVenueUserControl) //查询主键
	r.POST("/save", fcVenueUser.SaveFcVenueUserControl)          //保存
	r.POST("/update", fcVenueUser.UpdateFcVenueUserControl)      //修改
	r.POST("/delete", fcVenueUser.DeleteFcVenueUserControl)      //删除

	r.POST("/in", fcVenueUser.TransferIn)                                    //转入金额
	r.POST("/out", fcVenueUser.TransferOut)                                  //转出金额
	r.GET("/getBalancePage", fcVenueUser.FindPageFcVenueUserByUserIdControl) //获取账户场馆信息及余额
}
