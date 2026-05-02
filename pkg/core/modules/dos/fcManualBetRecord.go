package dos

type FcManualBetRecordReq struct {
	VenueCode string `gorm:"column:venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `     // 场馆code
	StartAt   string `gorm:"column:create_time" json:"create_time" form:"create_time" uri:"create_time" ` // 起点时间
}

type FcManualBetRecordResp struct {
	Code int    `gorm:"column:code" json:"code" form:"code" uri:"code" ` // 错误码：0成功
	Msg  string `gorm:"column:msg" json:"msg" form:"msg" uri:"msg" `     //提示信息
}
