// The build tag makes sure the stub is not built in the final build.

package fcUserSiteMessage

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserSiteMessage/save
func SaveFcUserSiteMessageControl(c *gin.Context) {
	var jsonp dos.FcUserSiteMessage
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, err := modules.SaveFcUserSiteMessage(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcUserSiteMessage/findPage
func FindPageFcUserSiteMessageControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcUserSiteMessage
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.MsgType = tool.Atoi(c.DefaultQuery("msg_type", ""))
	jsonp.MsgIdType = tool.Atoi(c.DefaultQuery("msg_id_type", ""))
	jsonp.NotifyType = tool.Atoi(c.DefaultQuery("notify_type", ""))
	jsonp.ReadStatus = tool.Atoi(c.DefaultQuery("read_status", ""))
	jsonp.DelStatus = tool.Atoi(c.DefaultQuery("del_status", ""))
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
	data, total := modules.FindPageFcUserSiteMessage(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserSiteMessage, &jsonp.PageTimeQuery, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserSiteMessage/findByKey
func FindByKeyFcUserSiteMessageControl(c *gin.Context) {
	var jsonp dos.FcUserSiteMessage
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
	data := modules.FindByKeyFcUserSiteMessage(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcUserSiteMessage/update
func UpdateFcUserSiteMessageControl(c *gin.Context) {
	var jsonp dos.FcUserSiteMessage
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcUserSiteMessage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserSiteMessage/delete
func DeleteFcUserSiteMessageControl(c *gin.Context) {
	var jsonp dos.FcUserSiteMessage
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

	updateBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		updateBy = userInfo.(*dos.AdminUser).UserName
	}

	row := modules.FindByKeyFcUserSiteMessageFirst(&jsonp)
	if row.Id == "" {
		tmpStr := fmt.Sprintf("id: %s query not exist", row.Id)
		global.G_LOG.Errorf(tmpStr)
		response.SuccessJSON(c, struct{}{})
		return
	}

	if !modules.CheckAdminUserMerchantPerms(c, row.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	//  如果是人工全局消息,则只更新
	if row.MsgIdType == 1 {
		updataMap := map[string]interface{}{}
		updataMap["del_status"] = 2
		updataMap["update_time"] = automaticType.Now()
		updataMap["update_by"] = updateBy

		err = global.G_DB.Model(&dos.FcUserSiteMessage{}).Where("id = ?", jsonp.Id).Updates(updataMap).Error
		if err != nil {
			global.G_LOG.Error(err)
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}

	} else {
		modules.DeleteFcUserSiteMessage(&jsonp)
	}

	response.SuccessJSON(c, struct{}{})
}
