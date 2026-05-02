package router

import (
	"bootpkg/cmd/api/controller/siteControl"
	"bootpkg/cmd/api/handler"
)

func init() {
	routersFun = append(routersFun, SiteRouter)
}

func SiteRouter() {
	r := routers.Group("/api/site/").Use(handler.UrlLocalCacheJson())
	r.GET("/announcement/top", siteControl.GetAnnouncementTop) //获取公告滚动播放
	r.GET("/banner/list", siteControl.GetBanners)              //获取Banner图片

	r.GET("/get/siteLink", siteControl.SiteBaseLink) //获取基本配置链接

	r.GET("/get/clientSettings", siteControl.ClientSettings)   //获取合营计划联系信息
	r.GET("/notifyMarquee/list", siteControl.GetNotifyMarquee) //跑马灯

	rConfirm := routers.Group("/api/site/user").Use(handler.AuthApiMiddleware())
	rConfirm.GET("/notify/count", siteControl.GetSiteUserNotifyCount)
	//rConfirm.GET("/notify/count", siteControl.GetSiteNotifyCount)             //获取需要阅读消息总数
	rConfirm.GET("/announcement/list", siteControl.GetUserAnnouncement) //获取通知列表
	//rConfirm.GET("/announcement/list", siteControl.GetAnnouncement)           //获取公告列表
	rConfirm.GET("/announcement/detail", siteControl.GetUserAnnouncementDetail) //获取公告详情
	//rConfirm.GET("/announcement/detail", siteControl.GetAnnouncementDetail)   //获取公告详情
	rConfirm.POST("/announcement/confirm", siteControl.GetUserAnnouncementConfirm) //公告阅读
	//rConfirm.GET("/announcement/confirm", siteControl.GetAnnouncementConfirm) //公告阅读
	rConfirm.POST("/announcement/del", siteControl.GetUserAnnouncementDel) //删除阅读公告
	//rConfirm.GET("/announcement/del", siteControl.GetAnnouncementDel)         //删除阅读公告

	rConfirm.GET("/notify/list", siteControl.GetUserNotifyMessage)           //获取通知列表
	rConfirm.GET("/notify/detail", siteControl.GetUserNotifyMessageDetail)   //获取通知详情
	rConfirm.GET("/notify/unread", siteControl.UserNotifyMessageUnreadCount) //获取用户未读取的通知总数
	rConfirm.POST("/notify/confirm", siteControl.UserNotifyMessageConfirm)   //消息阅读
	rConfirm.POST("/notify/del", siteControl.UserNotifyMessageDelete)        //删除消息

	rBulletin := routers.Group("/api/site/bulletin") // 首页公告
	rBulletin.GET("/list", siteControl.GetBulletin)

	rMsg := routers.Group("/api/site/message").Use(handler.AuthApiMiddleware()) // 站内信
	rMsg.GET("/list", siteControl.GetSiteMsgList)                               // 站内信列表
	rMsg.POST("/update/read", siteControl.UpdateSiteMsgReadStatus)              // 更新站内信为已读
	rMsg.POST("/del", siteControl.DelSiteMsg)                                   // 删除站内信
}
