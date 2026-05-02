package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcActivityConfig struct {
	BaseDos
	ConfigKey    string             `gorm:"column:config_key" json:"config_key" form:"config_key" uri:"config_key" `
	ConfigValue  string             `gorm:"column:config_value" json:"config_value" form:"config_value" uri:"config_value" `
	Desc         string             `gorm:"column:desc" json:"desc" form:"desc" uri:"desc" `
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcActivityConfig) TableName() string {
	return "fc_activity_config"
}
