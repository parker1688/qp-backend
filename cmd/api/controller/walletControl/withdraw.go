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
	"bootpkg/pkg/service/userTransfer"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func Withdraw(c *gin.Context) {
	var jsonp vo.WithdrawReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if jsonp.Amount <= 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_1"))
		return
	}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	withdrawLockKey := fmt.Sprintf(enmus.MEMBER_REDIS_WITHDRAW_APPLY_LOCK, userInfo.UserId)
	if !global.G_REDIS.SetNX(context.Background(), withdrawLockKey, "1", 5*time.Second).Val() {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作频繁，请稍后再试")
		return
	}
	defer global.G_REDIS.Del(context.Background(), withdrawLockKey)
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)

	//if userInfo.WalletPassword == "" {
	//	response.FailErrJSON(c, ecode.WALLET_PASSWORD_ERR, langs.GetWithLocaleGin(c, "message_2"))
	//	return
	//}
	////
	////if jsonp.WalletPassword == "" {
	////	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_3"))
	////	return
	////}
	////
	//if encrypt.Sha256(jsonp.WalletPassword+global.CONFIG.General.ApiSHA256Salt) != userInfo.WalletPassword {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_4"))
	//	return
	//}

	//var first dos.FcUserCodeBetLimit
	//global.G_DB.Model(&dos.FcUserCodeBetLimit{}).Where("user_id = ? and (status = 1 or status = 2)", userInfo.UserId).Take(&first)
	//if len(first.Id) > 0 {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_14", &langs.Replacements{"amount": tool.String(first.MinBetAmount)}))
	//	return
	//}

	//aw := getAlreadyWithdrawInfo(userInfo)
	//
	//if jsonp.Amount < aw.MinWithdrawAmount {
	//	response.FailErrJSON(c, ecode.MinWithdrawalAmount, langs.GetWithLocaleGin(c, "message_5", &langs.Replacements{"amount": tool.String(aw.MinWithdrawAmount)}))
	//	return
	//}
	//
	//newAmount := decimal.NewFromFloat(aw.Amount).Sub(decimal.NewFromFloat(jsonp.Amount))
	//if aw.Num-1 < 0 {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_6"))
	//	return
	//}
	//if newAmount.InexactFloat64() < 0 {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_7"))
	//	return
	//}

	//var count int64
	//global.G_DB.Model(&dos.FcOrderWithdraw{}).Where("user_id = ? AND status <= 1", userInfo.UserId).Count(&count)

	//if count > 0 {
	if modules.IsWithdrawOrderTodo(userInfo.UserId) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "您当前一笔订单处理中，请联系客服")
		return
	}

	/*report := modules.FindByKeyFcUserReportFirst(&dos.FcUserReport{
		UserId: userInfo.UserId,
	})*/
	depositWithdrawSubAmount := modules.GetDepositWithdrawSubAmount(userInfo.UserId) //decimal.NewFromFloat(report.RechargeAmount).Sub(decimal.NewFromFloat(report.WithdrawalAmount)).Truncate(2).InexactFloat64()
	orderSn := getOrderSn()
	withdraw := &dos.FcOrderWithdraw{
		UserId:                   userInfo.UserId,
		UserName:                 userInfo.UserName,
		Amount:                   jsonp.Amount,
		Status:                   enmus.OrderWithdrawStats_AuditWait,
		Ip:                       c.ClientIP(),
		CreateBy:                 userInfo.UserName,
		MerchantCode:             userInfo.MerchantCode,
		OrderSn:                  orderSn,
		Currency:                 jsonp.Currency,
		PreAmount:                jsonp.Amount,
		CreateTime:               automaticType.Now(),
		DepositWithdrawSubAmount: depositWithdrawSubAmount,
		AnotherPayStatus:         enmus.OrderWithdrawAnotherPayStats_None,
	}

	channelOutInfo := modules.FindByKeyFcPayChannelOutFirst(&dos.FcPayChannelOut{ChannelCode: jsonp.ChannelCode})

	//vipInfo := modules.FindByKeyFcVipFirst(&dos.FcVip{Level: userInfo.Level})
	if channelOutInfo.FeeRate > 0 {
		//手续费计算
		withdraw.FeeRate = channelOutInfo.FeeRate
		fee := decimal.NewFromFloat(withdraw.Amount).Mul(decimal.NewFromFloat(withdraw.FeeRate / 100))
		withdraw.Fee = fee.Truncate(2).InexactFloat64()
		preAmount := decimal.NewFromFloat(withdraw.Amount).Sub(decimal.NewFromFloat(withdraw.Fee))
		withdraw.PreAmount = preAmount.Truncate(2).InexactFloat64()
	}
	if jsonp.WithdrawType == enmus.ORDER_TYPE_BANK { //银行
		m := &dos.FcUserWithdrawBankBind{UserId: userInfo.UserId, Currency: jsonp.Currency}
		m.Id = jsonp.BankId
		bankInfo := modules.FindByKeyFcUserWithdrawBankBindFirst(m)
		if len(bankInfo.Id) == 0 {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "银行卡不存在")
			return
		}
		if bankInfo.Currency != jsonp.Currency {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "银行卡货币类型正确")
			return
		}
		//bankInfo.Decrypt()
		withdraw.Province = bankInfo.Province
		withdraw.City = bankInfo.City
		withdraw.BankAddress = bankInfo.BankAddress
		withdraw.AccountNumber = bankInfo.AccountNumber
		withdraw.AccountHolder = bankInfo.AccountHolder
		withdraw.AccountBankType = bankInfo.AccountBankType
		withdraw.AccountBankCode = bankInfo.AccountBankCode
		withdraw.Currency = bankInfo.Currency
		withdraw.OrderType = enmus.ORDER_TYPE_BANK
	} else if jsonp.WithdrawType == enmus.ORDER_TYPE_Virtual { //虚拟币
		m := &dos.FcUserWithdrawBlockchainBind{UserId: userInfo.UserId}
		m.Id = jsonp.VirtualId
		blockchainInfo := modules.FindByKeyFcUserWithdrawBlockchainBindFirst(m)
		if len(blockchainInfo.Id) == 0 {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "虚拟币不存在")
			return
		}
		//fx := modules.FindByKeyFcVirtualCurrencyFxFirst(&dos.FcVirtualCurrencyFx{
		//	OptType:       2,
		//	CurrencyChain: blockchainInfo.Blockchain,
		//	CurrencyName:  blockchainInfo.ContractType,
		//	CurrencyCode:  jsonp.Currency,
		//})
		//if len(fx.Id) == 0 {
		//	response.FailErrJSON(c, response.ERROR_PARAMETER, "汇率不存在")
		//	return
		//}
		//numD := decimal.NewFromFloat(jsonp.Amount).Div(decimal.NewFromFloat(fx.FxAmount))
		//num := numD.Truncate(2).InexactFloat64()
		withdraw.VirtualAddress = blockchainInfo.BlockchainAddress
		withdraw.VirtualType = blockchainInfo.ContractType
		withdraw.AccountHolder = blockchainInfo.RealName
		//withdraw.VirtualNum = num
		//withdraw.VirtualFx = fx.FxAmount
		withdraw.VirtualCurrencyChain = blockchainInfo.Blockchain
		withdraw.OrderType = enmus.ORDER_TYPE_Virtual
	} else if jsonp.WithdrawType == enmus.ORDER_TYPE_Online {
		onlineInfo := modules.FindByKeyFcUserWithdrawOnlineBindFirst(&dos.FcUserWithdrawOnlineBind{BaseDos: dos.BaseDos{Id: jsonp.OnlineId}})
		withdraw.AccountNumber = onlineInfo.AccountNumber
		withdraw.AccountHolder = onlineInfo.AccountHolder
		withdraw.OrderType = enmus.ORDER_TYPE_Online
	} else {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "提款类型不存在")
		return
	}
	isOk, err := userTransfer.UserWithdraw(withdraw)
	if err != nil || !isOk {
		response.FailErrJSON(c, response.ERROR_PARAMETER, response.ErrorLanguage(c, err))
		return
	}
	global.G_DB.Model(&dos.FcTranscation{}).Where("related_id=?", withdraw.OrderSn).Updates(map[string]interface{}{
		"status": 0,
	})
	//global.G_DB.Model(&dos.FcTranscation{}).Where("").Updates(map[string]interface{}{})
	response.SuccessMsgJSON(c, orderSn, "提现申请成功")
}

