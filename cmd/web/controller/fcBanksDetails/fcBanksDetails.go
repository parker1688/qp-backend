// The build tag makes sure the stub is not built in the final build.

package fcBanksDetails

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcBanksDetails/save
func SaveFcBanksDetailsControl(c *gin.Context) {
	var jsonp dos.FcBanksDetails
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

	jsonp.Encrypt()

	data, _ := modules.SaveFcBanksDetails(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBanksDetails/findPage
func FindPageFcBanksDetailsControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcBanksDetails
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.BankName = c.DefaultQuery("bank_name", "")
	jsonp.BankCode = c.DefaultQuery("bank_code", "")
	jsonp.MinLevel = tool.Atoi(c.DefaultQuery("min_level", ""))
	jsonp.MaxLevel = tool.Atoi(c.DefaultQuery("max_level", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.EntityAccountHolder = c.DefaultQuery("entity_account_holder", "")
	jsonp.EntityAccountBankName = c.DefaultQuery("entity_account_bank_name", "")
	jsonp.EntityAccountNumber = c.DefaultQuery("entity_account_number", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcBanksDetails(jsonp.PageNo, jsonp.PageSize, &jsonp.FcBanksDetails)

	for _, v := range data {
		v.Decrypt()
	}
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcBanksDetails/findByKey
func FindByKeyFcBanksDetailsControl(c *gin.Context) {
	var jsonp dos.FcBanksDetails
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
	data := modules.FindByKeyFcBanksDetails(&jsonp)
	for _, v := range data {
		v.Decrypt()
	}
	response.SuccessJSON(c, data)
}

// api: api/fcBanksDetails/update
func UpdateFcBanksDetailsControl(c *gin.Context) {
	var jsonp dos.FcBanksDetails
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
	jsonp.Encrypt()
	data := modules.UpdateFcBanksDetails(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBanksDetails/delete
func DeleteFcBanksDetailsControl(c *gin.Context) {
	var jsonp dos.FcBanksDetails
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
	data := modules.DeleteFcBanksDetails(&jsonp)
	response.SuccessJSON(c, data)
}
