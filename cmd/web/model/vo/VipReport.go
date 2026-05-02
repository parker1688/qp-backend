package vo

type VipReportReq struct {
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	MerchantName string `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `
}

type VipReportResp struct {
	VipData []VipReportData `json:"vip_data"`
}

type VipReportData struct {
	MerchantName string `json:"merchant_name"`
	UserCount    int    `json:"user_count"`
	VipCount     int    `json:"vip_count"`
	VipList      []int  `json:"list"`
}
