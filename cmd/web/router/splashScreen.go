package router

import (
	"bootpkg/cmd/web/controller/splashScreen"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, splashScreenRouter)
}

func splashScreenRouter() {
	r := routers.Group("/api/splashScreen").Use(handler.AuthMiddleware())

	r.GET("/findPage", splashScreen.FindPageSplashScreenControl)   //查询分页
	r.GET("/findByKey", splashScreen.FindByKeySplashScreenControl) //查询主键
	r.POST("/save", splashScreen.SaveSplashScreenControl)          //保存
	r.POST("/update", splashScreen.UpdateSplashScreenControl)      //修改
	r.POST("/delete", splashScreen.DeleteSplashScreenControl)      //删除
}
