package vo

import "bootpkg/pkg/core/modules/dos"

type GetVenusListResp struct {
	VenueCode   string            `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" ` // 场馆Code
	VenueName   string            `gorm:"venue_name" json:"venue_name" form:"venue_name" uri:"venue_name" `
	ImgIcon     string            `gorm:"img_icon" json:"img_icon" form:"img_icon" uri:"img_icon" `
	ImgBar      string            `gorm:"img_bar" json:"img_bar" form:"img_bar" uri:"img_bar" `
	Describe    string            `gorm:"describe" json:"describe" form:"describe" uri:"describe" `
	GameType    string            `gorm:"game_type" json:"game_type" form:"game_type" uri:"game_type" `
	IsMaintain  bool              `json:"is_maintain"` //维护 true 维护 false 未维护
	GameTypeImg []*dos.FcVenueImg `json:"game_type_img"`
}

type GetVenusBalancesResp struct {
	VenueCode string `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" ` // 场馆Code
	Amount    string `json:"amount"`                                                           //金额
}
