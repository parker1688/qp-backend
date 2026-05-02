// The build tag makes sure the stub is not built in the final build.

package fcAgentDomain

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcAgentDomain/save
func SaveFcAgentDomainControl(c *gin.Context) {
	var jsonp dos.FcAgentDomain
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, err := modules.SaveFcAgentDomain(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	response.SuccessJSON(c, data)
}

// api: api/fcAgentDomain/findPage
func FindPageFcAgentDomainControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcAgentDomain
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.AgentName = c.DefaultQuery("agent_name", "")
	jsonp.InviteCode = tool.Atoi(c.DefaultQuery("invite_code", ""))
	jsonp.Domain = c.DefaultQuery("domain", "")
	jsonp.ShortLink = c.DefaultQuery("short_link", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.AgentId = c.DefaultQuery("agent_id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.Type = tool.Atoi(c.DefaultQuery("type", "0"))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcAgentDomain(jsonp.PageNo, jsonp.PageSize, &jsonp.FcAgentDomain, c)

	list := []*dos.FcAgentDomainResp{}
	for _, v := range data {
		agentDomianResp := dos.FcAgentDomainResp{}
		tool.JsonMapper(v, &agentDomianResp)
		agentDomianResp.MerchantName = v.Merchant.MerchantName
		agentDomianResp.BannerImg = v.MerchantLink.BannerImg
		list = append(list, &agentDomianResp)
	}

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, list)
}

// api: api/fcAgentDomain/findByKey
func FindByKeyFcAgentDomainControl(c *gin.Context) {
	var jsonp dos.FcAgentDomain
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
	data := modules.FindByKeyFcAgentDomain(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcAgentDomain/update
func UpdateFcAgentDomainControl(c *gin.Context) {
	var jsonp dos.FcAgentDomain
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcAgentDomain(&jsonp)
	if !data {
		response.FailErrJSON(c, response.ERROR_SERVER, "更新失败")
		return
	}
	if data {
		// 同步客服链接
		modules.SyncFcAgentDomainCustomerLink(jsonp.MerchantCode, true)
	}

	response.SuccessJSON(c, data)
}

// api: api/fcAgentDomain/delete
func DeleteFcAgentDomainControl(c *gin.Context) {
	var jsonp dos.FcAgentDomain
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

	agentDomain := modules.FindByKeyFcAgentDomainFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, agentDomain.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcAgentDomain(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcAgentDomain/type/options
func TypeOptionsFcAgentDomainControl(c *gin.Context) {
	var jsonp struct {
		response.PageQuery
		dos.FcAgentDomain
	}
	err := c.ShouldBindQuery(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	if len(jsonp.MerchantCode) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户码为空")
		return
	}

	if jsonp.Type == enmus.AgentDomainType_Agent {
		// 是代理
		list, total := modules.GetFcAgentDomainAgentOptions(c, jsonp.PageQuery, &jsonp.FcAgentDomain)
		response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, list)
		return
	} else if jsonp.Type == enmus.AgentDomainType_Invite {
		// 是推广
		list, total := modules.GetFcAgentDomainInviteOptions(c, jsonp.PageQuery, &jsonp.FcAgentDomain)
		response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, list)
		return
	}

	response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
}
