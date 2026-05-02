package siteControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	commonResp "bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// api: api/site/message/list 站内信列表
func GetSiteMsgList(c *gin.Context) {
	jsonp := struct {
		commonResp.PageTimeQuery
		dos.FcUserMail
	}{}
	jsonp.PageTimeQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageTimeQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Type = tool.Atoi(c.DefaultQuery("type", "1"))
	jsonp.IsPopup = tool.Atoi(c.DefaultQuery("is_popup", "-1"))
	//jsonp.MsgType = tool.Atoi(c.DefaultQuery("msg_type", ""))
	jsonp.ReadStatus = tool.Atoi(c.DefaultQuery("read_status", "-1"))
	//jsonp.DelStatus = enmus.MailDelStats_No // 未删除

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	jsonp.UserId = userInfo.UserId
	jsonp.MerchantCode = userInfo.MerchantCode

	if jsonp.PageTimeQuery.PageNo == 1 {
		modules.DoUserMailAction(userInfo)
	}

	data, total := modules.FindPageFcUserMail(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserMail, &jsonp.PageTimeQuery, nil)
	newData := make([]*vo.FcUserMailResp, len(data))
	tool.JsonMapper(&data, &newData)

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, newData)
}

// 站内信已读
func UpdateSiteMsgReadStatus(c *gin.Context) {
	var jsonp dos.FcUserMail
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	upMail := dos.FcUserMail{}
	err = global.G_DB.WithContext(ctx).Model(&dos.FcUserMail{}).Select("is_keep").
		Where("user_id = ? AND id = ?", userInfo.UserId, jsonp.Id).Take(&upMail).Error
	if err != nil {
		global.G_LOG.Errorf("[UpdateSiteMsgReadStatus] Find user mail failed: userId=%s, mailId=%s, err=%s",
			userInfo.UserId, jsonp.Id, err.Error())
		response.FailErrJSON(c, response.ERROR_PARAMETER, "更新已读失败")
		return
	}

	if upMail.IsKeep == 0 { // 不保留则直接删除
		jsonp.UserId = userInfo.UserId
		modules.DelUserMail(&jsonp)
		response.SuccessMsgJSON(c, struct{}{}, "success")
		return
	}

	err = global.G_DB.WithContext(ctx).Model(&dos.FcUserMail{}).
		Where("user_id = ? AND id = ?", userInfo.UserId, jsonp.Id).
		Updates(map[string]interface{}{
			"read_status": enmus.MailStats_Readed,
			"is_popup":    0,
		}).Error
	if err != nil {
		global.G_LOG.Errorf("[UpdateSiteMsgReadStatus] Update user mail failed: userId=%s, mailId=%s, err=%s",
			userInfo.UserId, jsonp.Id, err.Error())
		response.FailErrJSON(c, response.ERROR_PARAMETER, "更新状态失败")
		return
	}

	response.SuccessMsgJSON(c, struct{}{}, "success")

	/*row := modules.FindByKeyFcUserSiteMessageFirst(&jsonp)
	if row.Id == "" {
		tmpStr := fmt.Sprintf("id: %s query not exist", row.Id)
		global.G_LOG.Errorf(tmpStr)
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	if row.UserId != userInfo.UserId {
		tmpStr := fmt.Sprintf("userName: %v userId: %s param UserId: %v not match", userInfo.UserName, row.UserId, userInfo.UserId)
		global.G_LOG.Errorf(tmpStr)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "user id not match")
		return
	}

	updataMap := map[string]interface{}{}
	updataMap["read_status"] = 2
	updataMap["update_time"] = automaticType.Now()
	updataMap["update_by"] = "userSelf"

	err = global.G_DB.Model(&dos.FcUserSiteMessage{}).Where("id = ?", jsonp.Id).Updates(updataMap).Error
	if err != nil {
		global.G_LOG.Error(err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	response.SuccessMsgJSON(c, struct{}{}, "success")
	return*/
}

// 站内信删除
func DelSiteMsg(c *gin.Context) {
	var jsonp dos.FcUserMail
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	jsonp.UserId = userInfo.UserId

	ret := modules.DelUserMail(&jsonp)
	if !ret {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "删除邮件失败")
		return
	}

	response.SuccessMsgJSON(c, struct{}{}, "success")

	/*
	   queryParam := dos.FcUserSiteMessage{}

	   	queryParam.Id = jsonp.Id
	   	row := modules.FindByKeyFcUserSiteMessageFirst(&queryParam)
	   	if row.Id == "" {
	   		tmpStr := fmt.Sprintf("id: %s query not exist", row.Id)
	   		global.G_LOG.Errorf(tmpStr)
	   		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
	   		return
	   	}

	   	if row.UserId != userInfo.UserId {
	   		tmpStr := fmt.Sprintf("userName: %v userId: %s param UserId: %v not match", userInfo.UserName, row.UserId, userInfo.UserId)
	   		global.G_LOG.Errorf(tmpStr)
	   		response.FailErrJSON(c, response.ERROR_PARAMETER, "user id not match")
	   		return
	   	}

	   	//  如果是人工全局消息,则只更新
	   	if row.MsgIdType == 1 {
	   		updataMap := map[string]interface{}{}
	   		updataMap["del_status"] = 2
	   		updataMap["update_time"] = automaticType.Now()
	   		updataMap["update_by"] = "userSelf"

	   		err = global.G_DB.Model(&dos.FcUserSiteMessage{}).Where("id = ?", jsonp.Id).Updates(updataMap).Error
	   		if err != nil {
	   			global.G_LOG.Error(err)
	   			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
	   			return
	   		}
	   	} else {
	   		modules.DeleteFcUserSiteMessage(&jsonp)
	   	}

	   	response.SuccessMsgJSON(c, struct{}{}, "success")
	   	return
	*/
}
