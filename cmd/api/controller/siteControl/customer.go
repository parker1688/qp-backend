package siteControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
)

// SiteBaseLink
//
//	@Description: 获取站点基本配置(客服链接, IOS下载地址、安卓下载地址)
//	@param c
func SiteBaseLink(c *gin.Context) {
	//domain := c.GetHeader(vo.Domain_Header)
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	data := modules.FindByKeyFcSiteLink(&dos.FcSiteLink{
		//Domain:       domain,
		MerchantCode: merchantCode,
	})
	var newData []*vo.SiteBaseLinkResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}

func ClientSettings(c *gin.Context) {
	data := modules.FindByKeyDictsDetail(&dos.DictsDetail{
		DictsTypeCode: "Cilent_System_Settings",
	})
	var newData []*vo.CilentSystemSettingsResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}
