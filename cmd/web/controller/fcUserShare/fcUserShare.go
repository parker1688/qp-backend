// The build tag makes sure the stub is not built in the final build.

package fcUserShare

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserShare/save
func SaveFcUserShareControl(c *gin.Context) {
	var jsonp dos.FcUserShare
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

	data, _ := modules.SaveFcUserShare(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserShare/findPage
func FindPageFcUserShareControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserShare
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.ShareLink = c.DefaultQuery("share_link", "")
	jsonp.ShareCode = c.DefaultQuery("share_code", "")
	jsonp.Quantity = tool.Atoi(c.DefaultQuery("quantity", ""))
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
	data, total := modules.FindPageFcUserShare(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserShare)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserShare/findByKey
func FindByKeyFcUserShareControl(c *gin.Context) {
	var jsonp dos.FcUserShare
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
	data := modules.FindByKeyFcUserShare(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserShare/update
func UpdateFcUserShareControl(c *gin.Context) {
	var jsonp dos.FcUserShare
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

	data := modules.UpdateFcUserShare(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserShare/delete
func DeleteFcUserShareControl(c *gin.Context) {
	var jsonp dos.FcUserShare
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
	data := modules.DeleteFcUserShare(&jsonp)
	response.SuccessJSON(c, data)
}