// GetWithdrawUserInfoToday
//
//	@Description: 今日剩余提款信息
//	@param c
func GetWithdrawUserInfoToday(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	//user := srv.GetUserMaterial(userInfo.UserName)
	//global.G_LOG.Infof("redis user: %s", tool.String(user))
	data := getAlreadyWithdrawInfo(userInfo)
	response.SuccessJSON(c, data)
}

func getAlreadyWithdrawInfo(userInfo *dos.FcUserMaterial) *vo.GetWithdrawUserInfoTodayResp {
	vip := modules.FindByKeyFcVipFirst(&dos.FcVip{
		Level: userInfo.Level,
	})
	//每日提现次数
	num := vip.DailyWithdrawalTimes
	//每日提现金额
	amount := vip.DailyWithdrawalAmount
	start, end := tool.TodayStartEndDate()
	var count int64
	global.G_DB.Model(&dos.FcOrderWithdraw{}).Where("user_id = ? and create_time >=? and create_time <= ? and status in (0,1,3)", userInfo.UserId, start, end).Count(&count)

	var sumAmount float64
	global.G_DB.Model(&dos.FcOrderWithdraw{}).Select("sum(amount)").Where("user_id = ? and create_time >=? and create_time <= ? and status in (0,1,3)", userInfo.UserId, start, end).Scan(&sumAmount)

	newAmount := decimal.NewFromFloat(amount).Sub(decimal.NewFromFloat(sumAmount))

	data := &vo.GetWithdrawUserInfoTodayResp{
		Amount:            newAmount.Truncate(2).InexactFloat64(),
		Num:               num - int(count),
		MinWithdrawAmount: vip.MinWithdrawAmount,
		MinRechargeAmount: vip.MinRechargeAmount,
		Vip:               userInfo.Level,
	}
	return data
}

