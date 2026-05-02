package venues

import (
	"bootpkg/pkg/service/venues/venueDetail"
)

type VenuePullOrderReq struct {
	StartTime    string //开始时间
	EndTime      string //结束时间
	MerchantCode string
	VenueCode    string
	Page         int
	PageSize     int
	GameType     string
}

// VenueLaunch
//
//	@Description: 用户运行游戏
//	@param req 用户请求结构体
//	@return *venueDetail.VenueLoginGameResponse 用户运行游戏返回
func VenuePullOrder(req *VenuePullOrderReq) *venueDetail.VenuePullOrderResponse {
	var venueResponse *venueDetail.VenuePullOrderResponse

	vMethod, _ := venuesLineConfig(req.MerchantCode, req.VenueCode)
	if vMethod == nil {
		//商户未配置线路
		return venueResponse
	}
	vData := vMethod.PullOrder(&venueDetail.VenuePullOrderRequest{
		Product:   req.VenueCode,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Page:      req.Page,
		PageSize:  req.PageSize,
		GameType:  req.GameType,
	})

	return vData
}
