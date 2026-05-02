package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserLevelWeekBounsRecord struct {
	BaseDos
	UserId       string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	BetType      int                `gorm:"column:bet_type" json:"bet_type" form:"bet_type" uri:"bet_type" ` // 1: 体育电竞 2：棋牌  3：彩票
	Level        int                `gorm:"column:level" json:"level" form:"level" uri:"level" `
	Bouns        float64            `gorm:"column:bouns" json:"bouns" form:"bouns" uri:"bouns" `                                      // 奖金
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcUserLevelWeekBounsRecord) TableName() string {
	return "fc_user_level_week_bouns_record"
}
