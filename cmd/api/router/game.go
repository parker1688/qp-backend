package router

import (
	"bootpkg/cmd/api/controller/gameControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, gameRouter)
}

func gameRouter() {

	r := routers.Group("/api/game/user").Use(handler.AuthApiMiddleware())
	r.GET("/history/list", gameControl.HistoryGame)     //历史游戏
	r.POST("/history/add", gameControl.HistoryGameAdd)  //添加历史游戏
	r.GET("/history/clear", gameControl.DelHistoryGame) //清空历史游戏

	r.POST("/collect/add", gameControl.CollectGameSlots)    //收藏游戏
	r.GET("/collect/list", gameControl.GetCollectGameSlots) //获取收藏游戏
	r.GET("/collect/del", gameControl.DelCollectGameSlots)  //删除收藏的游戏

	slots := routers.Group("/api/game/slots")   //.Use(handler.UrlLocalCacheJson())
	slots.GET("/get", gameControl.GetGameSlots) //获取电子游戏
	//slots.GET("/get", gameControl.GetAmbGameSlots) //获取电子游戏
	slots.GET("/query", gameControl.QueryGameSlots) //搜索电子游戏
}
