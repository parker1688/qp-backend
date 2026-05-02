package router

import (
	"bootpkg/cmd/web/controller/fcSiteBanner"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcSiteBannerRouter)
}

func fcSiteBannerRouter() {
	r := routers.Group("/api/fcSiteBanner").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcSiteBanner.FindPageFcSiteBannerControl)   //查询分页
	r.GET("/findByKey", fcSiteBanner.FindByKeyFcSiteBannerControl) //查询主键
	r.POST("/save", fcSiteBanner.SaveFcSiteBannerControl)          //保存
	r.POST("/update", fcSiteBanner.UpdateFcSiteBannerControl)      //修改
	r.POST("/delete", fcSiteBanner.DeleteFcSiteBannerControl)      //删除

	r.POST("/copy", fcSiteBanner.BannerCopy) //复制

}
