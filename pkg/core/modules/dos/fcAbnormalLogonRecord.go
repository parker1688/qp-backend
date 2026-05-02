package dos

type FcAbnormalLogonRecord struct {
	Key   string `gorm:"column:key" json:"key" form:"key" uri:"key"`
	Value string `gorm:"value" json:"value" form:"value" uri:"value"`
}

func (FcAbnormalLogonRecord) TableName() string {
	return "fc_abnormal_logon_record"
}
