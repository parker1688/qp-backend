package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenueMaintain struct {
	BaseDos
	VenueId       string             `gorm:"venue_id" json:"venue_id" form:"venue_id" uri:"venue_id" `         // 场馆ID
	VenueCode     string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" ` // 场馆code
	MaintainStart automaticType.Time `gorm:"maintain_start" json:"maintain_start" form:"maintain_start" uri:"maintain_start" `
	MaintainEnd   automaticType.Time `gorm:"maintain_end" json:"maintain_end" form:"maintain_end" uri:"maintain_end" `         // 维护结束时间
	CilentType    string             `gorm:"cilent_type" json:"cilent_type" form:"cilent_type" uri:"cilent_type" `             // 维护端：web,h5,android,ios
	AllowTransfer int                `gorm:"allow_transfer" json:"allow_transfer" form:"allow_transfer" uri:"allow_transfer" ` // 1:不允许转账 2:只允许转进场馆 3:只允许转出场馆 4:可以转进转出
	Remark        string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `
	CreateTime    automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime    automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode  string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcVenueMaintain) TableName() string {
	return "fc_venue_maintain"
}
