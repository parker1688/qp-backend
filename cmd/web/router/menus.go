package router

import (
	"bootpkg/cmd/web/controller/menus"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, menusRouter)
}

func menusRouter() {
	r := routers.Group("/api/menus").Use(handler.AuthMiddleware())

	r.GET("/findPage", menus.FindPageMenusControl)   //查询分页
	r.GET("/findByKey", menus.FindByKeyMenusControl) //查询主键
	r.POST("/save", menus.SaveMenusControl)          //保存
	r.POST("/update", menus.UpdateMenusControl)      //修改
	r.POST("/delete", menus.DeleteMenusControl)      //删除
	r.GET("/findByAll", menus.FindMenusAllControl)      //查询所有数据
}
