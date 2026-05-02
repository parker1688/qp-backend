// The build tag makes sure the stub is not built in the final build.

package fcGameRebate

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcGameRebate/save
func SaveFcGameRebateControl(c *gin.Context) {
	var jsonp dos.FcGameRebate
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
	if jsonp.MaxBetAmount < jsonp.MinBetAmount {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	data, err := modules.SaveFcGameRebate(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcGameRebate/findPage
func FindPageFcGameRebateControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcGameRebate
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.GameType = c.DefaultQuery("game_type", "")

	jsonp.Describe = c.DefaultQuery("describe", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total, err := modules.FindPageFcGameRebate(jsonp.PageNo, jsonp.PageSize, &jsonp.FcGameRebate)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcGameRebate/findByKey
func FindByKeyFcGameRebateControl(c *gin.Context) {
	var jsonp dos.FcGameRebate
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
	data := modules.FindByKeyFcGameRebate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcGameRebate/update
func UpdateFcGameRebateControl(c *gin.Context) {
	var jsonp dos.FcGameRebate
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
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.MaxBetAmount < jsonp.MinBetAmount {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	data := modules.UpdateFcGameRebate(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcGameRebate/delete
func DeleteFcGameRebateControl(c *gin.Context) {
	var jsonp dos.FcGameRebate
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
	data := modules.DeleteFcGameRebate(&jsonp)
	response.SuccessJSON(c, data)
}
