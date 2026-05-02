// The build tag makes sure the stub is not built in the final build.

package fcUserLevelWeekSetting

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserLevelWeekSetting/save
func SaveFcUserLevelWeekSettingControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekSetting
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

	data, err := modules.SaveFcUserLevelWeekSetting(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelWeekSetting/findPage
func FindPageFcUserLevelWeekSettingControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserLevelWeekSetting
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Level = tool.Atoi(c.DefaultQuery("level", ""))
	jsonp.BetType = tool.Atoi(c.DefaultQuery("bet_type", ""))

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
	data, total := modules.FindPageFcUserLevelWeekSetting(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserLevelWeekSetting)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserLevelWeekSetting/findByKey
func FindByKeyFcUserLevelWeekSettingControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekSetting
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
	data := modules.FindByKeyFcUserLevelWeekSetting(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelWeekSetting/update
func UpdateFcUserLevelWeekSettingControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekSetting
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

	data := modules.UpdateFcUserLevelWeekSetting(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelWeekSetting/delete
func DeleteFcUserLevelWeekSettingControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekSetting
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
	data := modules.DeleteFcUserLevelWeekSetting(&jsonp)
	response.SuccessJSON(c, data)
}
