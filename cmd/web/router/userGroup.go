package router

import (
	"bootpkg/cmd/web/controller/userGroup"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, userGroupRouter)
}

func userGroupRouter() {
	r := routers.Group("/api/userGroup").Use(handler.AuthMiddleware())

	r.GET("/findPage", userGroup.FindPageUserGroupControl)   //查询分页
	r.GET("/findByKey", userGroup.FindByKeyUserGroupControl) //查询主键
	r.POST("/save", userGroup.SaveUserGroupControl)          //保存
	r.POST("/update", userGroup.UpdateUserGroupControl)      //修改
	r.POST("/delete", userGroup.DeleteUserGroupControl)      //删除
}
