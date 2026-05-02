package vo

import "bootpkg/common/expands/automaticType"

type DashboardPicReq struct {
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	MerchantName string             `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `
	StartTime    automaticType.Time `gorm:"column:start_time" json:"start_time" form:"start_time" uri:"start_time" `
	EndTime      automaticType.Time `gorm:"column:end_time" json:"end_time" form:"end_time" uri:"end_time" `
	PreStartTime automaticType.Time `gorm:"column:pre_start_time" json:"pre_start_time" form:"pre_start_time" uri:"pre_start_time" `
	PreEndTime   automaticType.Time `gorm:"column:pre_end_time" json:"pre_end_time" form:"pre_end_time" uri:"pre_end_time" `
}

type DashboadPicResp struct {
	AddUserNum             int     `json:"add_user_num"`
	PreAddUserNum          int     `json:"pre_add_user_num"`
	LoginUserCount         int     `json:"login_user_count"`
	PreLoginUserCount      int     `json:"pre_login_user_count"`
	RechargeAmount         float64 `json:"recharge_amount"`
	PreRechargeAmount      float64 `json:"pre_recharge_amount"`
	WithdrawAmount         float64 `json:"withdraw_amount"`
	PreWithdrawAmount      float64 `json:"pre_withdraw_amount"`
	FirstRechargeAmount    float64 `json:"first_recharge_amount"`
	PreFirstRechargeAmount float64 `json:"pre_first_recharge_amount"`
	NetAmount              float64 `json:"net_amount"`
	PreNetAmount           float64 `json:"pre_net_amount"`
}

type DashboardHourReq struct {
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	MerchantName string `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `
}

type DashboadHourResp struct {
	RegisterNum []int `json:"register_num"`
}

type DashboardUserReq struct {
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	MerchantName string `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `
}

type DashboadUserResp struct {
	Num1     int `json:"num1"`
	Num2     int `json:"num2"`
	Num3     int `json:"num3"`
	Num4     int `json:"num4"`
	TotalNum int `json:"total_num"`
}
