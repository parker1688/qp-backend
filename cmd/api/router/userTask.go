package router

import (
	"bootpkg/cmd/api/controller/userTaskControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, userTaskRouter)
}

func userTaskRouter() {
	r := routers.Group("/api/userTask").Use(handler.AuthApiMiddleware())
	r.GET("/list", userTaskControl.UserTaskList)      // 用户任务列表
	r.POST("/reward", userTaskControl.UserTaskReward) // 用户任务奖励领取
}
