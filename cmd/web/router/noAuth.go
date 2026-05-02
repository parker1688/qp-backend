package router

import (
	"bootpkg/cmd/web/controller/adminUser"
	"bootpkg/cmd/web/controller/commonControl"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, loginTokenRouter)
}

func loginTokenRouter() {
	r := routers.Group("/api/auth/token/").Use(handler.AuthTokenMiddleware())
	r.POST("/isLogin", adminUser.AdminUserIsLogin)                             //是否登录
	r.GET("/menus", adminUser.AdminUserGetMenusByRole)                         //获取权限的菜单
	r.GET("/outlogin", adminUser.LoginOut)                                     //用户登出
	r.POST("/upload/5cc8019d300000980a055e76", commonControl.UploadFileSingle) //上传文件

	rNot := routers.Group("/api/auth/token/").Use(handler.AuthTokenNotSafeMiddleware())
	rNot.POST("/upassword", adminUser.UpdateAdminUserPassWord) //修改密码
}
