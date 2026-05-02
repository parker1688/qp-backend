package dos

type FcWelfareManage struct {
	BaseDos
	FlowMultiple int    `gorm:"column:flow_multiple" json:"flow_multiple" form:"flow_multiple" uri:"flow_multiple" ` // 打码倍数
	BonusType    int    `gorm:"column:bonus_type" json:"bonus_type" form:"bonus_type" uri:"bonus_type" `             // 交易工单类型
	Title        string `gorm:"column:title" json:"title" form:"title" uri:"title" `
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
}

func (FcWelfareManage) TableName() string {
	return "fc_welfare_manage"
}
