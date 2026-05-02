// The build tag makes sure the stub is not built in the final build.

package role

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/role/save
func SaveRoleControl(c *gin.Context) {
	var jsonp dos.Role
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 判断角色名是否重复
	if modules.CheckRoleNameRepeat(jsonp.RoleName) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "角色名已存在")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, _ := modules.SaveRole(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/role/findPage
func FindPageRoleControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.Role
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

	data, total := modules.FindPageRole(jsonp.PageNo, jsonp.PageSize, &jsonp.Role, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/role/findByKey
func FindByKeyRoleControl(c *gin.Context) {
	var jsonp dos.Role
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyRole(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/role/update
func UpdateRoleControl(c *gin.Context) {
	var jsonp dos.Role
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

	data := modules.UpdateRole(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/role/delete
func DeleteRoleControl(c *gin.Context) {
	var jsonp dos.Role
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.DeleteRole(&jsonp)
	response.SuccessJSON(c, data)
}

func FindRoleAllControl(c *gin.Context) {
	data := modules.FindRoleAll()
	response.SuccessJSON(c, data)
}

func UpdateRoleMenusControl(c *gin.Context) {
	var jsonp dos.Role
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.UpdateRoleMenus(&jsonp)
	response.SuccessJSON(c, data)
}

// 查询角色权限
func GetRolePermsControl(c *gin.Context) {
	var jsonp dos.Role
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

	var role dos.Role
	err = global.G_DB.Model(&dos.Role{}).Select("id", "perms_list").Where("id = ?", jsonp.Id).First(&role).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
	}

	response.SuccessJSON(c, struct {
		Id        string `json:"id"`
		PermsList string `json:"perms_list"`
	}{
		Id:        role.Id,
		PermsList: role.PermsList,
	})
}
