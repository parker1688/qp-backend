package vo

type AgentSaveReq struct {
	Status       int    `gorm:"column:status" json:"status" form:"status" uri:"status" ` // 读取状态 1正常 2 停用
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	MerchantName string `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" `
	Num          int    `gorm:"column:num" json:"num" form:"num" uri:"num" `
}
