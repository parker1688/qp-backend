package vo

type AgentSiteInfoResp struct {
	AppKey  string `gorm:"column:app_key" json:"app_key" form:"app_key" uri:"app_key" `     // 标识
	AppLink string `gorm:"column:app_link" json:"app_link" form:"app_link" uri:"app_link" ` // 地址
	ShowDoc string `gorm:"column:show_doc" json:"show_doc" form:"show_doc" uri:"show_doc" ` // 显示信息
}
