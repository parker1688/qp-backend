package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcSiteBanner struct {
	BaseDos
	BannerLink      string             `gorm:"column:banner_link" json:"banner_link" form:"banner_link" uri:"banner_link" `                         // banner图片地址
	BannerOtherType int                `gorm:"column:banner_other_type" json:"banner_other_type" form:"banner_other_type" uri:"banner_other_type" ` // 1 首页, 2 存款, 3 提款
	BannerHref      string             `gorm:"column:banner_href" json:"banner_href" form:"banner_href" uri:"banner_href" `                         // banner跳转地址
	Language        string             `gorm:"column:language" json:"language" form:"language" uri:"language" `                                     // 语言简码
	Sort            int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                                     // 排序从大到小
	BannerType      string             `gorm:"column:banner_type" json:"banner_type" form:"banner_type" uri:"banner_type" `                         // h5,web,android,ios
	CreateTime      automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `            // 创建时间
	CreateBy        string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                 // 创建人
	UpdateTime      automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `            // 修改时间
	UpdateBy        string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                 // 修改人
	MerchantCode    string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户code
}

func (FcSiteBanner) TableName() string {
	return "fc_site_banner"
}
