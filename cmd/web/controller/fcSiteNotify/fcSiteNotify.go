// The build tag makes sure the stub is not built in the final build.

package fcSiteNotify

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcSiteNotify/save
func SaveFcSiteNotifyControl(c *gin.Context) {
	var jsonp dos.FcSiteNotify
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

	data, _ := modules.SaveFcSiteNotify(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteNotify/findPage
func FindPageFcSiteNotifyControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcSiteNotify
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.NotifyType = c.DefaultQuery("notify_type", "")
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
	data, total := modules.FindPageFcSiteNotify(jsonp.PageNo, jsonp.PageSize, &jsonp.FcSiteNotify)
	var newData *dos.FcSiteNotify
	tool.JsonMapper(data, &newData)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcSiteNotify/findByKey
func FindByKeyFcSiteNotifyControl(c *gin.Context) {
	var jsonp dos.FcSiteNotify
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
	data := modules.FindByKeyFcSiteNotify(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteNotify/update
func UpdateFcSiteNotifyControl(c *gin.Context) {
	var jsonp dos.FcSiteNotify
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

	data := modules.UpdateFcSiteNotify(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteNotify/delete
func DeleteFcSiteNotifyControl(c *gin.Context) {
	var jsonp dos.FcSiteNotify
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
	data := modules.DeleteFcSiteNotify(&jsonp)
	response.SuccessJSON(c, data)
}
