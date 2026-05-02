package paymentOut

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/userTransfer"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// InsertOrderWithdrawAnotherPay
//
//	@Description: 新增代付订单
func InsertOrderWithdrawAnotherPay(vo *dos.FcOrderWithdraw, anotherPay *dos.FcOrderWithdrawPaymentOut) error {
	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		eRow := tx.Model(&dos.FcOrderWithdraw{}).Where("id = ? and status = ? and another_pay_status = ?",
			vo.Id,
			enmus.OrderWithdrawStats_AuditWait,
			enmus.OrderWithdrawAnotherPayStats_None).
			Updates(map[string]interface{}{
				"status":             enmus.OrderWithdrawStats_AuditApprove,
				"another_pay_status": enmus.OrderWithdrawAnotherPayStats_Progress,
			})
		if eRow.Error != nil {
			global.G_LOG.Errorf("[InsertOrderWithdrawAnotherPay] Update FcOrderWithdraw failed: id=%s, err=%s", vo.Id, eRow.Error.Error())
			return eRow.Error
		}
		if eRow.RowsAffected != 1 {
			global.G_LOG.Errorf("[InsertOrderWithdrawAnotherPay] Update FcOrderWithdraw failed: id=%s", vo.Id)
			return fmt.Errorf("更新提现订单(ID:%s)状态失败", vo.Id)
		}
		eRow = tx.Create(anotherPay)
		if eRow.Error != nil {
			global.G_LOG.Errorf("[InsertOrderWithdrawAnotherPay] Create FcOrderWithdrawPaymentOut failed: data=%+v, err=%s", anotherPay, eRow.Error.Error())
			return eRow.Error
		}
		if eRow.RowsAffected != 1 {
			global.G_LOG.Errorf("[InsertOrderWithdrawAnotherPay] Create FcOrderWithdrawPaymentOut failed: data=%+v", anotherPay)
			return fmt.Errorf("创建汇款申请订单失败")
		}
		return nil
	})
	return err
}

// OrderWithdrawAnotherPaySuccess
//
//	@Description: 代付成功
func OrderWithdrawAnotherPaySuccess(orderWithdrawAnotherPayId string, updateBy string) (bool, bool) {
	var m dos.FcOrderWithdrawPaymentOut
	global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).Where("id = ?", orderWithdrawAnotherPayId).Take(&m)
	if len(m.Id) == 0 {
		global.G_LOG.Errorf("[OrderWithdrawAnotherPaySuccess] Can't find order withdraw payment out data: id=%s", orderWithdrawAnotherPayId)
		return true, true
	}

	if m.WithdrawStatus == enmus.OrderWithdrawStats_ManualNo ||
		m.WithdrawStatus == enmus.OrderWithdrawStats_ManualYes {
		// 是人工操作就不继续处理
		return true, false
	}

	sets := map[string]interface{}{
		"status":    enmus.OrderWithdrawPaymentOutStats_Success,
		"update_by": updateBy,
	}

	if m.WithdrawStatus == enmus.OrderWithdrawStats_No { // 非人工打款状态则同步状态
		sets["withdraw_status"] = enmus.OrderWithdrawStats_Yes
	}

	var isOk bool
	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		row := tx.Model(&dos.FcOrderWithdrawPaymentOut{}).
			Where("id = ? and status = ?", m.Id, enmus.OrderWithdrawPaymentOutStats_Progress).
			Updates(sets).RowsAffected
		if row != 1 {
				return errors.New("update withdraw payment out status fail")
		}
		row = tx.Model(&dos.FcOrderWithdraw{}).
			Where("order_sn = ? and another_pay_status = ? and status = ?",
				m.OrderSn,
				enmus.OrderWithdrawAnotherPayStats_Progress,
				enmus.OrderWithdrawStats_AuditApprove,
			).
			Updates(
				map[string]interface{}{
					"another_pay_status": enmus.OrderWithdrawAnotherPayStats_Success,
					//"status":             3,
					"another_pay_time": tool.TimeNowString(),
					"callback_remark":  "打款成功",
					"auth_by":          updateBy,
					"auth_time":        automaticType.Time(time.Now()),
				}).RowsAffected
		if row != 1 {
				return errors.New("update withdraw another pay status fail")
		}

		tx.Model(&dos.FcTranscation{}).Where("related_id=?", m.OrderSn).Updates(map[string]interface{}{
			"status": 1,
		})

		// 如果存在手续费，单独扣除并写流水记录
		if m.Fee > 0 {
			if feeErr := userTransfer.UserAmountChange(tx, -m.Fee, userTransfer.TranServiceCharge, m.Currency,
				m.OrderSn, m.UserId, updateBy, m.OrderSn, "提款手续费"); feeErr != nil {
				global.G_LOG.Errorf("[OrderWithdrawAnotherPaySuccess] Deduct service fee failed: orderSn=%s, fee=%v, err=%s", m.OrderSn, m.Fee, feeErr.Error())
				return feeErr
			}
		}

		isOk = true
		return nil
	})
	if err != nil {
		global.G_LOG.Errorf("代付订单：%s err :%s", orderWithdrawAnotherPayId, err.Error())
	}
	return isOk, true
}

// OrderWithdrawAnotherPayFail
//
//	@Description: 代付失败
func OrderWithdrawAnotherPayFail(orderWithdrawAnotherPayId string, updateBy string) bool {
	return OrderWithdrawAnotherPayFailRemark(orderWithdrawAnotherPayId, "", 1, updateBy)
}

