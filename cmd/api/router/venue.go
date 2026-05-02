package router

import (
	"bootpkg/cmd/api/controller/venueControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, VenueRouter)
}
func VenueRouter() {
	r := routers.Group("/api/venue").Use(handler.AuthApiMiddleware())
	r.POST("/launch", venueControl.Launch)                          //场馆进入-免转
	r.POST("/register", venueControl.Register)                      //场馆注册
	r.POST("/balance", venueControl.Balance)                        //刷新余额
	r.POST("/transferConfirm", venueControl.TransferConfirm)        //场馆转账订单确认状态
	r.POST("/withdraw", venueControl.Withdraw)                      //场馆转账-转出
	r.POST("/deposit", venueControl.Deposit)                        //场馆转账-转入
	r.GET("/recover", venueControl.VenueRecover)                    //场馆一键回收上一家
	r.GET("/recoverall", venueControl.VenueRecoverAll)              //场馆一键回收
	r.GET("/gameRecord", venueControl.GameRecord)                   //注单列表
	r.GET("/gameRecordUnsettled", venueControl.GameRecordUnsettled) //注单列表未结算
	r.GET("/statis", venueControl.Statis)                           //注单分类
	r.GET("/new/batchBalance", venueControl.GetVenusBalances)       //获取所有场馆余额

	//rNew := routers.Group("/api/venue").Use(handler.UrlLocalCacheJson())
	rNew := routers.Group("/api/venue")
	rNew.GET("list", venueControl.GetVenusList) //场馆列表
}
