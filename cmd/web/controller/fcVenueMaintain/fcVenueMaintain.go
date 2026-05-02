// The build tag makes sure the stub is not built in the final build.

package fcVenueMaintain

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVenueMaintain/save
func SaveFcVenueMaintainControl(c *gin.Context) {
	var jsonp dos.FcVenueMaintain
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

	data, _ := modules.SaveFcVenueMaintain(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueMaintain/findPage
func FindPageFcVenueMaintainControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVenueMaintain
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.VenueId = c.DefaultQuery("venue_id", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.MaintainStart = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("maintain_start", "")))
	jsonp.MaintainEnd = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("maintain_end", "")))
	jsonp.CilentType = c.DefaultQuery("cilent_type", "")
	jsonp.AllowTransfer = tool.Atoi(c.DefaultQuery("allow_transfer", ""))
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
	data, total := modules.FindPageFcVenueMaintain(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVenueMaintain)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVenueMaintain/findByKey
func FindByKeyFcVenueMaintainControl(c *gin.Context) {
	var jsonp dos.FcVenueMaintain
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
	data := modules.FindByKeyFcVenueMaintain(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueMaintain/update
func UpdateFcVenueMaintainControl(c *gin.Context) {
	var jsonp dos.FcVenueMaintain
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

	data := modules.UpdateFcVenueMaintain(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueMaintain/delete
func DeleteFcVenueMaintainControl(c *gin.Context) {
	var jsonp dos.FcVenueMaintain
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
	data := modules.DeleteFcVenueMaintain(&jsonp)
	response.SuccessJSON(c, data)
}
