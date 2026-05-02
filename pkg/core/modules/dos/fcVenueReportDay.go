package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenueReportDay struct {
	BaseDos
	VenueName    string             `gorm:"column:venue_name" json:"venue_name" form:"venue_name" uri:"venue_name" `                  // 场馆名字
	VenueCode    string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                  // 场馆code
	ReportDate   automaticType.Time `gorm:"column:report_date" json:"report_date" form:"report_date" uri:"report_date" `              // 报表日期
	BetAmount    float64            `gorm:"column:bet_amount" json:"bet_amount" form:"bet_amount" uri:"bet_amount" `                  // 总流水
	BetCount     int                `gorm:"column:bet_count" json:"bet_count" form:"bet_count" uri:"bet_count" `                      // 投注笔数
	NetAmount    float64            `gorm:"column:net_amount" json:"net_amount" form:"net_amount" uri:"net_amount" `                  // 平台输赢
	ReturnRate   float64            `gorm:"column:return_rate" json:"return_rate" form:"return_rate" uri:"return_rate" `              // 返奖率
	VenueFeeRate float64            `gorm:"column:venue_fee_rate" json:"venue_fee_rate" form:"venue_fee_rate" uri:"venue_fee_rate" `  // 三方费率
	VenueFee     float64            `gorm:"column:venue_fee" json:"venue_fee" form:"venue_fee" uri:"venue_fee" `                      // 场馆费
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcVenueReportDay) TableName() string {
	return "fc_venue_report_day"
}
