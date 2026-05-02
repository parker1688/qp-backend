package walletControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
)

func TransactionOrder(c *gin.Context) {
	var jsonp struct {
		vo.TransactionOrderReq
		TimeType *int `json:"time_type"`
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	if jsonp.Current == 0 {
		jsonp.Current = 1
	}
	if jsonp.PageSize == 0 {
		jsonp.PageSize = 10
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo, ok := userInfoF.(*dos.FcUserMaterial)
	if !ok {
		response.FailErrJSON(c, response.ERROR_SERVER, "getUserInfoF fail")
		return
	}

	var data []*dos.FcTranscationExt
	query := global.G_DB.Model(&dos.FcTranscationExt{}).Where("user_id = ?", userInfo.UserId)
	if jsonp.FundingType > 0 {
		query = query.Where(" funding_type = ?", jsonp.FundingType)
	}
	if jsonp.TimeType == nil {
		if jsonp.StartAt != "" {
			query = query.Where("create_time >= ?", jsonp.StartAt)
		}
		if jsonp.EndAt != "" {
			query = query.Where("create_time <= ?", jsonp.EndAt)
		}
	} else {
		sTime, eTime := tool.GetDayRange(time.Now(), *jsonp.TimeType)
		query.Where("create_time BETWEEN ? AND ?", sTime, eTime)
	}

	query = query.Order("create_time desc")
	var total int64
	query.Count(&total)
	query.Offset((jsonp.Current - 1) * jsonp.PageSize).Limit(jsonp.PageSize).
		Preload("VenueTransfer").Preload("OrderManageOpt").Find(&data)

	dataMap := map[string]*dos.FcTranscationExt{}
	for _, v := range data {
		dataMap[v.Id] = v
	}

	newData := make([]*vo.TransactionOrderResp, 0, len(data))
	tool.JsonMapper(&data, &newData)
	for _, v := range newData {
		if val, ok := dataMap[v.Id]; ok {
			v.VenueCode = val.VenueTransfer.VenueCode
			v.OptType = val.VenueTransfer.OptType
			v.TrsType = val.OrderManageOpt.TrsType
		}
	}

	response.SuccessPageJSON(c, jsonp.Current, jsonp.PageSize, total, newData)
}
