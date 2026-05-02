package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcBonus struct {
	BaseDos
	UserId       string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" ` // 用户名
	Amount       float64            `gorm:"amount" json:"amount" form:"amount" uri:"amount" `             // 红利金额
	Currency     string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `     // 货币
	BonusType    int                `gorm:"bonus_type" json:"bonus_type" form:"bonus_type" uri:"bonus_type" `
	Remark       string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcBonus) TableName() string {
	return "fc_bonus"
}
