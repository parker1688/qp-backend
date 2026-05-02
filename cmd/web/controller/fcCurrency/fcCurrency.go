// The build tag makes sure the stub is not built in the final build.

package fcCurrency

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcCurrency/save
func SaveFcCurrencyControl(c *gin.Context) {
	var jsonp dos.FcCurrency
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

	data, _ := modules.SaveFcCurrency(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcCurrency/findPage
func FindPageFcCurrencyControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcCurrency
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Name = c.DefaultQuery("name", "")
	jsonp.Code = c.DefaultQuery("code", "")

	jsonp.Icon = c.DefaultQuery("icon", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcCurrency(jsonp.PageNo, jsonp.PageSize, &jsonp.FcCurrency)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcCurrency/findByKey
func FindByKeyFcCurrencyControl(c *gin.Context) {
	var jsonp dos.FcCurrency
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
	data := modules.FindByKeyFcCurrency(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcCurrency/update
func UpdateFcCurrencyControl(c *gin.Context) {
	var jsonp dos.FcCurrency
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

	data := modules.UpdateFcCurrency(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcCurrency/delete
func DeleteFcCurrencyControl(c *gin.Context) {
	var jsonp dos.FcCurrency
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
	data := modules.DeleteFcCurrency(&jsonp)
	response.SuccessJSON(c, data)
}
