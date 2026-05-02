package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcMerchantLink struct {
	BaseDos
	Link         string             `gorm:"column:link" json:"link" form:"link" uri:"link" `                                          // 网站
	Alias        string             `gorm:"column:alias" json:"alias" form:"alias" uri:"alias" `                                      // 别名
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	MerchantName string             `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `      // 商户名称
	LogoImg      string             `gorm:"column:logo_img" json:"logo_img" form:"logo_img" uri:"logo_img"`                           // logo图片
	BannerImg    string             `gorm:"column:banner_img" json:"banner_img" form:"banner_img" uri:"banner_img"`                   // banner图片
	HomeEntryIcons string             `gorm:"column:home_entry_icons;type:text" json:"home_entry_icons" form:"home_entry_icons" uri:"home_entry_icons"`                   // 首页入口图标配置 JSON (name/venue_type/image/imageOK)
}

func (FcMerchantLink) TableName() string {
	return "fc_merchant_link"
}
