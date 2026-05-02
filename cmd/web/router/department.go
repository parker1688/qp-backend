package router

import (
	"bootpkg/cmd/web/controller/department"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, departmentRouter)
}

func departmentRouter() {
	r := routers.Group("/api/department").Use(handler.AuthMiddleware())

	r.GET("/findPage", department.FindPageDepartmentControl)   //查询分页
	r.GET("/findByKey", department.FindByKeyDepartmentControl) //查询主键
	r.POST("/save", department.SaveDepartmentControl)          //保存
	r.POST("/update", department.UpdateDepartmentControl)      //修改
	r.POST("/delete", department.DeleteDepartmentControl)      //删除
	r.GET("/findByAll", department.FindDepartmentAllControl)      //查询所有数据
}
