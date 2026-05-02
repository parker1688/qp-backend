package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcSiteLink struct {
	BaseDos
	AppKey        string             `gorm:"column:app_key" json:"app_key" form:"app_key" uri:"app_key" `                              // 标识
	AppLink       string             `gorm:"column:app_link" json:"app_link" form:"app_link" uri:"app_link" `                          // 链接
	Content       string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 描述
	ContentDetail string             `gorm:"column:content_detail" json:"content_detail" form:"content_detail" uri:"content_detail" `  // 描述详情
	CreateTime    automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime    automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode  string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	ImgLink       string             `gorm:"column:img_link" json:"img_link" form:"img_link" uri:"img_link" `                          // 图片地址
	DownloadLink  string             `gorm:"column:download_link" json:"download_link" form:"download_link" uri:"download_link" `      // 下载地址
	//Domain       string             `gorm:"column:domain" json:"domain" form:"domain" uri:"domain" `                                  // 绑定站点
}

func (FcSiteLink) TableName() string {
	return "fc_site_link"
}
