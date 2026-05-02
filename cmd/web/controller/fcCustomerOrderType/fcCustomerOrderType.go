// The build tag makes sure the stub is not built in the final build.

package fcCustomerOrderType

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcCustomerOrderType/save
func SaveFcCustomerOrderTypeControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrderType
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
		jsonp.UpdateBy = jsonp.CreateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}

	data, err := modules.SaveFcCustomerOrderType(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcCustomerOrderType/findPage
func FindPageFcCustomerOrderTypeControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcCustomerOrderType
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.BonusType = tool.Atoi(c.DefaultQuery("bonus_type", ""))
	jsonp.FlowMultiple = tool.Atoi(c.DefaultQuery("flow_multiple", ""))
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Remark = c.DefaultQuery("remark", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcCustomerOrderType(jsonp.PageNo, jsonp.PageSize, &jsonp.FcCustomerOrderType)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcCustomerOrderType/findByKey
func FindByKeyFcCustomerOrderTypeControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrderType
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
	data := modules.FindByKeyFcCustomerOrderType(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcCustomerOrderType/update
func UpdateFcCustomerOrderTypeControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrderType
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

	data := modules.UpdateFcCustomerOrderType(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcCustomerOrderType/delete
func DeleteFcCustomerOrderTypeControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrderType
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
	data := modules.DeleteFcCustomerOrderType(&jsonp)
	response.SuccessJSON(c, data)
}
