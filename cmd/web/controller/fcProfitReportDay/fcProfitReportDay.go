// The build tag makes sure the stub is not built in the final build.

package fcProfitReportDay

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcProfitReportDay/save
func SaveFcProfitReportDayControl(c *gin.Context) {
	var jsonp dos.FcProfitReportDay
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

	data, err := modules.SaveFcProfitReportDay(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcProfitReportDay/findPage
func FindPageFcProfitReportDayControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcProfitReportDay
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.ReportDate = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("report_date", "")))
	jsonp.Online = tool.Atoi(c.DefaultQuery("online", ""))

	jsonp.Register = tool.Atoi(c.DefaultQuery("register", ""))

	jsonp.NewMemberRechargeNum = tool.Atoi(c.DefaultQuery("new_member_recharge_num", ""))

	jsonp.RelativelyNewMemberRechargeNum = tool.Atoi(c.DefaultQuery("relatively_new_member_recharge_num", ""))

	jsonp.OldMemberRechargeNum = tool.Atoi(c.DefaultQuery("old_member_recharge_num", ""))

	jsonp.NewMemberWithdrawNum = tool.Atoi(c.DefaultQuery("new_member_withdraw_num", ""))

	jsonp.NewMemberWithdrawCount = tool.Atoi(c.DefaultQuery("new_member_withdraw_count", ""))
	jsonp.RelativelyNewMemberWithdrawNum = tool.Atoi(c.DefaultQuery("relatively_new_member_withdraw_num", ""))

	jsonp.RelativelyNewMemberWithdrawCount = tool.Atoi(c.DefaultQuery("relatively_new_member_withdraw_count", ""))
	jsonp.OldMemberWithdrawNum = tool.Atoi(c.DefaultQuery("old_member_withdraw_num", ""))

	jsonp.OldMemberWithdrawCount = tool.Atoi(c.DefaultQuery("old_member_withdraw_count", ""))

	jsonp.TotalBetCount = tool.Atoi(c.DefaultQuery("total_bet_count", ""))

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
	data, total := modules.FindPageFcProfitReportDay(jsonp.PageNo, jsonp.PageSize, &jsonp.FcProfitReportDay)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcProfitReportDay/findByKey
func FindByKeyFcProfitReportDayControl(c *gin.Context) {
	var jsonp dos.FcProfitReportDay
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
	data := modules.FindByKeyFcProfitReportDay(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcProfitReportDay/update
func UpdateFcProfitReportDayControl(c *gin.Context) {
	var jsonp dos.FcProfitReportDay
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

	data := modules.UpdateFcProfitReportDay(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcProfitReportDay/delete
func DeleteFcProfitReportDayControl(c *gin.Context) {
	var jsonp dos.FcProfitReportDay
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
	data := modules.DeleteFcProfitReportDay(&jsonp)
	response.SuccessJSON(c, data)
}
