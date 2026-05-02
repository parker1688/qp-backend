package response

import "bootpkg/common/ecode"

var (
	ERROR_PARAMETER = ecode.AddCode(501, "参数解析错误")
	ERROR_SERVER    = ecode.AddCode(504, "服务器错误")

	Merchant_Not_Exsit     = ecode.AddCode(1, "商户不存在")
	Params_Is_Bad          = ecode.AddCode(2, "参数错误")
	Params_Not_Unmarshal   = ecode.AddCode(3, "参数解析错误")
	Address_not_exsit      = ecode.AddCode(4, "地址不存在")
	Address_type_Only_One  = ecode.AddCode(5, "该类型地址不能重复添加")
	Member_not_exsit       = ecode.AddCode(6, "会员不存在")
	Payment_Strategy_Limit = ecode.AddCode(7, "支付风控限制")

	ERROR_DEFAULT = ecode.AddCode(500, "默认错误")
)
