package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcChannelBankImg struct {
	BaseDos
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                         // 1：启动  2 关闭
	ChannelCode  string             `gorm:"column:channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" ` // 渠道code
	PaymentCode  string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" ` // 通道code
	ChannelName  string             `gorm:"column:channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" ` // 渠道code
	PaymentName  string             `gorm:"column:payment_name" json:"payment_name" form:"payment_name" uri:"payment_name" ` // 通道code
	Icon         string             `gorm:"column:icon" json:"icon" form:"icon" uri:"icon" `
	IconPath     string             `gorm:"column:icon_path" json:"icon_path" form:"icon_path" uri:"icon_path" `
	Img          string             `gorm:"column:img" json:"img" form:"img" uri:"img" `
	ImgPath      string             `gorm:"column:img_path" json:"img_path" form:"img_path" uri:"img_path" `
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `
}

func (FcChannelBankImg) TableName() string {
	return "fc_channel_bank_img"
}
