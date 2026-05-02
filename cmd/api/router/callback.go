package router

import "bootpkg/cmd/api/controller/callbackControl"

func init() {
	routersFun = append(routersFun, CallbackRouter)
}

func CallbackRouter() {
	routers.GET("/api/authenticate", callbackControl.WUGVerifySession) // WUGDZ 回调

	routers.POST("/api/updateProductStatus", callbackControl.YinRunPayUpdateStatusCallback)  // yinRunPay 更改通道状态回调
	routers.POST("/api/queryProductList", callbackControl.YinRunPayQueryProductListCallback) // yinRunPay 更改通道状态回调

	r := routers.Group("/api/callback")
	r.POST("/pg/VerifySession", callbackControl.PGVerifySession) //PG回调

	r.POST("/pay/:paymentType", callbackControl.PaymentCallBack)       //充值回调
	r.POST("/payOut/:paymentType", callbackControl.PaymentCallBackOut) //代付回调
}
