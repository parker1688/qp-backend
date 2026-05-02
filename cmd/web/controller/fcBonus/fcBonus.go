// The build tag makes sure the stub is not built in the final build.

package fcBonus

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcBonus/save
func SaveFcBonusControl(c *gin.Context) {
	var jsonp dos.FcBonus
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

	data, _ := modules.SaveFcBonus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBonus/findPage
func FindPageFcBonusControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcBonus
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")

	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.BonusType = tool.Atoi(c.DefaultQuery("bonus_type", ""))
	jsonp.Remark = c.DefaultQuery("remark", "")
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
	data, total := modules.FindPageFcBonus(jsonp.PageNo, jsonp.PageSize, &jsonp.FcBonus)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcBonus/findByKey
func FindByKeyFcBonusControl(c *gin.Context) {
	var jsonp dos.FcBonus
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
	data := modules.FindByKeyFcBonus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBonus/update
func UpdateFcBonusControl(c *gin.Context) {
	var jsonp dos.FcBonus
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

	data := modules.UpdateFcBonus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBonus/delete
func DeleteFcBonusControl(c *gin.Context) {
	var jsonp dos.FcBonus
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
	data := modules.DeleteFcBonus(&jsonp)
	response.SuccessJSON(c, data)
}
