package dos

type FcClientLog struct {
	BaseDos
	AgentId      int    `gorm:"column:agent_id" json:"agent_id" form:"agent_id" uri:"agent_id" `                     // 推广id
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" ` // 商户code
	MerchantName string `gorm:"column:merchant_name" json:"merchant_name" form:"merchant_name" uri:"merchant_name" ` // 商户名
	IP           string `gorm:"column:ip" json:"ip" form:"ip" uri:"ip" `                                             // 访问ip
	Address      string `gorm:"column:address" json:"address" form:"address" uri:"address" `                         // 归属地
	Device       int    `gorm:"column:device" json:"device" form:"device" uri:"device" `                             // 设备：0h5,1ios,2安卓
	Download     int    `gorm:"column:download" json:"download" form:"download" uri:"download" `                     // 是否下载：0否1是
	Customer     int    `gorm:"column:customer" json:"customer" form:"customer" uri:"customer" `                     // 是否联系客服：0否1是
	VisitTime    string `gorm:"column:visit_time" json:"visit_time" form:"visit_time" uri:"visit_time" `             // 访问时间
}

func (FcClientLog) TableName() string {
	return "fc_client_log"
}

type FcClientLogResp struct {
	TotalVisitCount    int `json:"total_visit_count"`    //总访问次数
	VisitCount1        int `json:"visit_count1"`         //ios
	VisitCount2        int `json:"visit_count2"`         //androw
	VisitCount3        int `json:"visit_count3"`         //other
	TotalDownloadCount int `json:"total_download_count"` //总下载次数
	DownloadCount1     int `json:"download_count1"`      //ios
	DownloadCount2     int `json:"download_count2"`      //androw
	CustomerCount      int `json:"customer_count"`       //联系客服量
}
