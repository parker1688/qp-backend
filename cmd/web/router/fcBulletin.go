package router

import (
	"bootpkg/cmd/web/controller/fcBulletin"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcBulletinRouter)
}

func fcBulletinRouter() {
	r := routers.Group("/api/fcBulletin").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcBulletin.FindPageFcBulletinControl)   //查询分页
	r.GET("/findByKey", fcBulletin.FindByKeyFcBulletinControl) //查询主键
	r.POST("/save", fcBulletin.SaveFcBulletinControl)          //保存
	r.POST("/update", fcBulletin.UpdateFcBulletinControl)      //修改
	r.POST("/delete", fcBulletin.DeleteFcBulletinControl)      //删除
}
