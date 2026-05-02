package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserCodeBetConfig struct {
	BaseDos
	FlagTitle    string             `gorm:"column:flag_title" json:"flag_title" form:"flag_title" uri:"flag_title" `                  // 打码主题
	Flag         string             `gorm:"column:flag" json:"flag" form:"flag" uri:"flag" `                                          // 打码标识（英文）
	Multiple     int                `gorm:"column:multiple" json:"multiple" form:"multiple" uri:"multiple" `                          // 打码倍数
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcUserCodeBetConfig) TableName() string {
	return "fc_user_code_bet_config"
}
