// The build tag makes sure the stub is not built in the final build.

package fcSiteBanner

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcSiteBanner/save
func SaveFcSiteBannerControl(c *gin.Context) {
	var jsonp dos.FcSiteBanner
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

	data, _ := modules.SaveFcSiteBanner(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteBanner/findPage
func FindPageFcSiteBannerControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcSiteBanner
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.BannerLink = c.DefaultQuery("banner_link", "")
	jsonp.BannerHref = c.DefaultQuery("banner_href", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.BannerType = c.DefaultQuery("banner_type", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.BannerOtherType = tool.Atoi(c.DefaultQuery("banner_other_type", ""))

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcSiteBanner(jsonp.PageNo, jsonp.PageSize, &jsonp.FcSiteBanner)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcSiteBanner/findByKey
func FindByKeyFcSiteBannerControl(c *gin.Context) {
	var jsonp dos.FcSiteBanner
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
	data := modules.FindByKeyFcSiteBanner(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteBanner/update
func UpdateFcSiteBannerControl(c *gin.Context) {
	var jsonp dos.FcSiteBanner
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

	data := modules.UpdateFcSiteBanner(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcSiteBanner/delete
func DeleteFcSiteBannerControl(c *gin.Context) {
	var jsonp dos.FcSiteBanner
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
	data := modules.DeleteFcSiteBanner(&jsonp)
	response.SuccessJSON(c, data)
}
