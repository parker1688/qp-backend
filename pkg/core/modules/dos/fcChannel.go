package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcChannel struct {
	BaseDos
	ChannelName  string             `gorm:"channel_name" json:"channel_name" form:"channel_name" uri:"channel_name" `
	ChannelCode  string             `gorm:"channel_code" json:"channel_code" form:"channel_code" uri:"channel_code" `          // 通道code
	Currency     string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                          // 货币类型，虚拟币就是协议
	Remark       string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `                                  // 备注
	Type         int                `gorm:"type" json:"type" form:"type" uri:"type" `                                          // 1:存款  2:提款
	IsBlockchain int                `gorm:"is_blockchain" json:"is_blockchain" form:"is_blockchain" uri:"is_blockchain" `      // 1:虚拟币  2：法币
	MinAmount    float64            `gorm:"min_amount" json:"min_amount" form:"min_amount" uri:"min_amount" `                  // 最小金额
	MaxAmount    float64            `gorm:"max_amount" json:"max_amount" form:"max_amount" uri:"max_amount" `                  // 最大金额
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
	MerchantCode string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcChannel) TableName() string {
	return "fc_channel"
}
