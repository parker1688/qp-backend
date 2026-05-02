package router

import (
	"bootpkg/cmd/web/controller/dicts"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, dictsRouter)
}

func dictsRouter() {
	r := routers.Group("/api/dicts").Use(handler.AuthMiddleware())

	r.GET("/findPage", dicts.FindPageDictsControl)   //查询分页
	r.GET("/findByKey", dicts.FindByKeyDictsControl) //查询主键
	r.POST("/save", dicts.SaveDictsControl)          //保存
	r.POST("/update", dicts.UpdateDictsControl)      //修改
	r.POST("/delete", dicts.DeleteDictsControl)      //删除
}
