package venues

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/venues/venueDetail"
	"bootpkg/pkg/service/venues/venuevo"
	"context"
	"fmt"
	"github.com/kirinlabs/utils"
	"time"
)

type VenueLaunchReq struct {
	UserName     string //用户名
	UserId       string //用户ID
	Currency     string //币种
	VenueCode    string //场馆Code
	MerchantCode string //商户Code(线路选择配置)
	IP           string //用户IP
	GameCode     string //单独进入游戏Code(可选)
	GType        string //游戏类型
	ClientType   string //客户端类型 h5、pc、android、 ios
	ReturnUrl    string //回调url
	CashierURL   string //存款地址
	TableId      string //桌号
	Token        string //token
	IsFree       bool   //试玩用户
	Language     string //语言
}

// VenueLaunch
//
//	@Description: 用户运行游戏
//	@param req 用户请求结构体
//	@return *venueDetail.VenueLoginGameResponse 用户运行游戏返回
func VenueLaunch(req *VenueLaunchReq) *venueDetail.VenueLoginGameResponse {
	ret := &venueDetail.VenueLoginGameResponse{
		Code: 0,
		Msg:  "",
	}
	vMethod, lineNum := venuesLineConfig(req.MerchantCode, req.VenueCode)
	if lineNum < 0 {
		ret.Code = 500
		ret.Msg = " line does not exist "
		return ret
	}

	venueUser := modules.FindByKeyFcVenueUserFirst(&dos.FcVenueUser{UserName: req.UserName, VenueCode: req.VenueCode, MerchantCode: req.MerchantCode, VenueLine: lineNum})
	if len(venueUser.Id) == 0 {
		venueRegistReq := &VenueRegisterReq{
			UserName:     req.UserName,
			UserId:       req.UserId,
			Currency:     req.Currency,
			VenueCode:    req.VenueCode,
			MerchantCode: req.MerchantCode,
			IP:           req.IP,
		}
		//用户未注册
		register := VenueRegister(venueRegistReq)
		if register.Code != 0 {
			tmpBytes, _ := tool.JsonMarshal(&register)
			global.G_LOG.Errorf("user=%s merchant=%s venue=%s register venue failed %s", req.UserName, req.MerchantCode, req.VenueCode, string(tmpBytes))
			ret.Code = 500
			ret.Msg = " register  fail, err:" + register.Msg
			return ret
		}
		venueUser = modules.FindByKeyFcVenueUserFirst(&dos.FcVenueUser{UserName: req.UserName, VenueCode: req.VenueCode, MerchantCode: req.MerchantCode, VenueLine: lineNum})
	}

	if len(venueUser.UserId) == 0 {
		global.G_LOG.Errorf("user=%s merchant=%s venue=%s venueUser not exist", req.UserName, req.MerchantCode, req.VenueCode)
		ret.Code = 500
		ret.Msg = "venue user is not find"
		return ret
	}

	if req.VenueCode == enmus.PGDZ || req.VenueCode == enmus.WUGDZ {
		tokenKey := ""
		if req.VenueCode == enmus.PGDZ {
			tokenKey = fmt.Sprintf(enmus.PG_CALLBACK_TOKEN_KEY, req.Token)
		} else if req.VenueCode == enmus.WUGDZ {
			tmpToken := req.MerchantCode + req.UserName + tool.GetRandStr(1, 16, false)
			tmpToken = tool.MD5([]byte(tmpToken))
			req.Token = tmpToken
			tokenKey = fmt.Sprintf(enmus.WUG_CALLBACK_TOKEN_KEY, tmpToken)
		}

		tokenData := venuevo.TokenData{
			UserName: venueUser.Account,
			NickName: venueUser.Account,
		}
		global.G_REDIS.Set(context.Background(), tokenKey, utils.Json(tokenData), 10*time.Minute)
	}

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

	venueReq := &venueDetail.VenueLoginGameRequest{
		UserId:     venueUser.UserId,
		UserName:   venueUser.Account,
		Password:   venueUser.Password,
		IP:         req.IP,
		GameCode:   req.GameCode,
		GType:      req.GType,
		ClientType: req.ClientType,
		Token:      token,
		ReturnUrl:  req.ReturnUrl,
		IsFree:     req.IsFree,
		CashierURL: req.CashierURL,
		Language:   req.Language,
		Product:    req.VenueCode,
		Currency:   req.Currency,
	}
	if req.VenueCode == enmus.LYQP {
		orderReq := vo.VenueGetOrderSnReq{}
		orderReq.VenueCode = req.VenueCode
		orderReq.UserName = venueReq.UserName
		venueReq.OrderSn = GetOrderSn(&orderReq)
		venueReq.MerchantCode = req.MerchantCode
	}

	vData := vMethod.LoginGame(venueReq)
	// 如果皇冠體育token失效，則需要再次獲取
	if req.VenueCode == enmus.HGTY {
		if vData.ThirdCode == "0002" {
			token, err = GetAndCreateToken(req.VenueCode)
			if err != nil {
				global.G_LOG.Errorf("user=%s venue=%s merchant=%s CreateHGTYToken err: %v",
					req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			} else {
				venueReq.Token = token
				vData = vMethod.LoginGame(venueReq)
			}
		}
	}
	return vData
}
