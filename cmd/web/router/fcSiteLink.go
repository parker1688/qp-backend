package router

import (
	"bootpkg/cmd/web/controller/fcSiteLink"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcSiteLinkRouter)
}

func fcSiteLinkRouter() {
	r := routers.Group("/api/fcSiteLink").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcSiteLink.FindPageFcSiteLinkControl)   //查询分页
	r.GET("/findByKey", fcSiteLink.FindByKeyFcSiteLinkControl) //查询主键
	r.POST("/save", fcSiteLink.SaveFcSiteLinkControl)          //保存
	r.POST("/update", fcSiteLink.UpdateFcSiteLinkControl)      //修改
	r.POST("/delete", fcSiteLink.DeleteFcSiteLinkControl)      //删除

	r.POST("/copy", fcSiteLink.TmplSetting) //复制配置

}
