// The build tag makes sure the stub is not built in the final build.

package fcPromotionInfo

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcPromotionInfo/save
func SaveFcPromotionInfoControl(c *gin.Context) {
	var jsonp dos.FcPromotionInfo
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, err1.Error())
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

	data, _ := modules.SaveFcPromotionInfo(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPromotionInfo/findPage
func FindPageFcPromotionInfoControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcPromotionInfo
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("StartAt", "")
	jsonp.EndAt = c.DefaultQuery("EndAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.PromotionType = tool.Atoi(c.DefaultQuery("promotion_type", ""))
	jsonp.GameType = tool.Atoi(c.DefaultQuery("game_type", ""))
	jsonp.PromotionImg = c.DefaultQuery("promotion_img", "")
	jsonp.StartTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("start_time", "")))
	jsonp.EndTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("end_time", "")))
	jsonp.Link = c.DefaultQuery("link", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.ClientType = c.DefaultQuery("client_type", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "0"))
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcPromotionInfo(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPromotionInfo, &jsonp.PageTimeQuery, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcPromotionInfo/findByKey
func FindByKeyFcPromotionInfoControl(c *gin.Context) {
	var jsonp dos.FcPromotionInfo
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
	data := modules.FindByKeyFcPromotionInfo(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcPromotionInfo/update
func UpdateFcPromotionInfoControl(c *gin.Context) {
	var jsonp dos.FcPromotionInfo
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
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcPromotionInfo(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPromotionInfo/update
func UpdateFcPromotionInfoStatusControl(c *gin.Context) {
	var jsonp dos.FcPromotionInfo
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

	if jsonp.Id == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	promotionInfo := modules.FindByKeyFcPromotionInfoFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, promotionInfo.MerchantCode) {
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

	paramMap := map[string]interface{}{
		"status":      jsonp.Status,
		"update_by":   jsonp.UpdateBy,
		"update_time": jsonp.UpdateTime,
	}
	err = global.G_DB.Model(dos.FcPromotionInfo{}).Where(`id = ?`, jsonp.Id).Updates(paramMap).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	response.SuccessJSON(c, struct{}{})
}

// api: api/fcPromotionInfo/delete
func DeleteFcPromotionInfoControl(c *gin.Context) {
	var jsonp dos.FcPromotionInfo
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

	promotionInfo := modules.FindByKeyFcPromotionInfoFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, promotionInfo.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcPromotionInfo(&jsonp)
	response.SuccessJSON(c, data)
}
