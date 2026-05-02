package vo

import "bootpkg/common/expands/automaticType"

type VenueCreateRequest struct {
	VenueName    string `json:"venue_name"`
	VenueCode    string `json:"venue_code"`
	Status       int    `json:"status"`
	VenueType    int    `json:"venue_type"`
	MerchantCode string `json:"merchant_code"`
	GameTypeImg  []struct {
		GameType string `json:"game_type"`
		ImgIcon  string `json:"img_icon"`
		ImgBar   string `json:"img_bar"`
		LinkIcon string `json:"link_icon"`
		LinkBar  string `json:"link_bar"`
	} `json:"game_type_img"`
}

type VenueUpdateRequest struct {
	Id           string `json:"id"`
	VenueName    string `json:"venue_name"`
	VenueCode    string `json:"venue_code"`
	Status       int    `json:"status"`
	VenueType    int    `json:"venue_type"`
	MerchantCode string `json:"merchant_code"`
	GameTypeImg  []struct {
		GameType string `json:"game_type"`
		ImgIcon  string `json:"img_icon"`
		ImgBar   string `json:"img_bar"`
		LinkIcon string `json:"link_icon"`
		LinkBar  string `json:"link_bar"`
	} `json:"game_type_img"`
}

type MerchantVenueCreateRequest struct {
	VenueId           string             `json:"venue_id"`
	VenueName         string             `json:"venue_name"`
	VenueCode         string             `json:"venue_code"`
	Status            int                `json:"status"`
	VenueType         int                `json:"venue_type"`
	GameType          string             `json:"game_type"`
	MerchantCode      string             `json:"merchant_code"`
	VenueFeeRate      float64            `gorm:"column:venue_fee_rate" json:"venue_fee_rate" form:"venue_fee_rate" uri:"venue_fee_rate" ` // 场馆费率
	InVenueCode       int                `gorm:"column:in_venue_code" json:"in_venue_code" form:"in_venue_code" uri:"in_venue_code" `     // 钱包转入状态 0 正常 1 维护
	OutVenueCode      int                `gorm:"column:out_venue_code" json:"out_venue_code" form:"out_venue_code" uri:"out_venue_code" `
	MaintainStartTime automaticType.Time `gorm:"column:maintain_start_time" json:"maintain_start_time" form:"maintain_start_time" uri:"maintain_start_time" ` // 维护开始时间
	MaintainEndTime   automaticType.Time `gorm:"column:maintain_end_time" json:"maintain_end_time" form:"maintain_end_time" uri:"maintain_end_time" `         // 维护结束时间
	Describe          string             `gorm:"column:describe" json:"describe" form:"describe" uri:"describe" `
	GameTypeImg       []struct {
		GameType string `json:"game_type"`
		ImgIcon  string `json:"img_icon"`
		ImgBar   string `json:"img_bar"`
		LinkIcon string `json:"link_icon"`
		LinkBar  string `json:"link_bar"`
	} `json:"game_type_img"`
}

type MerchantVenueUpdateRequest struct {
	Id                string             `json:"id"`
	VenueId           string             `json:"venue_id"`
	VenueName         string             `json:"venue_name"`
	VenueCode         string             `json:"venue_code"`
	Status            int                `json:"status"`
	VenueType         int                `json:"venue_type"`
	GameType          string             `json:"game_type"`
	MerchantCode      string             `json:"merchant_code"`
	VenueFeeRate      float64            `gorm:"column:venue_fee_rate" json:"venue_fee_rate" form:"venue_fee_rate" uri:"venue_fee_rate" ` // 场馆费率
	InVenueCode       int                `gorm:"column:in_venue_code" json:"in_venue_code" form:"in_venue_code" uri:"in_venue_code" `     // 钱包转入状态 0 正常 1 维护
	OutVenueCode      int                `gorm:"column:out_venue_code" json:"out_venue_code" form:"out_venue_code" uri:"out_venue_code" `
	MaintainStartTime automaticType.Time `gorm:"column:maintain_start_time" json:"maintain_start_time" form:"maintain_start_time" uri:"maintain_start_time" ` // 维护开始时间
	MaintainEndTime   automaticType.Time `gorm:"column:maintain_end_time" json:"maintain_end_time" form:"maintain_end_time" uri:"maintain_end_time" `         // 维护结束时间
	Describe          string             `gorm:"column:describe" json:"describe" form:"describe" uri:"describe" `
	GameTypeImg       []struct {
		GameType     string `json:"game_type"`
		ImgIcon      string `json:"img_icon"`
		ImgBar       string `json:"img_bar"`
		LinkIcon     string `json:"link_icon"`
		LinkBar      string `json:"link_bar"`
		MerchantCode string `json:"merchant_code"`
	} `json:"game_type_img"`
}
