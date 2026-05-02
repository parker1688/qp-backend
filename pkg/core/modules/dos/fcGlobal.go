package dos

type FcGlobal struct {
	Key   string `gorm:"column:key" json:"key" form:"key" uri:"key"`
	Value string `gorm:"value" json:"value" form:"value" uri:"value"`
}

func (FcGlobal) TableName() string {
	return "fc_global"
}
