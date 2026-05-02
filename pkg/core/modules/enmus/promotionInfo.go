package enmus

const (
	// 活动配置类型
	ActivityTypes_HealthPack      = 11 // 回血包
	ActivityTypes_RedEnvelopeRain = 12 // 红包雨

	// 活动配置状态
	ActivityStats_Opening = 1 // 开启
	ActivityStats_Closed  = 2 // 关闭

	// 活动周期
	ActivityCycle_Forever = 0 // 永久
	ActivityCycle_Daily   = 1 // 自然日
	ActivityCycle_Week    = 2 // 自然周（每周一）
	ActivityCycle_Month   = 3 // 自然月（每月1号）

	// 用户活动状态
	UserActivityStats_None     = 0 // 不可领
	UserActivityStats_Reward   = 1 // 可领取
	UserActivityStats_Rewarded = 2 // 已领取
)
