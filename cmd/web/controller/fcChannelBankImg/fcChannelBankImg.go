// The build tag makes sure the stub is not built in the final build.

package fcChannelBankImg

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcChannelBankImg/save
func SaveFcChannelBankImgControl(c *gin.Context) {
	var jsonp dos.FcChannelBankImg
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

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, err := modules.SaveFcChannelBankImg(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcChannelBankImg/findPage
func FindPageFcChannelBankImgControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcChannelBankImg
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")
	jsonp.ChannelName = c.DefaultQuery("channel_name", "")
	jsonp.PaymentName = c.DefaultQuery("payment_name", "")
	jsonp.Icon = c.DefaultQuery("icon", "")
	jsonp.IconPath = c.DefaultQuery("icon_path", "")
	jsonp.Img = c.DefaultQuery("img", "")
	jsonp.ImgPath = c.DefaultQuery("img_path", "")
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
	data, total := modules.FindPageFcChannelBankImg(jsonp.PageNo, jsonp.PageSize, &jsonp.FcChannelBankImg, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcChannelBankImg/findByKey
func FindByKeyFcChannelBankImgControl(c *gin.Context) {
	var jsonp dos.FcChannelBankImg
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
	data := modules.FindByKeyFcChannelBankImg(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcChannelBankImg/update
func UpdateFcChannelBankImgControl(c *gin.Context) {
	var jsonp dos.FcChannelBankImg
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

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcChannelBankImg(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcChannelBankImg/delete
func DeleteFcChannelBankImgControl(c *gin.Context) {
	var jsonp dos.FcChannelBankImg
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

	/*channelBankImg := modules.FindByKeyFcChannelBankImgFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, channelBankImg.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}*/

	data := modules.DeleteFcChannelBankImg(&jsonp)
	response.SuccessJSON(c, data)
}
