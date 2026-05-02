package router

import (
	"bootpkg/cmd/web/controller/fcChannelBankImg"
	"bootpkg/cmd/web/handler"
)

func init() {
	routersFun = append(routersFun, fcChannelBankImgRouter)
}

func fcChannelBankImgRouter() {
	r := routers.Group("/api/fcChannelBankImg").Use(handler.AuthMiddleware())

	r.GET("/findPage", fcChannelBankImg.FindPageFcChannelBankImgControl)   //查询分页
	r.GET("/findByKey", fcChannelBankImg.FindByKeyFcChannelBankImgControl) //查询主键
	r.POST("/save", fcChannelBankImg.SaveFcChannelBankImgControl)          //保存
	r.POST("/update", fcChannelBankImg.UpdateFcChannelBankImgControl)      //修改
	r.POST("/delete", fcChannelBankImg.DeleteFcChannelBankImgControl)      //删除
}
