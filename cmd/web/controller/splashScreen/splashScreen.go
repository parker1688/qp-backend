// The build tag makes sure the stub is not built in the final build.

package splashScreen

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

// api: api/splashScreen/save
func SaveSplashScreenControl(c *gin.Context) {
	var jsonp dos.SplashScreen
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

	if modules.CheckSplashScreenMerchant(jsonp.MerchantCode, "") {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户已存在")
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

	data, _ := modules.SaveSplashScreen(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/splashScreen/findPage
func FindPageSplashScreenControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.SplashScreen
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageSplashScreen(jsonp.PageNo, jsonp.PageSize, &jsonp.SplashScreen, c)

	list := []*dos.SplashScreenResp{}
	for _, v := range data {
		splashScreenResp := dos.SplashScreenResp{}
		tool.JsonMapper(v, &splashScreenResp)
		splashScreenResp.MerchantName = v.Merchant.MerchantName
		list = append(list, &splashScreenResp)
	}

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, list)
}

// api: api/splashScreen/findByKey
func FindByKeySplashScreenControl(c *gin.Context) {
	var jsonp dos.SplashScreen
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

	data := modules.FindByKeySplashScreen(&jsonp, c)

	var list []*dos.SplashScreenResp
	for _, v := range data {
		splashScreenResp := dos.SplashScreenResp{}
		tool.JsonMapper(v, &splashScreenResp)
		splashScreenResp.MerchantName = v.Merchant.MerchantName
		list = append(list, &splashScreenResp)
	}

	response.SuccessJSON(c, list)
}

// api: api/splashScreen/update
func UpdateSplashScreenControl(c *gin.Context) {
	var jsonp dos.SplashScreen
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

	if modules.CheckSplashScreenMerchant(jsonp.MerchantCode, jsonp.Id) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户已存在")
		return
	}

	merchant := modules.FindByKeySplashScreenFirst(&dos.SplashScreen{
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

	data := modules.UpdateSplashScreen(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/splashScreen/delete
func DeleteSplashScreenControl(c *gin.Context) {
	var jsonp dos.SplashScreen
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

	merchant := modules.FindByKeySplashScreenFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteSplashScreen(&jsonp)
	response.SuccessJSON(c, data)
}
