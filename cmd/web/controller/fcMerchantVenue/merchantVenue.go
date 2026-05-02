package fcMerchantVenue

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
	"strings"
)

// CopyVenue
//
//	@Description: 复制商户场馆配置
//	@param c
func CopyVenue(c *gin.Context) {
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
	venueInfo := modules.FindByKeyFcMerchantFirst(&dos.FcMerchant{
		MerchantCode: jsonp.MerchantCode,
	})

	venues := modules.FindByKeyFcMerchantVenue(&dos.FcMerchantVenue{
		MerchantCode: jsonp.MerchantCodeCopy,
	})
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	for _, v := range venues {
		d := modules.FindByKeyFcMerchantVenueFirst(&dos.FcMerchantVenue{
			MerchantCode: venueInfo.MerchantCode,
			VenueCode:    v.VenueCode,
		})
		if len(d.Id) > 0 {
			continue
		}
		modules.SaveFcMerchantVenue(&dos.FcMerchantVenue{
			MerchantCode: jsonp.MerchantCode,
			VenueId:      v.VenueId,
			VenueCode:    v.VenueCode,
			Status:       v.Status,
			ConfigId:     v.ConfigId,
			VenueName:    v.VenueName,
			VenueFeeRate: v.VenueFeeRate,
			ConfigAlias:  v.ConfigAlias,
			Currency:     v.Currency,
			CreateBy:     userInfo.CreateBy,
		})
	}
	response.SuccessJSON(c, true)
}
