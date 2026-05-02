// The build tag makes sure the stub is not built in the final build.

package dicts

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/dicts/save
func SaveDictsControl(c *gin.Context) {
	var jsonp dos.Dicts
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, _ := modules.SaveDicts(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dicts/findPage
func FindPageDictsControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.Dicts
	}{}
	err := c.ShouldBindQuery(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if jsonp.PageSize < 1 {
		jsonp.PageSize = 10
	}
	if jsonp.PageNo < 1 {
		jsonp.PageNo = 1
	}

	data, total := modules.FindPageDicts(jsonp.PageNo, jsonp.PageSize, &jsonp.Dicts)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/dicts/findByKey
func FindByKeyDictsControl(c *gin.Context) {
	var jsonp dos.Dicts
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	data := modules.FindByKeyDicts(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dicts/update
func UpdateDictsControl(c *gin.Context) {
	var jsonp dos.Dicts
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateDicts(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dicts/delete
func DeleteDictsControl(c *gin.Context) {
	var jsonp dos.Dicts
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.DeleteDicts(&jsonp)
	response.SuccessJSON(c, data)
}
