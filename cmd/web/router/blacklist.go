package router

import (
	"bootpkg/cmd/web/controller/blacklist"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, blacklistRouter)
}

func blacklistRouter() {
	r := routers.Group("/api/blacklist").Use(handler.AuthMiddleware())

	r.GET("/findPage", blacklist.FindPageBlacklistControl)   //查询分页
	r.GET("/findByKey", blacklist.FindByKeyBlacklistControl) //查询主键
	r.POST("/save", blacklist.SaveBlacklistControl)          //保存
	r.POST("/update", blacklist.UpdateBlacklistControl)      //修改
	r.POST("/delete", blacklist.DeleteBlacklistControl)      //删除
}
