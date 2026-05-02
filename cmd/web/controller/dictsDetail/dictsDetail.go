// The build tag makes sure the stub is not built in the final build.

package dictsDetail

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/dictsDetail/save
func SaveDictsDetailControl(c *gin.Context) {
	var jsonp dos.DictsDetail
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

	data, _ := modules.SaveDictsDetail(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dictsDetail/findPage
func FindPageDictsDetailControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.DictsDetail
	}{}
	err := c.ShouldBindQuery(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageDictsDetail(jsonp.PageNo, jsonp.PageSize, &jsonp.DictsDetail)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/dictsDetail/findByKey
func FindByKeyDictsDetailControl(c *gin.Context) {
	var jsonp dos.DictsDetail
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyDictsDetail(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dictsDetail/update
func UpdateDictsDetailControl(c *gin.Context) {
	var jsonp dos.DictsDetail
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

	data := modules.UpdateDictsDetail(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/dictsDetail/delete
func DeleteDictsDetailControl(c *gin.Context) {
	var jsonp dos.DictsDetail
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.DeleteDictsDetail(&jsonp)
	response.SuccessJSON(c, data)
}


// api: api/dictsDetail/all
func FindByKeyDictsDetailAll(c *gin.Context) {
	var jsonp dos.DictsDetail
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyDictsDetailAll(&jsonp)
	response.SuccessJSON(c, data)
}
