package venues

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/userTransfer"
	"bootpkg/pkg/service/venues/venueDetail"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type VenueWithdrawRequest struct {
	UserName     string //用户名
	UserId       string //用户ID
	Currency     string //币种
	VenueCode    string //场馆Code
	MerchantCode string //商户Code(线路选择配置)
	IP           string //用户IP
	OptBy        string //操作人员

	Amount      float64 //转账金额
	IsAllAmount bool    //是否所有金额
}

// VenueWithdraw
//
//	@Description: 用户转出场馆
//	@param req 场馆操作
//	@return *venueDetail.VenueResponse 结构体
func VenueWithdraw(req *VenueWithdrawRequest) *venueDetail.VenueWithdrawResponse {
	ret := &venueDetail.VenueWithdrawResponse{
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
		//用户未注册
		return ret
	}

	newAmount := vMethod.AmountLimitFix(req.Amount, req.Currency)
	if newAmount == 0 && !req.IsAllAmount {
		return ret
	}

	venueBalanceReq := &venueDetail.VenueGetUserBalanceRequest{
		UserName: venueUser.Account,
		Password: venueUser.Password,
		Currency: req.Currency,
		Product:  req.VenueCode,
	}
	var err error
	token := ""
	if req.VenueCode == enmus.HGTY { // 如果是 HGTY 需要獲取 token
		token, err = GetAndCreateToken(req.VenueCode)
		if err != nil {
			global.G_LOG.Errorf("user=%s venue=%s merchant=%s GetAndCreateHGTYToken err: %v",
				req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			ret.Code = 508
			ret.Msg = err.Error()
			return ret
		} else {
			venueBalanceReq.Token = token
		}
	}
	balance := vMethod.GetUserBalance(venueBalanceReq)

	// 如果皇冠體育token失效，則需要再次獲取
	if req.VenueCode == enmus.HGTY {
		if balance.ThirdCode == "0002" {
			token, err = CreateToken(req.VenueCode)
			if err != nil {
				global.G_LOG.Errorf("user=%s venue=%s merchant=%s CreateHGTYToken err: %v",
					req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			} else {
				venueBalanceReq.Token = token
				balance = vMethod.GetUserBalance(venueBalanceReq)
			}
		}
	}

	if balance.Data.Amount == 0 {
		modules.DelVenueEntryRecordVal(req.UserId)
		ret.Code = 507
		ret.Msg = " venue amount is 0 "
		return ret
	}

	if req.IsAllAmount {
		newAmount = balance.Data.Amount
	}

	newAmount = vMethod.AmountLimitFix(newAmount, req.Currency)
	if newAmount == 0 {
		return ret
	}

	subAmount := decimal.NewFromFloat(balance.Data.Amount).Sub(decimal.NewFromFloat(newAmount))

	if subAmount.InexactFloat64() < 0 { //场馆余额不足
		ret.Code = 508
		ret.Msg = "not enough venue"
		return ret
	}

	orderSn := vMethod.GetOrderNo()
	if req.VenueCode == enmus.LYQP {
		orderReq := vo.VenueGetOrderSnReq{}
		orderReq.VenueCode = req.VenueCode
		orderReq.UserName = venueUser.Account
		orderSn = GetOrderSn(&orderReq)
	}

	transfer := &dos.FcVenueTransfer{
		OrderSn:      orderSn,
		VenueCode:    req.VenueCode,
		VenueAccount: venueUser.Account,
		VenueLine:    lineNum,
		UserName:     req.UserName,
		Currency:     req.Currency,
		Amount:       newAmount,
		OptType:      2, //转出
		Ip:           req.IP,
		CreateBy:     req.OptBy,
		Status:       0, //待处理
		UserId:       req.UserId,
		MerchantCode: req.MerchantCode,
		CreateTime:   automaticType.Now(),
		UpdateTime:   automaticType.Now(),
	}
	var isOk bool
	//事务操作金额
	global.G_DB.Transaction(func(tx *gorm.DB) error {
		eRow := tx.Create(transfer)
		if eRow.RowsAffected != 1 {
			return eRow.Error
		} else {
			isOk = true
		}
		return eRow.Error
	})
	if !isOk {
		ret.Code = 504
		ret.Msg = "amount operate fail"
		return ret
	}

	venueReq := &venueDetail.VenueWithdrawRequest{
		UserName: venueUser.Account,
		Password: venueUser.Password,
		Currency: req.Currency,
		Amount:   newAmount,
		OrderSn:  orderSn,
		Product:  req.VenueCode,
		Token:    token,
	}

	withdraw := vMethod.Withdraw(venueReq)
	// 如果皇冠體育token失效，則需要再次獲取
	if req.VenueCode == enmus.HGTY {
		if withdraw.ThirdCode == "0002" {
			token, err = CreateToken(req.VenueCode)
			if err != nil {
				global.G_LOG.Errorf("user=%s venue=%s merchant=%s CreateHGTYToken err: %v",
					req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			} else {
				venueReq.Token = token
				withdraw = vMethod.Withdraw(venueReq)
			}
		}
	}

	if withdraw.Code == venueDetail.Withdraw_FAIL_CODE {
		//转账失败金额操作
		global.G_DB.Transaction(func(tx *gorm.DB) error {
			eRow := tx.Model(&dos.FcVenueTransfer{}).Where("id = ? and status=0", transfer.Id).Update("status", 2)
			return eRow.Error
		})

		return withdraw
	} else if withdraw.Code == venueDetail.Withdraw_SUCCESS_CODE {
		//成功操作
		global.G_DB.Transaction(func(tx *gorm.DB) error {
			eRow := tx.Model(&dos.FcVenueTransfer{}).Where("id = ? and status=0", transfer.Id).Update("status", 1)
			if eRow.RowsAffected != 1 {
				return eRow.Error
			}
			err := userTransfer.UserAmountChange(tx, newAmount, userTransfer.TranAmountConvert, req.Currency, "<===【"+venueUser.VenueCode+"】"+orderSn+" 【Success】", req.UserId, req.OptBy, orderSn,
				modules.GetFcTranscationFundingSubType(enmus.FundingTypePlatform, "withdraw")+venueUser.VenueCode)

			modules.DelVenueEntryRecordVal(req.UserId)

			return err
		})
	}
	global.G_REDIS.Set(context.Background(), fmt.Sprintf("venue-status:%v", req.UserId), "", 360*24*time.Hour)
	return withdraw
}

func VenueWithdrawTimeoutAsync(req []*VenueWithdrawRequest) {
	wg := &sync.WaitGroup{}
	for _, v := range req {
		wg.Add(1)
		go func(req *VenueWithdrawRequest) {
			defer wg.Done()
			VenueWithdraw(req)
		}(v)
	}
	tool.WaitTimeout(wg, 3*time.Second) //等待5秒
}
