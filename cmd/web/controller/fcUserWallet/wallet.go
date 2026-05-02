package fcUserWallet

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/channelData"
	"bootpkg/pkg/service/userTransfer"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

var (
	ManualRechargeMap = map[int]int{
		1: enmus.Recharge_Order_Type_Wx,
		2: enmus.Recharge_Order_Type_Bank,
		3: enmus.Recharge_Order_Type_Alipay,
		4: enmus.Recharge_Order_Type_Wallet,
		5: enmus.Recharge_Order_Type_NumberCNY,
		6: enmus.Recharge_Order_Type_Virtual,
	}

	ManualBonusMap = map[int]int{
		//30: enmus.Promotion_Official_Bonus,  // 官方赠送
		//31: enmus.Promotion_Official_Recoup, // 官方补偿
		51: enmus.Promotion_Official_Deduct, // 福利扣除
	}
)

func WalletAmountOpt(c *gin.Context) {
	var jsonp dos.FcOrderManageOpt
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if jsonp.Amount == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "amount 0")
		return
	}
	// 默认为 CNY
	if jsonp.Currency == "" {
		jsonp.Currency = enmus.CNY
	}

	if (jsonp.TrsType == 40 || jsonp.TrsType == 41) &&
		modules.CheckVenueEntryRecordValNoCache(jsonp.UserId) {
		response.FailErrJSON(c, response.ERROR_SERVER, "该用户有场馆金额未转出无法进行扣除操作")
		return
	}

	m := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: jsonp.UserId,
	})
	if m.UserId == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "userId is not exist")
		return
	}

	userWallet := modules.FindByKeyFcUserWalletFirst(&dos.FcUserWallet{UserId: jsonp.UserId, Currency: jsonp.Currency})
	userMaterial := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{UserId: jsonp.UserId})

	// 判断当前账号只能操作对应商户
	if !modules.CheckAdminUserMerchantPerms(c, userMaterial.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "无权限处理该商户")
		return
	}

	if len(userWallet.Id) == 0 {
		//userMaterial := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{UserId: jsonp.UserId})
		userWallet = &dos.FcUserWallet{
			UserId:        userMaterial.UserId,
			UserName:      userMaterial.UserName,
			Currency:      jsonp.Currency,
			TotalAmount:   0,
			AvaAmount:     0,
			FronzenAmount: 0,
			IsLock:        0,
			MerchantCode:  userMaterial.MerchantCode,
		}
		modules.SaveFcUserWallet(userWallet)
	}
	var userName string
	userInfo, ok := c.Get("UserInfo")
	if ok {
		loginAdmin := userInfo.(*dos.AdminUser)
		userName = loginAdmin.UserName
		global.G_LOG.Infof("[WalletAmountOpt] session admin userName=%s id=%s accountType=%d trsType=%d scoreType=%d amount=%.2f",
			loginAdmin.UserName, loginAdmin.Id, loginAdmin.AccountType, jsonp.TrsType, jsonp.ScoreType, jsonp.Amount)

		if jsonp.ScoreType == 1 && (jsonp.TrsType == 30 || jsonp.TrsType == 31 || jsonp.TrsType == 1 || jsonp.TrsType == 2 ||
			jsonp.TrsType == 3 || jsonp.TrsType == 4 || jsonp.TrsType == 5 || jsonp.TrsType == 6) {
			// 超级管理员（或默认admin）不受额度限制
			needLimitCheck := !(loginAdmin.AccountType == 1 || loginAdmin.UserName == "admin")
			if needLimitCheck {
				uid := loginAdmin.Id
				adminUser := dos.AdminUser{}
				err1 := global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", uid).Take(&adminUser).Error
				if err1 != nil {
					response.FailErrJSON(c, response.ERROR_PARAMETER, "账号异常")
					return
				}

				// 超级管理员不受额度限制
				if adminUser.AccountType != 1 {
					global.G_LOG.Infof("[WalletAmountOpt] limit check userName=%s id=%s accountType=%d total=%.2f limitPer=%.2f cur=%.2f",
						adminUser.UserName, adminUser.Id, adminUser.AccountType, adminUser.TotalAmount, adminUser.LimitPertimeAmount, adminUser.CurAmount)
					totalAmount := adminUser.TotalAmount
					curAmount := adminUser.CurAmount
					limitAmount := adminUser.LimitPertimeAmount
					if jsonp.Amount > limitAmount {
						response.FailErrJSON(c, response.ERROR_PARAMETER, "单笔额度限制，请联系管理员")
						return
					}
					if jsonp.Amount+curAmount > totalAmount {
						response.FailErrJSON(c, response.ERROR_PARAMETER, "总额度不足，请联系管理员")
						return
					}
					curAmount = curAmount + jsonp.Amount
					adminUser.CurAmount = curAmount
					global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", adminUser.Id).Update("cur_amount", curAmount)
				}
			}
		}
	}

	trsAmountType := userTransfer.TranManageAdd // 默认为加钱
	amount := jsonp.Amount                      // 默认上分
	if jsonp.ScoreType == 2 {                   // 下分
		amount = -amount
		trsAmountType = userTransfer.TranManageReduce
	}
	beforeAmount := userWallet.AvaAmount
	afterAmount := userWallet.AvaAmount + amount
	if afterAmount < 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "amount less than 0")
		return
	}

	nowTime := automaticType.Now()

	flowAmount := jsonp.Amount * float64(jsonp.FlowMultiple)
	orderManageOpt := &dos.FcOrderManageOpt{
		Currency:          jsonp.Currency,
		Amount:            jsonp.Amount,
		TranscationBefore: beforeAmount,
		TranscationAfter:  afterAmount,
		Status:            1, // 成功
		TrsType:           jsonp.TrsType,
		ScoreType:         jsonp.ScoreType,
		FlowMultiple:      jsonp.FlowMultiple,
		FlowAmount:        flowAmount,
		CreateBy:          userName,
		UpdateBy:          userName,
		MerchantCode:      userWallet.MerchantCode,
		UserId:            userWallet.UserId,
		UserName:          userWallet.UserName,
		Remarks:           jsonp.Remarks,
		CreateTime:        nowTime,
		UpdateTime:        nowTime,
	}

	// 如果是人工存款类型，则需要进行写充值表
	bonusOk := false // 是否需要写福利, false 不需要
	var bonusType int
	deposit := dos.FcOrderDeposit{}
	fop := dos.FcOrderPromotion{}
	rechargeOrderType, manualOk := ManualRechargeMap[jsonp.TrsType]
	if manualOk {
		if jsonp.ScoreType == 2 { // 人工存款不能为下分
			response.FailErrJSON(c, response.ERROR_PARAMETER, "param err")
			return
		}

		trsAmountType = userTransfer.TranManual
		orderSn := getOrderSn()

		//处理用户添加金额订单
		deposit = dos.FcOrderDeposit{
			UserId:          userWallet.UserId,
			UserName:        userWallet.UserName,
			OrderSn:         orderSn,
			Amount:          jsonp.Amount,
			FactAmount:      jsonp.Amount,
			Status:          enmus.ORDER_YES_STATUS, //三方回调处理
			Remark:          jsonp.Remarks,
			DepositRemark:   jsonp.Remarks,
			Currency:        jsonp.Currency,
			Ip:              c.ClientIP(),
			CreateBy:        userName,
			UpdateBy:        userName,
			MerchantCode:    userWallet.MerchantCode,
			InviteCode:      userMaterial.AgentInviteCode,
			ChannelCode:     enmus.Manual_Recharge_ChannelCode,
			PaymentCode:     enmus.Manual_Recharge_PayCode,
			PaymentName:     enmus.Manual_Recharge_PayName,
			PayAliasName:    enmus.Manual_Recharge_PayAliasName,
			PaymentId:       enmus.Manual_Recharge_PayId,
			OrderType:       enmus.Recharge_Order_Type_Manual,
			OrderSecondType: rechargeOrderType,
			Level:           m.Level, //用户等级区分新老用户
			PayTime:         nowTime,
			CreateTime:      nowTime,
			AuthBy:          userName,
			AuthTime:        nowTime,
		}
	} else {
		// 客服补偿和官方赠送需要写红利表
		bonusType, bonusOk = ManualBonusMap[jsonp.TrsType]
		//applyRate := 1.0
		//if bonusType == enmus.Promotion_Official_Bonus || bonusType == enmus.Promotion_Official_Recoup {
		//	applyRate = 0
		//	amount = 0
		//}
		if bonusOk {
			orderSn := tool.SnowflakeId()
			fop = dos.FcOrderPromotion{
				ApplyAmount:  amount,
				OrderSn:      orderSn,
				AppleRate:    0,
				ApplyType:    bonusType,
				Amount:       amount,
				Status:       enmus.ORDER_YES_STATUS,
				UserName:     m.UserName,
				UserId:       m.UserId,
				TurnOver:     0,
				MerchantCode: m.MerchantCode,
				Remake:       jsonp.Remarks,
				Currency:     jsonp.Currency,
				CreateBy:     userName,
				UpdateBy:     userName,
			}
		}
	}

	global.G_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(orderManageOpt).Error
		if err != nil {
			global.G_LOG.Error("WalletAmountOpt UserAmountChange username: %v err: %v", userWallet.UserName, err)
			return err
		}

		//global.G_LOG.Infof("[WalletAmountOpt]orderManageOpt.Id = %s", orderManageOpt.Id)

		err = userTransfer.UserAmountChange2(tx, amount, trsAmountType, jsonp.Currency, jsonp.Remarks, jsonp.UserId, userName, "", orderManageOpt.Id,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypeOther, fmt.Sprintf("trstype_%d", jsonp.TrsType)))
		if err != nil {
			return err
		}

		// 如果是人工存款
		if manualOk {
			err = tx.Create(&deposit).Error
			if err != nil {
				return err
			}
		}

		// 如果是发红利，需要写红利表
		if bonusOk {
			err = tx.Create(&fop).Error
			if err != nil {
				return err
			}
		}

		return err
	})
	if err != nil {
		global.G_LOG.Error("WalletAmountOpt UserAmountChange username: %v err: %v", userWallet.UserName, err)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	// 如果是人工存款, 还需要判断是否是首存
	if manualOk {
		//充值成功系统邮件提示
		go modules.SendSystemMail(userWallet.UserId, userWallet.MerchantCode, enmus.MailType_RechargeSuccess, amount)

		err = userTransfer.FirstDepositSave(&deposit)
		if err != nil {
			global.G_LOG.Infof("WalletAmountOpt FirstDepositSave username: %v err: %v", userWallet.UserName, err)
		}
	}

	// 发送存款信息
	if global.CONFIG.Mq.IsInit && manualOk {
		msg := &channelData.UserRechargeMessage{
			UserId:        m.UserId,
			UserName:      m.UserName,
			OrderSn:       deposit.OrderSn,
			DepositTime:   deposit.CreateTime.String(),
			DepositAmount: deposit.Amount,
			T:             time.Now().UnixMicro(),
		}
		errR := channelData.SendUserRecharge(msg)
		if errR != nil {
			global.G_LOG.Errorf("WalletAmountOpt SendUserRecharge kafka err: %v , msg: %s", errR, tool.String(msg))
		}
	}

	if global.CONFIG.Mq.IsInit && bonusOk {
		promoData := channelData.UserPromotionMessage{}
		promoData.UserId = m.UserId
		promoData.UserName = m.UserName
		promoData.OrderSn = fop.OrderSn
		promoData.ForceStatus = 1
		promoData.T = time.Now().UnixMicro()
		promoData.PromotionTime = fop.CreateTime.String()
		promoData.PromotionAmount = amount
		promoData.PromotionType = 1

		// 发送红利消息给消息队列
		err = channelData.SendUserPromotion(&promoData)
		if err != nil {
			global.G_LOG.Errorf("WalletAmountOpt SendUserPromotion kafka err: %v , msg: %s", err, tool.String(promoData))
		}
	}

	modules.DelVenueEntryRecordVal(jsonp.UserId)

	// ============= 充值任务（累计充值）以下
	channelMap := map[int]string{
		// trs_type => channel
		1: "wx",
		2: "bank",
		3: "alipay",
		4: "wallet",
		5: "number_cny",
		6: "USDT-CR20",
	}
	if jsonp.TrsType >= 1 && jsonp.TrsType <= 6 {
		modules.DoUserTaskAction(userMaterial.UserId, []modules.TaskActionParam{
			{
				Type:        enmus.DailyTaskType_Recharge,
				Subtype:     enmus.DailyTaskSubType_Pay,
				Amount:      jsonp.Amount,
				ChannelCode: channelMap[jsonp.TrsType],
			},
		}, true)
	}
	// ============= 充值任务（累计充值）以上

	// ============= 活动（回血包）以下
	/*if jsonp.TrsType >= 1 && jsonp.TrsType <= 6 {
		go func() {
			remaingAmount, ratio := modules.GetUserRemainingAmountAndRechargeRatio(userMaterial.UserId)
			modules.DoUserActivityAction(enmus.ActivityTypes_HealthPack,
				userMaterial.UserId, []modules.ActivityActionParam{
					{
						RechargeBalanceRatio: ratio,
						Balance:              remaingAmount,
						RegTime:              modules.GetFcUserMaterialRegTime(userMaterial.UserId),
						FirstRechargeAmount:  modules.GetUserFirstRechargeAmount(userMaterial.UserId),
					},
				})
		}()
	}*/
	// ============= 活动（回血包）以上

	response.SuccessJSON(c, true)
}

