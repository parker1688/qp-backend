// The build tag makes sure the stub is not built in the final build.

package fcVip

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVip/save
func SaveFcVipControl(c *gin.Context) {
	var jsonp dos.FcVip
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
		jsonp.UpdateBy = jsonp.CreateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}
	if jsonp.DailyWithdrawalTimes < 1 {
		jsonp.DailyWithdrawalTimes = 99999
	}
	if jsonp.DailyWithdrawalAmount < 1 {
		jsonp.DailyWithdrawalTimes = 99999999
	}

	data, err := modules.SaveFcVip(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcVip/findPage
func FindPageFcVipControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVip
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.VipName = c.DefaultQuery("vip_name", "")
	jsonp.Level = tool.Atoi(c.DefaultQuery("level", ""))

	jsonp.DailyWithdrawalTimes = tool.Atoi(c.DefaultQuery("daily_withdrawal_times", ""))

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.Rank = c.DefaultQuery("rank", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcVip(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVip)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVip/findByKey
func FindByKeyFcVipControl(c *gin.Context) {
	var jsonp dos.FcVip
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
	data := modules.FindByKeyFcVip(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVip/update
func UpdateFcVipControl(c *gin.Context) {
	var jsonp dos.FcVip
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

	fmt.Printf("weeklugift %f", jsonp.WeeklyGift)

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.DailyWithdrawalTimes < 1 {
		jsonp.DailyWithdrawalTimes = 99999
	}
	if jsonp.DailyWithdrawalAmount < 1 {
		jsonp.DailyWithdrawalTimes = 99999999
	}

	data := modules.UpdateFcVip(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVip/delete
func DeleteFcVipControl(c *gin.Context) {
	var jsonp dos.FcVip
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
	data := modules.DeleteFcVip(&jsonp)
	response.SuccessJSON(c, data)
}
