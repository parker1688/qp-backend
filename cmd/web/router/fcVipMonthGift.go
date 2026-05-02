package router

import (
	"bootpkg/cmd/web/controller/fcVipMonthGift"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVipMonthGiftRouter)
}

func fcVipMonthGiftRouter() {
	r := routers.Group("/api/fcVipMonthGift").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVipMonthGift.FindPageFcVipMonthGiftControl)   //查询分页
	r.GET("/findByKey", fcVipMonthGift.FindByKeyFcVipMonthGiftControl) //查询主键
	r.POST("/save", fcVipMonthGift.SaveFcVipMonthGiftControl)          //保存
	r.POST("/update", fcVipMonthGift.UpdateFcVipMonthGiftControl)      //修改
	r.POST("/delete", fcVipMonthGift.DeleteFcVipMonthGiftControl)      //删除
}
