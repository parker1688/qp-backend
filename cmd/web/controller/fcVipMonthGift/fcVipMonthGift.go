// The build tag makes sure the stub is not built in the final build.

package fcVipMonthGift

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVipMonthGift/save
func SaveFcVipMonthGiftControl(c *gin.Context) {
	var jsonp dos.FcVipMonthGift
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

	data, err := modules.SaveFcVipMonthGift(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcVipMonthGift/findPage
func FindPageFcVipMonthGiftControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVipMonthGift
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Month = c.DefaultQuery("month", "")
	jsonp.VipName = c.DefaultQuery("vip_name", "")
	jsonp.Level = tool.Atoi(c.DefaultQuery("level", ""))
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
	data, total := modules.FindPageFcVipMonthGift(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVipMonthGift, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVipMonthGift/findByKey
func FindByKeyFcVipMonthGiftControl(c *gin.Context) {
	var jsonp dos.FcVipMonthGift
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
	data := modules.FindByKeyFcVipMonthGift(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcVipMonthGift/update
func UpdateFcVipMonthGiftControl(c *gin.Context) {
	var jsonp dos.FcVipMonthGift
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

	data := modules.UpdateFcVipMonthGift(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVipMonthGift/delete
func DeleteFcVipMonthGiftControl(c *gin.Context) {
	var jsonp dos.FcVipMonthGift
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

	vipMonth := modules.FindByKeyFcVipMonthGiftFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, vipMonth.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcVipMonthGift(&jsonp)
	response.SuccessJSON(c, data)
}