func OrderWithdrawAnotherPayFailRemark(orderWithdrawAnotherPayId, remark string, status int, updateBy string) bool {
	var m dos.FcOrderWithdrawPaymentOut
	global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}).Where("id = ?", orderWithdrawAnotherPayId).Take(&m)
	if len(m.Id) == 0 {
		global.G_LOG.Errorf("[OrderWithdrawAnotherPayFailRemark] Can't find order withdraw payment out data failed: id=%s",
			orderWithdrawAnotherPayId)
		return true
	}

	if m.WithdrawStatus == enmus.OrderWithdrawStats_ManualNo ||
		m.WithdrawStatus == enmus.OrderWithdrawStats_ManualYes {
		// 是人工操作就不继续处理
		return true
	}

	var vo dos.FcOrderWithdraw
	global.G_DB.Model(&dos.FcOrderWithdraw{}).Where("order_sn = ?", m.OrderSn).Take(&vo)
	if len(vo.Id) == 0 {
		global.G_LOG.Errorf("[OrderWithdrawAnotherPayFailRemark] Can't find order withdraw data failed: id=%s, orderSn=%s",
			orderWithdrawAnotherPayId, m.OrderSn)
		return true
	}

	var isOk bool
	global.G_DB.Transaction(func(tx *gorm.DB) error {
		callbackRemark := remark
		if len(remark) == 0 {
			callbackRemark = "财务退回"
		}

		row := tx.Model(&dos.FcOrderWithdrawPaymentOut{}).
			Where("id = ? and status = ?", m.Id, status).
			Updates(map[string]interface{}{
				"status":      enmus.OrderWithdrawPaymentOutStats_Failed,
				"remark":      callbackRemark,
				"update_by":   updateBy,
				"update_time": automaticType.Now(),
			}).RowsAffected
		if row != 1 {
			return errors.New("update withdraw payment out status fail")
		}

		row = tx.Model(&dos.FcOrderWithdraw{}).
			Where("order_sn = ? and another_pay_status = ? and status = ?",
				m.OrderSn,
				enmus.OrderWithdrawAnotherPayStats_Progress,
				enmus.OrderWithdrawStats_AuditApprove).
			Updates(map[string]interface{}{
				"another_pay_status": enmus.OrderWithdrawAnotherPayStats_Failed,
				"callback_remark":    callbackRemark,
				"another_pay_time":   tool.TimeNowString(),
				"auth_by":            updateBy,
				"auth_time":          automaticType.Time(time.Now()),
			}).RowsAffected
		if row != 1 {
			return errors.New("update withdraw another pay status fail")
		}

		// 退回用户提款金额
		err := userTransfer.UserAmountChange(tx, vo.Amount, userTransfer.TranWithdrawReject, vo.Currency, vo.OrderSn, vo.UserId, vo.CreateBy, vo.OrderSn,
			modules.GetFcTranscationFundingSubType(enmus.FundingTypeWithdraw, "orderreturn"))
		if err != nil {
			global.G_LOG.Errorf("[OrderWithdrawAnotherPayFailRemark] userTransfer.UserAmountChange failed: %s", err.Error())
			return err
		}

		tx.Model(&dos.FcTranscation{}).Where("related_id = ? AND funding_subtype = ?",
			vo.OrderSn, modules.GetFcTranscationFundingSubType(enmus.FundingTypeWithdraw, "orderreturn")).
			Updates(map[string]interface{}{
				"remark": userTransfer.GetRemarkByOrderType(&vo),
			})

		err = tx.Model(&dos.FcTranscation{}).Where("related_id = ? AND funding_type = ?", m.OrderSn, int(userTransfer.TranWithdraw)).
			Updates(map[string]interface{}{
				"status": 2,
			}).Error
		if err != nil {
			global.G_LOG.Errorf("[OrderWithdrawAnotherPayFailRemark] Update transcation status failed: %s", err.Error())
			return err
		}

		// 如果存在手续费，单独扣除并写流水记录
		if m.Fee > 0 {
			if feeErr := userTransfer.UserAmountChange(tx, -m.Fee, userTransfer.TranServiceCharge, m.Currency,
				m.OrderSn, m.UserId, updateBy, m.OrderSn, "提款手续费"); feeErr != nil {
				global.G_LOG.Errorf("[OrderWithdrawAnotherPaySuccess] Deduct service fee failed: orderSn=%s, fee=%v, err=%s", m.OrderSn, m.Fee, feeErr.Error())
				return feeErr
			}
		}

		isOk = true
		return nil
	})

	return isOk
}

// GetDictConfigPaymentOutWaitDuration - 获取提款订单排队时长（秒）
// @returns int
func GetDictConfigPaymentOutWaitDuration() int64 {
	rebateCostRate := modules.FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTypeCode: "withdraw_sendingwaittime",
		DictsTag:      "withdraw_sendingwaittime",
	})

	rebateCostRateDefault := int64(30)

	if rebateCostRate != nil {
		val, err := strconv.ParseInt(rebateCostRate.DictsValue, 10, 64)
		if err != nil {
			//global.G_LOG.Errorf("[GetDictConfigPaymentOutWaitDuration] ParseFloat failed: %v", rebateCostRate.DictsValue)
			return rebateCostRateDefault
		}

		if val < 0 {
			global.G_LOG.Errorf("[GetDictConfigPaymentOutWaitDuration] Configure dict wrong: %v", val)
			return rebateCostRateDefault
		}

		return val
	} else {
		global.G_LOG.Error("[GetDictConfigPaymentOutWaitDuration] Can't find withdraw_sendingwaittime.withdraw_sendingwaittime")
		return rebateCostRateDefault
	}
}
