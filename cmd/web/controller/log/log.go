package log

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AdminUserLoginLog(c *gin.Context){
	jsonp := struct {
		response.PageQuery
		dos.LoginLog
	}{}
	err := c.ShouldBindQuery(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageLoginLog(jsonp.PageNo, jsonp.PageSize, &jsonp.LoginLog)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

func FindPageActionControl(c *gin.Context){
	jsonp := struct {
		response.PageQuery
		dos.UserAction
	}{}
	err := c.ShouldBindQuery(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageUserAction(jsonp.PageNo, jsonp.PageSize, &jsonp.UserAction)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}