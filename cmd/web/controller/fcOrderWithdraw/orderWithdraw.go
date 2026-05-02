package fcOrderWithdraw

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/channelData"
	"bootpkg/pkg/service/paymentOut"
	"bootpkg/pkg/service/userTransfer"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func OrderWithdrawOk(c *gin.Context) {
	var jsonp *dos.FcOrderWithdraw
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	order := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	if order == nil || len(order.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, order.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}
	jsonp.OrderSn = order.OrderSn
	jsonp.Remark = "已出款"
	affect, err := userTransfer.UserWithdrawSuccess(jsonp)
	if !affect {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
		return
	}
	channelData.SendUserWithdrawal(&channelData.UserWithdrawalMessage{
		UserId:           order.UserId,
		UserName:         order.UserName,
		OrderSn:          order.OrderSn,
		WithdrawalAmount: order.Amount,
	})
	response.SuccessJSON(c, affect)
}

func OrderWithdrawNo(c *gin.Context) {
	var jsonp *dos.FcOrderWithdraw
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	//jsonp = modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: jsonp.Id}})

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	order := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	if order == nil || len(order.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, order.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	affect, err := userTransfer.UserWithdrawReject(jsonp)
	if !affect {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
		return
	}
	response.SuccessJSON(c, affect)
}

// 用户提款审核通过接口（新）
func OrderWithdrawAudit2(c *gin.Context) {
	var jsonp *vo.OrderWithdrawAnotherPayReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfo, ok := c.Get("UserInfo")
	var userName string
	if ok {
		userName = userInfo.(*dos.AdminUser).UserName
	}
	order := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{
		BaseDos: dos.BaseDos{
			Id: jsonp.Id,
		},
	})
	if order.Id != jsonp.Id {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,订单不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, order.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	var channelCode string
	var thirdCode string
	switch order.OrderType {
	case 1:
		channelCode = enmus.Another_Bank
	case 3:
		channelCode = enmus.Another_Virtual
		thirdCode = modules.GetFcPaymentOutThirdCode(order.VirtualType)
	case 4:
		channelCode = enmus.Another_AliPay
	}

	/*var paymentData dos.FcPaymentOut
	global.G_DB.Model(&dos.FcPaymentOut{}).Where("channel_code=? AND status=1", channelCode).First(&paymentData)

	if channelCode != paymentData.ChannelCode { //渠道不符合类型
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败: 渠道类型不符合")
		return
	}*/

	paymentData, _ := modules.GetFcPaymentOutData(channelCode, thirdCode)

	if len(paymentData.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,无可用通道")
		return
	}

	m := paymentOut.GetPaymentOutChannel(channelCode)
	if m == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败: 通道不存在")
		return
	}

	order.Decrypt()
	preStatus := m.PreWithdraw(paymentOut.PreOtherPaymentOutParam{
		WithdrawBankType: order.AccountBankType,
		WithdrawBankCode: order.AccountBankCode,
	})
	if preStatus.Code != 200 { //预检测失败
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败: %v", preStatus.ErrorMsg))
		return
	}

	fee := decimal.NewFromFloat(order.PreAmount).Mul(decimal.NewFromFloat(paymentData.FeeRate / 100)).Truncate(2).InexactFloat64()

	anotherPay := &dos.FcOrderWithdrawPaymentOut{
		OrderSn:                  order.OrderSn,
		UserId:                   order.UserId,
		UserName:                 order.UserName,
		Amount:                   order.PreAmount,
		Status:                   enmus.OrderWithdrawPaymentOutStats_Prepare,
		MerchantCode:             order.MerchantCode,
		Currency:                 order.Currency,
		OrderType:                order.OrderType,
		FeeRate:                  paymentData.FeeRate,
		Fee:                      fee,
		ChannelCode:              paymentData.ChannelCode,
		PaymentId:                paymentData.Id,
		PaymentCode:              paymentData.PaymentCode,
		CreateTime:               automaticType.Now(),
		CreateBy:                 userName,
		ThirdCode:                paymentData.ThirdCode,
		WithdrawAmount:           order.Amount,
		WithdrawStatus:           enmus.OrderWithdrawStats_No,
		DepositWithdrawSubAmount: order.DepositWithdrawSubAmount,
	}

	err = paymentOut.InsertOrderWithdrawAnotherPay(order, anotherPay)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	response.SuccessJSON(c, true)
}

