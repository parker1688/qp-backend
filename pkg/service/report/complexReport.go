package report

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ComplexReport struct {
	StartAt      automaticType.Time
	EndAt        automaticType.Time
	UserId       string
	MerchantCode string
	MerchantName string
}

func applyComplexReportTimeMerchantFilters(query *gorm.DB, r ComplexReport, timeColumn string) *gorm.DB {
	if !r.StartAt.Timer().IsZero() {
		query = query.Where(timeColumn+">=?", r.StartAt)
	}

	if !r.EndAt.Timer().IsZero() {
		query = query.Where(timeColumn+"<?", r.EndAt)
	}

	if len(r.MerchantCode) > 0 {
		query = query.Where("merchant_code =?", r.MerchantCode)
	}

	return query
}

func buildComplexReportBetRecordQuery(query *gorm.DB, r ComplexReport) *gorm.DB {
	query = applyComplexReportTimeMerchantFilters(query, r, "bet_time")

	if len(r.UserId) > 0 {
		query = query.Where("user_id=?", r.UserId)
	}

	return query
}

// 累计输赢
func (r ComplexReport) BetWin() float64 {
	var betWin float64
	query := buildComplexReportBetRecordQuery(global.G_DB.Model(&dos.FcBetRecord{}), r)
	query.Select("sum(net_amount) as betWin").Scan(&betWin)
	return betWin
}

// 有效投注
func (r ComplexReport) BetAmount() float64 {
	var betAmount float64
	query := buildComplexReportBetRecordQuery(global.G_DB.Model(&dos.FcBetRecord{}), r)
	query.Select("sum(valid_betamount) as betAmount").Scan(&betAmount)
	return betAmount
}

// 游戏杀率
func (r ComplexReport) GameKillRate(report *dos.FcComplexReport) float64 {
	if report.BetAmount == 0 {
		return 0
	}
	return decimal.NewFromFloat(report.BetWin).Div(decimal.NewFromFloat(report.BetAmount)).Truncate(4).InexactFloat64()
}

// 返水
func (r ComplexReport) RebateAmount() float64 {
	var rebateAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcUserRebateRecords{}), r, "create_time")
	query = query.Where("status =?", 1)
	query.Select("sum(bonus_amount) as rebateAmount").Scan(&rebateAmount)
	return rebateAmount
}

// 投注倍数 - 有效投注/充值汇总，分母为0时，最终值做0处理
func (r ComplexReport) BetMultiple(report *dos.FcComplexReport) float64 {
	if report.TotalDepositAmount == 0 {
		return 0
	}

	return decimal.NewFromFloat(report.BetAmount).Div(decimal.NewFromFloat(report.TotalDepositAmount)).Truncate(4).InexactFloat64()
}

// 注册人数
func (r ComplexReport) RegisterNum() int {
	var registerNum int
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcUserMaterial{}), r, "create_time")

	query.Select("count(1) as registerNum").Scan(&registerNum)
	return registerNum
}

// 首充人数
func (r ComplexReport) FirstDepositNum() int {
	var firstDepositNum int
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Table(dos.FcOrderDeposit{}.TableName()+" as a"), r, "create_time")
	query = query.Where("a.status =?", 3)
	query = query.Where("not EXISTS(select * from  fc_order_deposit as b where create_time <? AND a.user_id=b.user_id and b.status =?)", r.StartAt, 3)
	query.Select("count(DISTINCT(user_id)) as registerNum").Scan(&firstDepositNum)
	return firstDepositNum
}

// 注册充值率 - 首充人数/注册人数，分母为0时，最终值做0处理
func (r ComplexReport) RegisterDepositRate(report *dos.FcComplexReport) float64 {
	if report.RegisterNum == 0 {
		return 0
	}

	return decimal.NewFromFloat(float64(report.FirstDepositNum)).Div(decimal.NewFromFloat(float64(report.RegisterNum))).Truncate(4).InexactFloat64()
}

// 充值人数
func (r ComplexReport) DepositNum() int {
	var depositNum int
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("status =?", 3)
	query.Select("count(DISTINCT(user_id)) as registerNum").Scan(&depositNum)
	return depositNum
}

// 充值笔数
func (r ComplexReport) DepositCount() int {
	var depositNum int
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("status =?", 3)
	query.Select("count(1) as depositNum").Scan(&depositNum)
	return depositNum
}

