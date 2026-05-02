package venues

import (
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/service/venues/venueDetail"
)

type VenueTransferConfirmRequest struct {
	UserName     string //游戏账号
	VenueCode    string //场馆Code
	MerchantCode string //商户Code(线路选择配置)
	IP           string //用户IP
	OrderSn      string //转账金额
}

func VenueTransferConfirm(req *VenueTransferConfirmRequest) *venueDetail.VenueResponse {
	ret := &venueDetail.VenueResponse{
		Code: venueDetail.TransferConfirm_Processing_CODE,
		Msg:  "",
	}
	//获取商户配置线路
	vMethod, lineNum := venuesLineConfig(req.MerchantCode, req.VenueCode)
	if vMethod == nil {
		//商户未配置线路
		return ret
	}
	venueUser := modules.FindByKeyFcVenueUserFirst(&dos.FcVenueUser{UserName: req.UserName, VenueCode: req.VenueCode, MerchantCode: req.MerchantCode, VenueLine: lineNum})
	if len(venueUser.Id) == 0 {
		//用户未注册
		return ret
	}

	//evo 如果要接入需要传三方订单号，暂时没接入
	ret = vMethod.TransferConfirm(&venueDetail.VenueTransferConfirmRequest{
		UserName: venueUser.Account,
		Password: venueUser.Password,
		IP:       req.IP,
		OrderSn:  req.OrderSn,
		Product:  req.VenueCode,
	})
	return ret
}
