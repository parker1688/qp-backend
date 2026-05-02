package venueControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/venues"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Register(c *gin.Context) {
	token := c.GetHeader("Token")
	var jsonp vo.VenueRegisterReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}

	venueRegisterReq := &venues.VenueRegisterReq{
		UserName:     userInfo.UserName,
		UserId:       userInfo.UserId,
		VenueCode:    jsonp.VenueCode,
		MerchantCode: userInfo.MerchantCode,
		Token:        token,
		Currency:     jsonp.Currency,
		Language:     language,
	}
	venueRegisterResp := venues.VenueRegister(venueRegisterReq)
	if venueRegisterResp.Code != 0 {
		response.FailErrCodeJSON(c, venueRegisterResp.Code, venueRegisterResp.Msg, nil)
		return
	}
	response.SuccessMsgJSON(c, struct{}{}, "注册成功")
}

func Balance(c *gin.Context) {
	var jsonp vo.VenueBalanceReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	key := fmt.Sprintf("venue_balance:%s:%s:%s", userInfo.UserId, jsonp.VenueCode, jsonp.Currency)
	amount := global.G_REDIS.Get(context.Background(), key).Val()
	if amount == "" {
		venueBalanceReq := &venues.VenueBalanceReq{
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			Currency:     jsonp.Currency,
			VenueCode:    jsonp.VenueCode,
			MerchantCode: userInfo.MerchantCode,
		}
		venueBalanceResp := venues.VenueBalance(venueBalanceReq)
		if venueBalanceResp.Code != 0 {
			response.FailErrCodeJSON(c, venueBalanceResp.Code, venueBalanceResp.Msg, nil)
			return
		}
		global.G_REDIS.Set(context.Background(), key, venueBalanceResp.Data.Amount, 3*time.Minute)
		response.SuccessMsgJSON(c, vo.VenueBalanceResp{
			Balance: venueBalanceResp.Data.Amount,
		}, "操作成功")
	} else {
		amount2, err2 := strconv.ParseFloat(amount, 64)
		if err2 != nil {
			amount2 = 0.0
		}
		response.SuccessMsgJSON(c, vo.VenueBalanceResp{
			Balance: amount2,
		}, "操作成功")
	}
}

func TransferConfirm(c *gin.Context) {
	var jsonp vo.VenueTransferConfirmReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	venueTransferConfirmReq := &venues.VenueTransferConfirmRequest{
		UserName:     userInfo.UserName,
		VenueCode:    jsonp.VenueCode,
		MerchantCode: userInfo.MerchantCode,
		OrderSn:      jsonp.OrderSn,
	}
	venueTransferConfirmResp := venues.VenueTransferConfirm(venueTransferConfirmReq)
	if venueTransferConfirmResp.Code != 0 {
		response.FailErrCodeJSON(c, venueTransferConfirmResp.Code, venueTransferConfirmResp.Msg, nil)
		return
	}

	response.SuccessMsgJSON(c, venueTransferConfirmResp.Code, "操作成功")
}

func Deposit(c *gin.Context) {
	var jsonp vo.VenueTransferReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	venueDepositRequest := &venues.VenueDepositRequest{
		UserName:     userInfo.UserName,
		UserId:       userInfo.UserId,
		VenueCode:    jsonp.VenueCode,
		MerchantCode: userInfo.MerchantCode,
		IP:           c.ClientIP(),
		OptBy:        userInfo.UserName,
		Amount:       jsonp.Amount,
		Currency:     jsonp.Currency,
	}
	venueDepositResp := venues.VenueDeposit(venueDepositRequest)
	msg := venueDepositResp.Msg
	if venueDepositResp.Code != 0 {
		response.FailErrCodeJSON(c, venueDepositResp.Code, venueDepositResp.Msg, nil)
		return
	}
	response.SuccessMsgJSON(c, venueDepositResp.Code, msg)
}

