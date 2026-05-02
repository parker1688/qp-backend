// The build tag makes sure the stub is not built in the final build.

package fcSiteMessage

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcSiteMessage/save
/*func SaveFcSiteMessageControl(c *gin.Context) {
	var req vo.SiteMsgSaveReq
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(req)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	if req.MsgType == 0 {
		req.MsgType = 1 // 默认未一般消息
	}

	createBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		createBy = userInfo.(*dos.AdminUser).UserName

	}
	nowTime := automaticType.Now()

	//global.G_LOG.Infof("SaveFcSiteMessageControl----------------------------------0: %v", req.NotifyType)
	jsonp := dos.FcSiteMessage{}
	tool.JsonMapper(&req, &jsonp)
	jsonp.CreateBy = createBy
	jsonp.UpdateBy = createBy
	jsonp.CreateTime = nowTime
	jsonp.UpdateTime = nowTime

	//global.G_LOG.Infof("SaveFcSiteMessageControl---------------------------------1: %v", jsonp.NotifyType)

	data, err := modules.SaveFcSiteMessage(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 如果是发送给所有商户的所有人
	msgDataArr := []*vo.SiteMsgVO{}
	if jsonp.NotifyType == 1 {
		if jsonp.MerchantCode == "" { // 发送给所有商户
			queryParam := &dos.FcMerchant{}
			queryParam.Status = 1
			merchantArr := modules.FindByKeyFcMerchant(queryParam, nil)
			for _, v := range merchantArr {
				// 将消息 push 到 kafka
				msgData := vo.SiteMsgVO{}
				tool.JsonMapper(&jsonp, &msgData)
				msgData.MsgId = jsonp.Id
				msgData.MerchantCode = v.MerchantCode
				msgData.MsgIdType = jsonp.MsgIdType
				msgDataArr = append(msgDataArr, &msgData)
			}
		} else { // 只发送给某一个商户的所有人
			msgData := vo.SiteMsgVO{}
			tool.JsonMapper(&jsonp, &msgData)
			msgData.MsgId = jsonp.Id
			msgDataArr = append(msgDataArr, &msgData)
		}
	} else if req.UserIds != "" { // 表示单个发送, user
		userIdArr := strings.Split(req.UserIds, ",")
		for _, v := range userIdArr {
			userName := ""
			err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", v).Pluck("user_name", &userName).Error
			if err != nil {
				global.G_LOG.Errorf("SaveFcSiteMessageControl get user name err: %s", err.Error())
				continue
			}

			msgData := vo.SiteMsgVO{}
			tool.JsonMapper(&jsonp, &msgData)
			msgData.UserId = v
			msgData.UserName = userName
			msgData.MsgId = jsonp.Id
			msgDataArr = append(msgDataArr, &msgData)
		}
	} else {
		// 将消息 push 到 kafka
		msgData := vo.SiteMsgVO{}
		tool.JsonMapper(&jsonp, &msgData)
		msgData.MsgId = jsonp.Id
		msgDataArr = append(msgDataArr, &msgData)
	}

	for i := range msgDataArr {
		msgData := msgDataArr[i]
		msgDataBytes, err := tool.JsonMarshal(msgData)
		err = channelData.SendUserSiteMsgData(string(msgDataBytes))
		if err != nil {
			global.G_LOG.Errorf("SendUserSiteMsgData merchant: %v msg: %s, err: %s", msgData.MerchantCode, string(msgDataBytes), err.Error())
		}
	}

	response.SuccessJSON(c, data)
}*/

