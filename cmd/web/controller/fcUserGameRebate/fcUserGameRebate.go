// The build tag makes sure the stub is not built in the final build.

package fcUserGameRebate

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserGameRebate/save
func SaveFcUserGameRebateControl(c *gin.Context) {
	var jsonp dos.FcUserGameRebate
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

	data, err := modules.SaveFcUserGameRebate(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcUserGameRebate/findPage
func FindPageFcUserGameRebateControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserGameRebate
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Day = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("day", "")))
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.GameType = c.DefaultQuery("game_type", "")

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
	data, total := modules.FindPageFcUserGameRebate(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserGameRebate)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserGameRebate/findByKey
func FindByKeyFcUserGameRebateControl(c *gin.Context) {
	var jsonp dos.FcUserGameRebate
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
	data := modules.FindByKeyFcUserGameRebate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserGameRebate/update
func UpdateFcUserGameRebateControl(c *gin.Context) {
	var jsonp dos.FcUserGameRebate
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

	data := modules.UpdateFcUserGameRebate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserGameRebate/delete
func DeleteFcUserGameRebateControl(c *gin.Context) {
	var jsonp dos.FcUserGameRebate
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
	data := modules.DeleteFcUserGameRebate(&jsonp)
	response.SuccessJSON(c, data)
}
