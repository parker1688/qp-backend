package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserShare struct {
	UserId       string             `gorm:"user_id;primary_key" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `                      // 用户名
	ShareLink    string             `gorm:"share_link" json:"share_link" form:"share_link" uri:"share_link" `                  // 分享链接
	ShareCode    string             `gorm:"share_code" json:"share_code" form:"share_code" uri:"share_code" `                  // 分享code
	Quantity     int                `gorm:"quantity" json:"quantity" form:"quantity" uri:"quantity" `                          // 累计邀请数量
	CreateTime   automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcUserShare) TableName() string {
	return "fc_user_share"
}
