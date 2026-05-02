package walletControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/langs"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/payment"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func init() {

}

const (
	DEPOSIT_NO_KEY = "DEPOSIT_NO:%s"
)

func GetRechargeChannel(c *gin.Context) {
	var jsonp vo.UserPayChannelReq
	_ = c.ShouldBind(&jsonp)
	//if err != nil {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
	//	return
	//}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)

	var data []*dos.FcPayChannel
	query := global.G_DB.Model(&dos.FcPayChannel{})
	query.Order("sort desc").Where(`status = 1 and merchant_code = ? `, userInfo.MerchantCode)
	query.Find(&data)
	retData := make([]*vo.RechargeChannelRespVO, 0, len(data))
	for _, v := range data {
		//global.G_LOG.Infof("username: %s merchant_code: %s channel_code: %s level: %v", userInfo.UserName, userInfo.MerchantCode, v.ChannelCode, userInfo.Level)
		if v.MaxLevel >= userInfo.Level && v.MinLevel <= userInfo.Level {
			retData = append(retData, &vo.RechargeChannelRespVO{
				Id:          v.Id,
				ChannelName: v.ChannelName,
				ChannelCode: v.ChannelCode,
				ChannelType: v.ChannelType,
				Currency:    v.Currency,
				MinAmount:   v.MinAmount,
				MaxAmount:   v.MaxAmount,
				Icon:        v.Icon,
				Sort:        v.Sort,
				Hot:         v.Hot,

				AmountRange:            v.AmountRange,
				InputAmountDisplay:     v.InputAmountDisplay,
				InputNameDisplay:       v.InputNameDisplay,
				InputVirtualPayAddress: v.InputVirtualPayAddress,
				InputVirtualPayShow:    v.InputVirtualPayShow,
			})
		}
	}
	response.SuccessJSON(c, retData)
}

func GetRechargeChannelSetting(c *gin.Context) {
	var jsonp vo.UserPayChannelSettingReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	resp := payment.GetPayChannelSetting(jsonp.ChannelCode)
	data := &vo.UserPayChannelSettingResp{
		AmountRange:             resp.AmountRange,
		InputAmountDisplay:      resp.InputAmountDisplay,
		InputNameDisplay:        resp.InputNameDisplay,
		InputVirtualPayAddress:  resp.InputVirtualPayAddress,
		InputVirtualPayShow:     resp.InputVirtualPayShow,
		InputVirtualPayShowList: resp.InputVirtualPayShowList,
	}
	response.SuccessJSON(c, data)
}

