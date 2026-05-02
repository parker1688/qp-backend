package vo

type UserRechargeConfirmReq struct {
	PaymentId int     `json:"payment_id" form:"payment_id" uri:"payment_id" `                          //通道ID
	Amount    float64 `gorm:"id;primary_key;AUTO_INCREMENT" json:"amount" form:"amount" uri:"amount" ` //充值金额
}
