// The build tag makes sure the stub is not built in the final build.

package department

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/department/save
func SaveDepartmentControl(c *gin.Context) {
	var jsonp dos.Department
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

	// 判断部门名称是否重复
	if modules.CheckDepartmentName(jsonp.DepartmentName, "") {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "已存在该部门")
		return
	}

	data, _ := modules.SaveDepartment(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/department/findPage
func FindPageDepartmentControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.Department
	}{}
	err := c.ShouldBindQuery(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	if c.DefaultQuery("status", "-1") == "-1" {
		jsonp.Status = -1
	}

	if jsonp.PageSize < 1 {
		jsonp.PageSize = 10
	}
	if jsonp.PageNo < 1 {
		jsonp.PageNo = 1
	}

	data, total := modules.FindPageDepartment(jsonp.PageNo, jsonp.PageSize, &jsonp.Department)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/department/findByKey
func FindByKeyDepartmentControl(c *gin.Context) {
	var jsonp dos.Department
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyDepartment(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/department/update
func UpdateDepartmentControl(c *gin.Context) {
	var validateJsonp vo.DepartmentRequest
	var jsonp dos.Department

	err := c.ShouldBind(&validateJsonp)
	err1 := validator.New().Struct(validateJsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, err1.Error())
		return
	}

	tool.JsonMapper(validateJsonp, &jsonp)

	if modules.CheckDepartmentName(validateJsonp.DepartmentName, validateJsonp.Id) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "已存在该部门")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateDepartment(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/department/delete
func DeleteDepartmentControl(c *gin.Context) {
	var jsonp dos.Department
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.DeleteDepartment(&jsonp)
	response.SuccessJSON(c, data)
}

func FindDepartmentAllControl(c *gin.Context) {
	data := modules.FindDepartmentAll()
	response.SuccessJSON(c, data)
}