// OrderWithdrawBatchNo
//
//	@Description: 批量拒绝
//	@param c
func OrderWithdrawBatchNo(c *gin.Context) {
	jsonp := struct {
		Ids            []string `json:"ids"`             //ID集合
		CallbackRemark string   `json:"callback_remark"` //拒绝理由
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	var count int
	for _, v := range jsonp.Ids {
		order := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: v}})
		if order == nil || len(order.Id) == 0 {
			continue
		}
		if !modules.CheckAdminUserMerchantPerms(c, order.MerchantCode) {
			continue
		}

		out := &dos.FcOrderWithdraw{
			BaseDos:        dos.BaseDos{Id: v},
			UpdateBy:       userInfo.UserName,
			Remark:         jsonp.CallbackRemark,
			CallbackRemark: jsonp.CallbackRemark,
		}
		b, _ := userTransfer.UserWithdrawReject(out)
		if b {
			count++
		}
	}
	if count == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请刷新后重试")
		return
	}
	//affect, err := userTransfer.UserWithdrawReject(jsonp)
	//if !affect {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
	//	return
	//}
	response.SuccessJSON(c, true)
}

// OrderWithdrawBatchAudit
//
//	@Description: 批量审核通过
//	@param c
func OrderWithdrawBatchAudit(c *gin.Context) {
	jsonp := struct {
		Ids []string `json:"ids"` //ID集合
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)

	var count int
	for _, v := range jsonp.Ids {
		order := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{BaseDos: dos.BaseDos{Id: v}})
		if order == nil || len(order.Id) == 0 {
			continue
		}
		if !modules.CheckAdminUserMerchantPerms(c, order.MerchantCode) {
			continue
		}

		out := &dos.FcOrderWithdraw{
			BaseDos:  dos.BaseDos{Id: v},
			UpdateBy: userInfo.UserName,
		}
		b := userTransfer.UserWithdrawAuditSuccess(out)
		if b {
			count++
		}
	}
	if count == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请刷新后重试")
		return
	}
	//affect := userTransfer.UserWithdrawAuditSuccess(jsonp)
	//if !affect {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试"))
	//	return
	//}
	response.SuccessJSON(c, true)
}

