package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcMerchantVenue struct {
	BaseDos
	MerchantCode      string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                         // 商户code
	VenueId           string             `gorm:"column:venue_id" json:"venue_id" form:"venue_id" uri:"venue_id" `                                             // 场馆ID
	VenueCode         string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                                     // 场馆Code
	Status            int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                                     // 1:正常 2:下线 3:维护
	MaintainStartTime automaticType.Time `gorm:"column:maintain_start_time" json:"maintain_start_time" form:"maintain_start_time" uri:"maintain_start_time" ` // 维护开始时间
	MaintainEndTime   automaticType.Time `gorm:"column:maintain_end_time" json:"maintain_end_time" form:"maintain_end_time" uri:"maintain_end_time" `         // 维护结束时间
	ConfigId          string             `gorm:"column:config_id" json:"config_id" form:"config_id" uri:"config_id" `                                         // 三方场馆配置json格式
	VenueName         string             `gorm:"column:venue_name" json:"venue_name" form:"venue_name" uri:"venue_name" `                                     // 场馆名字
	ConfigAlias       string             `gorm:"column:config_alias" json:"config_alias" form:"config_alias" uri:"config_alias" `                             // 线路别名
	Currency          string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                                             // 货币：多个用英文逗号（,）隔开
	CreateTime        automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                    // 创建时间
	CreateBy          string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                         // 创建人
	UpdateTime        automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                    // 修改时间
	UpdateBy          string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                         // 修改人
	VenueFeeRate      float64            `gorm:"column:venue_fee_rate" json:"venue_fee_rate" form:"venue_fee_rate" uri:"venue_fee_rate" `                     // 场馆费率
	InVenueCode       int                `gorm:"column:in_venue_code" json:"in_venue_code" form:"in_venue_code" uri:"in_venue_code" `                         // 钱包转入状态 0 正常 1 维护
	OutVenueCode      int                `gorm:"column:out_venue_code" json:"out_venue_code" form:"out_venue_code" uri:"out_venue_code" `                     // 钱包转出状态 0 正常 1 维护
	GameType          string             `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                                         // 场馆类型    live：真人 sport：体育 slot：老虎机
	ImgIcon           string             `gorm:"column:img_icon" json:"img_icon" form:"img_icon" uri:"img_icon" `                                             // 图片
	ImgBar            string             `gorm:"column:img_bar" json:"img_bar" form:"img_bar" uri:"img_bar" `                                                 // 图片
	Describe          string             `gorm:"column:describe" json:"describe" form:"describe" uri:"describe" `                                             // 描述
}

func (FcMerchantVenue) TableName() string {
	return "fc_merchant_venue"
}
