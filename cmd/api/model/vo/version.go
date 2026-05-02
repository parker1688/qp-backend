package vo

type VersionInfoReq struct {
}

type VersionInfoResp struct {
	Version string `gorm:"version" json:"version" form:"version" uri:"version" ` // 版本号
	CdnUrl  string `gorm:"cdn_url" json:"cdn_url" form:"cdn_url" uri:"cdn_url" ` // 排序
}
