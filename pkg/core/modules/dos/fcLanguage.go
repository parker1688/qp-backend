package dos

type FcLanguage struct {
	BaseDos
	Language string `gorm:"language" json:"language" form:"language" uri:"language" `
	Code     string `gorm:"code" json:"code" form:"code" uri:"code" `
}

func (FcLanguage) TableName() string {
	return "fc_language"
}
