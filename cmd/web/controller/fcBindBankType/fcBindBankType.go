// The build tag makes sure the stub is not built in the final build.

package fcBindBankType

import (
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcBindBankType/save
func SaveFcBindBankTypeControl(c *gin.Context) {
	var jsonp dos.FcBindBankType
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

	data, _ := modules.SaveFcBindBankType(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBindBankType/findPage
func FindPageFcBindBankTypeControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcBindBankType
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.BankName = c.DefaultQuery("bank_name", "")
	jsonp.BankCode = c.DefaultQuery("bank_code", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.BankImg = c.DefaultQuery("bank_img", "")
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcBindBankType(jsonp.PageNo, jsonp.PageSize, &jsonp.FcBindBankType, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcBindBankType/findByKey
func FindByKeyFcBindBankTypeControl(c *gin.Context) {
	var jsonp dos.FcBindBankType
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
	data := modules.FindByKeyFcBindBankType(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcBindBankType/update
func UpdateFcBindBankTypeControl(c *gin.Context) {
	var jsonp dos.FcBindBankType
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

	bindBankType := modules.FindByKeyFcBindBankTypeFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, bindBankType.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcBindBankType(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBindBankType/delete
func DeleteFcBindBankTypeControl(c *gin.Context) {
	var jsonp dos.FcBindBankType
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

	bindBankType := modules.FindByKeyFcBindBankTypeFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, bindBankType.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcBindBankType(&jsonp)
	response.SuccessJSON(c, data)
}
