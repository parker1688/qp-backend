package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenueImg struct {
	BaseDos
	VenueName    string             `gorm:"column:venue_name" json:"venue_name" form:"venue_name" uri:"venue_name" `                  // 场馆名字
	VenueCode    string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                  // 场馆code
	GameType     string             `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                      // 1：转账钱包   2：单一钱包
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	ImgIcon      string             `gorm:"column:img_icon" json:"img_icon" form:"img_icon" uri:"img_icon" `                          // 图片
	ImgBar       string             `gorm:"column:img_bar" json:"img_bar" form:"img_bar" uri:"img_bar" `                              // 导航栏图片
	LinkIcon     string             `gorm:"column:link_icon" json:"link_icon" form:"link_icon" uri:"link_icon" `                      // 图片地址
	LinkBar      string             `gorm:"column:link_bar" json:"link_bar" form:"link_bar" uri:"link_bar" `                          // 导航栏地址
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
}

func (FcVenueImg) TableName() string {
	return "fc_venue_img"
}