func UserPaymentChannel(c *gin.Context) {
	var jsonp vo.UserRechargeReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	rechargeLockKey := fmt.Sprintf(enmus.MEMBER_REDIS_RECHARGE_APPLY_LOCK, userInfo.UserId)
	if !global.G_REDIS.SetNX(context.Background(), rechargeLockKey, "1", 5*time.Second).Val() {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作频繁，请稍后再试")
		return
	}
	defer global.G_REDIS.Del(context.Background(), rechargeLockKey)
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code

	// 支付限单（风控处理）
	if !modules.CheckUserPaymentStrategy(userInfo.UserId, jsonp.ChannelCode,
		userInfo.MerchantCode) {
		response.FailJSON(c, response.Payment_Strategy_Limit)
		return
	}

	if jsonp.ChannelCode == "" {
		global.G_LOG.Errorf("username: %s merchant_code: %s channel_code is empty", userInfo.MerchantCode, jsonp.ChannelCode)
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_28"))
		return
	}
	if jsonp.PaymentCode == "" {
		global.G_LOG.Errorf("username: %s merchant_code: %s channel_code is empty", userInfo.MerchantCode, jsonp.ChannelCode)
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_28"))
		return
	}
	if jsonp.PayId == "" {
		global.G_LOG.Errorf("username: %s merchant_code: %s channel_code is empty", userInfo.MerchantCode, jsonp.ChannelCode)
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_28"))
		return
	}

	// 获取商户的某个渠道
	var payChannel *dos.FcPayChannel
	global.G_DB.Model(&dos.FcPayChannel{}).Where("status = 1 and channel_code=? and merchant_code=?", jsonp.ChannelCode, userInfo.MerchantCode).Take(&payChannel)
	if !(payChannel.MaxLevel >= userInfo.Level && payChannel.MinLevel <= userInfo.Level) {
		global.G_LOG.Errorf("username: %s merchant_code: %s level: %v channel_code: %v min_level: %v max_level: %v",
			userInfo.UserName, userInfo.MerchantCode, userInfo.Level, jsonp.ChannelCode, payChannel.MinLevel, payChannel.MaxLevel)
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_28"))
		return
	}
	if jsonp.Amount < payChannel.MinAmount || jsonp.Amount > payChannel.MaxAmount {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("%s: %v ~ %v", langs.GetWithLocaleGin(c, "message_25"), payChannel.MinAmount, payChannel.MaxAmount))
		return
	}

	//// 获取商户的某个渠道的支付通道
	//var data []*dos.FcPayment
	//query := global.G_DB.Model(&dos.FcPayment{})
	//query.Order("sort desc").Where(`status = 1 and channel_code=? and merchant_code=?`, jsonp.ChannelCode, userInfo.MerchantCode)
	//query.Find(&data)
	//
	//var paymentData *dos.FcPayment
	////var minAmount float64 // 最小金额
	////选择通道
	//for i, v := range data {
	//	// 判断用户等级是否合格
	//	if userInfo.Level > v.MaxLevel || userInfo.Level < v.MinLevel {
	//		continue
	//	}
	//
	//	// 是否在通道的金额范围之内
	//	if jsonp.Amount > v.MaxAmount || jsonp.Amount < v.MinAmount {
	//		continue
	//	}
	//
	//	//tmpMinAmount := v.MinAmount
	//	//if tmpMinAmount > v.MinAmount {
	//	//	minAmount = tmpMinAmount
	//	//}
	//
	//	if len(v.AmountRange) > 0 { //固定金额通道
	//		aRange := tool.ToFloat64Zero(v.AmountRange)
	//		if len(aRange) < 1 {
	//			continue
	//		}
	//
	//		var fOk bool
	//		for _, av := range aRange {
	//			if av == jsonp.Amount {
	//				fOk = true
	//				break
	//			}
	//		}
	//
	//		if !fOk {
	//			continue
	//		} else {
	//			////判断今日总存款
	//			//if getIsPaymentChannel(v.Id, v.DayMaxAmount) {
	//			//	paymentData = v
	//			//	break
	//
	//			paymentData = data[i]
	//		}
	//	}
	//
	//	////判断今日总存款
	//	//if jsonp.Amount >= v.MinAmount && jsonp.Amount <= v.MaxAmount && getIsPaymentChannel(v.Id, v.DayMaxAmount) {
	//	//	paymentData = v
	//	//	break
	//	//}
	//}
	//
	//if paymentData == nil {
	//	global.G_LOG.Errorf("not found available payment data")
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_27"))
	//	return
	//}
	//if minAmount > jsonp.Amount {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_26")+":"+tool.String(minAmount))
	//	return
	//}

	// 获取商户的某个渠道的支付通道
	var paymentData *dos.FcPayment
	query := global.G_DB.Model(&dos.FcPayment{})
	query.Order("sort desc").Where(`status = 1 and channel_code=? and merchant_code=? and payment_code = ? and pay_id = ? and min_level <= ? and max_level >= ?`,
		jsonp.ChannelCode, userInfo.MerchantCode, jsonp.PaymentCode, jsonp.PayId, userInfo.Level, userInfo.Level).Scan(&paymentData)
	err = query.Take(&paymentData).Error
	if err != nil {
		global.G_LOG.Errorf("GetPayment FcPayment username: %v merchantCode: %v channelCode: %v query payment err: %v", userInfo.UserName, userInfo.MerchantCode, jsonp.ChannelCode, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if jsonp.Amount < paymentData.MinAmount || jsonp.Amount > paymentData.MaxAmount {
		global.G_LOG.Errorf("UserPaymentChannel GetFcPayment username: %s merchant_code: %s level: %v channel_code: %v paymentCode: %v minAmount: %v maxAmount: %v",
			userInfo.UserName, userInfo.MerchantCode, userInfo.Level, jsonp.ChannelCode, jsonp.PaymentCode, paymentData.MinAmount, paymentData.MaxAmount)
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("%s: %v ~ %v", langs.GetWithLocaleGin(c, "message_25"), paymentData.MinAmount, paymentData.MaxAmount))
		return
	}

	// USDT需要进行汇率转换
	jsonp.Amount, err = payment.YRPayAmountFxCovert(paymentData.ChannelCode == "USDT-CR20", jsonp.Amount)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	// 支付工厂是否存在该通道
	m := payment.GetRechargeChannel(paymentData.PaymentCode)
	if m == nil {
		global.G_LOG.Errorf("UserPaymentChannel username: %v paymentCode: %v is not available", userInfo.UserName, paymentData.PaymentCode)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "支付通道不可用")
		return
	}

	orderSn := getOrderSn()
	bonusAmount := jsonp.Amount * paymentData.BonusRate
	factAmount := jsonp.Amount + bonusAmount
	mResp := m.ChannelRecharge(paymentData, payment.OtherPaymentParam{
		OrderSn:      orderSn,
		UserInfo:     userInfo,
		Amount:       jsonp.Amount,
		BonusRate:    paymentData.BonusRate,
		BonusAmount:  bonusAmount,
		FactAmount:   factAmount,
		DepositName:  jsonp.DepositName,
		Currency:     payChannel.Currency,
		MerchantCode: userInfo.MerchantCode,
		ReturnUrl:    jsonp.ReturnUrl,
		ClientIp:     tool.ClientIP(c),
	})
	if !mResp.Success {
		global.G_LOG.Errorf("UserPaymentChannel username: %v amount: %v channelName: %v payName: %v payCode: %v err: %v",
			userInfo.UserName, jsonp.Amount, paymentData.ChannelName, paymentData.PaymentName, paymentData.PaymentCode, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, mResp.ErrorMsg)
		return
	}

	payId, err := strconv.Atoi(paymentData.PayId)
	if err != nil {
		payId = 0
	}

	//处理用户添加金额订单
	deposit := &dos.FcOrderDeposit{
		UserId:          userInfo.UserId,
		UserName:        userInfo.UserName,
		OrderSn:         orderSn,
		Amount:          jsonp.Amount,
		BonusRate:       paymentData.BonusRate,
		BonusAmount:     bonusAmount,
		FactAmount:      factAmount,
		Status:          enmus.Order_STATUS_PENDING_PAY, //三方回调处理
		DepositRemark:   jsonp.DepositRemark,
		Currency:        payChannel.Currency,
		Ip:              c.ClientIP(),
		CreateBy:        userInfo.UserName,
		MerchantCode:    userInfo.MerchantCode,
		InviteCode:      userInfo.AgentInviteCode,
		ChannelCode:     jsonp.ChannelCode,
		PaymentCode:     paymentData.PaymentCode,
		PaymentName:     paymentData.PaymentName,
		PayAliasName:    paymentData.PayAliasName,
		PaymentId:       payId,
		OrderType:       payChannel.ChannelType,
		OrderSecondType: payChannel.ChannelType,
		Level:           userInfo.Level,   //用户等级区分新老用户
		ActivityId:      jsonp.ActivityId, //用户选择参与的活动(可以自动选择)
		CreateTime:      automaticType.Time(time.Now()),
	}
	if paymentData.FeeRate > 0 {
		deposit.FeeRate = paymentData.FeeRate
		deposit.Fee = decimal.NewFromFloat(deposit.Amount).Mul(decimal.NewFromFloat(paymentData.FeeRate / 100)).Truncate(2).InexactFloat64()
	}

	switch payChannel.ChannelType {
	case enmus.Recharge_Order_Type_Wx:
	case enmus.Recharge_Order_Type_Bank:
	case enmus.Recharge_Order_Type_Alipay:
	case enmus.Recharge_Order_Type_NumberCNY:
	case enmus.Recharge_Order_Type_Virtual:
		//deposit.VirtualType = mResp.Data.VirtualCoin.CurrencyName
		//deposit.VirtualCurrencyChain = mResp.Data.VirtualCoin.CurrencyChain
		//deposit.VirtualAddress = mResp.Data.VirtualCoin.ToAddr
		//deposit.VirtualFx = mResp.Data.VirtualCoin.Fx
		//deposit.VirtualNum = mResp.Data.VirtualCoin.Num
		//deposit.VirtualPayAddress = jsonp.VirtualPayAddress
	}

	if paymentData.PaymentCode == payment.PaymentCode_LOCAL_BANK || paymentData.PaymentCode == payment.PaymentCode_LOCAL_VIRTUAL {
		//deposit.Status = enmus.ORDER_PENDING_STATUS //财务处理
	}

	//if paymentData.PaymentCode == payment.PaymentCode_LOCAL_BANK && mResp.Data.Style == payment.STYLE_BANK { //本地银行卡
	//	//deposit.OrderType = enmus.ORDER_TYPE_BANK
	//	deposit.EntityAccountNumber = mResp.Data.Bank.EntityAccountNumber
	//	deposit.EntityAccountBankName = mResp.Data.Bank.EntityAccountBankName
	//	deposit.EntityAccountHolder = mResp.Data.Bank.EntityAccountHolder
	//} else if paymentData.PaymentCode == payment.PaymentCode_LOCAL_VIRTUAL && mResp.Data.Style == payment.STYLE_VirtualCoin { //本地虚拟币
	//	//deposit.OrderType = enmus.ORDER_TYPE_Virtual
	//	deposit.VirtualType = mResp.Data.VirtualCoin.CurrencyName
	//	deposit.VirtualCurrencyChain = mResp.Data.VirtualCoin.CurrencyChain
	//	deposit.VirtualAddress = mResp.Data.VirtualCoin.ToAddr
	//	deposit.VirtualFx = mResp.Data.VirtualCoin.Fx
	//	deposit.VirtualNum = mResp.Data.VirtualCoin.Num
	//	deposit.VirtualPayAddress = jsonp.VirtualPayAddress
	//} else if mResp.Data.Style == payment.STYLE_Qr {
	//	deposit.DepositRemark += ":" + mResp.Data.Qr.QrImg
	//	mResp.Data.Qr.Expiration = 5
	//} else {
	//	deposit.OrderType = enmus.ORDER_TYPE_Other
	//}

	isSuc, _ := modules.SaveFcOrderDeposit(deposit)
	if !isSuc {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "记录订单失败,请重试")
		return
	}

	addPaymentChannel(paymentData.Id, jsonp.Amount)
	mResp.Data.OrderSn = orderSn
	mResp.Data.Amount = jsonp.Amount
	response.SuccessJSON(c, mResp.Data)
}

