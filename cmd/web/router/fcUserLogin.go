package router

import (
	"bootpkg/cmd/web/controller/fcUserLogin"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserLoginRouter)
}

func fcUserLoginRouter() {
	r := routers.Group("/api/fcUserLogin").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserLogin.FindPageFcUserLoginControl)   //查询分页
	r.GET("/findByKey", fcUserLogin.FindByKeyFcUserLoginControl) //查询主键
	r.POST("/save", fcUserLogin.SaveFcUserLoginControl)          //保存
	r.POST("/update", fcUserLogin.UpdateFcUserLoginControl)      //修改
	r.POST("/delete", fcUserLogin.DeleteFcUserLoginControl)      //删除

	r.POST("/resetPwd", fcUserLogin.ResetPassword)     //重置密码
	r.POST("/update/status", fcUserLogin.UpdateStatus) //修改禁用状态
}
