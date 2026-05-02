// The build tag makes sure the stub is not built in the final build.

package fcComplexReport

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcComplexReport/save
func SaveFcComplexReportControl(c *gin.Context) {
	var jsonp dos.FcComplexReport
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

	data, err := modules.SaveFcComplexReport(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcComplexReport/findPage
func FindPageFcComplexReportControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcComplexReport
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Day = c.DefaultQuery("day", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.MerchantName = c.DefaultQuery("merchant_name", "")

	jsonp.RegisterNum = tool.Atoi(c.DefaultQuery("register_num", ""))
	jsonp.FirstDepositNum = tool.Atoi(c.DefaultQuery("first_deposit_num", ""))

	jsonp.DepositNum = tool.Atoi(c.DefaultQuery("deposit_num", ""))
	jsonp.DepositCount = tool.Atoi(c.DefaultQuery("deposit_count", ""))

	jsonp.NewUserDepositCount = tool.Atoi(c.DefaultQuery("new_user_deposit_count", ""))
	jsonp.LoginNum = tool.Atoi(c.DefaultQuery("login_num", ""))
	jsonp.WithdrawNum = tool.Atoi(c.DefaultQuery("withdraw_num", ""))
	jsonp.BetNum = tool.Atoi(c.DefaultQuery("bet_num", ""))

	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcComplexReport(jsonp.PageNo, jsonp.PageSize, &jsonp.FcComplexReport, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcComplexReport/findByKey
func FindByKeyFcComplexReportControl(c *gin.Context) {
	var jsonp dos.FcComplexReport
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
	data := modules.FindByKeyFcComplexReport(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcComplexReport/update
func UpdateFcComplexReportControl(c *gin.Context) {
	var jsonp dos.FcComplexReport
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

	data := modules.UpdateFcComplexReport(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcComplexReport/delete
func DeleteFcComplexReportControl(c *gin.Context) {
	var jsonp dos.FcComplexReport
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

	fcComplexReport := modules.FindByKeyFcComplexReportFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, fcComplexReport.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcComplexReport(&jsonp)
	response.SuccessJSON(c, data)
}
