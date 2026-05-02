// The build tag makes sure the stub is not built in the final build.

package fcNotifyTemplate

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcNotifyTemplate/save
func SaveFcNotifyTemplateControl(c *gin.Context) {
	var validateJsonp vo.NotifyTemplateRequest
	var jsonp dos.FcNotifyTemplate
	err := c.ShouldBind(&validateJsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, err.Error())
		return
	}
	err1 := validator.New().Struct(validateJsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, err1.Error())
		return
	}

	tool.JsonMapper(validateJsonp, &jsonp)

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}

	data, err := modules.SaveFcNotifyTemplate(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcNotifyTemplate/findPage
func FindPageFcNotifyTemplateControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcNotifyTemplate
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.TemplateContent = c.DefaultQuery("template_content", "")
	jsonp.NotifyFlag = c.DefaultQuery("notify_flag", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcNotifyTemplate(jsonp.PageNo, jsonp.PageSize, &jsonp.FcNotifyTemplate, &jsonp.PageTimeQuery)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcNotifyTemplate/findByKey
func FindByKeyFcNotifyTemplateControl(c *gin.Context) {
	var jsonp dos.FcNotifyTemplate
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
	data := modules.FindByKeyFcNotifyTemplate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcNotifyTemplate/update
func UpdateFcNotifyTemplateControl(c *gin.Context) {
	var jsonp dos.FcNotifyTemplate
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

	data := modules.UpdateFcNotifyTemplate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcNotifyTemplate/delete
func DeleteFcNotifyTemplateControl(c *gin.Context) {
	var jsonp dos.FcNotifyTemplate
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
	data := modules.DeleteFcNotifyTemplate(&jsonp)
	response.SuccessJSON(c, data)
}
