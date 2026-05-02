// The build tag makes sure the stub is not built in the final build.

package fcUserLevelUpBounsRecord

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserLevelUpBounsRecord/save
func SaveFcUserLevelUpBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelUpBounsRecord
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

	data, err := modules.SaveFcUserLevelUpBounsRecord(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelUpBounsRecord/findPage
func FindPageFcUserLevelUpBounsRecordControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserLevelUpBounsRecord
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
	data, total := modules.FindPageFcUserLevelUpBounsRecord(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserLevelUpBounsRecord)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserLevelUpBounsRecord/findByKey
func FindByKeyFcUserLevelUpBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelUpBounsRecord
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
	data := modules.FindByKeyFcUserLevelUpBounsRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelUpBounsRecord/update
func UpdateFcUserLevelUpBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelUpBounsRecord
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

	data := modules.UpdateFcUserLevelUpBounsRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserLevelUpBounsRecord/delete
func DeleteFcUserLevelUpBounsRecordControl(c *gin.Context) {
	var jsonp dos.FcUserLevelUpBounsRecord
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
	data := modules.DeleteFcUserLevelUpBounsRecord(&jsonp)
	response.SuccessJSON(c, data)
}
