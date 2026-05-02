package siteControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/setnotify"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	user_site_notify_unread = 1
	user_site_notify_readed = 2
	user_site_notify_delete = 3
)

/*
获取公告列表
*/
func GetUserAnnouncement(c *gin.Context) {
	jsonp := struct {
		ClassType int `json:"class_type" form:"class_type" uri:"class_type" ` // 平台公告分类类型 1. 公告 2.  赛事  3. 充提
	}{}
	jsonp.ClassType = tool.Atoi(c.DefaultQuery("class_type", "1"))
	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	isLook := global.G_REDIS.SetNX(context.Background(), "GetUserAnnouncement", "1", 10*time.Second).Val()
	if isLook {
		setnotify.SynUserSiteNotify(userInfo, clientType, language)
	}
	data := modules.FindByKeyFcUserSiteNotify(&dos.FcUserSiteNotify{
		NotifyType: clientType,
		Language:   language,
		//MerchantCode: merchantCode,
		ClassType: jsonp.ClassType,
		UserId:    userInfo.UserId,
	})
	var newData []*vo.AnnouncementResp
	for _, v := range data {
		if v.Status == user_site_notify_delete {
			continue
		}
		newData = append(newData, &vo.AnnouncementResp{
			Id:         v.Id,
			Title:      v.Title,
			TitleImg:   v.TitleImg,
			ClassType:  v.ClassType,
			CreateTime: v.CreateTime,
			Content:    v.Content,
			Status:     v.Status,
		})

	}
	response.SuccessJSON(c, newData)
}

func GetUserAnnouncementDetail(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	m := &dos.FcUserSiteNotify{
		NotifyType: clientType,
		Language:   language,
		//MerchantCode: merchantCode,
	}
	m.Id = id
	data := modules.FindByKeyFcUserSiteNotifyFirst(m)
	newData := &vo.AnnouncementDetailResp{
		Id:         data.Id,
		Title:      data.Title,
		Content:    data.Content,
		CreateTime: data.CreateTime,
	}

	response.SuccessJSON(c, newData)
}

func GetUserAnnouncementConfirm(c *gin.Context) {
	jsonp := struct {
		Ids []string `json:"ids" form:"ids" uri:"ids" ` // 阅读消息
	}{}

	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	for _, id := range jsonp.Ids {
		m := &dos.FcUserSiteNotify{}
		m.Id = id
		row := modules.FindByKeyFcUserSiteNotifyFirst(m)
		if len(row.Id) > 0 && userInfo.UserId == row.UserId {
			row.Status = user_site_notify_readed
			modules.UpdateFcUserSiteNotify(row)
		}

	}
	response.SuccessJSON(c, true)
}

func GetUserAnnouncementDel(c *gin.Context) {
	jsonp := struct {
		Ids []string `json:"ids" form:"ids" uri:"ids" ` // 消息IDS
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	for _, id := range jsonp.Ids {
		m := &dos.FcUserSiteNotify{}
		m.Id = id
		row := modules.FindByKeyFcUserSiteNotifyFirst(m)
		if len(row.Id) > 0 && userInfo.UserId == row.UserId {
			row.Status = user_site_notify_delete
			modules.UpdateFcUserSiteNotify(row)
		}
	}
	response.SuccessJSON(c, true)
}

func GetSiteUserNotifyCount(c *gin.Context) {
	jsonp := struct {
		ClassType int `json:"class_type" form:"class_type" uri:"class_type" ` // 平台公告分类类型 1. 公告 2.  赛事  3. 充提
	}{}
	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	setnotify.SynUserSiteNotify(userInfo, "", "")
	data := modules.FindByKeyFcUserSiteNotify(&dos.FcUserSiteNotify{
		NotifyType: clientType,
		Language:   language,
		//MerchantCode: merchantCode,
		ClassType: jsonp.ClassType,
		UserId:    userInfo.UserId,
		Status:    user_site_notify_unread,
	})

	//1. 公告 2.  赛事  3. 充提  9. 用户通知消息
	unReadCount := make(map[string]int64, 5)
	unReadCount["0"] = 0 //防止异常
	unReadCount["1"] = 0
	unReadCount["2"] = 0
	unReadCount["3"] = 0
	unReadCount["9"] = int64(len(data)) //用户通知消息未读总数

	for _, v := range data {
		classType := tool.String(v.ClassType)
		count, ok := unReadCount[classType]
		if ok {
			unReadCount[classType] = count + 1
		} else {
			unReadCount[classType] = 1
		}
	}
	response.SuccessJSON(c, unReadCount)
}
