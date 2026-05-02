package router

import (
	"bootpkg/cmd/web/controller/fcUserRemark"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserRemarkRouter)
}

func fcUserRemarkRouter() {
	r := routers.Group("/api/fcUserRemark").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserRemark.FindPageFcUserRemarkControl)   //查询分页
	r.GET("/findByKey", fcUserRemark.FindByKeyFcUserRemarkControl) //查询主键
	r.POST("/save", fcUserRemark.SaveFcUserRemarkControl)          //保存
	r.POST("/update", fcUserRemark.UpdateFcUserRemarkControl)      //修改
	r.POST("/delete", fcUserRemark.DeleteFcUserRemarkControl)      //删除
}
