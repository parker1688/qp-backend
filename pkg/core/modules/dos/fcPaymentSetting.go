package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcPaymentSetting struct {
	BaseDos
	PaymentCode  string             `gorm:"column:payment_code" json:"payment_code" form:"payment_code" uri:"payment_code" `          // 通道code
	PKey         string             `gorm:"column:p_key" json:"p_key" form:"p_key" uri:"p_key" `                                      // 支付Key键
	PValue       string             `gorm:"column:p_value" json:"p_value" form:"p_value" uri:"p_value" `                              // 支付Value值
	Sort         int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序：值越大越靠前
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	Remark       string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                  // 备注
}

func (FcPaymentSetting) TableName() string {
	return "fc_payment_setting"
}
