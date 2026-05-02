package userTransfer

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	vos "bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/channelData"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// CallbackDepositSuccess
//
//	@Description: 回调用户存款通过
//	@param vo 订单信息
//	@return bool true 成功
func CallbackDepositSuccess(voData *dos.FcOrderDeposit) (bool, error) {
	var vo *dos.FcOrderDeposit
	err1 := global.G_DB.Model(&dos.FcOrderDeposit{}).Where("id = ?", voData.Id).Take(&vo).Error
	global.G_LOG.Infof("sreach order deposit: %v  ------- %v --------- %v", tool.String(voData), tool.String(vo), err1)
	if len(vo.Id) == 0 {
		return false, errors.New("order does not exist")
	}

	// 如果是已成功不需要修改
	if vo.Status == enmus.ORDER_YES_STATUS {
		return true, nil
	}
	if vo.Status != enmus.Order_STATUS_PENDING_PAY {
		return false, errors.New("please refresh and try again")
	}
	voData.PayTime = vo.PayTime
	voData.Status = enmus.ORDER_YES_STATUS
	var isOk bool
	var err error

	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		// 待处理状态才能修改
		err = callbackUpdateDepositSuccessInfo(tx, voData)
		if err != nil {
			return err
		}
		//修改用户金额表
		err = UserAmountChange(tx, vo.Amount, TranOnline, vo.Currency, vo.OrderSn, vo.UserId, vo.UpdateBy, "",
			modules.GetFcTranscationFundingSubType(enmus.FundingTypeDesposit, vo.ChannelCode))
		if err != nil {
			return err
		}

		isOk = true
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return false, err
	}

	//设置最后充值成功缓存
	if isOk {
		global.G_REDIS.Set(context.Background(), fmt.Sprintf(enmus.RECHARGESUCCESSINFO, vo.UserId), tool.String(&vos.RechargeSuccessInfoVO{
			Success:  true,
			Amount:   tool.String(vo.Amount),
			Currency: vo.Currency,
		}), 30*time.Hour)

		// 存款的一些数据处理
		go DepositLinkHandle(vo)
	}

	if global.CONFIG.Mq.IsInit && isOk {
		msg := &channelData.UserRechargeMessage{
			UserId:        vo.UserId,
			UserName:      vo.UserName,
			OrderSn:       vo.OrderSn,
			DepositTime:   vo.CreateTime.String(),
			DepositAmount: vo.Amount,
			T:             time.Now().UnixMicro(),
		}
		errR := channelData.SendUserRecharge(msg)
		if errR != nil {
			global.G_LOG.Errorf(" send kafka err: %v , msg: %s", errR, tool.String(msg))
		}
	}
	return isOk, err
}

func callbackUpdateDepositSuccessInfo(tx *gorm.DB, vo *dos.FcOrderDeposit) error {
	nowTimeStr := automaticType.Now().String()
	updates := map[string]interface{}{
		"remark":    vo.Remark,
		"update_by": vo.UpdateBy,
		"status":    vo.Status,
		"auth_by":   vo.AuthBy,
		"auth_time": nowTimeStr,
	}
	if vo.PayTime.Timer().IsZero() {
		updates["pay_time"] = nowTimeStr
	}

	if vo.OrderType == enmus.Recharge_Order_Type_Bank {
		updates["remitter_account_holder"] = vo.RemitterAccountHolder
		updates["remitter_account_bank_name"] = vo.RemitterAccountBankName
		updates["remitter_account_number"] = vo.RemitterAccountNumber
	} else if vo.OrderType == enmus.Recharge_Order_Type_Virtual {
		updates["virtual_pay_address"] = vo.VirtualPayAddress
		updates["virtual_pay_no"] = vo.VirtualPayNo
		updates["virtual_pay_amount"] = vo.VirtualPayAmount
	}
	eRow := tx.Model(vo).Where("id = ? AND status = ?", vo.Id, enmus.Order_STATUS_PENDING_PAY).Updates(updates)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected == 0 {
		return errors.New("update deposit status fail")
	}
	return nil
}

func updateDepositSuccessInfo(tx *gorm.DB, vo *dos.FcOrderDeposit) error {
	nowTimeStr := automaticType.Now().String()
	updates := map[string]interface{}{
		"remark":    vo.Remark,
		"update_by": vo.UpdateBy,
		"status":    vo.Status,
		"auth_by":   vo.AuthBy,
		"auth_time": nowTimeStr,
	}
	if vo.PayTime.Timer().IsZero() {
		updates["pay_time"] = nowTimeStr
	}

	eRow := tx.Model(vo).Where("id = ? AND status in ?", vo.Id, []int{enmus.Order_STATUS_PENDING_PAY, enmus.ORDER_STATUS_WAIT}).Updates(updates)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected == 0 {
		return errors.New("update deposit status fail")
	}
	return nil
}

