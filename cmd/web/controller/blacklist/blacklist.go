// The build tag makes sure the stub is not built in the final build.

package blacklist

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

// api: api/blacklist/save
func SaveBlacklistControl(c *gin.Context) {
	var jsonp dos.Blacklist
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
		jsonp.CreateTime = automaticType.Time(time.Now())
		jsonp.UpdateTime = jsonp.CreateTime
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}

	data, _ := modules.SaveBlacklist(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/blacklist/findPage
func FindPageBlacklistControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.Blacklist
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Type = tool.Atoi(c.DefaultQuery("type", "0"))
	jsonp.Value = c.DefaultQuery("value", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageBlacklist(jsonp.PageNo, jsonp.PageSize, &jsonp.Blacklist)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/blacklist/findByKey
func FindByKeyBlacklistControl(c *gin.Context) {
	var jsonp dos.Blacklist
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
	data := modules.FindByKeyBlacklist(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/blacklist/update
func UpdateBlacklistControl(c *gin.Context) {
	var jsonp dos.Blacklist
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
		jsonp.UpdateTime = automaticType.Time(time.Now())
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateBlacklist(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/blacklist/delete
func DeleteBlacklistControl(c *gin.Context) {
	var jsonp dos.Blacklist
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
	data := modules.DeleteBlacklist(&jsonp)
	response.SuccessJSON(c, data)
}
