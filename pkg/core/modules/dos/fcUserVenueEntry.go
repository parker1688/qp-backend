package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserVenueEntry struct {
	UserId    string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	VenueCode string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `
	CreatedAt automaticType.Time `gorm:"created_at;default:null" json:"created_at" form:"created_at" uri:"created_at" ` // 创建时间
}

func (FcUserVenueEntry) TableName() string {
	return "fc_user_venue_entry"
}
