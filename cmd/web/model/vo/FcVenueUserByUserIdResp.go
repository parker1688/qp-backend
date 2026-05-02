package vo

import "bootpkg/common/expands/automaticType"

type FcVenueUserByUserIdResp struct {
	Id         string             `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	VenueId    string             `gorm:"venue_id" json:"venue_id" form:"venue_id" uri:"venue_id" ` // 场馆ID
	VenueCode  string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `
	UserId     string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName   string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 用户账号
	Account    string             `gorm:"account" json:"account" form:"account" uri:"account" `                              //三方账号
	Currency   string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                          // 币种
	CreateTime automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	Amount     string             `json:"amount"`                                                                            //场馆余额
}