func ClearWalletAmountOpt(c *gin.Context) {
	var jsonp vo.WalletAmountOptRequest
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userWallet := modules.FindByKeyFcUserWalletFirst(&dos.FcUserWallet{UserId: jsonp.UserId, Currency: jsonp.Currency})
	userMaterial := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{UserId: jsonp.UserId})

	if !modules.CheckAdminUserMerchantPerms(c, userMaterial.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	if len(userWallet.Id) == 0 {

		userWallet = &dos.FcUserWallet{
			UserId:        userMaterial.UserId,
			UserName:      userMaterial.UserName,
			Currency:      jsonp.Currency,
			TotalAmount:   0,
			AvaAmount:     0,
			FronzenAmount: 0,
			IsLock:        0,
			MerchantCode:  userMaterial.MerchantCode,
		}
		modules.SaveFcUserWallet(userWallet)
	}
	var userName string
	userInfo, ok := c.Get("UserInfo")
	if ok {
		userName = userInfo.(*dos.AdminUser).UserName
	}
	if userWallet.AvaAmount <= 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "amount 0")
		return
	}
	global.G_DB.Transaction(func(tx *gorm.DB) error {
		opt := userTransfer.TranManageReduce
		err = userTransfer.UserAmountChange(tx, -userWallet.AvaAmount, opt, jsonp.Currency, jsonp.Remark, jsonp.UserId, userName, "", "")
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcUserSiteMessage/findPage
func FindPageFcOrderManageOptControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcOrderManageOpt
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.TrsType = tool.Atoi(c.DefaultQuery("trs_type", ""))
	jsonp.ScoreType = tool.Atoi(c.DefaultQuery("score_type", ""))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageOrderManageOpt(jsonp.PageNo, jsonp.PageSize, &jsonp.FcOrderManageOpt, &jsonp.PageTimeQuery, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

func getOrderSn() string {
	orderNoKey := time.Now().Format("20060102")
	orderNoPre := time.Now().Format("20060102150405")
	orderNoAdd := global.G_REDIS.Incr(context.Background(), fmt.Sprintf(enmus.DEPOSIT_NO_KEY, orderNoKey)).Val()
	orderSn := orderNoPre + tool.RandString(3) + tool.String(orderNoAdd)
	return orderSn
}