func Withdraw(c *gin.Context) {
	var jsonp vo.VenueTransferReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	venueWithdrawRequest := &venues.VenueWithdrawRequest{
		UserName:     userInfo.UserName,
		UserId:       userInfo.UserId,
		VenueCode:    jsonp.VenueCode,
		MerchantCode: userInfo.MerchantCode,
		IP:           c.ClientIP(),
		OptBy:        userInfo.UserName,
		Amount:       jsonp.Amount,
		Currency:     jsonp.Currency,
	}
	venueDepositResp := venues.VenueWithdraw(venueWithdrawRequest)
	msg := venueDepositResp.Msg

	if venueDepositResp.Code != 0 {
		response.FailErrCodeJSON(c, venueDepositResp.Code, msg, nil)
		return
	}
	response.SuccessMsgJSON(c, venueDepositResp.Code, msg)
}

func Launch(c *gin.Context) {
	var jsonp vo.VenueLaunchReq

	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	entryVenueCode := modules.GetVenueEntryRecordVal(userInfo.UserId)
	if len(entryVenueCode) > 0 && entryVenueCode != jsonp.VenueCode {
		response.FailErrJSON(c, response.ERROR_SERVER, fmt.Sprintf("上个场馆资金还未转出暂时无法进入新场馆%v", entryVenueCode))
		return
	}

	venuesInfo := venues.GetVenuesStatus(userInfo.MerchantCode, jsonp.VenueCode)
	switch venuesInfo.Status {
	case 3:
		response.FailErrCodeJSON(c, 500, fmt.Sprintf("maintenance : %v  - %v", venuesInfo.MaintainStartTime, venuesInfo.MaintainEndTime), nil)
		return
	case 2:
		response.FailErrCodeJSON(c, 500, "Game suspended service", nil)
		return
	}
	venuesTypeInfo := venues.GetVenuesType(jsonp.VenueCode)
	//开启免转功能
	if userInfo.ManualTransferWallet == false && !userInfo.IsFree && venuesTypeInfo.VenueType == enmus.Venue_Type_Wallet_Transfer {
		venueCode := global.G_REDIS.Get(context.Background(), fmt.Sprintf("venue-status:%v", userInfo.UserId)).Val()
		if venueCode != "" && jsonp.VenueCode != venueCode {
			var withdrawRequest []*venues.VenueWithdrawRequest
			withdrawRequest = append(withdrawRequest, &venues.VenueWithdrawRequest{
				UserName:     userInfo.UserName,
				UserId:       userInfo.UserId,
				VenueCode:    venueCode,
				MerchantCode: userInfo.MerchantCode,
				IP:           c.ClientIP(),
				OptBy:        "",
				Amount:       0,
				Currency:     jsonp.Currency,
				IsAllAmount:  true, //金额全部回收
			})
			venues.VenueWithdrawTimeoutAsync(withdrawRequest)
		}

		venueDepositRequest := &venues.VenueDepositRequest{
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			VenueCode:    jsonp.VenueCode,
			MerchantCode: userInfo.MerchantCode,
			IP:           c.ClientIP(),
			OptBy:        userInfo.UserName,
			Amount:       0,
			Currency:     jsonp.Currency,
			IsAllAmount:  true, //全部金额
		}
		if venuesInfo.InVenueCode == 0 {
			venues.VenueDeposit(venueDepositRequest)
		}
	}
	token := c.GetHeader("Token")
	venueLaunchReq := &venues.VenueLaunchReq{
		UserName:     userInfo.UserName,
		UserId:       userInfo.UserId,
		VenueCode:    jsonp.VenueCode,
		MerchantCode: userInfo.MerchantCode,
		IP:           c.ClientIP(),
		GameCode:     jsonp.GameCode,
		GType:        tool.String(jsonp.Gtype),
		Currency:     jsonp.Currency,
		ReturnUrl:    jsonp.ReturnUrl,
		CashierURL:   jsonp.CashierURL,
		TableId:      jsonp.TableId,
		Token:        token,
		IsFree:       bool(userInfo.IsFree),
		Language:     language,
	}

	venueLaunchResp := venues.VenueLaunch(venueLaunchReq)

	if venueLaunchResp.Code != 0 {
		response.FailErrCodeJSON(c, venueLaunchResp.Code, venueLaunchResp.Msg, nil)
		return
	}
	global.G_REDIS.Set(context.Background(), fmt.Sprintf("venue-status:%s", userInfo.UserId), jsonp.VenueCode, 360*24*time.Hour)
	response.SuccessMsgJSON(c, venueLaunchResp.Data, "Success")

}
