package dos

type FcPgdz struct {
	BaseDos
	Name string `gorm:"column:name" json:"name" form:"name" uri:"name" `
}

func (FcPgdz) TableName() string {
	return "fc_pgdz"
}
