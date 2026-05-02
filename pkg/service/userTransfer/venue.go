package userTransfer

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/pkg/core/modules/dos"
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func lockUserWallet(tx *gorm.DB, currency, userId string) (dos.FcUserWallet, error) {
	var wallet dos.FcUserWallet
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&dos.FcUserWallet{}).
		Where("currency=? and user_id=?", currency, userId).
		Take(&wallet).Error
	if err != nil {
		return wallet, err
	}
	if len(wallet.Id) == 0 {
		return wallet, errors.New(currency + " wallet does not exist")
	}
	return wallet, nil
}

// UserVenueAmountChange
//
//	@Description: 用户转入场馆金额 - 锁定金额
//	@param tx 事务tx
//	@param amount 处理金额(正金额)
//	@param tranType 帐变类型
//	@param remark 备注
//	@param userId 用户ID
//	@param currency 币种类型
//	@param createBy 创建人
//	@return error 错误信息
func UserVenueAmountChange(tx *gorm.DB, amount float64, funding TransactionAmountType, currency string, remark string, userId string, createBy string, relatedId string, fundingSubType string) error {
	wallet, err := lockUserWallet(tx, currency, userId)
	if err != nil {
		return err
	}
	beforeAmount := decimal.NewFromFloat(wallet.AvaAmount)
	//减
	afterAmount := beforeAmount.Sub(decimal.NewFromFloat(amount)).Truncate(2).InexactFloat64()
	if afterAmount < 0 || amount <= 0 {
		return errors.New("amount after the change is less than 0")
	}
	//可用金额转入冻结金额
	eRow := tx.Exec("update fc_user_wallet set ava_amount=ava_amount-?,fronzen_amount=fronzen_amount+? where id=? and ava_amount-?>=0", amount, amount, wallet.Id, amount)
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
		Amount:            -amount,
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

// UserVenueAmountChangeCallback
//
//	@Description: 用户转入场馆金额 - 失败
//	@param tx 事务tx
//	@param amount 处理金额(正金额)
//	@param tranType 帐变类型
//	@param remark 备注
//	@param userId 用户ID
//	@param currency 币种类型
//	@param createBy 创建人
//	@return error 错误信息
func UserVenueAmountChangeCallback(tx *gorm.DB, amount float64, funding TransactionAmountType, currency string, remark string, userId string, createBy string, relatedId string, fundingSubType string) error {
	wallet, err := lockUserWallet(tx, currency, userId)
	if err != nil {
		return err
	}
	beforeAmount := decimal.NewFromFloat(wallet.AvaAmount)
	//加
	afterAmount := beforeAmount.Add(decimal.NewFromFloat(amount)).Truncate(2).InexactFloat64()
	if afterAmount < 0 || amount <= 0 {
		return errors.New("amount after the change is less than 0")
	}
	//可用金额转入冻结金额
	eRow := tx.Exec("update fc_user_wallet set ava_amount=ava_amount+?,fronzen_amount=fronzen_amount-? where id=? and fronzen_amount-?>=0", amount, amount, wallet.Id, amount)
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

// UserVenueAmountConfirmChange
//
//	@Description: 用户转入场馆金额 - 确认成功
//	@param tx 事务tx
//	@param amount 处理金额(正金额)
//	@param tranType 帐变类型
//	@param remark 备注
//	@param userId 用户ID
//	@param currency 币种类型
//	@param createBy 创建人
//	@return error 错误信息
func UserVenueAmountConfirmChange(tx *gorm.DB, amount float64, funding TransactionAmountType, currency string, remark string, userId string, createBy string) error {
	wallet, err := lockUserWallet(tx, currency, userId)
	if err != nil {
		return err
	}
	//可用金额转入冻结金额
	eRow := tx.Exec("update fc_user_wallet set fronzen_amount=fronzen_amount-? where id=? and fronzen_amount-?>=0", amount, wallet.Id, amount)
	if eRow.Error != nil {
		return eRow.Error
	}
	if eRow.RowsAffected != 1 {
		return errors.New("update wallet err")
	}
	return nil
}
