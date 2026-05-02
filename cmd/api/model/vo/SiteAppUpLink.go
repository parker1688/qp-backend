package vo

type SiteAppUpLink struct {
	AppVersion string `gorm:"column:app_version" json:"app_version" form:"app_version" uri:"app_version" ` // APP标识
	AppLink    string `gorm:"column:app_link" json:"app_link" form:"app_link" uri:"app_link" `             // app
	Content    string `gorm:"column:content" json:"content" form:"content" uri:"content" `                 // 描述更新内容
	Forcibly   int    `gorm:"column:forcibly" json:"forcibly" form:"forcibly" uri:"forcibly" `             // 1 非强制更新 1 强制更新
}
