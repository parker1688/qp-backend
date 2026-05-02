// The build tag makes sure the stub is not built in the final build.

package guide

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/guide/save
func SaveGuideControl(c *gin.Context) {
	var jsonp dos.Guide
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

	data, _ := modules.SaveGuide(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/guide/findPage
func FindPageGuideControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.Guide
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.Key = c.DefaultQuery("key", "")
	jsonp.Name = c.DefaultQuery("name", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageGuide(jsonp.PageNo, jsonp.PageSize, &jsonp.Guide, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/guide/findByKey
func FindByKeyGuideControl(c *gin.Context) {
	var jsonp dos.Guide
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
	data := modules.FindByKeyGuide(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/guide/update
func UpdateGuideControl(c *gin.Context) {
	var jsonp dos.Guide
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

	data := modules.UpdateGuide(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/guide/delete
func DeleteGuideControl(c *gin.Context) {
	var jsonp dos.Guide
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

	data := modules.DeleteGuide(&jsonp)
	response.SuccessJSON(c, data)
}
