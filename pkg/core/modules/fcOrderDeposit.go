// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var orderDepositCache = cache.New(3*time.Minute, 30*time.Minute)

func SaveFcOrderDeposit(vo *dos.FcOrderDeposit) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcOrderDeposit(page, pageSize int, vo *dos.FcOrderDeposit, pageQuery response.PageTimeQuery, c *gin.Context) (ret []*dos.FcOrderDeposit, total int64, sumAmount float64) {
	page, pageSize = response.NormalizePage(page, pageSize)
	query := global.G_DB.Model(&dos.FcOrderDeposit{})
	query1 := global.G_DB.Model(&dos.FcOrderDeposit{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
		query1 = query1.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
		query1 = query1.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
		query1 = query1.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
		query1 = query1.Where("currency = ?", vo.Currency)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
		query1 = query1.Where("order_sn = ?", vo.OrderSn)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
		query1 = query1.Where("status = ?", vo.Status)
	}

	if len(vo.EntityAccountHolder) > 0 {
		query = query.Where("entity_account_holder = ?", vo.EntityAccountHolder)
		query1 = query1.Where("entity_account_holder = ?", vo.EntityAccountHolder)
	}

	if len(vo.EntityAccountBankName) > 0 {
		query = query.Where("entity_account_bank_name = ?", vo.EntityAccountBankName)
	}

	if len(vo.EntityAccountNumber) > 0 {
		query = query.Where("entity_account_number = ?", vo.EntityAccountNumber)
	}

	if len(vo.RemitterAccountHolder) > 0 {
		query = query.Where("remitter_account_holder = ?", vo.RemitterAccountHolder)
	}

	if len(vo.RemitterAccountBankName) > 0 {
		query = query.Where("remitter_account_bank_name = ?", vo.RemitterAccountBankName)
	}

	if len(vo.RemitterAccountNumber) > 0 {
		query = query.Where("remitter_account_number = ?", vo.RemitterAccountNumber)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.DepositRemark) > 0 {
		query = query.Where("deposit_remark = ?", vo.DepositRemark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
		query1 = query1.Where("ip = ?", vo.Ip)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
		query1 = query1.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
		query1 = query1.Where("update_by = ?", vo.UpdateBy)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
		query1 = query1.Where("merchant_code = ?", vo.MerchantCode)
	}

	if vo.ChannelId > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
		query1 = query1.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
		query1 = query1.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.PaymentId > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
		query1 = query1.Where("payment_code = ?", vo.PaymentCode)
	}

	if len(pageQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageQuery.StartAt)
		query1 = query1.Where("create_time >= ?", pageQuery.StartAt)
	}
	if len(pageQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageQuery.EndAt)
		query1 = query1.Where("create_time <= ?", pageQuery.EndAt)
	}

	if len(pageQuery.LastStartAt) > 0 {
		query = query.Where("pay_time >= ?", pageQuery.StartAt)
		query1 = query1.Where("pay_time >= ?", pageQuery.StartAt)
	}
	if len(pageQuery.LastEndAt) > 0 {
		query = query.Where("pay_time <= ?", pageQuery.EndAt)
		query1 = query1.Where("pay_time <= ?", pageQuery.EndAt)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, 0, 0
		}

		if query1, ok = QueryAdminUserMerchantCodes(c, query1, vo.MerchantCode); !ok {
			return ret, 0, 0
		}
	}

	var count int64
	sumAmount = 0.00
	query1.Select("sum(amount) as sumAmount").Scan(&sumAmount)
	query.Count(&count)
	var dataSlice []*dos.FcOrderDeposit
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count, sumAmount
}

