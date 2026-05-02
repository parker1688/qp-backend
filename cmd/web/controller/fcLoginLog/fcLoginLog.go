// The build tag makes sure the stub is not built in the final build.

package fcLoginLog

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcLoginLog/save
func SaveFcLoginLogControl(c *gin.Context) {
	var jsonp dos.FcLoginLog
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

	data, _ := modules.SaveFcLoginLog(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcLoginLog/findPage
func FindPageFcLoginLogControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcLoginLog
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.PageTimeQuery.StartAt = c.DefaultQuery("startAt", "")
	jsonp.PageTimeQuery.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Ip = c.DefaultQuery("ip", "")
	jsonp.Version = c.DefaultQuery("version", "")
	jsonp.ClientType = c.DefaultQuery("client_type", "")
	jsonp.VisitorId = c.DefaultQuery("visitor_id", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcLoginLog(jsonp.PageNo, jsonp.PageSize, &jsonp.FcLoginLog, &jsonp.PageTimeQuery)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcLoginLog/findByKey
func FindByKeyFcLoginLogControl(c *gin.Context) {
	var jsonp dos.FcLoginLog
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
	data := modules.FindByKeyFcLoginLog(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcLoginLog/update
func UpdateFcLoginLogControl(c *gin.Context) {
	var jsonp dos.FcLoginLog
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

	data := modules.UpdateFcLoginLog(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcLoginLog/delete
func DeleteFcLoginLogControl(c *gin.Context) {
	var jsonp dos.FcLoginLog
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
	data := modules.DeleteFcLoginLog(&jsonp)
	response.SuccessJSON(c, data)
}
