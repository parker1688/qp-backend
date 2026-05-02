package router

import (
	"bootpkg/cmd/web/controller/fcVipWeekGift"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcVipWeekGiftRouter)
}

func fcVipWeekGiftRouter() {
	r := routers.Group("/api/fcVipWeekGift").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcVipWeekGift.FindPageFcVipWeekGiftControl)   //查询分页
	r.GET("/findByKey", fcVipWeekGift.FindByKeyFcVipWeekGiftControl) //查询主键
	r.POST("/save", fcVipWeekGift.SaveFcVipWeekGiftControl)          //保存
	r.POST("/update", fcVipWeekGift.UpdateFcVipWeekGiftControl)      //修改
	r.POST("/delete", fcVipWeekGift.DeleteFcVipWeekGiftControl)      //删除
}
