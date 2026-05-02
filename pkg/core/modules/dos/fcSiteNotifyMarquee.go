package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcSiteNotifyMarquee struct {
	BaseDos
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 公告名称
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 0：正常 1：下架
	Frequency    int                `gorm:"column:frequency" json:"frequency" form:"frequency" uri:"frequency" `                      // 频率
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 公告内容
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcSiteNotifyMarquee) TableName() string {
	return "fc_site_notify_marquee"
}
