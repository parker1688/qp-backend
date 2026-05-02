package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenueTransfer struct {
	BaseDos
	OrderSn      string             `gorm:"column:order_sn" json:"order_sn" form:"order_sn" uri:"order_sn" `                          // 场馆转账订单号
	VenueCode    string             `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                  // 场馆Code
	VenueAccount string             `gorm:"column:venue_account" json:"venue_account" form:"venue_account" uri:"venue_account" `      // 场馆账户
	VenueLine    int                `gorm:"column:venue_line" json:"venue_line" form:"venue_line" uri:"venue_line" `                  // 线路ID
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 用户名称
	Currency     string             `gorm:"column:currency" json:"currency" form:"currency" uri:"currency" `                          // 币种简码
	Amount       float64            `gorm:"column:amount" json:"amount" form:"amount" uri:"amount" `                                  // 操作金额
	OptType      int                `gorm:"column:opt_type" json:"opt_type" form:"opt_type" uri:"opt_type" `                          // 1 转入 2 转出
	Ip           string             `gorm:"column:ip" json:"ip" form:"ip" uri:"ip" `                                                  // 执行IP
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 0 处理中 1 成功 2 失败
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
}

func (FcVenueTransfer) TableName() string {
	return "fc_venue_transfer"
}
