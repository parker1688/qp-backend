package dos

import (
	"bootpkg/common/expands/automaticType"
)

type DomainDistribution struct {
	BaseDos
	DomainLink string             `gorm:"column:domain_link" json:"domain_link" form:"domain_link" uri:"domain_link" `
	DomainType string             `gorm:"column:domain_type" json:"domain_type" form:"domain_type" uri:"domain_type" `              // 1. 预埋域名 2. API域名 3. websocket 域名
	CreateTime automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy   string             `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                      // 创建人
	UpdateTime automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy   string             `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                      // 修改人
	Sort       int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 权重排序
	MinLevel   int                `gorm:"column:min_level" json:"min_level" form:"min_level" uri:"min_level" `                      // 最小权重等级
	MaxLevel   int                `gorm:"column:max_level" json:"max_level" form:"max_level" uri:"max_level" `                      // 最大权重等级
}

func (DomainDistribution) TableName() string {
	return "domain_distribution"
}
