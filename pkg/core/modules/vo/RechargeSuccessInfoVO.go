package vo

type RechargeSuccessInfoVO struct {
	Success  bool   `json:"success,omitempty"`
	Amount   string `json:"amount,omitempty" json:"amount,omitempty"`
	Currency string `json:"currency,omitempty" json:"currency,omitempty"`
}
