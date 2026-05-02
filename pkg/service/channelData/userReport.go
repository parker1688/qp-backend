package channelData

import (
	"bootpkg/common/tool"
	"errors"
)

func safeSend(topic, payload string) error {
	if producerOtp == nil {
		return errors.New("kafka producer not initialized")
	}
	err := producerOtp.Send(topic, payload)
	if err != nil {
		err = producerOtp.Send(topic, payload)
	}
	return err
}

type UserRechargeMessage struct {
	UserId        string  `json:"user_id"`        //用户ID
	UserName      string  `json:"user_name"`      //用户名
	OrderSn       string  `json:"order_sn"`       //存款单号
	DepositTime   string  `json:"deposit_time"`   //存款时间
	DepositAmount float64 `json:"deposit_amount"` //存款金额
	T             int64   `json:"t"`              //时间戳微妙
	ForceStatus   int     `json:"force_status"`   //强制状态  1 强制更新  0 默认
}

// SendUserRecharge
//
//	@Description: 发送充值成功用户金额
//	@param s 消息体
//	@return error 错误信息
func SendUserRecharge(s *UserRechargeMessage) error {
	return safeSend(Kakfa_Topic_User_recharge, tool.String(s))
}

// SendUserWithdrawal
//
//	@Description: 发送提款成功用户金额
//	@param s 消息体
//	@return error 错误信息
func SendUserWithdrawal(s *UserWithdrawalMessage) error {
	return safeSend(Kakfa_Topic_User_Withdrawal, tool.String(s))
}

type UserWithdrawalMessage struct {
	UserId           string  `json:"user_id"`           //用户ID
	UserName         string  `json:"user_name"`         //用户名
	OrderSn          string  `json:"order_sn"`          //存款单号
	WithdrawalTime   string  `json:"withdrawal_time"`   //提款时间
	WithdrawalAmount float64 `json:"withdrawal_amount"` //提款金额
	T                int64   `json:"t"`                 //时间戳微妙
	ForceStatus      int     `json:"force_status"`      //强制状态  1 强制更新  0 默认
}

// SendUserWithdrawal
//
//	@Description: 发送优惠成功用户金额
//	@param s 消息体
//	@return error 错误信息
func SendUserPromotion(s *UserPromotionMessage) error {
	return safeSend(Kakfa_Topic_User_Promotion, tool.String(s))
}

type UserPromotionMessage struct {
	UserId          string  `json:"user_id"`          //用户ID
	UserName        string  `json:"user_name"`        //用户名
	OrderSn         string  `json:"order_sn"`         //单号
	PromotionTime   string  `json:"promotion_time"`   //优惠时间
	PromotionAmount float64 `json:"promotion_amount"` //优惠金额
	T               int64   `json:"t"`                //时间戳微妙
	ForceStatus     int     `json:"force_status"`     //强制状态  1 强制更新  0 默认
	PromotionType   int     `json:"promotion_type"`   // 类型，1:  红利信息，2： 返水信息
}
