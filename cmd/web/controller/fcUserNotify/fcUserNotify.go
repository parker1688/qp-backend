// The build tag makes sure the stub is not built in the final build.

package fcUserNotify

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserNotify/save
func SaveFcUserNotifyControl(c *gin.Context) {
	var jsonp dos.FcUserNotify
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
	playerName := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: jsonp.UserId,
	})
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}
	if playerName.UserId == "" {
		response.FailErrJSON(c, response.Member_not_exsit, "会员不存在")
		return
	}
	jsonp.UserId = playerName.UserId
	jsonp.UserName = playerName.UserName
	data, _ := modules.SaveFcUserNotify(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserNotify/findPage
func FindPageFcUserNotifyControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserNotify
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcUserNotify(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserNotify)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserNotify/findByKey
func FindByKeyFcUserNotifyControl(c *gin.Context) {
	var jsonp dos.FcUserNotify
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
	data := modules.FindByKeyFcUserNotify(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserNotify/update
func UpdateFcUserNotifyControl(c *gin.Context) {
	var jsonp dos.FcUserNotify
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

	data := modules.UpdateFcUserNotify(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserNotify/delete
func DeleteFcUserNotifyControl(c *gin.Context) {
	var jsonp dos.FcUserNotify
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
	data := modules.DeleteFcUserNotify(&jsonp)
	response.SuccessJSON(c, data)
}
