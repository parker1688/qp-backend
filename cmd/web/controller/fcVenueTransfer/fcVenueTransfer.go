// The build tag makes sure the stub is not built in the final build.

package fcVenueTransfer

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVenueTransfer/save
func SaveFcVenueTransferControl(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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

	data, _ := modules.SaveFcVenueTransfer(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueTransfer/findPage
func FindPageFcVenueTransferControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcVenueTransfer
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.OrderSn = c.DefaultQuery("order_sn", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.VenueLine = tool.Atoi(c.DefaultQuery("venue_line", ""))
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")

	jsonp.OptType = tool.Atoi(c.DefaultQuery("opt_type", ""))
	jsonp.Ip = c.DefaultQuery("ip", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "-1"))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	if len(c.DefaultQuery("startAt", "")) == 0 {
		jsonp.CreateTime = automaticType.Time(time.Now().Add(-1 * time.Minute)) //3分钟未处理的转账
	}

	data, total := modules.FindPageFcVenueTransfer(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVenueTransfer, jsonp.PageTimeQuery, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVenueTransfer/findByKey
func FindByKeyFcVenueTransferControl(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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
	data := modules.FindByKeyFcVenueTransfer(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueTransfer/update
func UpdateFcVenueTransferControl(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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

	venueTransfer := modules.FindByKeyFcVenueTransferFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, venueTransfer.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcVenueTransfer(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueTransfer/delete
func DeleteFcVenueTransferControl(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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

	venueTransfer := modules.FindByKeyFcVenueTransferFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, venueTransfer.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcVenueTransfer(&jsonp)
	response.SuccessJSON(c, data)
}
