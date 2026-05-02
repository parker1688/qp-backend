package opRecord

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/opRecord/findPage
func FindPageOpRecord(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.OpRecord
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageOpRecord(jsonp.PageNo, jsonp.PageSize, &jsonp.OpRecord, &jsonp.PageTimeQuery, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcBetRecord/findByKey
func FindByKeyFcBetRecordControl(c *gin.Context) {
	var jsonp dos.FcBetRecord
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
	data := modules.FindByKeyFcBetRecord(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcBetRecord/update
func UpdateFcBetRecordControl(c *gin.Context) {
	var jsonp dos.FcBetRecord
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcBetRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBetRecord/delete
func DeleteFcBetRecordControl(c *gin.Context) {
	var jsonp dos.FcBetRecord
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

	betRecord := modules.FindByKeyFcBetRecordFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, betRecord.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcBetRecord(&jsonp)
	response.SuccessJSON(c, data)
}
