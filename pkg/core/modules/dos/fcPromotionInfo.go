package dos

import (
	"bootpkg/common/expands/automaticType"
)

type DateRangeDataAmountsField struct {
	AmountScope []float64 `json:"area"`  // 金额范围
	Prob        float64   `json:"trans"` // 概率
}

type DateRangeDataField struct {
	TimeScope []string                    `json:"sj"` // 时间范围
	Amounts   []DateRangeDataAmountsField `json:"other"`
}

type FcPromotionInfo struct {
	BaseDos
	ClientType           string              `gorm:"column:client_type" json:"client_type" form:"client_type" uri:"client_type" `                                                       // h5,web,android,ios
	Language             string              `gorm:"column:language" json:"language" form:"language" uri:"language" `                                                                   // 语言简码
	MerchantCode         string              `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `                                               // 商户code
	PromotionType        int                 `gorm:"column:promotion_type" json:"promotion_type" form:"promotion_type" uri:"promotion_type" `                                           // 活动类型 1. 限时活动 2. 存款活动 3 日常活动
	GameType             int                 `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                                                               // 游戏类型 0 默认显示 1 体育 2 真人 3 棋牌 4 电子 5 捕鱼  6  彩票  7  区块链
	Title                string              `gorm:"column:title" json:"title" form:"title" uri:"title" validate:"max=10"`                                                              // 活动标题
	Content              string              `gorm:"column:content" json:"content" form:"content" uri:"content" `                                                                       // 活动详情
	Sort                 int                 `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                                                                   // 排序(最大越靠前)
	PromotionImg         string              `gorm:"column:promotion_img" json:"promotion_img" form:"promotion_img" uri:"promotion_img" `                                               // 活动图片
	H5Img                string              `gorm:"column:h5_img" json:"h5_img" form:"h5_img" uri:"h5_img" `                                                                           // 活动h5图片
	Link                 string              `gorm:"column:link" json:"link" form:"link" uri:"link" `                                                                                   // 活动链接
	H5Link               string              `gorm:"column:h5_link" json:"h5_link" form:"h5_link" uri:"h5_link" `                                                                       // 活动h5链接
	CreateTime           automaticType.Time  `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `                                          // 创建时间
	CreateBy             string              `gorm:"column:create_by" json:"create_by" form:"create_by" uri:"create_by" `                                                               // 创建人
	UpdateTime           automaticType.Time  `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `                                          // 修改时间
	UpdateBy             string              `gorm:"column:update_by" json:"update_by" form:"update_by" uri:"update_by" `                                                               // 修改人
	StartTime            automaticType.Time  `gorm:"column:start_time" json:"start_time" form:"start_time" uri:"start_time" `                                                           // 开始时间
	EndTime              automaticType.Time  `gorm:"column:end_time" json:"end_time" form:"end_time" uri:"end_time" `                                                                   // 结束时间
	Status               int                 `gorm:"column:status" json:"status" form:"status" uri:"status" `                                                                           // 状态 1:开启 2:关闭
	H5Content            string              `gorm:"column:h5_content" json:"h5_content" form:"h5_content" uri:"h5_content" `                                                           // 活动h5详情
	StageContent         string              `gorm:"column:stage_content" json:"stage_content" form:"stage_content" uri:"stage_content" `                                               // 阶段详情
	GiftStyle            int                 `gorm:"column:gift_style;default:0" json:"gift_style" form:"gift_style" uri:"gift_style" `                                                 // 礼包样式
	RechargeBalanceRatio float64             `gorm:"column:recharge_balance_ratio;default:0" json:"recharge_balance_ratio" form:"recharge_balance_ratio" uri:"recharge_balance_ratio" ` // 充值余额比
	Balance              float64             `gorm:"column:balance;default:0" json:"balance" form:"balance" uri:"balance" `                                                             // 余额
	FirstRechargeAmount  float64             `gorm:"column:first_recharge_amount;default:0" json:"first_recharge_amount" form:"first_recharge_amount" uri:"first_recharge_amount" `     // 首充金额
	BonusAmount          float64             `gorm:"column:bonus_amount;default:0" json:"bonus_amount" form:"bonus_amount" uri:"bonus_amount" `                                         // 奖励金额
	RegStartTime         *automaticType.Time `gorm:"column:reg_start_time;default:null" json:"reg_start_time" form:"reg_start_time" uri:"reg_start_time" `                              // 注册开始时间
	RegEndTime           *automaticType.Time `gorm:"column:reg_end_time;default:null" json:"reg_end_time" form:"reg_end_time" uri:"reg_end_time" `                                      // 注册结束时间
	Cycle                int                 `gorm:"column:cycle;default:0" json:"cycle" form:"cycle" uri:"cycle" `                                                                     // 活动周期
	DateRangeData        string              `gorm:"column:date_range_data" json:"date_range_data" form:"date_range_data" uri:"date_range_data" `                                       // 时间段数据
}

