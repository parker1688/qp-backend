// The build tag makes sure the stub is not built in the final build.

package fcUserVipRecord

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcUserVipRecord/save
func SaveFcUserVipRecordControl(c *gin.Context) {
	var jsonp dos.FcUserVipRecord
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

	data, _ := modules.SaveFcUserVipRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserVipRecord/findPage
func FindPageFcUserVipRecordControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcUserVipRecord
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.BeforLevel = tool.Atoi(c.DefaultQuery("befor_level", ""))
	jsonp.Level = tool.Atoi(c.DefaultQuery("level", ""))
	jsonp.BeforVip = c.DefaultQuery("befor_vip", "")
	jsonp.Vip = c.DefaultQuery("vip", "")

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
	data, total := modules.FindPageFcUserVipRecord(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserVipRecord, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcUserVipRecord/findByKey
func FindByKeyFcUserVipRecordControl(c *gin.Context) {
	var jsonp dos.FcUserVipRecord
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
	data := modules.FindByKeyFcUserVipRecord(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcUserVipRecord/update
func UpdateFcUserVipRecordControl(c *gin.Context) {
	var jsonp dos.FcUserVipRecord
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

	data := modules.UpdateFcUserVipRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserVipRecord/delete
func DeleteFcUserVipRecordControl(c *gin.Context) {
	var jsonp dos.FcUserVipRecord
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

	userVipRecord := modules.FindByKeyFcUserVipRecordFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, userVipRecord.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcUserVipRecord(&jsonp)
	response.SuccessJSON(c, data)
}
