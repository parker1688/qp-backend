package router

import (
	"bootpkg/cmd/web/controller/dictsDetail"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, dictsDetailRouter)
}

func dictsDetailRouter() {
	r := routers.Group("/api/dictsDetail").Use(handler.AuthTokenMiddleware())
	r.GET("/findPage", dictsDetail.FindPageDictsDetailControl)   //查询分页
	r.GET("/findByKey", dictsDetail.FindByKeyDictsDetailControl) //查询主键
	r.POST("/save", dictsDetail.SaveDictsDetailControl)          //保存
	r.POST("/update", dictsDetail.UpdateDictsDetailControl)      //修改
	r.POST("/delete", dictsDetail.DeleteDictsDetailControl)      //删除
	r.GET("/findByAll", dictsDetail.FindByKeyDictsDetailAll)     //查询所有符合条件的数据
}
