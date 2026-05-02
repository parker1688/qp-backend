package userTransfer

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"errors"

	"gorm.io/gorm"
)

// PromotionOk
//
//	@Description: 优惠申请通过
//	@param voData
//	@return bool
//	@return error
func PromotionOk(voData *dos.FcOrderPromotion) (bool, error) {
	var vo *dos.FcOrderPromotion
	global.G_DB.Model(&dos.FcOrderPromotion{}).Where("id = ?", voData.Id).Take(&vo)
	if len(vo.Id) == 0 {
		return false, errors.New("order does not exist")
	}
	var isOk bool
	var err error
	global.G_DB.Transaction(func(tx *gorm.DB) error {
		//待处理才进行处理
		eRow := tx.Exec(`update fc_order_promotion set status=? where status=? and id=?`, enmus.ORDER_YES_STATUS, enmus.ORDER_PENDING_STATUS, vo.Id)
		if eRow.Error != nil {
			err = eRow.Error
			return err
		}
		if eRow.RowsAffected != 1 {
			err = errors.New("update promotion status fail")
			return err
		}
		//打码
		if err != nil {
			return err
		}
		//修改用户金额表
		err = UserAmountChange(tx, vo.Amount, TranDiscount, vo.Currency, vo.OrderSn, vo.UserId, vo.UpdateBy, "", "")
		if err != nil {
			return err
		}
		isOk = true
		// 返回 nil 提交事务
		return nil
	})
	return isOk, err
}
