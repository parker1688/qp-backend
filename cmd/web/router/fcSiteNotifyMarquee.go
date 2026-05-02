package router

import (
	"bootpkg/cmd/web/controller/fcSiteNotifyMarquee"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcSiteNotifyMarqueeRouter)
}

func fcSiteNotifyMarqueeRouter() {
	r := routers.Group("/api/fcSiteNotifyMarquee").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcSiteNotifyMarquee.FindPageFcSiteNotifyMarqueeControl)   //查询分页
	r.GET("/findByKey", fcSiteNotifyMarquee.FindByKeyFcSiteNotifyMarqueeControl) //查询主键
	r.POST("/save", fcSiteNotifyMarquee.SaveFcSiteNotifyMarqueeControl)          //保存
	r.POST("/update", fcSiteNotifyMarquee.UpdateFcSiteNotifyMarqueeControl)      //修改
	r.POST("/delete", fcSiteNotifyMarquee.DeleteFcSiteNotifyMarqueeControl)      //删除
}