// 新户首充金额
func (r ComplexReport) FirstDepositAmount() float64 {
	var firstDepositAmount float64
	var userIds []string

	query := global.G_DB.Table(dos.FcOrderDeposit{}.TableName() + " as a")
	if !r.StartAt.Timer().IsZero() {
		query = query.Where("create_time>=?", r.StartAt)
	}

	if !r.EndAt.Timer().IsZero() {
		query = query.Where("create_time<?", r.EndAt)
	}

	if len(r.MerchantCode) > 0 {
		query = query.Where("merchant_code =?", r.MerchantCode)
	}
	query = query.Where("a.status =?", 3)
	query = query.Where("not EXISTS(select * from  fc_order_deposit as b where create_time <? AND a.user_id=b.user_id and b.status =?)", r.StartAt, 3)
	query.Select("DISTINCT(user_id) as registerNum").Scan(&userIds)

	for _, userId := range userIds {
		var userFirstDepositAmount float64
		depositQuery := global.G_DB.Model(&dos.FcOrderDeposit{}).Where("user_id=? and  status =? ", userId, 3)
		if !r.StartAt.Timer().IsZero() {
			depositQuery = depositQuery.Where("create_time>=?", r.StartAt)
		}

		if !r.EndAt.Timer().IsZero() {
			depositQuery = depositQuery.Where("create_time<?", r.EndAt)
		}

		depositQuery = depositQuery.Order("create_time asc").Limit(1)
		depositQuery.Select("amount").Scan(&userFirstDepositAmount)
		if userFirstDepositAmount > 0 {
			firstDepositAmount = decimal.NewFromFloat(firstDepositAmount).Add(decimal.NewFromFloat(userFirstDepositAmount)).Truncate(2).InexactFloat64()
		}
	}
	return firstDepositAmount
}

// 新户充值金额
func (r ComplexReport) NewUserDepositAmount() float64 {
	var newUserDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Table(dos.FcOrderDeposit{}.TableName()+" as a"), r, "create_time")
	query = query.Where("a.status =?", 3)
	query = query.Where("not EXISTS(select * from  fc_order_deposit as b where create_time <? AND a.user_id=b.user_id and b.status =?)", r.StartAt, 3)
	query.Select("sum(amount) as newUserDepositAmount").Scan(&newUserDepositAmount)
	return newUserDepositAmount
}

func (r ComplexReport) NewUserFirstDepositAmount() float64 {
	var amount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Table(dos.FcOrderDeposit{}.TableName()+" as a"), r, "create_time")
	query = query.Where("a.status =?", 3)
	query = query.Where("not EXISTS(select * from  fc_order_deposit as b where create_time <? AND a.user_id=b.user_id and b.status =?)", r.StartAt, 3)
	query = query.Order("create_time asc").Limit(1)
	query.Select("amount").Scan(&amount)
	return amount
}

// 新户充值笔数
func (r ComplexReport) NewUserDepositCount() int {
	var newUserDepositCount int
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Table(dos.FcOrderDeposit{}.TableName()+" as a"), r, "create_time")
	query = query.Where("a.status =?", 3)
	query = query.Where("not EXISTS(select * from  fc_order_deposit as b where create_time <? AND a.user_id=b.user_id and b.status =?)", r.StartAt, 3)
	query.Select("count(1) as newUserDepositCount").Scan(&newUserDepositCount)
	return newUserDepositCount
}

// 登录人数
func (r ComplexReport) LoginNum() int {
	var loginNum int
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcLoginLog{}), r, "create_time")

	query.Select("count(DISTINCT(user_name)) as loginNum").Scan(&loginNum)
	return loginNum
}

// 提现人数
func (r ComplexReport) WithdrawNum() int {
	var withdrawNum int

	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}), r, "create_time")
	query = query.Where("status =?", 3)
	query.Select("count(DISTINCT(user_id)) as withdrawNum").Scan(&withdrawNum)
	return withdrawNum
}

// 投注人数
func (r ComplexReport) BetNum() int {
	var betNum int

	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcBetRecord{}), r, "bet_time")

	query.Select("count(DISTINCT(user_id)) as betNum").Scan(&betNum)
	return betNum
}

func (r ComplexReport) PromotionAmount() float64 {
	var promotionAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderPromotion{}), r, "create_time")
	query = query.Where("status =?", 3)
	query.Select("sum(amount) as promotionAmount").Scan(&promotionAmount)

	promotionAmount += modules.GetDailyBonusPromotionTotalByDate(r.MerchantCode, r.StartAt, r.EndAt)
	promotionAmount += modules.GetUserTaskPromotionTotalByDate(r.MerchantCode, r.StartAt, r.EndAt)
	promotionAmount += modules.GetUserActivityPromotionTotalByDate(r.MerchantCode, r.StartAt, r.EndAt)
	return promotionAmount
}

