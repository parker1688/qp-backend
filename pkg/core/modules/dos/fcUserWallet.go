package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserWallet struct {
	BaseDos
	UserId        string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName      string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Currency      string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                          // 货币类型
	TotalAmount   float64            `gorm:"total_amount" json:"total_amount" form:"total_amount" uri:"total_amount" `          // 总金额
	AvaAmount     float64            `gorm:"ava_amount" json:"ava_amount" form:"ava_amount" uri:"ava_amount" `                  // 可用金额
	FronzenAmount float64            `gorm:"fronzen_amount" json:"fronzen_amount" form:"fronzen_amount" uri:"fronzen_amount" `  // 冻结金额
	IsLock        int                `gorm:"is_lock" json:"is_lock" form:"is_lock" uri:"is_lock" `                              // 1禁用 2正常
	CreateTime    automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime    automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode  string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcUserWallet) TableName() string {
	return "fc_user_wallet"
}
