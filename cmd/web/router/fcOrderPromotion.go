package router

import (
	"bootpkg/cmd/web/controller/fcOrderPromotion"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcOrderPromotionRouter)
}

func fcOrderPromotionRouter() {
	r := routers.Group("/api/fcOrderPromotion").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcOrderPromotion.FindPageFcOrderPromotionControl)   //查询分页
	r.GET("/findByKey", fcOrderPromotion.FindByKeyFcOrderPromotionControl) //查询主键
	r.POST("/save", fcOrderPromotion.SaveFcOrderPromotionControl)          //保存
	r.POST("/update", fcOrderPromotion.UpdateFcOrderPromotionControl)      //修改
	r.POST("/delete", fcOrderPromotion.DeleteFcOrderPromotionControl)      //删除

	r.POST("/no", fcOrderPromotion.UpdatePromotionNo)  //拒绝
	r.POST("/yes", fcOrderPromotion.UpdatePromotionOK) //通过
}