// OrderWithdrawInfo
//
//	@Description: 获取提款订单列表
//	@param c
func OrderWithdrawInfo(c *gin.Context) {
	var jsonp struct {
		vo.OrderWithdrawInfoReq
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
	if jsonp.PageIndex == 0 {
		jsonp.PageIndex = 1
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)

	var data []*dos.FcOrderWithdraw
	query := global.G_DB.Model(&dos.FcOrderWithdraw{}).
		Where("user_id = ? ", userInfo.UserId).
		Order("create_time desc")

	if jsonp.TimeType == nil {
		if jsonp.StartTime != "" {
			query.Where("create_time >=?", jsonp.StartTime)
		}

		if jsonp.EndTime != "" {
			query.Where("create_time <=?", jsonp.EndTime)
		}
	} else {
		sTime, eTime := tool.GetDayRange(time.Now(), *jsonp.TimeType)
		query.Where("create_time BETWEEN ? AND ?", sTime, eTime)
	}

	// 1 处理中；2 失败；3 成功
	if jsonp.Status > -1 {
		/*if jsonp.Status == 1 {
			query.Where("status in(?)", []int{0, 1})
		} else {
			query.Where("status =?", jsonp.Status)
		}*/
		switch jsonp.Status {
		case 1:
			// 待审核 或 审核通过且无代付或代付中
			query.Where("status = ? OR (status = ? AND another_pay_status in ?)",
				enmus.OrderWithdrawStats_AuditWait,
				enmus.OrderWithdrawStats_AuditApprove,
				[]int{
					enmus.OrderWithdrawAnotherPayStats_None,
					enmus.OrderWithdrawAnotherPayStats_Progress,
				})
		case 2:
			// 审核未通过 或 代付失败
			query.Where("status = ? OR (status = ? AND another_pay_status = ?)",
				enmus.OrderWithdrawStats_AuditReject,
				enmus.OrderWithdrawStats_AuditApprove,
				enmus.OrderWithdrawAnotherPayStats_Failed)
		case 3:
			// 审核通过且代付成功
			query.Where("status = ? AND another_pay_status = ?",
				enmus.OrderWithdrawStats_AuditApprove,
				enmus.OrderWithdrawAnotherPayStats_Success)
		}
	}

	if jsonp.PageIndex == 0 {
		jsonp.PageIndex = 1
	}

	if jsonp.PageSize == 0 {
		jsonp.PageSize = 10
	}

	var total int64
	query.Count(&total)
	query.Offset((jsonp.PageIndex - 1) * jsonp.PageSize).Limit(jsonp.PageSize).
		Scan(&data)

	// 调整发给前端的状态
	for i, v := range data {
		if v.Status == enmus.OrderWithdrawStats_AuditWait ||
			(v.Status == enmus.OrderWithdrawStats_AuditApprove &&
				(v.AnotherPayStatus == enmus.OrderWithdrawAnotherPayStats_None ||
					v.AnotherPayStatus == enmus.OrderWithdrawAnotherPayStats_Progress)) {
			data[i].Status = 1 // 处理中
		} else if v.Status == enmus.OrderWithdrawStats_AuditReject ||
			(v.Status == enmus.OrderWithdrawStats_AuditApprove &&
				v.AnotherPayStatus == enmus.OrderWithdrawAnotherPayStats_Failed) {
			data[i].Status = 2 // 提款失败
		} else if v.Status == enmus.OrderWithdrawStats_AuditApprove &&
			v.AnotherPayStatus == enmus.OrderWithdrawAnotherPayStats_Success {
			data[i].Status = 3 // 提款成功
		}
	}

	newData := make([]*vo.OrderWithdrawInfoResp, 0, len(data))
	tool.JsonMapper(data, &newData)

	response.SuccessPageJSON(c, jsonp.PageIndex, jsonp.PageSize, total, newData)

}

func FlowProgressBar(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	var first dos.FcUserCodeBetLimit
	global.G_DB.Model(&dos.FcUserCodeBetLimit{}).Where("user_id = ? and (status = 1 or status = 2)", userInfo.UserId).Take(&first)
	if len(first.Id) == 0 {
		response.SuccessJSON(c, 100) //100%进度条
		return
	}
	d := decimal.NewFromFloat(first.FinishBetAmount).Div(decimal.NewFromFloat(first.MinBetAmount))
	a := d.Truncate(2).InexactFloat64()
	if a >= 1 {
		response.SuccessJSON(c, 100) //100%进度条
		return
	}
	retAmount := decimal.NewFromFloat(a * 100).Truncate(2).InexactFloat64()
	response.SuccessJSON(c, retAmount) //进度
}

func GetWithdrawChannel(c *gin.Context) {

	//userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	//userInfo := userInfoF.(*dos.FcUserMaterial)
	//merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)

	var data []*dos.FcPayChannelOut
	query := global.G_DB.Model(&dos.FcPayChannelOut{})
	//query.Order("sort desc").Where(`status = 1 and merchant_code = ? `, userInfo.MerchantCode)
	query.Order("sort desc").Where(`status = 1  `)
	query.Find(&data)
	//retData := make([]*vo.RechargeChannelRespVO, 0, len(data))

	response.SuccessJSON(c, data)
}

func GetWithdrawPaymentOut(c *gin.Context) {
	var jsonp vo.PaymentOutInfoReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	var data []*dos.FcPaymentOut
	query := global.G_DB.Model(&dos.FcPaymentOut{})
	//query.Order("sort desc").Where(`status = 1 and merchant_code = ? `, userInfo.MerchantCode)
	query.Order("sort desc").Where(`status = 1  AND channel_code=?`, jsonp.ChannelCode)
	query.Find(&data)

	newData := make([]*vo.PaymentOutInfoResp, len(data))
	tool.JsonMapper(data, &newData)
	for _, v := range newData {
		if v.PaymentCode != "" {
			img := modules.FindByKeyFcChannelBankImgFirst(&dos.FcChannelBankImg{PaymentCode: v.PaymentCode, Status: 1})
			v.Icon = img.Icon
			v.IconPath = img.IconPath
			v.Img = img.Img
			v.ImgPath = img.ImgPath
		}
	}
	response.SuccessJSON(c, newData)
}
