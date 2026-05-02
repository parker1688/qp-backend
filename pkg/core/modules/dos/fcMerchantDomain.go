package dos

type FcMerchantDomain struct {
	BaseDos
	MerchantCode string `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
	Domain       string `gorm:"domain" json:"domain" form:"domain" uri:"domain" `                             // 域名
	Status       int    `gorm:"status" json:"status" form:"status" uri:"status" `                             // 1:正常   2:禁止
}

func (FcMerchantDomain) TableName() string {
	return "fc_merchant_domain"
}