func FindByKeyFcOrderDeposit(vo *dos.FcOrderDeposit, c *gin.Context) []*dos.FcOrderDeposit {
	var data []*dos.FcOrderDeposit
	query := global.G_DB.Model(&dos.FcOrderDeposit{})

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

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.EntityAccountHolder) > 0 {
		query = query.Where("entity_account_holder = ?", vo.EntityAccountHolder)
	}

	if len(vo.EntityAccountBankName) > 0 {
		query = query.Where("entity_account_bank_name = ?", vo.EntityAccountBankName)
	}

	if len(vo.EntityAccountNumber) > 0 {
		query = query.Where("entity_account_number = ?", vo.EntityAccountNumber)
	}

	if len(vo.RemitterAccountHolder) > 0 {
		query = query.Where("remitter_account_holder = ?", vo.RemitterAccountHolder)
	}

	if len(vo.RemitterAccountBankName) > 0 {
		query = query.Where("remitter_account_bank_name = ?", vo.RemitterAccountBankName)
	}

	if len(vo.RemitterAccountNumber) > 0 {
		query = query.Where("remitter_account_number = ?", vo.RemitterAccountNumber)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.DepositRemark) > 0 {
		query = query.Where("deposit_remark = ?", vo.DepositRemark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
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

	if vo.ChannelId > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.PaymentId > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcOrderDepositFirst(vo *dos.FcOrderDeposit) *dos.FcOrderDeposit {
	var data *dos.FcOrderDeposit
	query := global.G_DB.Model(&dos.FcOrderDeposit{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}
	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}
	if len(vo.EntityAccountHolder) > 0 {
		query = query.Where("entity_account_holder = ?", vo.EntityAccountHolder)
	}

	if len(vo.EntityAccountBankName) > 0 {
		query = query.Where("entity_account_bank_name = ?", vo.EntityAccountBankName)
	}

	if len(vo.EntityAccountNumber) > 0 {
		query = query.Where("entity_account_number = ?", vo.EntityAccountNumber)
	}

	if len(vo.RemitterAccountHolder) > 0 {
		query = query.Where("remitter_account_holder = ?", vo.RemitterAccountHolder)
	}

	if len(vo.RemitterAccountBankName) > 0 {
		query = query.Where("remitter_account_bank_name = ?", vo.RemitterAccountBankName)
	}

	if len(vo.RemitterAccountNumber) > 0 {
		query = query.Where("remitter_account_number = ?", vo.RemitterAccountNumber)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if len(vo.DepositRemark) > 0 {
		query = query.Where("deposit_remark = ?", vo.DepositRemark)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
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

	if vo.ChannelId > 0 {
		query = query.Where("channel_id = ?", vo.ChannelId)
	}

	if len(vo.ChannelCode) > 0 {
		query = query.Where("channel_code = ?", vo.ChannelCode)
	}

	if vo.PaymentId > 0 {
		query = query.Where("payment_id = ?", vo.PaymentId)
	}

	if len(vo.PaymentCode) > 0 {
		query = query.Where("payment_code = ?", vo.PaymentCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcOrderDeposit(vo *dos.FcOrderDeposit) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":                    vo.UserId,
		"user_name":                  vo.UserName,
		"order_sn":                   vo.OrderSn,
		"amount":                     vo.Amount,
		"status":                     vo.Status,
		"entity_account_holder":      vo.EntityAccountHolder,
		"entity_account_bank_name":   vo.EntityAccountBankName,
		"entity_account_number":      vo.EntityAccountNumber,
		"remitter_account_holder":    vo.RemitterAccountHolder,
		"remitter_account_bank_name": vo.RemitterAccountBankName,
		"remitter_account_number":    vo.RemitterAccountNumber,
		"remark":                     vo.Remark,
		"deposit_remark":             vo.DepositRemark,
		"ip":                         vo.Ip,
		"update_by":                  vo.UpdateBy,
		"merchant_code":              vo.MerchantCode,
		"channel_id":                 vo.ChannelId,
		"channel_code":               vo.ChannelCode,
		"payment_id":                 vo.PaymentId,
		"payment_code":               vo.PaymentCode,
		"currency":                   vo.Currency,
	}).Error == nil
}

func DeleteFcOrderDeposit(vo *dos.FcOrderDeposit) bool {
	return global.G_DB.Model(&dos.FcOrderDeposit{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// CheckUserPaymentStrategy - 验证用户支付限单
// @param {string} userId
// @param {string} channelCode
// @param {string} merchantCode
// @returns bool
func CheckUserPaymentStrategy(userId string, channelCode string, merchantCode string) bool {
	now := time.Now()
	payStrategyLis := GetFcPayChannelStrategyList(channelCode, merchantCode)

	//global.G_LOG.Infof("[CheckUserPaymentStrategy] userId=%s, channelCode=%s, merchantCode=%s, payStrategyLis=%+v",
	//	userId, channelCode, merchantCode, payStrategyLis)

	for _, v := range payStrategyLis {
		var totalCount int64

		preNow := now.Unix() - v.Time*60

		rTime := GetUserPaymentStrategyRecordTime(userId, v.Time)
		if rTime != nil {
			if preNow < rTime.Unix() {
				// 前置时间比记录时间小则按记录时间为准
				preNow = rTime.Unix()
			} else {
				// 前置时间比记录时间大则删除记录时间
				DelUserPaymentStrategyRecordTime(userId, v.Time)
			}
		}

		err := global.G_DB.Model(&dos.FcOrderDeposit{}).
			Where("create_time BETWEEN ? AND ?", time.Unix(preNow, 0), now).
			Where("user_id = ?", userId).
			Where("channel_code = ?", channelCode).
			Where("merchant_code = ?", merchantCode).
			Where("status = ?", enmus.Order_STATUS_PENDING_PAY).
			Count(&totalCount).Error
		if err != nil {
			global.G_LOG.Errorf("[CheckUserPaymentStrategy] Find order deposit data failed: userId=%s, channelCode=%s, merchantCode=%s, err=%v",
				userId, channelCode, merchantCode, err.Error())
		}

		if totalCount >= v.Num {
			SetUserPaymentStrategyTriggerFlag(userId, v.Time, 1)
			return false
		}
	}

	return true
}

// GetUserPaymentStrategyTriggerFlagCacheKey - 获取用户支付限单策略触发标识缓存键
// @param {string} userId
// @param {int64} ts
// @returns string
func GetUserPaymentStrategyTriggerFlagCacheKey(userId string, ts int64) string {
	return fmt.Sprintf("userPaymentStrategyTriggerFlagCacheKey:%s:%d", userId, ts)
}

// GetUserPaymentStrategyTriggerFlag - 获取用户支付限单策略触发标识
// @param {string} userId
// @param {int64} ts
// @returns int
func GetUserPaymentStrategyTriggerFlag(userId string, ts int64) int {
	flag, ret := orderDepositCache.Get(GetUserPaymentStrategyTriggerFlagCacheKey(userId, ts))
	if ret {
		if v, ok := flag.(int); ok {
			return v
		}
	}

	return 0
}

// SetUserPaymentStrategyTriggerFlag - 设置用户支付限单策略触发标识
// @param {string} userId
// @param {int64} ts
// @returns
func SetUserPaymentStrategyTriggerFlag(userId string, ts int64, flag int) {
	orderDepositCache.Set(GetUserPaymentStrategyTriggerFlagCacheKey(userId, ts),
		flag,
		cache.NoExpiration)
}

// DelUserPaymentStrategyTriggerFlag - 删除用户支付限单策略触发标识
// @param {string} userId
// @param {int64} ts
// @returns
func DelUserPaymentStrategyTriggerFlag(userId string, ts int64) {
	orderDepositCache.Delete(GetUserPaymentStrategyTriggerFlagCacheKey(userId, ts))
}

// GetUserPaymentStrategyRecordTimeCacheKey - 获取用户支付限单策略记录时间缓存键
// @param {string} userId
// @param {int64} ts
// @returns string
func GetUserPaymentStrategyRecordTimeCacheKey(userId string, ts int64) string {
	return fmt.Sprintf("userPaymentStrategyRecordTimeCacheKey:%s:%d", userId, ts)
}

// GetUserPaymentStrategyRecordTime - 获取用户支付限单策略记录时间
// @param {string} userId
// @param {int64} ts
// @returns *time.Time
func GetUserPaymentStrategyRecordTime(userId string, ts int64) *time.Time {
	flag, ret := orderDepositCache.Get(GetUserPaymentStrategyRecordTimeCacheKey(userId, ts))
	if ret {
		if v, ok := flag.(time.Time); ok {
			return &v
		}
	}

	return nil
}

// SetUserPaymentStrategyRecordTime - 设置用户支付限单策略记录时间
// @param {string} userId
// @param {int64} ts
// @param {time.Time} rt
// @returns
func SetUserPaymentStrategyRecordTime(userId string, ts int64, rt time.Time) {
	orderDepositCache.Set(GetUserPaymentStrategyRecordTimeCacheKey(userId, ts),
		rt,
		cache.NoExpiration)
}

// DelUserPaymentStrategyRecordTime - 删除用户支付限单策略记录时间
// @param {string} userId
// @param {int64} ts
// @returns
func DelUserPaymentStrategyRecordTime(userId string, ts int64) {
	orderDepositCache.Delete(GetUserPaymentStrategyRecordTimeCacheKey(userId, ts))
}

// DoUserPaymentStrategyResetAction - 处理用户支付策略重置（仅在用户充值成功调用）
// @param {string} userId
// @param {string} channelCode
// @param {string} merchantCode
// @returns
func DoUserPaymentStrategyResetAction(userId string, channelCode string, merchantCode string) {
	payStrategyLis := GetFcPayChannelStrategyList(channelCode, merchantCode)
	for _, v := range payStrategyLis {
		if GetUserPaymentStrategyTriggerFlag(userId, v.Time) == 1 {
			// 说明出触发了限单
			SetUserPaymentStrategyRecordTime(userId, v.Time, time.Now())
			DelUserPaymentStrategyTriggerFlag(userId, v.Time)
		}
	}
}

// GetUserFirstRechargeAmount - 获取首充金额
// @param {string} userId
// @returns float64
func GetUserFirstRechargeAmount(userId string) float64 {
	data := dos.FcOrderDeposit{}
	err := global.G_DB.Model(&dos.FcOrderDeposit{}).
		Select("amount").
		Where("user_id = ?", userId).
		First(&data).
		Order("create_time").Error
	if err != nil {
		global.G_LOG.Errorf("[GetUserFirstRechargeAmount] Find user first recharge amount failed: userId=%s, err=%s",
			userId, err.Error())
		return 0
	}

	return data.Amount
}

// GetUserTodayRechargeAmount - 获取今日充值金额
// @param {string} userId
// @returns float64
func GetUserTodayRechargeAmount(userId string) float64 {
	nDate := time.Now().Format(tool.TimeDateLayout)
	sDate := nDate + " 00:00:00"
	eDate := nDate + " 23:59:59"

	var sumAmount float64

	err := global.G_DB.Model(&dos.FcOrderDeposit{}).
		Select("IFNULL(sum(amount), 0) AS sumAmount").
		Where("user_id = ?", userId).
		Where("create_time BETWEEN ? AND ?", sDate, eDate).
		Scan(&sumAmount).Error

	if err != nil {
		global.G_LOG.Errorf("[GetUserTodayRechargeAmount] Find user today recharge amount failed: userId=%s, err=%s",
			userId, err.Error())
		return 0
	}

	return sumAmount
}
