package vo

type GetPaymentBankReq struct {
	PaymentId int `json:"payment_id" form:"payment_id" uri:"payment_id" ` //通道ID
}
