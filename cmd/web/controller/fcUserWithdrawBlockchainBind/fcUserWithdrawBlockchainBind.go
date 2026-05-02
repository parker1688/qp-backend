// The build tag makes sure the stub is not built in the final build.

package fcUserWithdrawBlockchainBind

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserWithdrawBlockchainBind/save
func SaveFcUserWithdrawBlockchainBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBlockchainBind
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

	data, _ := modules.SaveFcUserWithdrawBlockchainBind(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawBlockchainBind/findPage
func FindPageFcUserWithdrawBlockchainBindControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserWithdrawBlockchainBind
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Blockchain = c.DefaultQuery("blockchain", "")
	jsonp.BlockchainAddress = c.DefaultQuery("blockchain_address", "")
	jsonp.ContractType = c.DefaultQuery("contract_type", "")
	jsonp.IsDefault = tool.Atoi(c.DefaultQuery("is_default", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
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
	data, total := modules.FindPageFcUserWithdrawBlockchainBind(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserWithdrawBlockchainBind, c)
	for _, v := range data {
		v.Decrypt()
	}
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserWithdrawBlockchainBind/findByKey
func FindByKeyFcUserWithdrawBlockchainBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBlockchainBind
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
	data := modules.FindByKeyFcUserWithdrawBlockchainBind2(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawBlockchainBind/update
func UpdateFcUserWithdrawBlockchainBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBlockchainBind
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

	data := modules.UpdateFcUserWithdrawBlockchainBind(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserWithdrawBlockchainBind/delete
func DeleteFcUserWithdrawBlockchainBindControl(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBlockchainBind
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

	userWithdrawBlockchainBind := modules.FindByKeyFcUserWithdrawBlockchainBindFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, userWithdrawBlockchainBind.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcUserWithdrawBlockchainBind(&jsonp)
	response.SuccessJSON(c, data)
}
