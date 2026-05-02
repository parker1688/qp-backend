package enmus

const (
	// 任务类型
	DailyTaskType_None     = 0 // 无类型（只用于枚举判断无实际含义）
	DailyTaskType_Daily    = 1 // 每日任务
	DailyTaskType_Recharge = 2 // 充值任务
	DailyTaskType_Elec     = 3 // 电子任务
	DailyTaskType_Chess    = 4 // 棋牌任务
	DailyTaskType_Fish     = 5 // 捕鱼任务
	DailyTaskType_Sport    = 6 // 体育任务

	// 任务目标
	DailyTaskSubType_Pay = 1 // 累计充值
	DailyTaskSubType_Bet = 2 // 累计投注
	DailyTaskSubType_Los = 3 // 累计亏损
	DailyTaskSubType_Win = 4 // 累计盈利

	// 任务状态
	DailyTaskStats_Normal = 1 // 进行中
	DailyTaskStats_Over   = 2 // 已结束

	// 任务周期
	DailyTaskCycle_Forever = 0 // 永久
	DailyTaskCycle_Daily   = 1 // 自然日
	DailyTaskCycle_Week    = 2 // 自然周（每周一）
	DailyTaskCycle_Month   = 3 // 自然月（每月1号）

	// 用户任务状态
	UserTaskStats_None     = 0 // 不可领
	UserTaskStats_Reward   = 1 // 可领取
	UserTaskStats_Rewarded = 2 // 已领取
)

var (
	EnumDailyTaskTypeMap = map[string]int{
		// game type => task type
		"elecgame": DailyTaskType_Elec,
		"chess":    DailyTaskType_Chess,
		"fish":     DailyTaskType_Fish,
		"sport":    DailyTaskType_Sport,
	}
)
