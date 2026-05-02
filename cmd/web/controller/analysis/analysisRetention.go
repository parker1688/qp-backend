// The build tag makes sure the stub is not built in the final build.

package analysis

import (
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"

	"github.com/gin-gonic/gin"
)

// api: api/analysisRetention/findByKey
func FindByKeyAnalysisRetentionControl(c *gin.Context) {
	jsonp := struct {
		MerchantCode string `json:"merchant_code"`
		InviteCode   int    `json:"invite_code"`
		response.PageTimeQuery
	}{}
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.InviteCode = tool.Atoi(c.DefaultQuery("invite_code", "0"))
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")

	if len(jsonp.StartAt) == 0 || len(jsonp.EndAt) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	data := modules.GetAnalysisRetentionData(jsonp.MerchantCode, jsonp.InviteCode,
		jsonp.StartAt, jsonp.EndAt)
	response.SuccessJSON(c, data)
}
