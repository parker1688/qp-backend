// The build tag makes sure the stub is not built in the final build.

package fcUserLevelWeekBounsRecord

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserLevelWeekBounsRecord/save
func SaveFcUserLevelWeekBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekBounsRecord
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

	data, err := modules.SaveFcUserLevelWeekBounsRecord(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelWeekBounsRecord/findPage
func FindPageFcUserLevelWeekBounsRecordControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserLevelWeekBounsRecord
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.BetType = tool.Atoi(c.DefaultQuery("bet_type", ""))
	jsonp.Level = tool.Atoi(c.DefaultQuery("level", ""))

	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcUserLevelWeekBounsRecord(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserLevelWeekBounsRecord)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserLevelWeekBounsRecord/findByKey
func FindByKeyFcUserLevelWeekBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekBounsRecord
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
	data := modules.FindByKeyFcUserLevelWeekBounsRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelWeekBounsRecord/update
func UpdateFcUserLevelWeekBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekBounsRecord
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

	data := modules.UpdateFcUserLevelWeekBounsRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelWeekBounsRecord/delete
func DeleteFcUserLevelWeekBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelWeekBounsRecord
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
	data := modules.DeleteFcUserLevelWeekBounsRecord(&jsonp)
	response.SuccessJSON(c, data)
}
