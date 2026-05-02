package ecode

var (
	PASSWORD_FAIL        = AddCode(100002, "密码错误")
	PASSWORD_ERR_NUM     = AddCode(100003, "密码错误次数过多")
	ACCOUNT_DISABLE      = AddCode(100004, "账号已被禁用")
	INPUT_VERFIYCODE     = AddCode(100004, "请输入验证码")
	VERFIYCODE_ERROR     = AddCode(100005, "验证码错误")
	WALLET_PASSWORD_ERR  = AddCode(100006, "支付密码错误")
	ACCOUNT_NOT_REGISTER = AddCode(100007, "账号没注册")
	MinWithdrawalAmount  = AddCode(100008, "最小提款金额")
	EAMIL_HAS_EXIST      = AddCode(100009, "邮箱已经绑定")
	VERIFY_CODE_ERR      = AddCode(100010, "验证码错误")
	LOGON_ABNORMAL       = AddCode(100011, "登录异常")
	ACCOUNT_IS_DISABLED  = AddCode(100012, "该账号已被停用，请联系管理员")
)
