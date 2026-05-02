// The build tag makes sure the stub is not built in the final build.

package fcSiteNotifyMarquee

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcSiteNotifyMarquee/save
func SaveFcSiteNotifyMarqueeControl(c *gin.Context) {
	var jsonp dos.FcSiteNotifyMarquee
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

	data, err := modules.SaveFcSiteNotifyMarquee(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcSiteNotifyMarquee/findPage
func FindPageFcSiteNotifyMarqueeControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcSiteNotifyMarquee
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.Frequency = tool.Atoi(c.DefaultQuery("frequency", ""))
	jsonp.Content = c.DefaultQuery("content", "")
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
	data, total := modules.FindPageFcSiteNotifyMarquee(jsonp.PageNo, jsonp.PageSize, &jsonp.FcSiteNotifyMarquee, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcSiteNotifyMarquee/findByKey
func FindByKeyFcSiteNotifyMarqueeControl(c *gin.Context) {
	var jsonp dos.FcSiteNotifyMarquee
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
	data := modules.FindByKeyFcSiteNotifyMarquee(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteNotifyMarquee/update
func UpdateFcSiteNotifyMarqueeControl(c *gin.Context) {
	var jsonp dos.FcSiteNotifyMarquee
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
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcSiteNotifyMarquee(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteNotifyMarquee/delete
func DeleteFcSiteNotifyMarqueeControl(c *gin.Context) {
	var jsonp dos.FcSiteNotifyMarquee
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

	siteNotifyMarquee := modules.FindByKeyFcSiteNotifyMarqueeFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, siteNotifyMarquee.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcSiteNotifyMarquee(&jsonp)
	response.SuccessJSON(c, data)
}