func (FcPromotionInfo) TableName() string {
	return "fc_promotion_info"
}

type FcPromotionInfoPage struct {
	BaseDos
	PromotionType int                `gorm:"column:promotion_type" json:"promotion_type" form:"promotion_type" uri:"promotion_type" `  // 优惠类型 1. 限时活动 2. 存款活动 3 日常活动
	GameType      int                `gorm:"column:game_type" json:"game_type" form:"game_type" uri:"game_type" `                      // 游戏类型 0 默认显示 1 体育 2 真人 3 棋牌 4 电子 5 捕鱼  6  彩票  7  区块链
	Title         string             `gorm:"column:title" json:"title" form:"title" uri:"title" `                                      // 优惠标题
	Status        int                `gorm:"column:status" json:"status" form:"status" uri:"status" `                                  // 状态 1:开启 2:关闭
	PromotionImg  string             `gorm:"column:promotion_img" json:"promotion_img" form:"promotion_img" uri:"promotion_img" `      // 优惠图片
	H5Img         string             `gorm:"column:h5_img" json:"h5_img" form:"h5_img" uri:"h5_img" `                                  // 优惠图片
	StartTime     automaticType.Time `gorm:"column:start_time" json:"start_time" form:"start_time" uri:"start_time" `                  // 开始时间
	EndTime       automaticType.Time `gorm:"column:end_time" json:"end_time" form:"end_time" uri:"end_time" `                          // 结束时间
	Link          string             `gorm:"column:link" json:"link" form:"link" uri:"link" `                                          // 优惠链接
	H5Link        string             `gorm:"column:h5_link" json:"h5_link" form:"h5_link" uri:"h5_link" `                              // h5活动链接
	CreateTime    automaticType.Time `gorm:"column:create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" ` // 创建时间
	CreateBy      string             `gorm:"column:create_by" json:"-" form:"create_by" uri:"create_by" `                              // 创建人
	UpdateTime    automaticType.Time `gorm:"column:update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" ` // 修改时间
	UpdateBy      string             `gorm:"column:update_by" json:"-" form:"update_by" uri:"update_by" `                              // 修改人
	Content       string             `gorm:"column:content" json:"content" form:"content" uri:"content" `                              // 优惠类型详情
	Sort          int                `gorm:"column:sort" json:"sort" form:"sort" uri:"sort" `                                          // 排序(最大越靠前)
	ClientType    string             `gorm:"column:client_type" json:"client_type" form:"client_type" uri:"client_type" `              // h5,web,android,ios
	Language      string             `gorm:"column:language" json:"language" form:"language" uri:"language" `                          // 语言简码
	MerchantCode  string             `gorm:"column:merchant_code" json:"merchant_code" form:"merchant_code" uri:"merchant_code" `      // 商户code
}

func (FcPromotionInfoPage) TableName() string {
	return "fc_promotion_info"
}
