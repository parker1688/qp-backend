package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcSiteNotify struct {
	BaseDos
	Title        string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 公告名称
	TitleImg     string             `gorm:"column:title_img" json:"title_img" form:"title_img" uri:"title_img" `                      // 公告圖片
	Content      string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 公告内容
	Language     string             `gorm:"column:language" json:"language" form:"language" uri:"language" `                          // 语言简码
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序从大到小
	NotifyType   string             `gorm:"column:notify_type" json:"notify_type" form:"notify_type" uri:"notify_type" `              // h5,web,android,ios
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	ClassType    int                `gorm:"column:class_type" json:"class_type" form:"class_type" uri:"class_type" `                  // 平台公告分类类型 1. 公告 2.  赛事  3. 充提 9. 用户通知消息
}

func (FcSiteNotify) TableName() string {
	return "fc_site_notify"
}
