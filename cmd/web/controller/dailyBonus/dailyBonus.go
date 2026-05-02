// The build tag makes sure the stub is not built in the final build.

package dailyBonus

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/dailyBonus/save
func SaveDailyBonusControl(c *gin.Context) {
	var jsonp dos.DailyBonus
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

	if modules.CheckDailyBonusMerchant(jsonp.MerchantCode, "") {
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

	data, _ := modules.SaveDailyBonus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dailyBonus/findPage
func FindPageDailyBonusControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.DailyBonus
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageDailyBonus(jsonp.PageNo, jsonp.PageSize, &jsonp.DailyBonus, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/dailyBonus/findByKey
func FindByKeyDailyBonusControl(c *gin.Context) {
	var jsonp dos.DailyBonus
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
	data := modules.FindByKeyDailyBonus(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/dailyBonus/update
func UpdateDailyBonusControl(c *gin.Context) {
	var jsonp dos.DailyBonus
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

	if modules.CheckDailyBonusMerchant(jsonp.MerchantCode, jsonp.Id) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户已存在")
		return
	}

	merchant := modules.FindByKeyDailyBonusFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateTime = automaticType.Time(time.Now())
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateDailyBonus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dailyBonus/delete
func DeleteDailyBonusControl(c *gin.Context) {
	var jsonp dos.DailyBonus
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

	merchant := modules.FindByKeyDailyBonusFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteDailyBonus(&jsonp)
	response.SuccessJSON(c, data)
}
