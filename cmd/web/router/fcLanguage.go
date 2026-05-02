package router

import (
	"bootpkg/cmd/web/controller/fcLanguage"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcLanguageRouter)
}

func fcLanguageRouter() {
	r := routers.Group("/api/fcLanguage").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcLanguage.FindPageFcLanguageControl)   //查询分页
	r.GET("/findByKey", fcLanguage.FindByKeyFcLanguageControl) //查询主键
	r.POST("/save", fcLanguage.SaveFcLanguageControl)          //保存
	r.POST("/update", fcLanguage.UpdateFcLanguageControl)      //修改
	r.POST("/delete", fcLanguage.DeleteFcLanguageControl)      //删除
}
