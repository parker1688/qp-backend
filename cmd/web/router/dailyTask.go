package router

import (
	"bootpkg/cmd/web/controller/dailyTask"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, dailyTaskRouter)
}

func dailyTaskRouter() {
	r := routers.Group("/api/dailyTask").Use(handler.AuthMiddleware())

	r.GET("/findPage", dailyTask.FindPageDailyTaskControl)   //查询分页
	r.GET("/findByKey", dailyTask.FindByKeyDailyTaskControl) //查询主键
	r.POST("/save", dailyTask.SaveDailyTaskControl)          //保存
	r.POST("/update", dailyTask.UpdateDailyTaskControl)      //修改
	r.POST("/delete", dailyTask.DeleteDailyTaskControl)      //删除
}
