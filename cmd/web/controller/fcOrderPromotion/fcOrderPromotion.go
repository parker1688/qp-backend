// The build tag makes sure the stub is not built in the final build.

package fcOrderPromotion

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcOrderPromotion/save
func SaveFcOrderPromotionControl(c *gin.Context) {
	var jsonp dos.FcOrderPromotion
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

	data, _ := modules.SaveFcOrderPromotion(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderPromotion/findPage
func FindPageFcOrderPromotionControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcOrderPromotion
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.OrderSn = c.DefaultQuery("order_sn", "")

	jsonp.ApplyType = tool.Atoi(c.DefaultQuery("apply_type", ""))

	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.TurnOver = tool.Atoi(c.DefaultQuery("turn_over", ""))

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcOrderPromotion(jsonp.PageNo, jsonp.PageSize, &jsonp.FcOrderPromotion)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcOrderPromotion/findByKey
func FindByKeyFcOrderPromotionControl(c *gin.Context) {
	var jsonp dos.FcOrderPromotion
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
	data := modules.FindByKeyFcOrderPromotion(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderPromotion/update
func UpdateFcOrderPromotionControl(c *gin.Context) {
	var jsonp dos.FcOrderPromotion
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

	data := modules.UpdateFcOrderPromotion(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderPromotion/delete
func DeleteFcOrderPromotionControl(c *gin.Context) {
	var jsonp dos.FcOrderPromotion
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
	data := modules.DeleteFcOrderPromotion(&jsonp)
	response.SuccessJSON(c, data)
}
