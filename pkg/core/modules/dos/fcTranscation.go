package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcTranscation struct {
	BaseDos
	UserId            string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName          string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Status            int                `gorm:"status" json:"status" form:"status" uri:"status" `                                                 // 1:通过  2：不通过
	Amount            float64            `gorm:"amount" json:"amount" form:"amount" uri:"amount" `                                                 // 金额
	TranscationBefore float64            `gorm:"transcation_before" json:"transcation_before" form:"transcation_before" uri:"transcation_before" ` // 账变前金额
	TranscationAfter  float64            `gorm:"transcation_after" json:"transcation_after" form:"transcation_after" uri:"transcation_after" `
	Remark            string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `
	FundingType       int                `gorm:"funding_type" json:"funding_type" form:"funding_type" uri:"funding_type" `                     // 资金类型：  1：公司入款  2：在线存款 3：手动存款 4：提款 5：优惠 6：返水 7：手续费 8：管理员添加 9：管理员扣除 10：额度转换
	CreateTime        automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `            // 创建时间
	CreateBy          string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                                 // 创建人
	UpdateTime        automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `            // 修改时间
	UpdateBy          string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                                 // 修改人
	MerchantCode      string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                 // 商户code
	Currency          string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                                     // 币种类型
	RelatedId         string             `gorm:"related_id" json:"related_id" form:"related_id" uri:"related_id" `                             // 关联id（订单号）
	ManualRelatedId   string             `gorm:"manual_related_id" json:"manual_related_id" form:"manual_related_id" uri:"manual_related_id" ` // 人工关联id
	FundingSubtype    string             `gorm:"funding_subtype" json:"funding_subtype" form:"funding_subtype" uri:"funding_subtype" `         // 资金子类型
}

func (FcTranscation) TableName() string {
	return "fc_transcation"
}

type FcTranscationSharding struct {
	BaseDos
	UserId            string             `gorm:"user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName          string             `gorm:"user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Status            int                `gorm:"status" json:"status" form:"status" uri:"status" `                                                 // 1:通过  2：不通过
	Amount            float64            `gorm:"amount" json:"amount" form:"amount" uri:"amount" `                                                 // 金额
	TranscationBefore float64            `gorm:"transcation_before" json:"transcation_before" form:"transcation_before" uri:"transcation_before" ` // 账变前金额
	TranscationAfter  float64            `gorm:"transcation_after" json:"transcation_after" form:"transcation_after" uri:"transcation_after" `
	Remark            string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `
	FundingType       int                `gorm:"funding_type" json:"funding_type" form:"funding_type" uri:"funding_type" `                           // 资金类型：  1：公司入款  2：在线存款 3：手动存款 4：提款 5：优惠 6：返水 7：手续费 8：管理员添加 9：管理员扣除 10：额度转换
	CreateTime        automaticType.Time `gorm:"create_time;default:null;type:datetime(3)" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy          string             `gorm:"create_by" json:"create_by" form:"create_by" uri:"create_by" `                                       // 创建人
	UpdateTime        automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                  // 修改时间
	UpdateBy          string             `gorm:"update_by" json:"update_by" form:"update_by" uri:"update_by" `                                       // 修改人
	MerchantCode      string             `gorm:"merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                       // 商户code
	Currency          string             `gorm:"currency" json:"currency" form:"currency" uri:"currency" `                                           // 币种类型
	FundingSubtype    string             `gorm:"funding_subtype" json:"funding_subtype" form:"funding_subtype" uri:"funding_subtype" `               // 资金子类型
}

func (FcTranscationSharding) TableName() string {
	return "fc_transcation"
}

type FcTranscationExt struct {
	FcTranscation
	VenueTransfer  FcVenueTransfer  `json:"venue_transfer" gorm:"foreignkey:OrderSn;references:RelatedId"`
	OrderManageOpt FcOrderManageOpt `json:"order_manage_opt" gorm:"foreignkey:Id;references:ManualRelatedId"`
}
