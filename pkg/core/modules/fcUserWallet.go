// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcUserWallet(vo *dos.FcUserWallet) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserWallet(page, pageSize int, vo *dos.FcUserWallet, c *gin.Context) (ret []*dos.FcUserWallet, total int64) {
	query := global.G_DB.Model(&dos.FcUserWallet{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.IsLock > 0 {
		query = query.Where("is_lock = ?", vo.IsLock)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserWallet
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserWallet(vo *dos.FcUserWallet, c *gin.Context) []*dos.FcUserWallet {
	var data []*dos.FcUserWallet
	query := global.G_DB.Model(&dos.FcUserWallet{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.IsLock > 0 {
		query = query.Where("is_lock = ?", vo.IsLock)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserWalletFirst(vo *dos.FcUserWallet) *dos.FcUserWallet {
	var data *dos.FcUserWallet
	query := global.G_DB.Model(&dos.FcUserWallet{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.IsLock > 0 {
		query = query.Where("is_lock = ?", vo.IsLock)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

func FindByUserIdsFcUserWallet(userIdArr []string) []*dos.FcUserWallet {
	var data []*dos.FcUserWallet
	query := global.G_DB.Model(&dos.FcUserWallet{})
	query.Where("user_id in ?", userIdArr)

	query.Find(&data)
	return data
}

// 根据主键Update
func UpdateFcUserWallet(vo *dos.FcUserWallet) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":        vo.UserId,
		"user_name":      vo.UserName,
		"currency":       vo.Currency,
		"total_amount":   vo.TotalAmount,
		"ava_amount":     vo.AvaAmount,
		"fronzen_amount": vo.FronzenAmount,
		"is_lock":        vo.IsLock,
		"create_by":      vo.CreateBy,
		"update_by":      vo.UpdateBy,
		"merchant_code":  vo.MerchantCode,
	}).Error == nil
}

func DeleteFcUserWallet(vo *dos.FcUserWallet) bool {
	return global.G_DB.Model(&dos.FcUserWallet{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func FindPageOrderManageOpt(page, pageSize int, vo *dos.FcOrderManageOpt, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcOrderManageOpt, total int64) {
	query := global.G_DB.Model(&dos.FcOrderManageOpt{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.TrsType > 0 {
		query = query.Where("trs_type = ?", vo.TrsType)
	}

	if vo.ScoreType > 0 {
		query = query.Where("score_type = ?", vo.ScoreType)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if pageTimeQuery.StartAt != "" {
		query = query.Where("create_time >= ?", pageTimeQuery.StartAt)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("create_time <= ?", pageTimeQuery.EndAt)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcOrderManageOpt
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

// GetFcUserWalletRemainingAmount - 获取用户余额
// @param {string} userId
// @returns float64
func GetFcUserWalletRemainingAmount(userId string) float64 {
	data := dos.FcUserWallet{}
	err := global.G_DB.Model(&dos.FcUserWallet{}).
		Select("ava_amount").
		Where("user_id = ?", userId).
		Take(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcUserWalletAvaAmount] Find user wallet failed: userId=%s, err=%s",
			userId, err.Error())
		return 0
	}

	return data.AvaAmount
}

// GetUserRemainingAmountAndRechargeRatio - 获取用户余额及充值比
func GetUserRemainingAmountAndRechargeRatio(userId string) (float64, float64) {
	remainingAmount := GetFcUserWalletRemainingAmount(userId)
	todayRechargeAmount := GetUserTodayRechargeAmount(userId)
	if todayRechargeAmount == 0 {
		return remainingAmount, 0
	}

	return remainingAmount, (remainingAmount / todayRechargeAmount)
}
