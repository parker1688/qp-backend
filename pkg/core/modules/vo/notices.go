package vo

const (
	NOTICES_DEPOSIT       = "deposit"
	NOTICES_WITHDRAW      = "withdraw"
	NOTICES_Promotion     = "promotion"
	NOTICES_VenueTransfer = "venueTransfer"
)

type NoticesProVO struct {
	UnreadCount int64   `json:"unreadCount"` //未读数量
	TotalCount  int64   `json:"totalCount"`  //总数量
	CountData   []int64 `json:"countData""`
	//Data        []NoticesProDetailsVO `json:"data"`        //消息详情
	//Deposit       NoticesProDetailVO    `json:"deposit"`       //消息详情
	//Withdraw      NoticesProDetailVO    `json:"withdraw"`      //消息详情
	//Promotion     NoticesProDetailVO    `json:"promotion"`     //消息详情
	//VenueTransfer NoticesProDetailVO    `json:"venueTransfer"` //消息详情
	//Data          []NoticesProDetailsVO `json:"data"`          //消息详情
}

type NoticesProDetailVO struct {
	UnreadCount int64 `json:"unreadCount"` //未读数量
	TotalCount  int64 `json:"totalCount"`  //总未数量
}

type NoticesProDetailsVO struct {
	ID       string `json:"id"`
	Avatar   string `json:"avatar"`
	Title    string `json:"title"`
	Datetime string `json:"datetime"`
	Type     string `json:"type"` //类型
}

type SiteMsgVO struct {
	MsgId        string `json:"msg_id"`                                                                      // 消息 id
	MsgIdType    int    `gorm:"column:msg_id_type" json:"msg_id_type" form:"msg_id_type" uri:"msg_id_type" ` // 消息 id 类型, 1: 人工消息，2:模板消息id
	NotifyType   int    `gorm:"column:notify_type" json:"notify_type" form:"notify_type" uri:"notify_type" ` // 通知类型 1 全局消息 2 部分通知
	UserId       string `gorm:"column:user_id" json:"user_id" form:"user_id" uri:"user_id" `
	UserName     string `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `
	Title        string `gorm:"column:title" json:"title" form:"title" uri:"title" `             // 消息标题
	Content      string `gorm:"column:content" json:"content" form:"content" uri:"content" `     // 消息内容
	MsgType      int    `gorm:"column:msg_type" json:"msg_type" form:"msg_type" uri:"msg_type" ` // 消息类型 1 普通信息 2 赛事  3 充值 4提款 5 红利
	Language     string `gorm:"column:language" json:"language" form:"language" uri:"language" ` // 语言简码
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
	CreateBy     string `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" ` // 创建人
	UpdateBy     string `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" ` // 修改人
}

type SiteMsgSaveReq struct {
	MsgType      int    `gorm:"column:msg_type" json:"msg_type" form:"msg_type" uri:"msg_type" `             // 消息类型 1 普通信息 2 赛事  3 充值 4提款 5 红利
	MsgIdType    int    `gorm:"column:msg_id_type" json:"msg_id_type" form:"msg_id_type" uri:"msg_id_type" ` // 消息 id 类型, 1: 人工全局消息，2:模板消息id
	NotifyType   int    `gorm:"column:notify_type" json:"notify_type" form:"notify_type" uri:"notify_type" ` // 通知类型 1 全局消息 2 部分通知
	UserIds      string `gorm:"column:user_ids" json:"user_ids" form:"user_ids" uri:"user_ids" `             // 多个 user_id 用 , 隔开
	Title        string `gorm:"column:title" json:"title" form:"title" uri:"title" `                         // 消息标题
	Content      string `gorm:"column:content" json:"content" form:"content" uri:"content" `                 // 消息内容
	MerchantCode string `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `
}
