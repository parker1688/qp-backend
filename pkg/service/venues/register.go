package venues

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/venues/venueDetail"
	"strings"
)

type VenueRegisterReq struct {
	UserName     string //用户名
	UserId       string //用户名
	Currency     string //币种
	VenueCode    string //场馆Code
	MerchantCode string // 商户Code(线路选择配置)
	Token        string
	Language     string //语言
	IP           string //客户端ip
}

// VenueRegister
//
//	@Description: 注册用户
//	@param req 请求结构体
//	@return *venueDetail.VenueResponse 返回信息
func VenueRegister(req *VenueRegisterReq) *venueDetail.VenueResponse {
	ret := &venueDetail.VenueResponse{
		Code: 0,
		Msg:  "",
	}

	//if req.VenueCode == enmus.BBIN { //bbin因为限制账号长度，用uid代替username
	//	req.UserName = req.UserId
	//}
	merchantVenue := modules.FindByKeyFcMerchantVenueFirst(&dos.FcMerchantVenue{VenueCode: req.VenueCode, MerchantCode: req.MerchantCode})
	if len(merchantVenue.Id) == 0 {
		global.G_LOG.Errorf("user=%s venue=%s merchant=%s FindByKeyFcMerchantVenueFirst not exist", req.UserName, req.VenueCode, req.MerchantCode)
		return ret
	}

	//获取商户配置线路
	vMethod, lineNum := venuesLineConfig(req.MerchantCode, req.VenueCode)
	if vMethod == nil {
		//商户未配置线路
		return ret
	}
	venueUser := modules.FindByKeyFcVenueUserFirst(&dos.FcVenueUser{UserName: req.UserName, VenueCode: req.VenueCode, MerchantCode: req.MerchantCode, VenueLine: lineNum})
	if len(venueUser.Id) > 0 {
		//用户已注册
		return ret
	}

	venueInfo := modules.FindByKeyFcVenueFirst(&dos.FcVenue{
		VenueCode: req.VenueCode,
	})
	merchant := modules.FindByKeyFcMerchantFirst(&dos.FcMerchant{
		MerchantCode: req.MerchantCode,
	})

	account := strings.ToLower(merchant.Prefix) + req.UserName
	if req.VenueCode == enmus.BBIN {
		account = req.UserId
	}
	switch req.VenueCode {
	case enmus.TYQP:
		if len(req.UserName) > 16 {
			req.UserName = req.UserName[len(req.UserName)-16:]
		}

	}

	password := tool.MakePassWord()
	var err error
	token := req.Token
	if req.VenueCode == enmus.HGTY { // 如果是 HGTY 需要獲取 token
		token, err = GetAndCreateToken(req.VenueCode)
		if err != nil {
			global.G_LOG.Errorf("user=%s venue=%s merchant=%s GetAndCreateHGTYToken err: %v",
				req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			return ret
		}
	}

	venueReq := &venueDetail.VenueCreateUserRequest{
		UserName: account,
		Password: password,
		Currency: req.Currency,
		Product:  req.VenueCode,
		Token:    token,
	}
	if req.VenueCode == enmus.LYQP {
		orderReq := vo.VenueGetOrderSnReq{}
		orderReq.VenueCode = req.VenueCode
		orderReq.UserName = venueReq.UserName
		venueReq.OrderSn = GetOrderSn(&orderReq)
		venueReq.MerchantCode = req.MerchantCode
	}

	ret = vMethod.CreateUser(venueReq)
	// 如果皇冠體育token失效，則需要再次獲取
	if req.VenueCode == enmus.HGTY {
		if ret.ThirdCode == "0002" {
			token, err = CreateToken(req.VenueCode)
			if err != nil {
				global.G_LOG.Errorf("user=%s venue=%s merchant=%s CreateHGTYToken err: %v",
					req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			} else {
				venueReq.Token = token
				ret = vMethod.CreateUser(venueReq)
			}
		}
	}

	if ret.Code == 0 {
		m := &dos.FcVenueUser{
			VenueCode:    venueInfo.VenueCode,
			VenueId:      venueInfo.Id,
			UserId:       req.UserId,
			UserName:     req.UserName,
			Account:      account,
			Password:     password,
			Currency:     req.Currency,
			MerchantCode: merchant.MerchantCode,
			VenueLine:    lineNum,
		}
		b, _ := modules.SaveFcVenueUser(m)
		if !b {
			global.G_LOG.Errorf(" VenueRegister save user fail : %v", tool.String(m))
		}
	} else {
		venueCreatUserRespBytes, _ := tool.JsonMarshal(&ret)
		global.G_LOG.Errorf("user=%s venue=%s merchant=%s venueCreatUserRespBytes=%s", req.UserName, req.VenueCode, req.MerchantCode, string(venueCreatUserRespBytes))
	}

	return ret
}