// getIsPaymentChannel
//
//	@Description: 判断是否达到每日最大金额
//	@param paymentId 通道ID
//	@param amount 支付金额
//	@param dayMaxAmount 每日最大金额
//	@return bool true: 可以充值  false 已经最大值
func getIsPaymentChannel(paymentId string, dayMaxAmount float64) bool {
	value := global.G_REDIS.Get(context.Background(), fmt.Sprintf(vo.Recharge_DayMaxAmount_KEY_R, paymentId))
	valueInt := tool.Int(value)
	if valueInt > 0 && valueInt >= int64(dayMaxAmount) {
		return false
	}
	return true
}

// addPaymentChannel
//
//	@Description: 存款渠道添加存款金额
//	@param paymentId 通道ID
//	@param addAmount 存款金额
func addPaymentChannel(paymentId string, addAmount float64) {
	key := fmt.Sprintf(vo.Recharge_DayMaxAmount_KEY_R, paymentId)
	global.G_REDIS.IncrBy(context.Background(), key, int64(addAmount))
	t := tool.TimeTomorrowTime()
	global.G_REDIS.Expire(context.Background(), key, t)
}

func getOrderSn() string {
	orderNoKey := time.Now().Format("20060102")
	orderNoPre := time.Now().Format("20060102150405")
	orderNoAdd := global.G_REDIS.Incr(context.Background(), fmt.Sprintf(DEPOSIT_NO_KEY, orderNoKey)).Val()
	orderSn := orderNoPre + tool.RandString(3) + tool.String(orderNoAdd)
	return orderSn
}

