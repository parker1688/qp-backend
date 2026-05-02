package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcUserCodeBetLimit struct {
	BaseDos
	UserId          string             `gorm:"column:user_id;primary_key" json:"user_id" form:"user_id" uri:"user_id" `                             // 用户Id
	UserName        string             `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `                                 // 用户账号
	Flag            string             `gorm:"column:flag" json:"flag" form:"flag" uri:"flag" `                                                     // 标识
	CodeAmount      float64            `gorm:"column:code_amount" json:"code_amount" form:"code_amount" uri:"code_amount" `                         // 打码金额
	MinBetAmount    float64            `gorm:"column:min_bet_amount" json:"min_bet_amount" form:"min_bet_amount" uri:"min_bet_amount" `             // 流水要求
	FinishBetAmount float64            `gorm:"column:finish_bet_amount" json:"finish_bet_amount" form:"finish_bet_amount" uri:"finish_bet_amount" ` // 完成流水
	TotalBetAmount  float64            `gorm:"column:total_bet_amount" json:"total_bet_amount" form:"total_bet_amount" uri:"total_bet_amount" `     // 总流水
	Status          int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                             // 1:待解锁 2：打码中  3：已解锁
	Remark          string             `gorm:"column:remark" json:"remark" form:"remark" uri:"remark" `                                             // 备注
	Multiple        int                `gorm:"column:multiple" json:"multiple" form:"multiple" uri:"multiple" `                                     // 流水倍数
	MerchantCode    string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户code
	CreateTime      automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `            // 创建时间
	CreateBy        string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                 // 创建人
	UpdateTime      automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `            // 修改时间
	UpdateBy        string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                 // 修改人
}

func (FcUserCodeBetLimit) TableName() string {
	return "fc_user_code_bet_limit"
}
