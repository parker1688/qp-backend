package dos

type FcVipRights struct {
	BaseDos
	VipId        string `gorm:"column:vip_id;primary_key" json:"vip_id" form:"vip_id" uri:"vip_id" `
	Level        int    `gorm:"column:level" json:"level" form:"level" uri:"level" `             // 1 ~ 10
	VipName      string `gorm:"column:vip_name" json:"vip_name" form:"vip_name" uri:"vip_name" ` // vip名称 VIP1~VIP10
	RightsId     int    `gorm:"column:rights_id;primary_key" json:"rights_id" form:"rights_id" uri:"rights_id" `
	RightsName   string `gorm:"column:rights_name" json:"rights_name" form:"rights_name" uri:"rights_name" `         // 权益名称
	RightsKey    string `gorm:"column:rights_key" json:"rights_key" form:"rights_key" uri:"rights_key" `             // 权益key
	RightsValue  string `gorm:"column:rights_value" json:"rights_value" form:"rights_value" uri:"rights_value" `     // 权益value
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
}

func (FcVipRights) TableName() string {
	return "fc_vip_rights"
}
