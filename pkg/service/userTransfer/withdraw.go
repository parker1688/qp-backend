package userTransfer

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetRemarkByOrderType(vo *dos.FcOrderWithdraw) string {
	var remark string
	switch vo.OrderType {
	case enmus.ORDER_TYPE_BANK:
		remark = vo.AccountBankType
	case enmus.ORDER_TYPE_Virtual:
		remark = vo.VirtualCurrencyChain
	case enmus.ORDER_TYPE_Online:
		remark = "支付宝"
	}

	return remark
}

func UserWithdraw(vo *dos.FcOrderWithdraw) (bool, error) {
	var isOk bool
	var err error
	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		//修改用户金额表
		err = UserAmountChange(tx, -vo.Amount, TranWithdraw, vo.Currency, GetRemarkByOrderType(vo), vo.UserId, vo.CreateBy, vo.OrderSn,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypeWithdraw, fmt.Sprintf("ordertype_%d", vo.OrderType)))
		if err != nil {
			return err
		}
		eRow := tx.Create(vo)
		err = eRow.Error
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

func UserWithdrawReject(voData *dos.FcOrderWithdraw) (bool, error) {
	var vo *dos.FcOrderWithdraw
	global.G_DB.Model(&dos.FcOrderWithdraw{}).Where("id = ?", voData.Id).Take(&vo)
	if len(vo.Id) == 0 {
		global.G_LOG.Errorf("[UserWithdrawReject] Find order withdraw failed: id=%s", voData.Id)
		return false, errors.New("order does not exist")
	}
	if len(vo.CallbackRemark) > 0 {
		voData.CallbackRemark = vo.CallbackRemark
	}
	if len(vo.Remark) > 0 {
		voData.Remark = vo.Remark
	}
	var isOk bool
	var err error
	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		//待处理才进行处理
		eRow := tx.Exec(`update fc_order_withdraw set status=?,remark=?,update_time=?,update_by=? where id=? and status=?`,
			enmus.OrderWithdrawStats_AuditReject,
			voData.Remark,
			automaticType.Now(),
			voData.UpdateBy,
			vo.Id,
			enmus.OrderWithdrawStats_AuditWait)
		if eRow.Error != nil {
			err = eRow.Error
			return err
		}
		if eRow.RowsAffected != 1 {
			err = errors.New("update withdraw status fail")
			return err
		}
		//修改用户金额表
		err = UserAmountChange(tx, vo.Amount, TranWithdrawReject, vo.Currency, vo.OrderSn, vo.UserId, vo.CreateBy, vo.OrderSn,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypeWithdraw, "orderrefuse"))
		if err != nil {
			return err
		}
		isOk = true
		// 返回 nil 提交事务
		tx.Model(&dos.FcTranscation{}).Where("related_id=? AND funding_type = ?",
			vo.OrderSn, int(TranWithdraw)).
			Updates(map[string]interface{}{
				"status": 2,
			})

		tx.Model(&dos.FcTranscation{}).Where("related_id = ? AND funding_subtype = ?",
			vo.OrderSn, modules.GetFcTranscationFundingSubType(enmus.FundingTypeWithdraw, "orderrefuse")).
			Updates(map[string]interface{}{
				"remark": GetRemarkByOrderType(vo),
			})

		return nil
	})
	if err != nil {
		return false, err
	}
	//提款失败系统邮件提示
	go modules.SendSystemMail(vo.UserId, vo.MerchantCode, enmus.MailType_WithdrawFail, vo.Amount)
	return isOk, err
}

func UserWithdrawSuccess(vo *dos.FcOrderWithdraw) (bool, error) {
	var isOk bool
	var err error
	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		//审核通过才能改状态
		err = updateWithdrawSuccess(tx, vo)
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

// UserWithdrawAuditSuccess
//
//	@Description: 用户审核通过
//	@param vo
//	@return bool
//	@return error
func UserWithdrawAuditSuccess(vo *dos.FcOrderWithdraw) bool {
	affect := global.G_DB.Exec(`update fc_order_withdraw set status=?,update_by=?,remark=? where status=? and id=?`,
		enmus.ORDER_STATUS_WAIT, vo.UpdateBy, vo.Remark, enmus.ORDER_PENDING_STATUS, vo.Id).RowsAffected
	return affect == 1
}

func updateWithdrawSuccess(tx *gorm.DB, vo *dos.FcOrderWithdraw) error {
	updates := map[string]interface{}{
		"remark":    vo.Remark,
		"update_by": vo.UpdateBy,
		"status":    enmus.ORDER_YES_STATUS,
	}
	if vo.OrderType == enmus.ORDER_TYPE_Virtual {
		updates["virtual_pay_address"] = vo.VirtualPayAddress
		updates["virtual_pay_no"] = vo.VirtualPayNo
		updates["virtual_pay_amount"] = vo.VirtualPayAmount
	}
	tx.Model(&dos.FcTranscation{}).Where("related_id=?", vo.OrderSn).Updates(map[string]interface{}{
		"status": 1,
	})

	eRow := tx.Model(vo).Where(`id = ? and status=?`, vo.Id, enmus.ORDER_STATUS_WAIT).Updates(updates)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected != 1 {
		return errors.New("update withdraw status fail")
	}
	return nil

}

// VenueWithdrawSuccess
//
//	@Description: 场馆提款成功
//	@param vo
func VenueWithdrawSuccess(voData *dos.FcVenueTransfer) (bool, error) {
	var vo *dos.FcVenueTransfer
	global.G_DB.Model(&dos.FcVenueTransfer{}).Where("id = ?", voData.Id).Take(&vo)
	if len(vo.Id) == 0 {
		return false, errors.New("order does not exist")
	}
	var isOk bool
	var err error
	err = global.G_DB.Transaction(func(tx *gorm.DB) error {
		eRow := tx.Model(&dos.FcVenueTransfer{}).Where("id = ? and status=0", vo.Id).Updates(map[string]interface{}{
			"update_by": voData.UpdateBy,
			"status":    1,
		})
		if eRow.Error != nil {
			return eRow.Error
		}
		if eRow.RowsAffected != 1 {
			return errors.New("update venue withdraw status fail")
		}
		//修改用户金额表
		err = UserAmountChange(tx, vo.Amount, TranWithdraw, vo.Currency, "<===【"+vo.VenueCode+"】"+vo.OrderSn+" 【Success】", vo.UserId, vo.CreateBy, vo.OrderSn,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypePlatform, "withdraw")+vo.VenueCode)
		if err != nil {
			return err
		}
		isOk = true
		return err
	})
	if err != nil {
		return false, err
	}
	return isOk, err
}

