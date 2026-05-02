package fcSiteBanner

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"strings"
)

func BannerCopy(c *gin.Context) {
	var jsonp struct {
		MerchantCode     string `json:"merchant_code"`
		MerchantCodeCopy string `json:"merchant_code_copy"`
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	jsonp.MerchantCode = strings.TrimSpace(jsonp.MerchantCode)
	jsonp.MerchantCodeCopy = strings.TrimSpace(jsonp.MerchantCodeCopy)
	if len(jsonp.MerchantCodeCopy) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "请输入复制商户的名称")
		return
	}

	siteBanners := modules.FindByKeyFcSiteBanner(&dos.FcSiteBanner{
		MerchantCode: jsonp.MerchantCodeCopy,
	})
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	for _, v := range siteBanners {
		d := modules.FindByKeyFcSiteBannerFirst(&dos.FcSiteBanner{
			MerchantCode: jsonp.MerchantCode,
			BannerLink:   v.BannerLink,
		})
		if len(d.Id) > 0 {
			continue
		}
		modules.SaveFcSiteBanner(&dos.FcSiteBanner{
			BannerLink:      v.BannerLink,
			BannerOtherType: v.BannerOtherType,
			BannerHref:      v.BannerHref,
			Language:        v.Language,
			Sort:            v.Sort,
			BannerType:      v.BannerType,
			MerchantCode:    jsonp.MerchantCode,
			CreateBy:        userInfo.UserName,
		})
	}
	response.SuccessJSON(c, true)
}
