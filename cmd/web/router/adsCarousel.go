package router

import (
	"bootpkg/cmd/web/controller/adsCarousel"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, adsCarouselRouter)
}

func adsCarouselRouter() {
	r := routers.Group("/api/adsCarousel").Use(handler.AuthMiddleware())

	r.GET("/findPage", adsCarousel.FindPageAdsCarouselControl)   //查询分页
	r.GET("/findByKey", adsCarousel.FindByKeyAdsCarouselControl) //查询主键
	r.POST("/save", adsCarousel.SaveAdsCarouselControl)          //保存
	r.POST("/update", adsCarousel.UpdateAdsCarouselControl)      //修改
	r.POST("/delete", adsCarousel.DeleteAdsCarouselControl)      //删除
}
