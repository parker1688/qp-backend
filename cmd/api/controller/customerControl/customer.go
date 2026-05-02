package customerControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func FindCustomer(c *gin.Context) {
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var FcSiteLink []*dos.FcSiteLink
	err := global.G_DB.WithContext(ctx).Model(&dos.FcSiteLink{}).
		Where("merchant_code=?  and app_key in ?",
			merchantCode,
			[]string{"facebookCustomer", "whatsappCustomer", "twitterCustomer", "telegramCustomer", "skypeCustomer"}).
		Find(&FcSiteLink).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	FcSiteLinkMap := make(map[string][]*dos.FcSiteLink, len(FcSiteLink))
	for _, v := range FcSiteLink {
		if _, ok := FcSiteLinkMap[v.AppKey]; !ok {
			FcSiteLinkMap[v.AppKey] = []*dos.FcSiteLink{v}
		} else {
			FcSiteLinkMap[v.AppKey] = append(FcSiteLinkMap[v.AppKey], v)
		}

	}
	response.SuccessJSON(c, FcSiteLinkMap)
}
