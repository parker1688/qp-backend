// The build tag makes sure the stub is not built in the final build.

package fcWelfareManage

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcWelfareManage/save
func SaveFcWelfareManageControl(c *gin.Context) {
	var jsonp dos.FcWelfareManage
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

	/*if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	data, err := modules.SaveFcWelfareManage(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcWelfareManage/findPage
func FindPageFcWelfareManageControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcWelfareManage
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.FlowMultiple = tool.Atoi(c.DefaultQuery("flow_multiple", ""))
	jsonp.BonusType = tool.Atoi(c.DefaultQuery("bonus_type", ""))
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcWelfareManage(jsonp.PageNo, jsonp.PageSize, &jsonp.FcWelfareManage, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcWelfareManage/findByKey
func FindByKeyFcWelfareManageControl(c *gin.Context) {
	var jsonp dos.FcWelfareManage
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
	data := modules.FindByKeyFcWelfareManage(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcWelfareManage/update
func UpdateFcWelfareManageControl(c *gin.Context) {
	var jsonp dos.FcWelfareManage
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

	/*if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	data := modules.UpdateFcWelfareManage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcWelfareManage/delete
func DeleteFcWelfareManageControl(c *gin.Context) {
	var jsonp dos.FcWelfareManage
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

	/*welfareManage := modules.FindByKeyFcWelfareManageFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, welfareManage.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	data := modules.DeleteFcWelfareManage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcWelfareManage/update/flowMultiple
func UpdateFcWelfareManageFlowMultipleControl(c *gin.Context) {
	var jsonp dos.FcWelfareManage
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

	/*welfareManage := modules.FindByKeyFcWelfareManageFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, welfareManage.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	err = global.G_DB.Model(&dos.FcWelfareManage{}).Where("id = ?", jsonp.Id).Update("flow_multiple", jsonp.FlowMultiple).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	response.SuccessJSON(c, true)
}