// UpdateDepositOrderStatus
//
//	@Description: 待支付改为待处理
//	@param c
func UpdateDepositOrderStatus(c *gin.Context) {
	var jsonp vo.DepositOrderStatusReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	eRow := global.G_DB.Model(&dos.FcOrderDeposit{}).
		Where("user_id = ? and status = ? and order_sn = ?", userInfo.UserId, enmus.Order_STATUS_PENDING_PAY, jsonp.OrderSn).
		Update("status", enmus.ORDER_STATUS_WAIT)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, eRow.Error.Error())
		return
	}
	if eRow.RowsAffected != 1 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单状态已变更，请刷新后重试")
		return
	}
	response.SuccessJSON(c, true)
}

// GetDepositOrderStatus
//
//	@Description: 获取订单状态
//	@param c
func GetDepositOrderStatus(c *gin.Context) {
	var jsonp vo.DepositOrderStatusReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	var data dos.FcOrderDeposit
	global.G_DB.Model(&dos.FcOrderDeposit{}).Where("user_id = ? and order_sn = ?", userInfo.UserId, jsonp.OrderSn).Take(&data)
	response.SuccessJSON(c, data.Status)
}

// OrderDepositInfo
//
//	@Description: 获取存款订单列表
//	@param c
func OrderDepositInfo(c *gin.Context) {
	var jsonp struct {
		vo.OrderDepositInfoReq
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
	userInfo := userInfoF.(*dos.FcUserMaterial)

	var data []*dos.FcOrderDeposit
	query := global.G_DB.Model(&dos.FcOrderDeposit{}).
		Where("user_id = ?", userInfo.UserId).
		Order("create_time desc")

	if jsonp.Status > 0 {
		query = query.Where("status = ?", jsonp.Status)
	}
	if jsonp.TimeType == nil {
		if jsonp.StartAt != "" {
			query.Where("create_time >=?", jsonp.StartAt)
		}

		if jsonp.EndAt != "" {
			query.Where("create_time <=?", jsonp.EndAt)
		}
	} else {
		sTime, eTime := tool.GetDayRange(time.Now(), *jsonp.TimeType)
		query.Where("create_time BETWEEN ? AND ?", sTime, eTime)
	}

	var total int64
	query.Count(&total)
	query.Offset((jsonp.Current - 1) * jsonp.PageSize).Limit(jsonp.PageSize).
		Scan(&data)

	newData := make([]*vo.OrderDepositInfoResp, 0, len(data))
	tool.JsonMapper(data, &newData)
	response.SuccessPageJSON(c, jsonp.Current, jsonp.PageSize, total, newData)
}

func GetPayment(c *gin.Context) {
	var jsonp vo.GetPaymentReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code

	if jsonp.ChannelCode == "" {
		global.G_LOG.Errorf("username: %s merchant_code: %s channel_code is empty", userInfo.MerchantCode, jsonp.ChannelCode)
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_28"))
		return
	}

	// 获取商户的某个渠道
	var payChannel *dos.FcPayChannel
	err = global.G_DB.Model(&dos.FcPayChannel{}).Where("status = 1 and channel_code=? and merchant_code=? and min_level <= ? and max_level >= ?", jsonp.ChannelCode, userInfo.MerchantCode, userInfo.Level, userInfo.Level).Take(&payChannel).Error
	if err != nil {
		global.G_LOG.Errorf("FcPayChannel username: %s merchant_code: %s level: %v channel_code: %v min_level: %v max_level: %v err: %v",
			userInfo.UserName, userInfo.MerchantCode, userInfo.Level, jsonp.ChannelCode, payChannel.MinLevel, payChannel.MaxLevel, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_28"))
		return
	}

	// 获取商户的某个渠道的支付通道
	var data []*dos.FcPayment
	query := global.G_DB.Model(&dos.FcPayment{})
	query.Order("sort desc").Where(`status = 1 and channel_code=? and merchant_code=? and min_level <= ? and max_level >= ?`, jsonp.ChannelCode, userInfo.MerchantCode, userInfo.Level, userInfo.Level).Scan(&data)
	err = query.Find(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.SuccessJSON(c, []struct{}{})
			return
		}

		global.G_LOG.Errorf("GetPayment FcPayment username: %v merchantCode: %v channelCode: %v query payment err: %v", userInfo.UserName, userInfo.MerchantCode, jsonp.ChannelCode, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	payDataArr := make([]*vo.GetPaymentResp, 0, len(data))
	tool.JsonMapper(&data, &payDataArr)
	response.SuccessJSON(c, payDataArr)
}

func GetPaymentDetail(c *gin.Context) {
	var jsonp vo.GetPaymentDetailReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G) //商户Code

	// 获取通道
	var paymentData *dos.FcPayment
	err = global.G_DB.Model(&dos.FcPayment{}).Where("id = ? and status = 1", jsonp.Id).Take(&paymentData).Error
	if err != nil {
		global.G_LOG.Errorf("GetPaymentDetail FcPayment id: %v username: %s merchant_code: %s  err: %v",
			jsonp.Id, userInfo.UserName, userInfo.MerchantCode, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	resp := vo.GetPaymentResp{}
	tool.JsonMapper(paymentData, &resp)
	response.SuccessJSON(c, resp)
}
