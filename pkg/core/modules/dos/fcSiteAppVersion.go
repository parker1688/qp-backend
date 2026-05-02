package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcSiteAppVersion struct {
	BaseDos
	AppKey       string             `gorm:"column:app_key" json:"app_key" form:"app_key" uri:"app_key" `                              // APP标识
	AppVersion   string             `gorm:"column:app_version" json:"app_version" form:"app_version" uri:"app_version" `              // APP标识
	AppLink      string             `gorm:"column:app_link" json:"app_link" form:"app_link" uri:"app_link" `                          // app
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 描述更新内容
	Forcibly     int                `gorm:"column:forcibly" json:"forcibly" form:"forcibly" uri:"forcibly" `                          // 1 非强制更新 1 强制更新
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcSiteAppVersion) TableName() string {
	return "fc_site_app_version"
}
