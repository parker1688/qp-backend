// The build tag makes sure the stub is not built in the final build.

package domainDistribution

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/domainDistribution/save
func SaveDomainDistributionControl(c *gin.Context) {
	var jsonp dos.DomainDistribution
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
	}

	data, _ := modules.SaveDomainDistribution(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/domainDistribution/findPage
func FindPageDomainDistributionControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.DomainDistribution
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.DomainLink = c.DefaultQuery("domain_link", "")
	jsonp.DomainType = c.DefaultQuery("domain_type", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.MinLevel = tool.Atoi(c.DefaultQuery("min_level", ""))
	jsonp.MaxLevel = tool.Atoi(c.DefaultQuery("max_level", ""))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageDomainDistribution(jsonp.PageNo, jsonp.PageSize, &jsonp.DomainDistribution)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/domainDistribution/findByKey
func FindByKeyDomainDistributionControl(c *gin.Context) {
	var jsonp dos.DomainDistribution
	err := c.ShouldBind(&jsonp)
	err1 := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	data := modules.FindByKeyDomainDistribution(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/domainDistribution/update
func UpdateDomainDistributionControl(c *gin.Context) {
	var jsonp dos.DomainDistribution
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

	data := modules.UpdateDomainDistribution(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/domainDistribution/delete
func DeleteDomainDistributionControl(c *gin.Context) {
	var jsonp dos.DomainDistribution
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
	data := modules.DeleteDomainDistribution(&jsonp)
	response.SuccessJSON(c, data)
}
