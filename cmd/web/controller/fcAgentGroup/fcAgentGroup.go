package fcAgentGroup

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func SaveFcAgentGroup(c *gin.Context) {
	//jsonpList := []dos.FcAgentGroup{}
	jsonpList := struct {
		Data []dos.FcAgentGroup `json:"data" form:"data" uri:"data" ` // 阅读消息
	}{}

	err := c.ShouldBind(&jsonpList)
	//err1 := validator.New().Struct(jsonpList)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data1 := true
	userInfo, ok := c.Get("UserInfo")
	merchantCodes := ""
	if ok {
		merchantCodes = userInfo.(*dos.AdminUser).MerchantCodes
	}
	//global.G_LOG.Info("fc agent group -------------------------------1:%v", jsonpList)

	for _, jsonp := range jsonpList.Data {
		//global.G_LOG.Info("fc agent group -------------------------------2:%v", jsonp)

		res := filterMerchantCodes(merchantCodes, jsonp.InviteCode)
		if !res {
			response.FailErrJSON(c, response.Merchant_Not_Exsit, "invite_code不在商户组权限内！")
			return
		}
		data, err := modules.SaveFcAgentGroup(&jsonp)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}
		//global.G_LOG.Info("fc agent group -------------------------------3:%v", data)
		data1 = data1 && data
	}
	//global.G_LOG.Info("fc agent group -------------------------------4:%v", data1)
	response.SuccessJSON(c, data1)
}

func FindByKeyFcAgentGroup(c *gin.Context) {
	jsonp := dos.FcAgentGroup{}
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
	merchanteCodes := ""
	if ok {
		merchanteCodes = userInfo.(*dos.AdminUser).MerchantCodes
	}
	data := modules.FindByKeyFcAgentGroup(&jsonp, merchanteCodes)
	response.SuccessJSON(c, data)
}

func ListFcAgentGroup(c *gin.Context) {
	jsonp := dos.FcAgentGroup{}
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
	merchanteCodes := ""
	if ok {
		merchanteCodes = userInfo.(*dos.AdminUser).MerchantCodes
	}
	data := modules.FindByKeyFcAgentGroup(&jsonp, merchanteCodes)
	response.SuccessJSON(c, data)
}

func UpdateFcAgentGroup(c *gin.Context) {
	jsonp := dos.FcAgentGroup{}
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
	merchanteCodes := ""
	if ok {
		merchanteCodes = userInfo.(*dos.AdminUser).MerchantCodes
	}
	data := modules.UpdateFcAgentGroup(&jsonp, merchanteCodes)
	response.SuccessJSON(c, data)
}

func DeleteFcAgentGroup(c *gin.Context) {
	jsonp := dos.FcAgentGroup{}
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
	merchanteCodes := ""
	if ok {
		merchanteCodes = userInfo.(*dos.AdminUser).MerchantCodes
	}
	data := modules.DeleteFcAgentGroup(&jsonp, merchanteCodes)
	response.SuccessJSON(c, data)
}

func filterMerchantCodes(merchantCodes string, id int) bool {
	query := global.G_DB.Model(&dos.FcAgent{})
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	query.Where("merchant_code in (?) and invite_code = ?", merchantCodes2, id)
	count := 0
	query.Select("count(1) as count").Scan(&count)
	return count > 0
}
