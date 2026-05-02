package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenueUser struct {
	BaseDos
	VenueId      string             `gorm:"venue_id" json:"venue_id" form:"venue_id" uri:"venue_id" ` // 场馆ID
	VenueCode    string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `
	UserId       string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 用户账号
	Account      string             `gorm:"account" json:"account" form:"account" uri:"account" `                              //三方账号
	Password     string             `gorm:"password" json:"password" form:"password" uri:"password" `                          //三方密码
	Currency     string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                          // 币种
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	VenueLine    int                `gorm:"venue_line" json:"venue_line" form:"venue_line" uri:"venue_line" `                  // 商户code
}

func (FcVenueUser) TableName() string {
	return "fc_venue_user"
}
