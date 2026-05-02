package vo

type CilentSystemSettingsResp struct {
	DictsTag string `gorm:"column:dicts_tag" json:"dicts_tag" form:"dicts_tag" uri:"dicts_tag" `         // 标识
	AppLink  string `gorm:"column:dicts_value" json:"dicts_value" form:"dicts_value" uri:"dicts_value" ` // 地址
}
