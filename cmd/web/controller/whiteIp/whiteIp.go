// The build tag makes sure the stub is not built in the final build.

package whiteIp

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/whiteIp/save
func SaveWhiteIpControl(c *gin.Context) {
	var jsonp dos.WhiteIp
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

	data, _ := modules.SaveWhiteIp(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/whiteIp/findPage
func FindPageWhiteIpControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.WhiteIp
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.IpAddr = c.DefaultQuery("ip_addr", "")
	jsonp.Remarks = c.DefaultQuery("remarks", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageWhiteIp(jsonp.PageNo, jsonp.PageSize, &jsonp.WhiteIp)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/whiteIp/findByKey
func FindByKeyWhiteIpControl(c *gin.Context) {
	var jsonp dos.WhiteIp
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
	data := modules.FindByKeyWhiteIp(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/whiteIp/update
func UpdateWhiteIpControl(c *gin.Context) {
	var jsonp dos.WhiteIp
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

	data := modules.UpdateWhiteIp(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/whiteIp/delete
func DeleteWhiteIpControl(c *gin.Context) {
	var jsonp dos.WhiteIp
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
	data := modules.DeleteWhiteIp(&jsonp)
	response.SuccessJSON(c, data)
}
