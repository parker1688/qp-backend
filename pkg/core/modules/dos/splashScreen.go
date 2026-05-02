package dos

import (
	"bootpkg/common/expands/automaticType"
)

type SplashScreen struct {
	BaseDos
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	LogoImg      string             `gorm:"column:logo_img" json:"logo_img" form:"logo_img" uri:"logo_img" `                          // logo图片
	BannerImg    string             `gorm:"column:banner_img" json:"banner_img" form:"banner_img" uri:"banner_img" `                  // banner图片
	ScreenImg    string             `gorm:"column:screen_img" json:"screen_img" form:"screen_img" uri:"screen_img" `                  // 屏幕图片
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
}

func (SplashScreen) TableName() string {
	return "splash_screen"
}

type SplashScreenEx struct {
	SplashScreen
	Merchant FcMerchant `json:"merchant" gorm:"foreignkey:MerchantCode;references:MerchantCode"`
}

type SplashScreenResp struct {
	SplashScreen
	MerchantName string `json:"merchant_name"`
}
