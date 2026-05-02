package router

import (
	"bootpkg/cmd/web/controller/role"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, roleRouter)
}

func roleRouter() {
	r := routers.Group("/api/role").Use(handler.AuthMiddleware())

	r.GET("/findPage", role.FindPageRoleControl)        //查询分页
	r.GET("/findByKey", role.FindByKeyRoleControl)      //查询主键
	r.POST("/save", role.SaveRoleControl)               //保存
	r.POST("/update", role.UpdateRoleControl)           //修改
	r.POST("/delete", role.DeleteRoleControl)           //删除
	r.GET("/findByAll", role.FindRoleAllControl)        //查询所有数据
	r.POST("/updateMenus", role.UpdateRoleMenusControl) //更新角色菜单信息
	r.GET("/permslist", role.GetRolePermsControl)       //获取角色权限列表
}
