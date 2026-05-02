package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcMerchantProfitReportDay struct {
	BaseDos
	Profit          float64            `gorm:"column:profit" json:"profit" form:"profit" uri:"profit" `                                         // 利润
	OnlineRecharge  float64            `gorm:"column:online_recharge" json:"online_recharge" form:"online_recharge" uri:"online_recharge" `     // 在线充值
	OfflineRecharge float64            `gorm:"column:offline_recharge" json:"offline_recharge" form:"offline_recharge" uri:"offline_recharge" ` // 线下充值
	OnlineWithdraw  float64            `gorm:"column:online_withdraw" json:"online_withdraw" form:"online_withdraw" uri:"online_withdraw" `     // 代付提款
	OfflineWithdraw float64            `gorm:"column:offline_withdraw" json:"offline_withdraw" form:"offline_withdraw" uri:"offline_withdraw" ` // 线下提款
	VenueCost       float64            `gorm:"column:venue_cost" json:"venue_cost" form:"venue_cost" uri:"venue_cost" `                         // 场馆成本
	RechargeCost    float64            `gorm:"column:recharge_cost" json:"recharge_cost" form:"recharge_cost" uri:"recharge_cost" `             // 充值成本
	ReportDate      automaticType.Time `gorm:"column:report_date" json:"report_date" form:"report_date" uri:"report_date" `                     // 统计日期
	MerchantCode    string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `             // 商户code
	CreateTime      automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `        // 创建时间
	CreateBy        string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                             // 创建人
	UpdateTime      automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `        // 修改时间
	UpdateBy        string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                             // 修改人
}

func (FcMerchantProfitReportDay) TableName() string {
	return "fc_merchant_profit_report_day"
}
