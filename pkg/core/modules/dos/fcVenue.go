package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenue struct {
	BaseDos
	VenueName  string             `gorm:"column:venue_name" json:"venue_name" form:"venue_name" uri:"venue_name" `                  // 场馆名字
	VenueCode  string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                  // 场馆code
	Language   string             `gorm:"column:language" json:"language" form:"language" uri:"language" `                          // 语言：多个用英文逗号（,）隔开
	Status     int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 状态  1:上线  2:下线
	VenueType  int                `gorm:"column:venue_type" json:"venue_type" form:"venue_type" uri:"venue_type" `                  // 1：转账钱包   2：单一钱包
	Currency   string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                          // 货币：多个用英文逗号（,）隔开
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcVenue) TableName() string {
	return "fc_venue"
}
