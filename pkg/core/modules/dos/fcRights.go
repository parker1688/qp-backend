package dos

type FcRights struct {
	BaseDos
	RightsName string `gorm:"column:rights_name" json:"rights_name" form:"rights_name" uri:"rights_name" ` // 权益名字
	RightsDesc string `gorm:"column:rights_desc" json:"rights_desc" form:"rights_desc" uri:"rights_desc" ` // 权益说明
	Vkey       string `gorm:"column:vkey" json:"vkey" form:"vkey" uri:"vkey" `                             // 权益英文代码
	Value      string `gorm:"column:value" json:"value" form:"value" uri:"value" `                         // 值
}

func (FcRights) TableName() string {
	return "fc_rights"
}
