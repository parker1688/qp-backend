package siteControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func GetNotifyMarquee(c *gin.Context) {
	merchantCode := modules.GetAgentDomainMerchantCodeByHeader(c)
	data := modules.FindByKeyFcSiteNotifyMarquee(&dos.FcSiteNotifyMarquee{
		MerchantCode: merchantCode,
		Status:       1,
	}, nil)
	var newData []*vo.NotifyMarqueeResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}
