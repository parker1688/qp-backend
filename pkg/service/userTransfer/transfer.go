package userTransfer

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/langs"
	"bootpkg/pkg/core/modules/dos"
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// UserAmountChange
//
//	@Description: 用户金额变动
//	@param tx 事务tx
//	@param amount 处理金额 +/-
//	@param tranType 帐变类型
//	@param remark 备注
//	@param userId 用户ID
//	@param currency 币种类型
//	@param createBy 创建人
//	@return error 错误信息
func UserAmountChange(tx *gorm.DB, amount float64, funding TransactionAmountType, currency string, remark string, userId string, createBy string, relatedId string, fundingSubType string) error {
	var wallet dos.FcUserWallet
	tx.Model(&dos.FcUserWallet{}).Where("currency=? and user_id=?", currency, userId).Take(&wallet)
	if len(wallet.Id) == 0 {
		return ecode.NewErrorCode("amount_change_1", currency+" 币种钱包不存在", &langs.Replacements{"currency": currency})
	}
	//预算金额
	beforeAmount := decimal.NewFromFloat(wallet.AvaAmount)
	afterAmount := beforeAmount.Add(decimal.NewFromFloat(amount)).Truncate(2).InexactFloat64()
	if afterAmount < 0 {
		return ecode.NewErrorCode("amount_change_2", "钱包金额不足")
	}
	//修改可用金额
	eRow := tx.Exec("update fc_user_wallet set ava_amount=ava_amount+? where id=? and ava_amount+?>=0", amount, wallet.Id, amount)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected != 1 {
		return errors.New("update wallet err")
	}

	fcTrans := &dos.FcTranscation{
		UserId:            wallet.UserId,
		UserName:          wallet.UserName,
		Status:            1, //1:通过  2：不通过
		Amount:            amount,
		TranscationBefore: wallet.AvaAmount,
		TranscationAfter:  afterAmount,
		Remark:            remark,
		FundingType:       int(funding),
		CreateBy:          createBy,
		MerchantCode:      wallet.MerchantCode,
		Currency:          currency,
		RelatedId:         relatedId,
		CreateTime:        automaticType.Time(time.Now()),
		FundingSubtype:    fundingSubType,
	}
	//新增
	eRow = tx.Create(fcTrans)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected != 1 {
		return errors.New("save transaction err")
	}
	return nil
}

func UserAmountChange2(tx *gorm.DB, amount float64, funding TransactionAmountType, currency string, remark string,
	userId string, createBy string, relatedId string, manualRelatedId string, fundingSubType string) error {
	var wallet dos.FcUserWallet
	tx.Model(&dos.FcUserWallet{}).Where("currency=? and user_id=?", currency, userId).Take(&wallet)
	if len(wallet.Id) == 0 {
		return ecode.NewErrorCode("amount_change_1", currency+" 币种钱包不存在", &langs.Replacements{"currency": currency})
	}
	//预算金额
	beforeAmount := decimal.NewFromFloat(wallet.AvaAmount)
	afterAmount := beforeAmount.Add(decimal.NewFromFloat(amount)).Truncate(2).InexactFloat64()
	if afterAmount < 0 {
		return ecode.NewErrorCode("amount_change_2", "钱包金额不足")
	}
	//修改可用金额
	eRow := tx.Exec("update fc_user_wallet set ava_amount=ava_amount+? where id=? and ava_amount+?>=0", amount, wallet.Id, amount)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected != 1 {
		return errors.New("update wallet err")
	}
	fcTrans := &dos.FcTranscation{
		UserId:            wallet.UserId,
		UserName:          wallet.UserName,
		Status:            1, //1:通过  2：不通过
		Amount:            amount,
		TranscationBefore: wallet.AvaAmount,
		TranscationAfter:  afterAmount,
		Remark:            remark,
		FundingType:       int(funding),
		CreateBy:          createBy,
		MerchantCode:      wallet.MerchantCode,
		Currency:          currency,
		RelatedId:         relatedId,
		ManualRelatedId:   manualRelatedId,
		CreateTime:        automaticType.Time(time.Now()),
		FundingSubtype:    fundingSubType,
	}
	//新增
	eRow = tx.Create(fcTrans)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected != 1 {
		return errors.New("save transaction err")
	}
	return nil
}
