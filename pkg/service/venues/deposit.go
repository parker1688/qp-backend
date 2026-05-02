package venues

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/userTransfer"
	"bootpkg/pkg/service/venues/venueDetail"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type VenueDepositRequest struct {
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

// VenueDeposit
//
//	@Description: 用户转入场馆
//	@param req 场馆操作
//	@return *venueDetail.VenueResponse 结构体
func VenueDeposit(req *VenueDepositRequest) *venueDetail.VenueDepositResponse {
	ret := &venueDetail.VenueDepositResponse{
		Code: 0,
		Msg:  "",
	}
	vMethod, lineNum := venuesLineConfig(req.MerchantCode, req.VenueCode)
	if lineNum < 0 {
		ret.Code = 500
		ret.Msg = " line does not exist "
		return ret
	}
	//global.G_LOG.Infof("venue-deposit--------------------------------1:%v", req.UserName)
	venueUser := modules.FindByKeyFcVenueUserFirst(&dos.FcVenueUser{UserName: req.UserName, VenueCode: req.VenueCode, MerchantCode: req.MerchantCode, VenueLine: lineNum})
	if len(venueUser.Id) == 0 {
		//用户未注册
		register := VenueRegister(&VenueRegisterReq{
			UserName:     req.UserName,
			UserId:       req.UserId,
			Currency:     req.Currency,
			VenueCode:    req.VenueCode,
			MerchantCode: req.MerchantCode,
		})
		if register.Code != 0 {
			ret.Code = 501
			ret.Msg = " register fail, err:" + register.Msg
			return ret
		}
		venueUser = modules.FindByKeyFcVenueUserFirst(&dos.FcVenueUser{UserName: req.UserName, VenueCode: req.VenueCode, MerchantCode: req.MerchantCode, VenueLine: lineNum})
	}
	newAmount := vMethod.AmountLimitFix(req.Amount, req.Currency)
	if newAmount == 0 && !req.IsAllAmount {
		ret.Code = 502
		ret.Msg = " amount is 0 "
		return ret
	}
	//global.G_LOG.Infof("venue-deposit--------------------------------2:%v", venueUser)
	wallet := modules.FindByKeyFcUserWalletFirst(&dos.FcUserWallet{
		UserId:   req.UserId,
		UserName: req.UserName,
		Currency: req.Currency,
	})
	if wallet.AvaAmount == 0 {
		ret.Code = 503
		ret.Msg = " amount is 0 "
		return ret
	}
	if req.IsAllAmount { //一键转入所有余额
		newAmount = wallet.AvaAmount
	}
	newAmount = vMethod.AmountLimitFix(newAmount, req.Currency)
	if newAmount == 0 {
		ret.Code = 502
		ret.Msg = " amount is 0 "
		return ret
	}

	subAmount := decimal.NewFromFloat(wallet.AvaAmount).Sub(decimal.NewFromFloat(newAmount)) //钱包金额不足
	if subAmount.InexactFloat64() < 0 {
		ret.Code = 503
		ret.Msg = "not enough wallet"
		return ret
	}
	orderSn := vMethod.GetOrderNo()
	//global.G_LOG.Infof("venue-deposit--------------------------------3:%v", req.UserName)
	var err error
	token := ""
	if req.VenueCode == enmus.HGTY { // 如果是 HGTY 需要獲取 token
		token, err = GetAndCreateToken(req.VenueCode)
		if err != nil {
			global.G_LOG.Errorf("user=%s venue=%s merchant=%s GetAndCreateHGTYToken err: %v",
				req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			ret.Code = 503
			ret.Msg = err.Error()
			return ret
		}
	} else if req.VenueCode == enmus.LYQP {
		orderReq := vo.VenueGetOrderSnReq{}
		orderReq.VenueCode = req.VenueCode
		orderReq.UserName = venueUser.Account
		orderSn = GetOrderSn(&orderReq)
	}

	transfer := &dos.FcVenueTransfer{
		OrderSn:      orderSn,
		VenueCode:    req.VenueCode,
		VenueLine:    lineNum,
		UserName:     req.UserName,
		Currency:     req.Currency,
		Amount:       newAmount,
		OptType:      1,
		Ip:           req.IP,
		CreateBy:     req.OptBy,
		Status:       0, //待处理
		UserId:       req.UserId,
		VenueAccount: venueUser.Account,
		MerchantCode: req.MerchantCode,
		CreateTime:   automaticType.Now(),
	}
	var isOk bool
	//事务操作金额
	global.G_DB.Transaction(func(tx *gorm.DB) error {
		eRow := tx.Create(transfer)
		if eRow.RowsAffected != 1 {
			return eRow.Error
		}
		//转入场馆
		err := userTransfer.UserVenueAmountChange(tx, newAmount, userTransfer.TranAmountConvert, req.Currency, "==>【"+venueUser.VenueCode+"】"+orderSn, req.UserId, req.OptBy, orderSn,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypePlatform, "deposit")+venueUser.VenueCode)
		if err == nil {
			isOk = true
		}
		return err
	})
	if !isOk {
		ret.Code = 504
		ret.Msg = "amount operate fail"
		return ret
	}

	venueReq := &venueDetail.VenueDepositRequest{
		UserName: venueUser.Account,
		Password: venueUser.Password,
		Currency: req.Currency,
		Amount:   newAmount,
		OrderSn:  orderSn,
		Product:  req.VenueCode,
		Token:    token,
	}
	deposit := vMethod.Deposit(venueReq)
	// 如果皇冠體育token失效，則需要再次獲取
	if req.VenueCode == enmus.HGTY {
		if deposit.ThirdCode == "0002" {
			token, err = CreateToken(req.VenueCode)
			if err != nil {
				global.G_LOG.Errorf("user=%s venue=%s merchant=%s CreateHGTYToken err: %v",
					req.UserName, req.VenueCode, req.MerchantCode, err.Error())
			} else {
				venueReq.Token = token
				deposit = vMethod.Deposit(venueReq)
			}
		}
	}

	//if deposit.Code == venueDetail.Deposit_FAIL_CODE || deposit.Code == venueDetail.Deposit_Processing_CODE {
	if deposit.Code == venueDetail.Deposit_FAIL_CODE {
		//转账失败金额操作
		global.G_DB.Transaction(func(tx *gorm.DB) error {
			eRow := tx.Model(&dos.FcVenueTransfer{}).Where("id = ? and status=0", transfer.Id).Update("status", 2)
			if eRow.RowsAffected != 1 {
				return eRow.Error
			}
			err := userTransfer.UserVenueAmountChangeCallback(tx, newAmount, userTransfer.TranAmountConvert, req.Currency, "==>【"+venueUser.VenueCode+"】"+orderSn+" 【Fail】", req.UserId, req.OptBy, orderSn,
				modules.GetFcTranscationFundingSubType(enmus.FundingTypePlatform, "deposit")+venueUser.VenueCode)
			return err
		})
		return deposit
	} else if deposit.Code == venueDetail.Deposit_SUCCESS_CODE {
		//成功操作
		global.G_DB.Transaction(func(tx *gorm.DB) error {
			eRow := tx.Model(&dos.FcVenueTransfer{}).Where("id = ? and status=0", transfer.Id).Update("status", 1)
			if eRow.RowsAffected != 1 {
				return eRow.Error
			}
			err := userTransfer.UserVenueAmountConfirmChange(tx, newAmount, userTransfer.TranAmountConvert, req.Currency, "==>【"+venueUser.VenueCode+"】"+orderSn+" 【Success】", req.UserId, req.OptBy)

			if err == nil {
				modules.SetVenueEntryRecordVal(req.UserId, req.VenueCode)
			}

			return err
		})
	}
	return deposit
}