// api: api/fcSiteMessage/save 人工邮件发送
func SaveFcSiteMessageControl(c *gin.Context) {
	var jsonp dos.Mail
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	if len(jsonp.UserIds) > 0 {
		ids := strings.Split(jsonp.UserIds, ",")
		data := []dos.FcUserMaterial{}
		err = global.G_DB.Model(&dos.FcUserMaterial{}).Select("user_id").
			Where("user_id in ?", ids).Find(&data).Error
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}

		includeMap := map[string]int{}
		for _, v := range data {
			includeMap[v.UserId] = 1
		}

		errIds := []string{}
		for _, id := range ids {
			if _, ok := includeMap[id]; !ok {
				errIds = append(errIds, id)
			}
		}

		if len(errIds) > 0 {
			response.FailErrJSON(c, response.ERROR_PARAMETER,
				fmt.Sprintf("存在不存在的玩家id: %s", strings.Join(errIds, ",")))
			return
		}
	}

	var createBy string
	var updateBy string
	userInfo, ok := c.Get("UserInfo")
	if ok {
		createBy = userInfo.(*dos.AdminUser).UserName
		updateBy = createBy
	}

	mailRecord := dos.Mail{
		MerchantCode: jsonp.MerchantCode,
		UserIds:      jsonp.UserIds,
		Type:         jsonp.Type,
		Title:        jsonp.Title,
		Content:      jsonp.Content,
		IsPopup:      0,
		IsKeep:       1,
		CreateTime:   automaticType.Now(),
		CreateBy:     createBy,
		UpdateTime:   automaticType.Now(),
		UpdateBy:     updateBy,
		Status:       0,
	}

	if jsonp.MerchantCode == "" || jsonp.UserIds == "" {
		// 全部商户的全部玩家（用户登录或打开邮箱动态处理）
		mailRecord.Status = 1
		err := modules.SaveMail(&mailRecord)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}
	} else {
		/*if jsonp.UserIds == "" {
			// 指定商户的全部玩家（用户登录或打开邮箱动态处理）
			mailRecord.Status = 1
			err := modules.SaveMail(&mailRecord)
			if err != nil {
				response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
				return
			}
		} else {*/
		// 指定商户的指定玩家（直接发用户邮箱）
		err := modules.SaveMail(&mailRecord)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}

		go func() {
			mails := []dos.FcUserMail{}
			userIds := strings.Split(jsonp.UserIds, ",")
			for _, userId := range userIds {
				mails = append(mails, dos.FcUserMail{
					MsgId:        mailRecord.Id,
					UserId:       userId,
					MerchantCode: jsonp.MerchantCode,
					Type:         enmus.MailType_Manual,
					Title:        jsonp.Title,
					Content:      jsonp.Content,
					IsPopup:      0,
					IsKeep:       1,
					CreateTime:   automaticType.Now(),
					ReadStatus:   enmus.MailStats_Unread,
					DelStatus:    enmus.MailDelStats_No,
				})
			}

			err = modules.SaveUserMailMulit(mails)
			if err != nil {
				global.G_LOG.Errorf("[SaveUserMailMulit] save user mail failed: data=%+v, err=%s",
					mails, err.Error())
			}
		}()
		//}
	}

	response.SuccessJSON(c, true)
}

