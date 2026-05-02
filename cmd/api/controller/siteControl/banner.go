package siteControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"github.com/gin-gonic/gin"
)

func GetBanners(c *gin.Context) {
	bannerOtherType := tool.Atoi(c.DefaultQuery("type", "1"))
	clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
	if len(clientType) == 0 {
		clientType = enmus.H5
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code
	data := modules.FindByKeyFcSiteBanner(&dos.FcSiteBanner{
		BannerType: clientType,
		//Language:        language,
		MerchantCode:    merchantCode,
		BannerOtherType: bannerOtherType,
	})
	var newData []*vo.BannerResp
	tool.JsonMapper(data, &newData)
	response.SuccessJSON(c, newData)
}