// UserDepositSuccess
//
//	@Description: 用户存款通过
//	@param vo 订单信息
//	@return bool true 成功
func UserDepositSuccess(voData *dos.FcOrderDeposit) (bool, error) {
	var vo *dos.FcOrderDeposit
	err1 := global.G_DB.Model(&dos.FcOrderDeposit{}).Where("id = ?", voData.Id).Take(&vo).Error
	global.G_LOG.Infof("sreach order deposit: %v  ------- %v --------- %v", tool.String(voData), tool.String(vo), err1)
	if len(vo.Id) == 0 {
		return false, errors.New("order does not exist")
	}
	var isOk bool
	var err error

	// 如果是已成功不需要修改
	if vo.Status == enmus.ORDER_YES_STATUS {
		return true, nil
	}

	voData.Status = enmus.ORDER_YES_STATUS

	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		// 更新为成功
		err = updateDepositSuccessInfo(tx, voData)
		if err != nil {
			return err
		}
		//修改用户金额表
		err = UserAmountChange(tx, vo.Amount, TranOnline, vo.Currency, vo.OrderSn, vo.UserId, voData.UpdateBy, "",
			modules.GetFcTranscationFundingSubType(enmus.FundingTypeDesposit, vo.ChannelCode))
		if err != nil {
			return err
		}

		isOk = true
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return false, err
	}

	//设置最后充值成功缓存
	if isOk {
		global.G_REDIS.Set(context.Background(), fmt.Sprintf(enmus.RECHARGESUCCESSINFO, vo.UserId), tool.String(&vos.RechargeSuccessInfoVO{
			Success:  true,
			Amount:   tool.String(vo.Amount),
			Currency: vo.Currency,
		}), 30*time.Hour)

		// 存款的一些数据处理
		go DepositLinkHandle(vo)
	}

	if global.CONFIG.Mq.IsInit && isOk {
		msg := &channelData.UserRechargeMessage{
			UserId:        vo.UserId,
			UserName:      vo.UserName,
			OrderSn:       vo.OrderSn,
			DepositTime:   vo.CreateTime.String(),
			DepositAmount: vo.Amount,
			T:             time.Now().UnixMicro(),
		}
		errR := channelData.SendUserRecharge(msg)
		if errR != nil {
			global.G_LOG.Errorf(" send kafka err: %v , msg: %s", errR, tool.String(msg))
		}
	}
	return isOk, err
}

// VenueDepositFail
//
//	@Description: 场馆存款失败,冻结金额退回
//	@param vo
func VenueDepositFail(voData *dos.FcVenueTransfer) (bool, error) {
	var vo *dos.FcVenueTransfer
	global.G_DB.Model(&dos.FcVenueTransfer{}).Where("id = ?", voData.Id).Take(&vo)
	if len(vo.Id) == 0 {
		return false, errors.New("order does not exist")
	}
	var isOk bool
	var err error
	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		//待处理状态才能修改
		affect := tx.Model(&dos.FcVenueTransfer{}).Where("id = ? and status = ?", vo.Id, 0).Updates(map[string]interface{}{
			"update_by": voData.UpdateBy,
			"status":    2,
		})
		if affect.RowsAffected == 0 {
			return errors.New(" rowsAffected 0")
		}
		//修改用户金额表
		err = UserVenueAmountChangeCallback(tx, vo.Amount, TranAmountConvert, vo.Currency, "==>【"+vo.VenueCode+"】"+vo.OrderSn+" 【Fail】", vo.UserId, vo.UpdateBy, vo.OrderSn,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypePlatform, "deposit")+vo.VenueCode)
		if err != nil {
			return err
		}
		isOk = true
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return false, err
	}
	return isOk, err
}

// VenueDepositSuccess
//
//	@Description: 场馆存款成功- 扣除冻结金额
//	@param vo
func VenueDepositSuccess(voData *dos.FcVenueTransfer) (bool, error) {
	var vo *dos.FcVenueTransfer
	global.G_DB.Model(&dos.FcVenueTransfer{}).Where("id = ?", voData.Id).Take(&vo)
	if len(vo.Id) == 0 {
		return false, errors.New("order does not exist")
	}
	var isOk bool
	var err error
	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		//待处理状态才能修改
		affect := tx.Model(&dos.FcVenueTransfer{}).Where("id = ? and status = ?", vo.Id, 0).Updates(map[string]interface{}{
			"update_by": voData.UpdateBy,
			"status":    1,
		})
		if affect.RowsAffected == 0 {
			return errors.New(" rowsAffected 0")
		}
		//扣除冻结金额
		err = UserVenueAmountConfirmChange(tx, vo.Amount, TranAmountConvert, vo.Currency, "==>【"+vo.VenueCode+"】"+vo.OrderSn+" 【Success】", vo.UserId, vo.UpdateBy)
		if err != nil {
			return err
		}
		isOk = true
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return false, err
	}

	if isOk {
		global.G_DB.Model(&dos.FcTranscation{}).Where("related_id=?", vo.OrderSn).
			Update("funding_subtype", modules.GetFcTranscationFundingSubType(enmus.FundingTypePlatform, "deposit")+vo.VenueCode)
	}

	return isOk, err
}

