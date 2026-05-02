package fcSiteLink

import (
	"bootpkg/common/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func TmplSetting(c *gin.Context) {
	var jsonp struct {
		Domain       string `json:"domain"`
		CopyDomain   string `json:"copy_domain"`
		MerchantCode string `json:"merchant_code"`
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	jsonp.Domain = strings.TrimSpace(jsonp.Domain)
	jsonp.CopyDomain = strings.TrimSpace(jsonp.CopyDomain)
	if len(jsonp.CopyDomain) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "请输入复制站点域名")
		return
	}

	//sites := modules.FindByKeyFcSiteLink(&dos.FcSiteLink{
	//	Domain: jsonp.CopyDomain,
	//})
	//userInfoF, _ := c.Get("UserInfo")
	//userInfo := userInfoF.(*dos.AdminUser)
	//for _, v := range sites {
	//	d := modules.FindByKeyFcSiteLinkFirst(&dos.FcSiteLink{
	//		AppKey: v.AppKey,
	//		Domain: v.AppKey,
	//	})
	//	if len(d.Id) > 0 {
	//		continue
	//	}
	//	modules.SaveFcSiteLink(&dos.FcSiteLink{
	//		AppKey:       v.AppKey,
	//		AppLink:      v.AppLink,
	//		Content:      v.Content,
	//		CreateBy:     userInfo.CreateBy,
	//		MerchantCode: jsonp.MerchantCode,
	//		Domain:       jsonp.Domain,
	//	})
	//}
	response.SuccessJSON(c, true)
}
