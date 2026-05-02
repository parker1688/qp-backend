// The build tag makes sure the stub is not built in the final build.

package adsCarousel

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/adsCarousel/save
func SaveAdsCarouselControl(c *gin.Context) {
	var jsonp dos.AdsCarousel
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
		jsonp.CreateTime = automaticType.Time(time.Now())
		jsonp.UpdateTime = jsonp.CreateTime
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}

	data, _ := modules.SaveAdsCarousel(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/adsCarousel/findPage
func FindPageAdsCarouselControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.AdsCarousel
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Key = c.DefaultQuery("key", "")
	jsonp.Name = c.DefaultQuery("name", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "-1"))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageAdsCarousel(jsonp.PageNo, jsonp.PageSize, &jsonp.AdsCarousel, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/adsCarousel/findByKey
func FindByKeyAdsCarouselControl(c *gin.Context) {
	var jsonp dos.AdsCarousel
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Key = c.DefaultQuery("key", "")
	jsonp.Name = c.DefaultQuery("name", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "-1"))

	data := modules.FindByKeyAdsCarousel(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/adsCarousel/update
func UpdateAdsCarouselControl(c *gin.Context) {
	var jsonp dos.AdsCarousel
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

	merchant := modules.FindByKeyAdsCarouselFirst(&dos.AdsCarousel{
		BaseDos: dos.BaseDos{Id: jsonp.Id},
	})
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateTime = automaticType.Time(time.Now())
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateAdsCarousel(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/adsCarousel/delete
func DeleteAdsCarouselControl(c *gin.Context) {
	var jsonp dos.AdsCarousel
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

	merchant := modules.FindByKeyAdsCarouselFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteAdsCarousel(&jsonp)
	response.SuccessJSON(c, data)
}
