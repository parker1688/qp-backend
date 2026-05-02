// The build tag makes sure the stub is not built in the final build.

package menus

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/menus/save
func SaveMenusControl(c *gin.Context) {
	var jsonp dos.Menus
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
	data, _ := modules.SaveMenus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/menus/findPage
func FindPageMenusControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.Menus
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

	data, total := modules.FindPageMenus(jsonp.PageNo, jsonp.PageSize, &jsonp.Menus)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/menus/findByKey
func FindByKeyMenusControl(c *gin.Context) {
	var jsonp dos.Menus
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyMenus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/menus/update
func UpdateMenusControl(c *gin.Context) {
	var jsonp dos.Menus
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
	data := modules.UpdateMenus(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/menus/delete
func DeleteMenusControl(c *gin.Context) {
	var jsonp dos.Menus
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.DeleteMenus(&jsonp)
	response.SuccessJSON(c, data)
}

func FindMenusAllControl(c *gin.Context) {
	data := modules.FindMenusAll()
	response.SuccessJSON(c, data)
}
