package enmus

const (
	GameOrderRedisKey        = "PullOrderGameOrderRedisKey::%s"
	TidbGameOrderRedisKey    = "TidbPullOrderGameOrderRedisKey::%s"
	GameNameRedisKey         = "PullOrderGameNameRedisKey::%s"
	VenueAccountUserRedisKey = "PullOrderVenueAccountUserRedisKey::%s::%s"
	VenuePGDZversionRedisKey = "PullOrderVenuePGDZversionRedisKey"
	UserBetRecordRecordKEY   = "UserBetRecordRecordVIPKEY:%s"
	UserUnlockBetAmountKEY   = "UserUnlockBetAmountKEYPKEY:%s"

	VenueVersionRedisKey = "VenueVersionRedisKey::%s" //捞取注单版本号缓存
	VenuLastUidRedisKey  = "VenuLastUidRedisKey::%s"  //捞取注单最后一个订单

	UserTotalRechargeAmountKey = "UserTotalRechargeAmountKey::%s" //用户充值总金额
	UserTotalBetAmountKey      = "UserTotalBetAmountKey::%s"
	UserVipNextLevelAmountKey  = "UserVipNextLevelAmount::%s"

	UserWeekTotalBetAmountKey  = "UserWeekTotalBetAmountKey::%s::%d::%d::%d" //按周统计流水
	UserMonthTotalBetAmountKey = "UserMonthTotalBetAmountKey::%s::%d::%d"    //按月统计流水

	UserVipNextLevelTotalBetAmountKey = "UserVipNextLevelTotalBetAmount::%s" //下一个VIP 累计投注额

	UserRebateFlow          = "User_RebateFlow:%s"            // 用户返水打码
	UserHisRebateFlow       = "User_His_RebateFlow:%s"        // 用户各游戏类型区间返水打码累积
	GetUserRebateFlowLock   = "Get_User_RebateFlow_Lock:%s"   // 获取用户打码量
	ApplyUserRebateFlowLock = "Apply_User_RebateFlow_Lock:%s" // 用户洗码

	Merchant_Agent_Invite_Code_Auto = "Merchant_Agent_Invite_Code_Auto"
	Member_Agent_Invite_Code_Auto   = "Member_Agent_Invite_Code_Auto"

	DEPOSIT_NO_KEY = "DEPOSIT_NO:%s"

	First_Deposite_UserId_Key = "First_Deposite_UserId_Key:%s" // 用户首存 key

	UserSiteMessageKey = "User_Site_Message:%s:%s" // 用户站内信 key,第一个为 msgId, 第二个为userId
)
