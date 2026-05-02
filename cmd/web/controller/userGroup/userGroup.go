// The build tag makes sure the stub is not built in the final build.

package userGroup

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/userGroup/save
func SaveUserGroupControl(c *gin.Context) {
	var jsonp dos.UserGroup
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

	if modules.CheckUserGroupName(jsonp.GroupName, "") {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "组名称已存在")
		return
	}

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	// 判断是否存在不同商户的用户id组
	err1 = modules.CheckUserGroupIdsByMerchantCode(jsonp.MerchantCode, jsonp.Data)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateTime = automaticType.Time(time.Now())
		jsonp.UpdateTime = jsonp.CreateTime
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}

	data, _ := modules.SaveUserGroup(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/userGroup/findPage
func FindPageUserGroupControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.UserGroup
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.GroupName = c.DefaultQuery("group_name", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageUserGroup(jsonp.PageNo, jsonp.PageSize, &jsonp.UserGroup, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/userGroup/findByKey
func FindByKeyUserGroupControl(c *gin.Context) {
	var jsonp dos.UserGroup
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
	data := modules.FindByKeyUserGroup(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/userGroup/update
func UpdateUserGroupControl(c *gin.Context) {
	var jsonp dos.UserGroup
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

	if modules.CheckUserGroupName(jsonp.GroupName, jsonp.Id) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "组名称已存在")
		return
	}

	merchant := modules.FindByKeyUserGroupFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	// 判断是否存在不同商户的用户id组
	err1 = modules.CheckUserGroupIdsByMerchantCode(jsonp.MerchantCode, jsonp.Data)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateTime = automaticType.Time(time.Now())
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateUserGroup(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/userGroup/delete
func DeleteUserGroupControl(c *gin.Context) {
	var jsonp dos.UserGroup
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

	merchant := modules.FindByKeyUserGroupFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteUserGroup(&jsonp)
	response.SuccessJSON(c, data)
}
