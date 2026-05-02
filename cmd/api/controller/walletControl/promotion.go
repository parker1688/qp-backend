package walletControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
)

// OrderPromotionInfo
//
//	@Description: 获取优惠信息
//	@param c
func OrderPromotionInfo(c *gin.Context) {
	var jsonp vo.OrderPromotionReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	if jsonp.PageIndex == 0 {
		jsonp.PageIndex = 1
	}
	if jsonp.PageSize == 0 {
		jsonp.PageSize = 10
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)

	var data []*dos.FcOrderPromotion
	query := global.G_DB.Model(&dos.FcOrderPromotion{}).
		Where("user_id = ? ", userInfo.UserId).
		Order("create_time desc")

	if jsonp.StartTime != "" {
		query.Where("create_time >=?", jsonp.StartTime)
	}

	if jsonp.EndTime != "" {
		query.Where("create_time <=?", jsonp.EndTime)
	}

	var total int64
	query.Count(&total)
	query.Offset((jsonp.PageIndex - 1) * jsonp.PageSize).Limit(jsonp.PageSize).
		Scan(&data)

	newData := make([]*vo.OrderPromotionResp, 0, len(data))
	tool.JsonMapper(data, &newData)
	response.SuccessPageJSON(c, jsonp.PageIndex, jsonp.PageSize, total, newData)
}
