package router

import (
	"bootpkg/cmd/web/controller/fcUserLevelWeekSetting"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcUserLevelWeekSettingRouter)
}

func fcUserLevelWeekSettingRouter() {
	r := routers.Group("/api/fcUserLevelWeekSetting").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcUserLevelWeekSetting.FindPageFcUserLevelWeekSettingControl)   //查询分页
	r.GET("/findByKey", fcUserLevelWeekSetting.FindByKeyFcUserLevelWeekSettingControl) //查询主键
	r.POST("/save", fcUserLevelWeekSetting.SaveFcUserLevelWeekSettingControl)          //保存
	r.POST("/update", fcUserLevelWeekSetting.UpdateFcUserLevelWeekSettingControl)      //修改
	r.POST("/delete", fcUserLevelWeekSetting.DeleteFcUserLevelWeekSettingControl)      //删除
}
