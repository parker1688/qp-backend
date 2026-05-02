// The build tag makes sure the stub is not built in the final build.

package fcUserWithdrawBankBind

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/vo"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserWithdrawBankBind/save
func SaveFcUserWithdrawBankBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBankBind
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

	data, _ := modules.SaveFcUserWithdrawBankBind(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawBankBind/findPage
func FindPageFcUserWithdrawBankBindControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserWithdrawBankBind
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Province = c.DefaultQuery("province", "")
	jsonp.City = c.DefaultQuery("city", "")
	jsonp.BankAddress = c.DefaultQuery("bank_address", "")
	jsonp.AccountNumber = c.DefaultQuery("account_number", "")
	jsonp.AccountHolder = c.DefaultQuery("account_holder", "")
	jsonp.AccountBankType = c.DefaultQuery("account_bank_type", "")
	jsonp.IsDefault = tool.Atoi(c.DefaultQuery("is_default", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.Currency = c.DefaultQuery("currency", "")
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
	jsonp.Encrypt()

	data, total := modules.FindPageFcUserWithdrawBankBind(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserWithdrawBankBind, c)
	for _, v := range data {
		v.Decrypt()
		p := &vo.FcPrivate{}
		p.AccountHolder = v.AccountHolder
		p.AccountNumber = v.AccountNumber
		tool.PrivateDataHandler(p)
	}
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserWithdrawBankBind/findByKey
func FindByKeyFcUserWithdrawBankBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBankBind
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
	data := modules.FindByKeyFcUserWithdrawBankBind2(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawBankBind/update
func UpdateFcUserWithdrawBankBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBankBind
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

	// =============
	userWithdrawBankBind := dos.FcUserWithdrawBankBind{}
	err = global.G_DB.Model(&dos.FcUserWithdrawBankBind{}).
		Select("merchant_code").Where("id=?", jsonp.Id).First(&userWithdrawBankBind).Error
	if err != nil {
		global.G_LOG.Errorf("[UpdateFcUserWithdrawBankBindControl] can not find user withdraw bank data err: %v", err.Error())
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	if !modules.CheckAdminUserMerchantPerms(c, userWithdrawBankBind.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}
	// =============

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	jsonp.NumberHash = tool.HMACSignatureSha256("sha256", jsonp.AccountNumber)
	jsonp.Encrypt()
	updateMap := map[string]interface{}{}
	updateMap["bank_address"] = jsonp.BankAddress
	updateMap["account_number"] = jsonp.AccountNumber
	updateMap["account_holder"] = jsonp.AccountHolder
	updateMap["account_bank_type"] = jsonp.AccountBankType
	updateMap["account_bank_code"] = jsonp.AccountBankCode
	updateMap["number_hash"] = jsonp.NumberHash
	updateMap["update_by"] = jsonp.UpdateBy
	updateMap["update_time"] = automaticType.Now()
	err = global.G_DB.Model(&dos.FcUserWithdrawBankBind{}).Where("id=?", jsonp.Id).Updates(updateMap).Error
	if err != nil {
		global.G_LOG.Errorf("updateFcUserWithdrawBankBind updateBy: %v error: %v", jsonp.UpdateBy, err.Error())
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	response.SuccessJSON(c, struct{}{})
}

// api: api/fcUserWithdrawBankBind/delete
func DeleteFcUserWithdrawBankBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBankBind
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

	userWithdrawBankBind := modules.FindByKeyFcUserWithdrawBankBindFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, userWithdrawBankBind.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcUserWithdrawBankBind(&jsonp)
	response.SuccessJSON(c, data)
}
