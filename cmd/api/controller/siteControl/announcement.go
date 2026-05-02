package siteControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"github.com/gin-gonic/gin"
)

func GetAnnouncement(c *gin.Context) {
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
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	data := modules.FindByKeyFcSiteNotify(&dos.FcSiteNotify{
		NotifyType:   clientType,
		Language:     language,
		MerchantCode: merchantCode,
		ClassType:    jsonp.ClassType,
	})
	var newData []*vo.AnnouncementResp

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	//查询已读的消息
	readIds := modules.FindByKeyFcSiteNotifyRead(&dos.FcSiteNotifyRead{
		UserId: userInfo.UserId,
	})
	readIdsMap := make(map[string]int, len(readIds))
	for _, v := range readIds {
		readIdsMap[v.SiteNotifyId] = v.Status
	}

	for _, v := range data {
		status, ok := readIdsMap[v.Id]
		if ok && status == 2 { //已经删除
			continue
		}
		var readStatus = 0     //阅读状态
		if ok && status == 1 { //阅读
			readStatus = 1
		}
		newData = append(newData, &vo.AnnouncementResp{
			Id:         v.Id,
			Title:      v.Title,
			ClassType:  v.ClassType,
			CreateTime: v.CreateTime,
			Status:     readStatus,
		})

	}
	response.SuccessJSON(c, newData)
}

func GetAnnouncementTop(c *gin.Context) {
	classType := tool.Atoi(c.DefaultQuery("class_type", "1")) // 平台公告分类类型 1. 公告 2.  赛事  3. 充提
	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	data := modules.FindByKeyFcSiteNotify(&dos.FcSiteNotify{
		//NotifyType: clientType,
		//Language:     language,
		MerchantCode: merchantCode,
		ClassType:    classType, //公告
	})
	newData := make([]*vo.AnnouncementResp, 0, len(data))
	for _, v := range data {
		newData = append(newData, &vo.AnnouncementResp{
			Id:         v.Id,
			Title:      v.Title,
			TitleImg:   v.TitleImg,
			ClassType:  v.ClassType,
			CreateTime: v.CreateTime,
			Status:     1,
			//Content:    v.Content,
		})
	}
	//tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}

func GetAnnouncementDetail(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	m := &dos.FcSiteNotify{
		NotifyType:   clientType,
		Language:     language,
		MerchantCode: merchantCode,
	}
	m.Id = id
	data := modules.FindByKeyFcSiteNotifyFirst(m)
	var newData *vo.AnnouncementDetailResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}

func GetAnnouncementConfirm(c *gin.Context) {
	jsonp := struct {
		Ids []string `json:"ids" form:"ids" uri:"ids" ` // 阅读消息
	}{}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	for _, v := range jsonp.Ids {
		modules.DeleteFcSiteNotifyRead(&dos.FcSiteNotifyRead{SiteNotifyId: v, UserId: userInfo.UserId})
		modules.SaveFcSiteNotifyRead(&dos.FcSiteNotifyRead{
			SiteNotifyId: v,
			UserId:       userInfo.UserId,
			Status:       1,
		})
	}
	response.SuccessJSON(c, true)
}

func GetAnnouncementDel(c *gin.Context) {
	jsonp := struct {
		Ids []string `json:"ids" form:"ids" uri:"ids" ` // 消息IDS
	}{}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	for _, v := range jsonp.Ids {
		modules.DeleteFcSiteNotifyRead(&dos.FcSiteNotifyRead{SiteNotifyId: v, UserId: userInfo.UserId})
		modules.SaveFcSiteNotifyRead(&dos.FcSiteNotifyRead{
			SiteNotifyId: v,
			UserId:       userInfo.UserId,
			Status:       2,
		})
	}
	response.SuccessJSON(c, true)
}

func GetSiteNotifyCount(c *gin.Context) {
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	//1. 公告 2.  赛事  3. 充提  9. 用户通知消息
	unReadCount := make(map[string]int64, 5)
	unReadCount["0"] = 0 //防止异常
	unReadCount["1"] = 0
	unReadCount["2"] = 0
	unReadCount["3"] = 0
	unReadCount["9"] = 0 //用户通知消息未读总数

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	var userCount int64
	global.G_DB.Model(&dos.FcUserNotify{}).Where("user_id = ? and status = 0", userInfo.UserId).Count(&userCount)
	unReadCount["9"] = userCount

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code

	//用户公告消息
	data := modules.FindByKeyFcSiteNotify(&dos.FcSiteNotify{
		Language:     language,
		MerchantCode: merchantCode,
	})
	readIds := modules.FindByKeyFcSiteNotifyRead(&dos.FcSiteNotifyRead{
		UserId: userInfo.UserId,
	})
	readIdsMap := make(map[string]struct{}, len(readIds))
	for _, v := range readIds {
		readIdsMap[v.SiteNotifyId] = struct{}{}
	}
	for _, v := range data {
		_, ok := readIdsMap[v.Id]
		if ok {
			continue
		}
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
