package req

type CommonReq struct {
	MerchantCode string `json:"merchant_code" form:"merchant_code" uri:"merchant_code"`
	Timestamp int64 `json:"timestamp" form:"timestamp" uri:"timestamp"`
	Data string `json:"data" form:"data" uri:"data"`
	Sign string `json:"sign" form:"sign" uri:"sign"`
}

type NotifyReq struct {
	CurrencyCode string `json:"currency_code" form:"currency_code" uri:"currency_code"`
	FromAddr string `json:"from_addr" form:"from_addr" uri:"from_addr"`
	ToAddr string `json:"to_addr" form:"to_addr" uri:"to_addr"`
	TxHash string `json:"tx_hash" form:"tx_hash" uri:"tx_hash"`
	Gas string `json:"gas" form:"gas" uri:"gas"`
	TransactionFee string `json:"transaction_fee" form:"transaction_fee" uri:"transaction_fee"`
	OrderSn string `json:"order_sn" form:"order_sn" uri:"order_sn"`
	Amount string `json:"amount" form:"amount" uri:"amount"`
	//RealAmount string `json:"real_amount" form:"real_amount" uri:"real_amount"`
	TransferUnixTime int64 `json:"transfer_unix_time" form:"transfer_unix_time" uri:"transfer_unix_time"`
	Status int `json:"status" form:"status" uri:"status"`
	Event  string `json:"event" form:"event" uri:"event"`
}

type NotifyRsep struct {
	Confirm bool  `json:"confirm" form:"confirm" uri:"confirm"`
}



type CommonResp struct {
	Code   int64  `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Data         string `json:"data"`
		MerchantCode string `json:"merchant_code"`
		Sign         string `json:"sign"`
	} `json:"result"`
	Success   bool  `json:"success"`
	Timestamp int64 `json:"timestamp"`
}