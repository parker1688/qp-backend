package enmus

const (
	MailStats_Unread = 1 // 未读
	MailStats_Readed = 2 // 已读

	// 用户删除状态（该标记用于是否显示在用户邮件列表中，但还要显示后台邮件列表）
	MailDelStats_No      = 0 // 未删除
	MailDelStats_Yes     = 1 // 已删除（用户删除）
	MailDelStats_Destroy = 2 // 被删除（后台删除）

	MailType_Manual          = 0 // 人工操作（个人消息）
	MailType_FirstLogin      = 1 // 首次登录（系统消息）
	MailType_RechargeSuccess = 2 // 充值成功（系统消息）
	MailType_WithdrawFail    = 3 // 提款失败（系统消息）
	MailType_WithdrawSuccess = 4 // 提款成功（系统消息）
)
