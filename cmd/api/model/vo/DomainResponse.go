package vo

type DomainResponse struct {
	DomainLink string `gorm:"column:domain_link" json:"domain_link" form:"domain_link" uri:"domain_link" `
	DomainType string `gorm:"column:domain_type" json:"domain_type" form:"domain_type" uri:"domain_type" ` // 1. 预埋域名 2. API域名 3. websocket 域名
	Sort       int    `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                             // 权重排序
}
