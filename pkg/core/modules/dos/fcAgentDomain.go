package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcAgentDomain struct {
	BaseDos
	AgentName    string             `gorm:"column:agent_name" json:"agent_name" form:"agent_name" uri:"agent_name" `
	InviteCode   int                `gorm:"column:invite_code" json:"invite_code" form:"invite_code" uri:"invite_code" `              // 用户邀请短码
	Domain       string             `gorm:"column:domain" json:"domain" form:"domain" uri:"domain" `                                  // 链接
	ShortLink    string             `gorm:"column:short_link" json:"short_link" form:"short_link" uri:"short_link" `                  // 推广短码
	CreateTime   automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy     string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime   automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy     string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	AgentId      string             `gorm:"column:agent_id" json:"agent_id" form:"agent_id" uri:"agent_id" `
	MerchantCode string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
	Status       int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                             // 读取状态 1正常 2 停用
	JumpLink     string             `gorm:"column:jump_link" json:"jump_link" form:"jump_link" uri:"jump_link" `                 // 跳转后域名
	Type         int                `gorm:"column:type" json:"type" form:"type" uri:"type"`                                      // 类型（1 官网 2 推广）
	CustomerLink string             `gorm:"column:customer_link" json:"customer_link" form:"customer_link" uri:"customer_link"`  // 客服链接
	IosLink      string             `gorm:"column:ios_link" json:"ios_link" form:"ios_link" uri:"ios_link"`                      // ios链接
	IosLink2     string             `gorm:"column:ios_link2" json:"ios_link2" form:"ios_link2" uri:"ios_link2"`                  // ios备用链接
	AndroidLink  string             `gorm:"column:android_link" json:"android_link" form:"android_link" uri:"android_link"`      // 安卓链接
	AndroidLink2 string             `gorm:"column:android_link2" json:"android_link2" form:"android_link2" uri:"android_link2"`  // 安卓备用链接
	BannerImg    string             `gorm:"column:banner_img" json:"banner_img" form:"banner_img" uri:"banner_img"`              // banner图
	LogoImg      string             `gorm:"column:logo_img" json:"logo_img" form:"logo_img" uri:"logo_img"`                      // logo图
}

func (FcAgentDomain) TableName() string {
	return "fc_agent_domain"
}

type FcAgentDomainExt struct {
	FcAgentDomain
	Merchant     FcMerchant     `json:"merchant" gorm:"foreignkey:MerchantCode;references:MerchantCode"`
	MerchantLink FcMerchantLink `json:"merchant_link" gorm:"foreignkey:MerchantCode;references:MerchantCode"`
}

type FcAgentDomainResp struct {
	FcAgentDomain
	MerchantName string `json:"merchant_name"`
}

type FcAgentDomainOptionResp struct {
	Id           string `json:"id"`
	CustomerLink string `json:"customer_link"`
	ShortLink    string `json:"short_link"`
	JumpLink     string `json:"jump_link"`
	Domain       string `json:"domain"`
}
