package router

import (
	"bootpkg/cmd/web/controller/fcLoginLog"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcLoginLogRouter)
}

func fcLoginLogRouter() {
	r := routers.Group("/api/fcLoginLog").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcLoginLog.FindPageFcLoginLogControl)   //查询分页
	r.GET("/findByKey", fcLoginLog.FindByKeyFcLoginLogControl) //查询主键
	r.POST("/save", fcLoginLog.SaveFcLoginLogControl)          //保存
	r.POST("/update", fcLoginLog.UpdateFcLoginLogControl)      //修改
	r.POST("/delete", fcLoginLog.DeleteFcLoginLogControl)      //删除
}
