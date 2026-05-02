package dos

type FcCurrency struct {
	BaseDos
	Name string  `gorm:"name" json:"name" form:"name" uri:"name" ` // 货币名称
	Code string  `gorm:"code" json:"code" form:"code" uri:"code" ` // 货币code
	Rate float64 `gorm:"rate" json:"rate" form:"rate" uri:"rate" ` // 汇率
	Icon string  `gorm:"icon" json:"icon" form:"icon" uri:"icon" ` // 图标
}

func (FcCurrency) TableName() string {
	return "fc_currency"
}
