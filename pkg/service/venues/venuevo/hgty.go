package venuevo

type HGTYCreatUserEncryptReq struct {
	Method    string `json:"method"`
	Password  string `json:"password"`
	Memname   string `json:"memname"`
	Currency  string `json:"currency"`
	Token     string `json:"token"`
	Timestamp string `json:"timestamp"`
}

type HGTYCreatUserReq struct {
	Request string `json:"Request"`
	HGTYCommonReq
}

type HGTYCommonReq struct {
	Method string `json:"Method"`
	AGID   string `json:"AGID"`
}

type HGTYCreateUserResp struct {
	Userdata  *HGTYCreateUserdata `json:"userdata"`
	Method    string              `json:"method"`
	Responeid string              `json:"responeid"`
	Respcode  string              `json:"respcode"`
	Status    string              `json:"status"`
	Timestamp string              `json:"timestamp"`
}

type HGTYCreateUserdata struct {
	UserGold     string `json:"user_gold"`
	UserUsername string `json:"user_username"`
	UserEnable   string `json:"user_enable"`
	UserCurrency string `json:"user_currency"`
}

type HGTYBalanceReq struct {
	Request string `json:"Request"`
	HGTYCommonReq
}

type HGTYBalanceEncryptReq struct {
	Method    string `json:"method"`
	Memname   string `json:"memname"`
	Token     string `json:"token"`
	Timestamp string `json:"timestamp"`
}

type HGTYBalanceResp struct {
	Memname   string `json:"memname"`
	Balance   string `json:"balance"`
	Method    string `json:"method"`
	Currency  string `json:"currency"`
	Responeid string `json:"responeid"`
	Respcode  string `json:"respcode"`
	Memid     string `json:"memid"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type HGTYDepositReq struct {
	Request string `json:"Request"`
	HGTYCommonReq
}

type HGTYDepositEncryptReq struct {
	Method    string `json:"method"`
	Memname   string `json:"memname"`
	Amount    string `json:"amount"`
	Payno     string `json:"payno"`
	Token     string `json:"token"`
	Timestamp string `json:"timestamp"`
}

type HGTYDepositResp struct {
	Method    string    `json:"method"`
	Responeid string    `json:"responeid"`
	Respcode  string    `json:"respcode"`
	Moneydata Moneydata `json:"moneydata"`
	Status    string    `json:"status"`
	Timestamp string    `json:"timestamp"`
}

type Moneydata struct {
	Gold     string `json:"gold"`
	Payno    string `json:"payno"`
	PayGold  string `json:"pay_gold"`
	Currency string `json:"currency"`
	PayWay   string `json:"pay_way"`
	Recid    string `json:"recid"`
	Username string `json:"username"`
}

type HGTYWithdrawReq struct {
	Request string `json:"Request"`
	HGTYCommonReq
}

type HGTYWithdrawEncryptReq struct {
	Method    string `json:"method"`
	Memname   string `json:"memname"`
	Amount    string `json:"amount"`
	Payno     string `json:"payno"`
	Token     string `json:"token"`
	Timestamp string `json:"timestamp"`
}

type HGTYWithdrawResp struct {
	Method    string    `json:"method"`
	Responeid string    `json:"responeid"`
	Respcode  string    `json:"respcode"`
	Moneydata Moneydata `json:"moneydata"`
	Status    string    `json:"status"`
	Timestamp string    `json:"timestamp"`
}

type HGTYLaunchGameReq struct {
	Request string `json:"Request"`
	HGTYCommonReq
}

type HGTYLaunchGameEncryptReq struct {
	Memname   string `json:"memname"`
	Password  string `json:"password"`
	Machine   string `json:"machine"`
	Langx     string `json:"langx"`
	IsSSL     string `json:"isSSL"`
	Token     string `json:"token"`
	Currency  string `json:"currency"`
	Remoteip  string `json:"remoteip"`
	Timestamp string `json:"timestamp"`
}

type HGTYLaunchGameResp struct {
	Timestamp     string `json:"timestamp"`
	Respcode      string `json:"respcode"`
	Responeid     string `json:"responeid"`
	Status        string `json:"status"`
	Launchgameurl string `json:"launchgameurl"`
	Method        string `json:"method"`
	Memname       string `json:"memname"`
	Memid         string `json:"memid"`
	MemToken      string `json:"memToken"`
	Remark        string `json:"remark"`
}

type HGTYChkTransInfoReq struct {
	Request string `json:"Request"`
	HGTYCommonReq
}

type HGTYChkTransInfoEncryptReq struct {
	Memname     string `json:"memname"`
	Transidtype string `json:"transidtype"`
	Transid     string `json:"transid"`
	Token       string `json:"token"`
	Timestamp   string `json:"timestamp"`
}

type HGTYChkTransInfoResp struct {
	Method    string       `json:"method"`
	Responeid string       `json:"responeid"`
	Respcode  string       `json:"respcode"`
	Transdata []*Transdata `json:"transdata"`
	Status    string       `json:"status"`
	Timestamp string       `json:"timestamp"`
}

type Transdata struct {
	Date     string `json:"date"`
	Gold     string `json:"gold"`
	Payno    string `json:"payno"`
	Memname  string `json:"memname"`
	Pay      string `json:"pay"`
	Currency string `json:"currency"`
	Id       string `json:"id"`
	Paycash  string `json:"paycash"`
	Recid    string `json:"recid"`
	Memid    string `json:"memid"`
}

type HGTYLoginReq struct {
	Request string `json:"Request"`
	HGTYCommonReq
}

type HGTYLoginEncryptReq struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
}

type HGTYLoginResp struct {
	Timestamp string `json:"timestamp"`
	Responeid string `json:"responeid"`
	Username  string `json:"username"`
	Respcode  string `json:"respcode"`
	Status    string `json:"status"`
	Token     string `json:"token"`
	Method    string `json:"method"`
	Aid       string `json:"aid"`
}
