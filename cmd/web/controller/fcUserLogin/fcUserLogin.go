// The build tag makes sure the stub is not built in the final build.

package fcUserLogin

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserLogin/save
func SaveFcUserLoginControl(c *gin.Context) {
	var jsonp dos.FcUserLogin
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
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, _ := modules.SaveFcUserLogin(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLogin/findPage
func FindPageFcUserLoginControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcUserLogin
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.PageTimeQuery.StartAt = c.DefaultQuery("startAt", "")
	jsonp.PageTimeQuery.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Password = c.DefaultQuery("password", "")
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
	data, total := modules.FindPageFcUserLogin(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserLogin, &jsonp.PageTimeQuery)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserLogin/findByKey
func FindByKeyFcUserLoginControl(c *gin.Context) {
	var jsonp dos.FcUserLogin
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
	data := modules.FindByKeyFcUserLogin(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLogin/update
func UpdateFcUserLoginControl(c *gin.Context) {
	var jsonp dos.FcUserLogin
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

	data := modules.UpdateFcUserLogin(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLogin/delete
func DeleteFcUserLoginControl(c *gin.Context) {
	var jsonp dos.FcUserLogin
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
	data := modules.DeleteFcUserLogin(&jsonp)
	response.SuccessJSON(c, data)
}
