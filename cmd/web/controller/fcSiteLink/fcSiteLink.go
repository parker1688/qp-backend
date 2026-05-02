// The build tag makes sure the stub is not built in the final build.

package fcSiteLink

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcSiteLink/save
func SaveFcSiteLinkControl(c *gin.Context) {
	var jsonp dos.FcSiteLink
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

	data, _ := modules.SaveFcSiteLink(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteLink/findPage
func FindPageFcSiteLinkControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcSiteLink
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.AppKey = c.DefaultQuery("app_key", "")
	jsonp.AppLink = c.DefaultQuery("app_link", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	//jsonp.Domain = c.DefaultQuery("domain", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcSiteLink(jsonp.PageNo, jsonp.PageSize, &jsonp.FcSiteLink)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcSiteLink/findByKey
func FindByKeyFcSiteLinkControl(c *gin.Context) {
	var jsonp dos.FcSiteLink
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
	data := modules.FindByKeyFcSiteLink(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteLink/update
func UpdateFcSiteLinkControl(c *gin.Context) {
	jsonp := struct {
		dos.FcSiteLink
	}{}
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

	data := modules.UpdateFcSiteLink(&jsonp.FcSiteLink)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteLink/delete
func DeleteFcSiteLinkControl(c *gin.Context) {
	var jsonp dos.FcSiteLink
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
	data := modules.DeleteFcSiteLink(&jsonp)
	response.SuccessJSON(c, data)
}
