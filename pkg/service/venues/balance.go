package venues

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/venues/venueDetail"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"sync"
	"time"
)

type VenueBalanceReq struct {
	UserName     string //用户名
	UserId       string //uid
	Currency     string //币种
	VenueCode    string //场馆Code
	MerchantCode string // 商户Code(线路选择配置)
}

// VenueBalance
//
//	@Description: 获取用户在场馆的余额
//	@param req 获取余额默认信息
//	@return *venueDetail.VenueGetUserBalanceResponse 返回信息
func VenueBalance(req *VenueBalanceReq) *venueDetail.VenueGetUserBalanceResponse {
	ret := &venueDetail.VenueGetUserBalanceResponse{
		Code: 0,
		Msg:  "",
		Data: struct {
			Amount float64 `json:"amount"`
		}{
			Amount: 0,
		},
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

	var err error
	token := ""
	if req.VenueCode == enmus.HGTY { // 如果是 HGTY 需要獲取 token
		token, err = GetAndCreateToken(req.VenueCode)
		if err != nil {
			global.G_LOG.Errorf("user=%s venue=%s merchant=%s GetAndCreateHGTYToken err: %v",
				req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			return ret
		}
	}

	venueReq := &venueDetail.VenueGetUserBalanceRequest{
		UserName: venueUser.Account,
		Password: venueUser.Password,
		Currency: req.Currency,
		Product:  req.VenueCode,
		Token:    token,
	}
	ret = vMethod.GetUserBalance(venueReq)
	// 如果皇冠體育token失效，則需要再次獲取
	if req.VenueCode == enmus.HGTY {
		if ret.ThirdCode == "0002" {
			token, err = CreateToken(req.VenueCode)
			if err != nil {
				global.G_LOG.Errorf("user=%s venue=%s merchant=%s CreateHGTYToken err: %v",
					req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			} else {
				venueReq.Token = token
				ret = vMethod.GetUserBalance(venueReq)
			}
		}
	}

	return ret
}

func VenueBalancesTimeoutAsync(req []*VenueBalanceReq) []*venueDetail.BatchVenueBalanceResponse {
	var venueBalance []*venueDetail.BatchVenueBalanceResponse
	wg := &sync.WaitGroup{}
	var mutex sync.Mutex
	for _, v := range req {
		wg.Add(1)
		go func(req *VenueBalanceReq) {
			defer wg.Done()
			ret := VenueBalance(req)
			mutex.Lock()
			venueBalance = append(venueBalance, &venueDetail.BatchVenueBalanceResponse{
				VenueCode: req.VenueCode,
				Amount:    ret.Data.Amount,
			})
			mutex.Unlock()
		}(v)
	}
	tool.WaitTimeout(wg, 3*time.Second) //等待3秒
	return venueBalance
}

// VenueBalancesAllAsync
//
//	@Description: 获取用户场馆所有余额
//	@param userInfo 用户信息
//	@param currency 币种
//	@return float64 总余额
func VenueBalancesAllAsync(userInfo *dos.FcUserMaterial, currency string) float64 {
	venueCode := global.G_REDIS.Get(context.Background(), fmt.Sprintf("venue-status:%v", userInfo.UserId)).Val()
	var venueBalanceResp []*venueDetail.BatchVenueBalanceResponse
	if venueCode != "" {
		var balanceRequest []*VenueBalanceReq
		balanceRequest = append(balanceRequest, &VenueBalanceReq{
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			VenueCode:    venueCode,
			Currency:     currency,
			MerchantCode: userInfo.MerchantCode,
		})
		venueBalanceResp = VenueBalancesTimeoutAsync(balanceRequest)
	}

	var sumAmount float64
	for _, v := range venueBalanceResp {
		sumAmount += decimal.NewFromFloat(v.Amount).Truncate(2).InexactFloat64()
	}
	return sumAmount
}
