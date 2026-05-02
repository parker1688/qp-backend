// The build tag makes sure the stub is not built in the final build.

package mailTemplate

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/mailTemplate/save
func SaveMailTemplateControl(c *gin.Context) {
	var jsonp dos.MailTemplate
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
		jsonp.CreateTime = automaticType.Time(time.Now())
		jsonp.UpdateTime = jsonp.CreateTime
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}

	data, _ := modules.SaveMailTemplate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/mailTemplate/findPage
func FindPageMailTemplateControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.MailTemplate
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageMailTemplate(jsonp.PageNo, jsonp.PageSize, &jsonp.MailTemplate)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/mailTemplate/findByKey
func FindByKeyMailTemplateControl(c *gin.Context) {
	var jsonp dos.MailTemplate
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
	data := modules.FindByKeyMailTemplate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/mailTemplate/update
func UpdateMailTemplateControl(c *gin.Context) {
	var jsonp dos.MailTemplate
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
		jsonp.UpdateTime = automaticType.Time(time.Now())
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateMailTemplate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/mailTemplate/delete
func DeleteMailTemplateControl(c *gin.Context) {
	var jsonp dos.MailTemplate
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

	data := modules.DeleteMailTemplate(&jsonp)
	response.SuccessJSON(c, data)
}
