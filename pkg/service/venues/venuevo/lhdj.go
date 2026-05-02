package venuevo

import "time"

type LHDJBalanceResp struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID              int       `json:"id"`
		OperatorID      string    `json:"operator_id"`
		MemberID        string    `json:"member_id"`
		Currency        string    `json:"currency"`
		Balance         float64   `json:"balance"`
		LastBetDatetime time.Time `json:"last_bet_datetime"`
		LastBetAmount   int       `json:"last_bet_amount"`
	} `json:"results"`
}

type LHDJDepositResp struct {
	Member          string  `json:"member"`
	OperatorID      int     `json:"operator_id"`
	Amount          float64 `json:"amount"`
	ReferenceNo     string  `json:"reference_no"`
	Currency        string  `json:"currency"`
	TransactionType string  `json:"transaction_type"`
	BalanceAmount   float64 `json:"balance_amount"`
}

type LHDJWithdrawResp struct {
	Member          string  `json:"member"`
	OperatorID      int     `json:"operator_id"`
	Amount          float64 `json:"amount"`
	ReferenceNo     string  `json:"reference_no"`
	Currency        string  `json:"currency"`
	TransactionType string  `json:"transaction_type"`
	BalanceAmount   float64 `json:"balance_amount"`
}
