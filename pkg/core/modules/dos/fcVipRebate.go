package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVipRebate struct {
	BaseDos
	Level        int                `gorm:"column:level" json:"level" form:"level" uri:"level" `                                      // VIP层级
	VipName      string             `gorm:"column:vip_name" json:"vip_name" form:"vip_name" uri:"vip_name" `                          // vip名称
	Chess        float64            `gorm:"column:chess" json:"chess" form:"chess" uri:"chess" `                                      // 棋牌返点
	Elecgame     float64            `gorm:"column:elecgame" json:"elecgame" form:"elecgame" uri:"elecgame" `                          // 电子游戏返点
	Zhenren      float64            `gorm:"column:zhenren" json:"zhenren" form:"zhenren" uri:"zhenren" `                              // 真人返点
	Sport        float64            `gorm:"column:sport" json:"sport" form:"sport" uri:"sport" `                                      // 体育返点
	Esport       float64            `gorm:"column:esport" json:"esport" form:"esport" uri:"esport" `                                  // 电竞返水
	Lottery      float64            `gorm:"column:lottery" json:"lottery" form:"lottery" uri:"lottery" `                              // 彩票返点
	Fish         float64            `gorm:"column:fish" json:"fish" form:"fish" uri:"fish" `                                          // 捕鱼
	RebateLimit  float64            `gorm:"column:rebate_limit" json:"rebate_limit" form:"rebate_limit" uri:"rebate_limit" `          // 返水上限
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
}

func (FcVipRebate) TableName() string {
	return "fc_vip_rebate"
}
