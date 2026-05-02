package router

import (
	"bootpkg/cmd/web/controller/guide"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, guideRouter)
}

func guideRouter() {
	r := routers.Group("/api/guide").Use(handler.AuthMiddleware())

	r.GET("/findPage", guide.FindPageGuideControl)   //查询分页
	r.GET("/findByKey", guide.FindByKeyGuideControl) //查询主键
	r.POST("/save", guide.SaveGuideControl)          //保存
	r.POST("/update", guide.UpdateGuideControl)      //修改
	r.POST("/delete", guide.DeleteGuideControl)      //删除
}
