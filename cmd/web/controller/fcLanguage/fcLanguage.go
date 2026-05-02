// The build tag makes sure the stub is not built in the final build.

package fcLanguage

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcLanguage/save
func SaveFcLanguageControl(c *gin.Context) {
	var jsonp dos.FcLanguage
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

	data, _ := modules.SaveFcLanguage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcLanguage/findPage
func FindPageFcLanguageControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcLanguage
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.Code = c.DefaultQuery("code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcLanguage(jsonp.PageNo, jsonp.PageSize, &jsonp.FcLanguage)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcLanguage/findByKey
func FindByKeyFcLanguageControl(c *gin.Context) {
	var jsonp dos.FcLanguage
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
	data := modules.FindByKeyFcLanguage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcLanguage/update
func UpdateFcLanguageControl(c *gin.Context) {
	var jsonp dos.FcLanguage
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

	data := modules.UpdateFcLanguage(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcLanguage/delete
func DeleteFcLanguageControl(c *gin.Context) {
	var jsonp dos.FcLanguage
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
	data := modules.DeleteFcLanguage(&jsonp)
	response.SuccessJSON(c, data)
}