func (r ComplexReport) AlipayDepositAmount() float64 {
	var alipayDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("channel_code =? and status=?", "alipay", 3)

	query.Select("sum(amount) as alipayDepositAmount").Scan(&alipayDepositAmount)
	return alipayDepositAmount
}

// 微信充值
func (r ComplexReport) WxDepositAmount() float64 {
	var wxDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("channel_code =? and status=?", "wx", 3)

	query.Select("sum(amount) as wxDepositAmount").Scan(&wxDepositAmount)
	return wxDepositAmount
}

// 银行卡充值
func (r ComplexReport) BankDepositAmount() float64 {
	var bankDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("channel_code =? and status=?", "bank", 3)

	query.Select("sum(amount) as bankDepositAmount").Scan(&bankDepositAmount)
	return bankDepositAmount
}

// 钱包充值
func (r ComplexReport) WalletDepositAmount() float64 {
	var walletDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("channel_code =? and status=?", "wallet", 3)

	query.Select("sum(amount) as bankDepositAmount").Scan(&walletDepositAmount)
	return walletDepositAmount
}

// 数字人民币充值
func (r ComplexReport) NumCnyDepositAmount() float64 {
	var numCnyDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("channel_code =? and status=?", "number_cny", 3)

	query.Select("sum(amount) as numCnyDepositAmount").Scan(&numCnyDepositAmount)
	return numCnyDepositAmount
}

// USDT充值
func (r ComplexReport) UsdtDepositAmount() float64 {
	var usdtDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("channel_code =? and status=?", "USDT-CR20", 3)

	query.Select("sum(amount) as numCnyDepositAmount").Scan(&usdtDepositAmount)
	return usdtDepositAmount
}

// 人工存款
func (r ComplexReport) AdminDepositAmount() float64 {
	var adminDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where("order_type =? and status=?", 20, 3)

	query.Select("sum(amount) as adminDepositAmount").Scan(&adminDepositAmount)
	return adminDepositAmount
}

// 充值汇总
func (r ComplexReport) TotalDepositAmount() float64 {
	var totalDepositAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderDeposit{}), r, "create_time")
	query = query.Where(" status=?", 3)

	query.Select("sum(amount) as totalDepositAmount").Scan(&totalDepositAmount)
	return totalDepositAmount
}

// 支付宝提现
func (r ComplexReport) AlipayWithdrawAmount() float64 {
	var alipayWithdrawAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}), r, "create_time")
	query = query.Where(" order_type=? AND status=?", 4, 3)
	query.Select("sum(amount) as alipayWithdrawAmount").Scan(&alipayWithdrawAmount)
	return alipayWithdrawAmount
}

// 银行卡提现
func (r ComplexReport) BankWithdrawAmount() float64 {
	var bankWithdrawAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}), r, "create_time")
	query = query.Where(" order_type=? AND status=?", 1, 3)
	query.Select("sum(amount) as bankWithdrawAmount").Scan(&bankWithdrawAmount)
	return bankWithdrawAmount
}

// 钱包提现
func (r ComplexReport) WalletWithdrawAmount() float64 {
	var walletWithdrawAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}), r, "create_time")
	query = query.Where(" order_type=? AND status=?", 3, 3)
	query.Select("sum(amount) as walletWithdrawAmount").Scan(&walletWithdrawAmount)
	return walletWithdrawAmount
}

func (r ComplexReport) UsdtWithdrawAmount() float64 {
	var usdtWithdrawAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}), r, "create_time")
	query = query.Where("order_type=? AND status=? AND third_code=?", 3, 3, "USDT")

	query.Select("sum(amount) as usdtWithdrawAmount").Scan(&usdtWithdrawAmount)
	return usdtWithdrawAmount
}

// 提现汇总
func (r ComplexReport) TotalWithdrawAmount() float64 {
	var totalWithdrawAmount float64
	query := applyComplexReportTimeMerchantFilters(global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}), r, "create_time")
	query = query.Where("status=?", 3)
	query.Select("sum(amount) as totalWithdrawAmount").Scan(&totalWithdrawAmount)
	return totalWithdrawAmount
}

// 充提差
func (r ComplexReport) DepositWithdrawSubAmount(report *dos.FcComplexReport) float64 {
	return decimal.NewFromFloat(report.TotalDepositAmount).Sub(decimal.NewFromFloat(report.TotalWithdrawAmount)).Truncate(2).InexactFloat64()
}

// 财务杀率
func (r ComplexReport) KillRate(report *dos.FcComplexReport) float64 {
	if report.TotalDepositAmount == 0 {
		return 0
	}
	return decimal.NewFromFloat(report.DepositWithdrawSubAmount).Div(decimal.NewFromFloat(report.TotalDepositAmount)).Truncate(4).InexactFloat64()
}
