// The build tag makes sure the stub is not built in the final build.

package fcFinanceReportDay

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcFinanceReportDay/save
func SaveFcFinanceReportDayControl(c *gin.Context) {
	var jsonp dos.FcFinanceReportDay
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

	data, err := modules.SaveFcFinanceReportDay(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcFinanceReportDay/findPage
func FindPageFcFinanceReportDayControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcFinanceReportDay
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.ReportDate = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("report_date", "")))

	jsonp.Online = tool.Atoi(c.DefaultQuery("online", ""))

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
	data, total := modules.FindPageFcFinanceReportDay(jsonp.PageNo, jsonp.PageSize, &jsonp.FcFinanceReportDay)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcFinanceReportDay/findByKey
func FindByKeyFcFinanceReportDayControl(c *gin.Context) {
	var jsonp dos.FcFinanceReportDay
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
	data := modules.FindByKeyFcFinanceReportDay(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcFinanceReportDay/update
func UpdateFcFinanceReportDayControl(c *gin.Context) {
	var jsonp dos.FcFinanceReportDay
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

	data := modules.UpdateFcFinanceReportDay(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcFinanceReportDay/delete
func DeleteFcFinanceReportDayControl(c *gin.Context) {
	var jsonp dos.FcFinanceReportDay
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
	data := modules.DeleteFcFinanceReportDay(&jsonp)
	response.SuccessJSON(c, data)
}