// UserWithdrawPaymentOutNo - 财务汇款未打款
func UserWithdrawPaymentOutNo(voData *dos.FcOrderWithdrawPaymentOut) error {
	var vo *dos.FcOrderWithdraw
	global.G_DB.Model(&dos.FcOrderWithdraw{}).Where("order_sn = ?", voData.OrderSn).Take(&vo)
	if len(vo.Id) == 0 {
		global.G_LOG.Errorf("[UserWithdrawPaymentOutNo] Find order withdraw failed: orderSn=%s", voData.OrderSn)
		return errors.New("order does not exist")
	}

	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		//待处理才进行处理
		eRow := tx.Exec(`update fc_order_withdraw_payment_out set withdraw_status = ?, remark = ?, update_by = ?, update_time = ? where id = ? AND status in ? AND withdraw_status = ?`,
			voData.WithdrawStatus,
			voData.Remark,
			voData.UpdateBy,
			automaticType.Now(),
			voData.Id,
			[]int{
				enmus.OrderWithdrawPaymentOutStats_Prepare,
				enmus.OrderWithdrawPaymentOutStats_Progress,
			},
			enmus.OrderWithdrawStats_No)
		if eRow.Error != nil {
			global.G_LOG.Errorf("[UserWithdrawPaymentOutNo] update order withdraw payment out failed: id=%s", voData.Id)
			return eRow.Error
		}
		if eRow.RowsAffected != 1 {
			return errors.New("update withdraw payment out status fail")
		}
		// 修改用户金额表
		err1 := UserAmountChange(tx, vo.Amount, TranWithdrawReject, vo.Currency, vo.OrderSn, vo.UserId, vo.CreateBy, vo.OrderSn,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypeWithdraw, "orderrefuse"))
		if err1 != nil {
			global.G_LOG.Errorf("[UserWithdrawPaymentOutNo] UserAmountChange failed: id=%s, err=%s", voData.Id, err1.Error())
			return err1
		}

		err1 = tx.Model(&dos.FcTranscation{}).Where("related_id = ? AND funding_type = ?", vo.OrderSn, int(TranWithdraw)).
			Updates(map[string]interface{}{
				"status": 2,
			}).Error
		if err1 != nil {
			global.G_LOG.Errorf("[UserWithdrawPaymentOutNo] Update transcation status failed: id=%s, orderSn=%s, err=%s", voData.Id, vo.OrderSn, err1.Error())
			return err1
		}

		tx.Model(&dos.FcTranscation{}).Where("related_id = ? AND funding_subtype = ?",
			vo.OrderSn, modules.GetFcTranscationFundingSubType(enmus.FundingTypeWithdraw, "orderrefuse")).
			Updates(map[string]interface{}{
				"remark": GetRemarkByOrderType(vo),
			})

		eRow = tx.Model(&dos.FcOrderWithdraw{}).Where("order_sn = ? AND another_pay_status = ?", vo.OrderSn, enmus.OrderWithdrawAnotherPayStats_Progress).
			Updates(map[string]interface{}{
				"another_pay_status": enmus.OrderWithdrawAnotherPayStats_Failed,
			})
		if eRow.Error != nil {
			global.G_LOG.Errorf("[UserWithdrawPaymentOutNo] Update order withdraw another status failed: id=%s, orderSn=%s, err=%s", voData.Id, vo.OrderSn, eRow.Error.Error())
			return eRow.Error
		}
		if eRow.RowsAffected != 1 {
			return errors.New("update withdraw another pay status fail")
		}

		return nil
	})
	return err
}
