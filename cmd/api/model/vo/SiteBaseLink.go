package vo

type SiteBaseLinkResp struct {
	AppKey        string `gorm:"column:app_key" json:"app_key" form:"app_key" uri:"app_key" `                             // 标识
	AppLink       string `gorm:"column:app_link" json:"app_link" form:"app_link" uri:"app_link" `                         // 链接
	Content       string `gorm:"column:content" json:"content" form:"content" uri:"content" `                             // 内容标题
	ContentDetail string `gorm:"column:content_detail" json:"content_detail" form:"content_detail" uri:"content_detail" ` // 内容详情
	ImgLink       string `gorm:"column:img_link" json:"img_link" form:"img_link" uri:"img_link" `                         // 图片地址
	DownloadLink  string `gorm:"column:download_link" json:"download_link" form:"download_link" uri:"download_link" `     // 下载地址
}
