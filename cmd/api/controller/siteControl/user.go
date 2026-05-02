package siteControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserNotifyMessage(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	data := modules.FindByKeyFcUserNotify(&dos.FcUserNotify{
		UserId: userInfo.UserId,
		//Language: language,
	})
	var newData []*vo.UserNotifyMessageResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}

func GetUserNotifyMessageDetail(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	id := c.DefaultQuery("id", "0")
	m := &dos.FcUserNotify{
		UserId: userInfo.UserId,
	}
	m.Id = id
	data := modules.FindByKeyFcUserNotifyFirst(m)
	newData := &vo.UserNotifyMessageDetailResp{
		Title:      data.Title,
		Content:    data.Content,
		CreateTime: data.CreateTime,
		Status:     data.Status,
	}

	response.SuccessJSON(c, newData)
}

func UserNotifyMessageUnreadCount(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var count int64
	err := global.G_DB.WithContext(ctx).Model(&dos.FcUserNotify{}).Where("user_id = ? and status = 0", userInfo.UserId).Count(&count).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	response.SuccessJSON(c, count)
}

func UserNotifyMessageConfirm(c *gin.Context) {
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	err = global.G_DB.WithContext(ctx).Model(&dos.FcUserNotify{}).Where("id in ? and user_id = ?", jsonp.Ids, userInfo.UserId).Update("status", 1).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	response.SuccessJSON(c, true)
}

func UserNotifyMessageDelete(c *gin.Context) {
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	err = global.G_DB.WithContext(ctx).Where("id in ? and user_id = ?", jsonp.Ids, userInfo.UserId).Delete(&dos.FcUserNotify{}).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	response.SuccessJSON(c, true)
}

func GetBulletin(c *gin.Context) {
	merchantCode := modules.GetAgentDomainMerchantCodeByHeader(c)
	data := modules.FindByKeyFcBulletin(&dos.FcBulletin{
		MerchantCode: merchantCode,
		IsDisplay:    1,
	})
	var newData []*vo.BulletinResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}
