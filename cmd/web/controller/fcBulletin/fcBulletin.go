// The build tag makes sure the stub is not built in the final build.

package fcBulletin

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcBulletin/save
func SaveFcBulletinControl(c *gin.Context) {
	var jsonp dos.FcBulletin
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
		jsonp.UpdateBy = jsonp.CreateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}

	jsonp.StartTime = automaticType.Now()
	jsonp.EndTime = automaticType.Now()

	data, err := modules.SaveFcBulletin(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcBulletin/findPage
func FindPageFcBulletinControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcBulletin
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.IsDisplay = tool.Atoi(c.DefaultQuery("is_display", ""))
	jsonp.BulletinType = tool.Atoi(c.DefaultQuery("bulletin_type", ""))
	jsonp.ContentType = tool.Atoi(c.DefaultQuery("content_type", ""))
	jsonp.BulletinImg = c.DefaultQuery("bulletin_img", "")
	jsonp.StartTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("start_time", "")))
	jsonp.EndTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("end_time", "")))
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
	data, total := modules.FindPageFcBulletin(jsonp.PageNo, jsonp.PageSize, &jsonp.FcBulletin, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcBulletin/findByKey
func FindByKeyFcBulletinControl(c *gin.Context) {
	var jsonp dos.FcBulletin
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

	data := modules.FindByKeyFcBulletin2(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcBulletin/update
func UpdateFcBulletinControl(c *gin.Context) {
	var jsonp dos.FcBulletin
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
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	data := modules.UpdateFcBulletin(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBulletin/delete
func DeleteFcBulletinControl(c *gin.Context) {
	var jsonp dos.FcBulletin
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

	bulletin := modules.FindByKeyFcBulletinFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, bulletin.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcBulletin(&jsonp)
	response.SuccessJSON(c, data)
}
