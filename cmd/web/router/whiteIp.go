package router

import (
	"bootpkg/cmd/web/controller/whiteIp"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, whiteIpRouter)
}

func whiteIpRouter() {
	r := routers.Group("/api/whiteIp").Use(handler.AuthMiddleware())

	r.GET("/findPage", whiteIp.FindPageWhiteIpControl)   //查询分页
	r.GET("/findByKey", whiteIp.FindByKeyWhiteIpControl) //查询主键
	r.POST("/save", whiteIp.SaveWhiteIpControl)          //保存
	r.POST("/update", whiteIp.UpdateWhiteIpControl)      //修改
	r.POST("/delete", whiteIp.DeleteWhiteIpControl)      //删除
}
