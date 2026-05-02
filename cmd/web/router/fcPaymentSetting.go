package router

import (
	"bootpkg/cmd/web/controller/fcPaymentSetting"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcPaymentSettingRouter)
}

func fcPaymentSettingRouter() {
	r := routers.Group("/api/fcPaymentSetting").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcPaymentSetting.FindPageFcPaymentSettingControl)   //查询分页
	r.GET("/findByKey", fcPaymentSetting.FindByKeyFcPaymentSettingControl) //查询主键
	r.POST("/save", fcPaymentSetting.SaveFcPaymentSettingControl)          //保存
	r.POST("/update", fcPaymentSetting.UpdateFcPaymentSettingControl)      //修改
	r.POST("/delete", fcPaymentSetting.DeleteFcPaymentSettingControl)      //删除
}
