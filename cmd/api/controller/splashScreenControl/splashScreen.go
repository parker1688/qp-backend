package splashScreenControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 开屏信息
func SplashScreenInfo(c *gin.Context) {
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	if len(merchantCode) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户码为空")
		return
	}

	data, err := modules.GetSplashScreenByMerchant(merchantCode)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER,
			fmt.Sprintf("未找到商户(%s)的开屏配置", merchantCode))
		return
	}

	var result struct {
		LogoImg   string `json:"logo_img"`
		BannerImg string `json:"banner_img"`
		ScreenImg string `json:"screen_img"`
	}

	tool.JsonMapper(data, &result)

	response.SuccessJSON(c, result)
}
