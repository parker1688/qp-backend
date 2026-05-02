// The build tag makes sure the stub is not built in the final build.

package fcAgent

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcAgent/save
//func SaveFcAgentControl(c *gin.Context) {
//	var jsonp dos.FcAgent
//	err := c.ShouldBind(&jsonp)
//	err1 := validator.New().Struct(jsonp)
//	if err != nil || err1 != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//
//	userInfo, ok := c.Get("UserInfo")
//	if ok {
//		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
//	}
//
//	data, err := modules.SaveFcAgent(&jsonp)
//	if err != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//	response.SuccessJSON(c, data)
//}

// api: api/fcAgent/save
func SaveFcAgentControl(c *gin.Context) {
	var jsonp vo.AgentSaveReq
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

	//var maxInviteCode int
	//
	//global.G_DB.Model(&dos.FcAgent{}).Select("max(invite_code) as maxInviteCode").Scan(&maxInviteCode)
	//if maxInviteCode == 0 {
	//	maxInviteCode = 600001
	//}

	for i := 0; i < jsonp.Num; i++ {
		//maxInviteCode := global.G_REDIS.Incr(context.Background(), enmus.Merchant_Agent_Invite_Code_Auto).Val()
		newInviteCode, err := modules.GetNextIdGeneral(modules.GKEY_INVITE_CODE_INCR)
		if err != nil {
			continue
		}
		agent := dos.FcAgent{InviteCode: int(newInviteCode)}
		userInfo, ok := c.Get("UserInfo")
		if ok {
			agent.CreateBy = userInfo.(*dos.AdminUser).UserName
			agent.UpdateBy = agent.CreateBy
		}
		if jsonp.Status == 0 {
			agent.Status = 1
		}
		agent.MerchantName = jsonp.MerchantName
		agent.MerchantCode = jsonp.MerchantCode

		_, err = modules.SaveFcAgent(&agent)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			continue
		}
	}

	response.SuccessJSON(c, "")
}

// api: api/fcAgent/findPage
func FindPageFcAgentControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcAgent
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.AgentName = c.DefaultQuery("agent_name", "")
	jsonp.InviteCode = tool.Atoi(c.DefaultQuery("invite_code", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcAgent(jsonp.PageNo, jsonp.PageSize, &jsonp.FcAgent, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcAgent/findByKey
func FindByKeyFcAgentControl(c *gin.Context) {
	var jsonp dos.FcAgent
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
	data := modules.FindByKeyFcAgent(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcAgent/update
func UpdateFcAgentControl(c *gin.Context) {
	var jsonp dos.FcAgent
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcAgent(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcAgent/delete
func DeleteFcAgentControl(c *gin.Context) {
	var jsonp dos.FcAgent
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

	agent := modules.FindByKeyFcAgentFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, agent.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcAgent(&jsonp)
	response.SuccessJSON(c, data)
}

func InviteCode(c *gin.Context) {
	var maxInviteCode int

	global.G_DB.Model(&dos.FcAgent{}).Select("max(invite_code) as maxInviteCode").Scan(&maxInviteCode)
	if maxInviteCode == 0 {
		response.SuccessJSON(c, 600001)
		return
	}

	response.SuccessJSON(c, maxInviteCode+1)
	return
}