// OrderWithdrawAnotherBatchPay
//
//	@Description: 批量申请代付
//	@param c
func OrderWithdrawAnotherBatchPay(c *gin.Context) {
	jsonp := struct {
		Ids         []string `json:"ids"`          //ID集合
		PaymentCode string   `json:"payment_code"` //选择支付通道
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)

	paymentData := modules.FindByKeyFcPaymentOutFirst(&dos.FcPaymentOut{
		PaymentCode: jsonp.PaymentCode,
	})
	if paymentData == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,无可用通道")
		return
	}
	var count int
	for _, v := range jsonp.Ids {
		order := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{
			BaseDos: dos.BaseDos{
				Id: v,
			},
		})
		if order == nil || len(order.Id) == 0 {
			continue
		}
		if !modules.CheckAdminUserMerchantPerms(c, order.MerchantCode) {
			continue
		}
		var channelCode string
		if order.OrderType == 1 {
			channelCode = enmus.Another_Bank
		} else if order.OrderType == 3 {
			channelCode = enmus.Another_Virtual
		}
		if channelCode != paymentData.ChannelCode { //渠道不符合类型
			continue
		}
		user := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{UserId: order.UserId})
		if user.Level > paymentData.MaxLevel || user.Level < paymentData.MinLevel { //等级不符合通道
			continue
		}
		if order.PreAmount > paymentData.MaxAmount || order.PreAmount < paymentData.MinAmount { //金额不符合通道
			continue
		}

		fee := decimal.NewFromFloat(order.PreAmount).Mul(decimal.NewFromFloat(paymentData.FeeRate / 100)).Truncate(2).InexactFloat64()

		anotherPay := &dos.FcOrderWithdrawPaymentOut{
			OrderSn:      order.OrderSn,
			UserId:       order.UserId,
			UserName:     order.UserName,
			Amount:       order.PreAmount,
			Status:       enmus.OrderWithdrawPaymentOutStats_Prepare,
			MerchantCode: order.MerchantCode,
			Currency:     order.Currency,
			OrderType:    order.OrderType,
			FeeRate:      paymentData.FeeRate,
			Fee:          fee,
			//ChannelId:    channelInfo.Id,
			ChannelCode:              paymentData.ChannelCode,
			PaymentId:                paymentData.Id,
			PaymentCode:              paymentData.PaymentCode,
			CreateTime:               automaticType.Now(),
			CreateBy:                 userInfo.UserName,
			DepositWithdrawSubAmount: order.DepositWithdrawSubAmount,
			ApplyAmount:              order.Amount,
			WithdrawStatus:           enmus.OrderWithdrawStats_No,
		}

		m := paymentOut.GetPaymentOutChannel(paymentData.PaymentCode)
		if m == nil {
			continue
		}

		order.Decrypt()
		preStatus := m.PreWithdraw(paymentOut.PreOtherPaymentOutParam{
			WithdrawBankType: order.AccountBankType,
			WithdrawBankCode: order.AccountBankCode,
		})
		if preStatus.Code != 200 { //预检测失败
			continue
		}

		err := paymentOut.InsertOrderWithdrawAnotherPay(order, anotherPay)
		if err == nil {
			count++
		}
	}
	if count == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请刷新后重试")
		return
	}

	response.SuccessJSON(c, true)
}

// OrderWithdrawGetAnotherPay
//
//	@Description: 获取符合代付渠道
//	@param c
func OrderWithdrawGetAnotherPay(c *gin.Context) {
	var jsonp *vo.OrderWithdrawAnotherPayReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	order := modules.FindByKeyFcOrderWithdrawFirst(&dos.FcOrderWithdraw{
		BaseDos: dos.BaseDos{
			Id: jsonp.Id,
		},
	})
	if order.Id != jsonp.Id {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,订单不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, order.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	var channelCode string
	if order.OrderType == 1 {
		channelCode = enmus.Another_Bank
	} else if order.OrderType == 3 {
		channelCode = enmus.Another_Virtual
	}
	if len(channelCode) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,没有符合的渠道")
		return
	}

	var data []*dos.FcPaymentOut
	query := global.G_DB.Model(&dos.FcPaymentOut{})
	query.Order("sort desc").Where(`merchant_code = ? and status = 1 and channel_code=?`, order.MerchantCode, channelCode)
	query.Find(&data)

	user := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{UserId: order.UserId})
	paymentData := make([]*dos.FcPaymentOut, 0)

	order.Decrypt()
	param := paymentOut.PreOtherPaymentOutParam{
		WithdrawBankType: order.AccountBankType,
		WithdrawBankCode: order.AccountBankCode,
	}

	//选择通道
	for _, v := range data {
		if user.Level <= v.MaxLevel && user.Level >= v.MinLevel {
			if order.PreAmount >= v.MinAmount && order.PreAmount <= v.MaxAmount {

				//预检测
				m := paymentOut.GetPaymentOutChannel(v.PaymentCode)
				if m == nil {
					continue
				}
				preStatus := m.PreWithdraw(param)
				if preStatus.Code != 200 { //预检测失败
					continue
				}

				paymentData = append(paymentData, v)
			}
		}
	}
	outJson := struct {
		Id          string              `json:"id,omitempty"`
		PaymentList []*dos.FcPaymentOut `json:"payment_list"`
	}{
		Id:          order.Id,
		PaymentList: paymentData,
	}

	response.SuccessJSON(c, outJson)
}
