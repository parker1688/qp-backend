// The build tag makes sure the stub is not built in the final build.

package fcVenueConfig

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVenueConfig/save
func SaveFcVenueConfigControl(c *gin.Context) {
	var jsonp dos.FcVenueConfig
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

	data, _ := modules.SaveFcVenueConfig(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueConfig/findPage
func FindPageFcVenueConfigControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVenueConfig
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Alias = c.DefaultQuery("alias", "")
	jsonp.VenueName = c.DefaultQuery("venue_name", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.Remark = c.DefaultQuery("remark", "")
	jsonp.Content = c.DefaultQuery("content", "")
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcVenueConfig(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVenueConfig)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVenueConfig/findByKey
func FindByKeyFcVenueConfigControl(c *gin.Context) {
	var jsonp dos.FcVenueConfig
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
	data := modules.FindByKeyFcVenueConfig(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueConfig/update
func UpdateFcVenueConfigControl(c *gin.Context) {
	var jsonp dos.FcVenueConfig
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

	data := modules.UpdateFcVenueConfig(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueConfig/delete
func DeleteFcVenueConfigControl(c *gin.Context) {
	var jsonp dos.FcVenueConfig
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
	data := modules.DeleteFcVenueConfig(&jsonp)
	response.SuccessJSON(c, data)
}
