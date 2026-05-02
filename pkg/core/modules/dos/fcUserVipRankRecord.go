package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserVipRankRecord struct {
	BaseDos
	UserId         string             `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName       string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `     // 用户名
	BeforRank      string             `gorm:"column:befor_rank" json:"befor_rank" form:"befor_rank" uri:"befor_rank" ` // 升级钱层级
	Level          int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                     // 升级后层级
	Rank           string             `gorm:"column:rank" json:"rank" form:"rank" uri:"rank" `
	TotalBetAmount float64            `gorm:"column:total_bet_amount" json:"total_bet_amount" form:"total_bet_amount" uri:"total_bet_amount" ` // 总投注额度
	CreateTime     automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `        // 创建时间
	CreateBy       string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                             // 创建人
	UpdateTime     automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `        // 修改时间
	UpdateBy       string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                             // 修改人
	MerchantCode   string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `             // 商户code
	Bonus          float64            `gorm:"column:bonus" json:"bonus" form:"bonus" uri:"bonus" `                                             // 奖金
}

func (FcUserVipRankRecord) TableName() string {
	return "fc_user_vip_rank_record"
}