// api: api/fcSiteMessage/findPage
func FindPageFcSiteMessageControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcSiteMessage
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.MsgType = tool.Atoi(c.DefaultQuery("msg_type", ""))
	jsonp.NotifyType = tool.Atoi(c.DefaultQuery("notify_type", ""))
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcSiteMessage(jsonp.PageNo, jsonp.PageSize, &jsonp.FcSiteMessage, &jsonp.PageTimeQuery)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcSiteMessage/findByKey
func FindByKeyFcSiteMessageControl(c *gin.Context) {
	var jsonp dos.FcSiteMessage
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	data := modules.FindByKeyFcSiteMessage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteMessage/update
func UpdateFcSiteMessageControl(c *gin.Context) {
	var jsonp dos.FcSiteMessage
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	data := modules.UpdateFcSiteMessage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteMessage/delete
func DeleteFcSiteMessageControl(c *gin.Context) {
	var jsonp dos.FcSiteMessage
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	data := modules.DeleteFcSiteMessage(&jsonp)
	response.SuccessJSON(c, data)
}

/////////////////////////////////////////////////

// api: api/fcSiteMessage/system/findPage 系统邮列表接口
func FindPageSystemMailControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.Mail
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Type = tool.Atoi(c.DefaultQuery("type", "-1"))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")

	data, total := modules.FindPageMail(jsonp.PageNo, jsonp.PageSize, &jsonp.Mail)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcSiteMessage/system/save 系统邮新增接口
func SaveSystemMailControl(c *gin.Context) {
	var jsonp dos.Mail
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 对用户id组格式判断
	if len(jsonp.UserIds) > 0 {
		ids := strings.Split(jsonp.UserIds, ",")
		data := []dos.FcUserMaterial{}
		err = global.G_DB.Model(&dos.FcUserMaterial{}).Select("user_id").
			Where("user_id in ?", ids).Find(&data).Error
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}

		includeMap := map[string]int{}
		for _, v := range data {
			includeMap[v.UserId] = 1
		}

		errIds := []string{}
		for _, id := range ids {
			if _, ok := includeMap[id]; !ok {
				errIds = append(errIds, id)
			}
		}

		if len(errIds) > 0 {
			response.FailErrJSON(c, response.ERROR_PARAMETER,
				fmt.Sprintf("存在不存在的玩家id: %s", strings.Join(errIds, ",")))
			return
		}
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateTime = automaticType.Now()
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateTime = jsonp.CreateTime
		jsonp.UpdateBy = jsonp.CreateBy
	}

	err = modules.SaveMail(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	response.SuccessJSON(c, true)
}

// api: api/fcSiteMessage/system/update 系统邮更新接口
func UpdateSystemMailControl(c *gin.Context) {
	var jsonp dos.Mail
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 对用户id组格式判断
	if len(jsonp.UserIds) > 0 {
		ids := strings.Split(jsonp.UserIds, ",")
		data := []dos.FcUserMaterial{}
		err = global.G_DB.Model(&dos.FcUserMaterial{}).Select("user_id").
			Where("user_id in ?", ids).Find(&data).Error
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}

		includeMap := map[string]int{}
		for _, v := range data {
			includeMap[v.UserId] = 1
		}

		errIds := []string{}
		for _, id := range ids {
			if _, ok := includeMap[id]; !ok {
				errIds = append(errIds, id)
			}
		}

		if len(errIds) > 0 {
			response.FailErrJSON(c, response.ERROR_PARAMETER,
				fmt.Sprintf("存在不存在的玩家id: %s", strings.Join(errIds, ",")))
			return
		}
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	err = modules.UpdateMail(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	response.SuccessJSON(c, true)
}

// api: api/fcSiteMessage/system/delete 系统邮件删除接口
func DeleteSystemMailControl(c *gin.Context) {
	var jsonp struct {
		Id string `json:"id" form:"id" uri:"id"`
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	err = modules.DelMail(jsonp.Id)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	response.SuccessJSON(c, true)
}

/////////////////////////////////////////////////

// api: api/fcSiteMessage/mail/findPage 邮件列表接口
func FindPageMailControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcUserMail
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Type = tool.Atoi(c.DefaultQuery("type", "-1"))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.DelStatus = -1
	jsonp.IsPopup = -1
	jsonp.ReadStatus = -1

	data, total := modules.FindPageFcUserMail(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserMail, &jsonp.PageTimeQuery, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcSiteMessage/mail/delete 邮件列表删除接口
func DeleteMailControl(c *gin.Context) {
	var jsonp struct {
		Ids []string `json:"ids" form:"ids" uri:"ids" ` // 消息IDS
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	ret := modules.SetUserMailDataByIds(jsonp.Ids, map[string]interface{}{
		"del_status": enmus.MailDelStats_Destroy, // 将删除状态改为销毁
	})
	if !ret {
		response.FailErrJSON(c, response.ERROR_SERVER, "删除失败")
		return
	}

	response.SuccessJSON(c, true)
}