// 存款的一些数据处理
func DepositLinkHandle(vo *dos.FcOrderDeposit) error {
	// 发送存款优惠
	err := DepositBonus(vo)
	if err != nil {
		global.G_LOG.Error("DepositBonus username: %v orderSn: %v bonusAmount: %v err: %v", vo.UserName, vo.OrderSn, vo.BonusAmount, err)
	}

	// 用户首存处理
	err = FirstDepositSave(vo)
	if err != nil {
		global.G_LOG.Error("FirstDepositSave username: %v orderSn: %v bonusAmount: %v err: %v", vo.UserName, vo.OrderSn, vo.BonusAmount, err)
	}

	return nil
}

// 存款优惠
func DepositBonus(vo *dos.FcOrderDeposit) error {
	if vo.BonusAmount <= 0.00 {
		return nil
	}

	// 判断用户是否有存款优惠, 如果有需要给用户送红利, 不需要流水
	tmpRemark := "充值优惠"
	currency := vo.Currency

	fop := &dos.FcOrderPromotion{
		ApplyAmount:  vo.Amount,
		OrderSn:      vo.OrderSn,
		AppleRate:    vo.BonusRate,
		ApplyType:    enmus.Promotion_Deposit,
		Amount:       vo.BonusAmount,
		Status:       enmus.ORDER_YES_STATUS,
		UserName:     vo.UserName,
		UserId:       vo.UserId,
		TurnOver:     0,
		MerchantCode: vo.MerchantCode,
		Remake:       tmpRemark,
		Currency:     currency,
	}

	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		//修改用户金额表
		err := UserAmountChange(tx, vo.BonusAmount, TranDiscount, currency, tmpRemark, vo.UserId, "", "",
			modules.GetFcTranscationFundingSubType(enmus.FundingTypePromotion, fmt.Sprintf("deposit_bonus_%d", enmus.Promotion_Deposit)))
		if err != nil {
			global.G_LOG.Errorf("RechargeBonus userTransfer.UserAmountChange userId:%s amount: %v err :%s", vo.UserId, vo.BonusAmount, err.Error())
			return err
		}

		// 写红利表
		err = tx.Create(fop).Error
		if err != nil {
			global.G_LOG.Errorf("RechargeBonus insert FcOrderPromotion userId:%s amount: %v err :%s", vo.UserId, vo.BonusAmount, err.Error())
			return err
		}
		return nil
	})

	if err != nil {
		global.G_LOG.Errorf("RechargeBonus FcOrderPromotion userId:%s amount: %v err :%s", vo.UserId, vo.BonusAmount, err.Error())
		return err
	}

	if global.CONFIG.Mq.IsInit {
		promoData := channelData.UserPromotionMessage{}
		promoData.UserId = vo.UserId
		promoData.UserName = vo.UserName
		promoData.OrderSn = vo.OrderSn
		promoData.ForceStatus = 1
		promoData.T = time.Now().UnixMicro()
		promoData.PromotionTime = vo.CreateTime.String()
		promoData.PromotionAmount = vo.BonusAmount
		promoData.PromotionType = 1 // 红利

		// 发送红利消息给消息队列
		err = channelData.SendUserPromotion(&promoData)
		if err != nil {
			global.G_LOG.Errorf(" SendUserPromotion kafka err: %v , msg: %s", err, tool.String(promoData))
		}
	}

	return nil
}

// 用户首存
func FirstDepositSave(vo *dos.FcOrderDeposit) error {
	// 判断用户是否首存, 先判断缓存，再判断数据库
	firstDepositeKey := fmt.Sprintf(enmus.First_Deposite_UserId_Key, vo.UserId)
	tmpFirstExist := global.G_REDIS.Exists(context.Background(), firstDepositeKey).Val()
	if tmpFirstExist > 0 {
		return nil
	}

	// 如果缓存不存在，则继续判断数据库中是否存在
	row := dos.FcFirstOrderDeposit{}
	err := global.G_DB.Model(&dos.FcFirstOrderDeposit{}).Where("user_id = ?", vo.UserId).First(&row).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果不存在则插入
			nowTime := automaticType.Now()
			row.CreateTime = nowTime
			row.PayTime = nowTime
			row.AuthTime = nowTime
			tool.JsonMapper(vo, &row)

			flag, _ := modules.SaveFcFirstOrderDeposit(&row)
			if !flag {
				global.G_LOG.Errorf("FirstDepositeSave create first deposite userId: %s orderSn: %v err: %v", vo.UserId, vo.OrderSn, err)
				return err
			}

			global.G_REDIS.SetEx(context.Background(), firstDepositeKey, 1, 24*time.Hour).Val()
			return nil
		}

		return err
	}

	global.G_REDIS.SetEx(context.Background(), firstDepositeKey, 1, 24*time.Hour).Val()

	return nil
}
